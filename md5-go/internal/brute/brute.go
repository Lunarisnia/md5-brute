package brute

import (
	"context"
	"errors"
)

type GoalTest func(guess string) bool

type Brute interface {
	SetTextLength(length uint) Brute
	SetGoalTest(goalTest GoalTest) Brute
	SetStartRune(start rune) Brute
	SetEndRune(end rune) Brute
	Crack(ctx context.Context) (string, error)
}

type brute struct {
	goalTest   GoalTest
	textLength uint

	startRune rune
	endRune   rune
}

func New() Brute {
	return &brute{
		startRune: rune(32),
		endRune:   rune(126),
	}
}

func (b brute) SetGoalTest(goalTest GoalTest) Brute {
	b.goalTest = goalTest
	return b
}

func (b brute) SetTextLength(length uint) Brute {
	b.textLength = length
	return b
}

func (b brute) SetStartRune(start rune) Brute {
	b.startRune = start
	return b
}

func (b brute) SetEndRune(end rune) Brute {
	b.endRune = end
	return b
}

// NOTE: Its fair game to have the length, special conditions like only lowercase character set before brute forcing it
// NOTE: https://www.youtube.com/watch?v=7U-RbOKanYs&t=389s
// NOTE: basically a binary counter for chars
// NOTE: AAA > AAB > AAC > ABA > ABB > ABC > ACA > ACB > ACC >
// NOTE: BAA > BAB > BAC > BBA > BBB > BBC > BCA > BCB > BCC >
// NOTE: CAA > CAB > CAC > CBA > CBC > CCA > CCB > CCC

// NOTE: AAA > AAB > ABA > ABB > BAA > BAB > BBB
// NOTE: AA > AB > BA > BB
// NOTE: Crack the password until the goal test return true
func (b brute) Crack(ctx context.Context) (string, error) {
	// NOTE: Lowercase: 97-122 - Uppercase: 65-90
	// NOTE: 32(or 33)-126: https://www.w3schools.com/charsets/ref_utf_basic_latin.asp
	initialRune := b.startRune
	lastRune := b.endRune
	// fmt.Println(initialRune, lastRune)

	var headNode *Node
	var tailNode *Node
	text := make([]*Node, 0)
	for i := range b.textLength {
		node := NewNode(initialRune, initialRune, lastRune)
		if i == 0 {
			headNode = &node
		} else {
			headNode.Append(&node)
		}
		text = append(text, &node)
		if b.textLength-i == 1 {
			tailNode = &node
		}
	}

	for {
		select {
		case <-ctx.Done():
			return "", errors.New("stopped")
		default:
			packed := headNode.Pack()
			result := b.goalTest(packed)
			if result {
				return packed, nil
			}
			// fmt.Println(packed)
			tailNode.Increment()

			same := true
			for _, r := range headNode.Pack() {
				if r != lastRune {
					same = false
				}
			}
			if same {
				packed := headNode.Pack()
				result := b.goalTest(packed)
				if result {
					return packed, nil
				}
				return "", errors.New("not found")
			}

		}
	}
}

type Node struct {
	Value      rune
	StartValue rune
	EndValue   rune
	Prev       *Node
	Next       *Node
}

func NewNode(value rune, start rune, end rune) Node {
	return Node{
		Value:      value,
		StartValue: start,
		EndValue:   end,
		Prev:       nil,
		Next:       nil,
	}
}

func (n *Node) ConnectNext(node *Node) {
	n.Next = node
}

func (n *Node) ConnectPrevious(node *Node) {
	n.Prev = node
}

func (n *Node) Append(newNode *Node) {
	var node *Node = n
	if node.Next == nil {
		node.ConnectNext(newNode)
		newNode.ConnectPrevious(node)
		return
	}

	for node.Next != nil {
		node = node.Next
		if node.Next == nil {
			node.ConnectNext(newNode)
			newNode.ConnectPrevious(node)
			return
		}
	}
}

func (n *Node) increment(node *Node) {
	if node.Value+1 > node.EndValue {
		if node.Prev != nil {
			node.increment(node.Prev)
		}
		node.Value = node.StartValue
	} else {
		node.Value += 1
	}
}

func (n *Node) Increment() {
	n.increment(n)
}

func (n *Node) Pack() string {
	var node *Node = n
	text := make([]rune, 0)

	for node != nil {
		text = append(text, node.Value)
		node = node.Next
	}

	return string(text)
}
