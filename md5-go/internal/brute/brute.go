package brute

type GoalTest func(guess string) bool

type Brute interface {
}

type brute struct {
	goalTest GoalTest
}

func New() Brute {
	return &brute{}
}

func (b brute) SetGoalTest(goalTest GoalTest) Brute {
	b.goalTest = goalTest
	return b
}

// NOTE: Its fair game to have the length, special conditions like only lowercase character set before brute forcing it
// NOTE: https://www.youtube.com/watch?v=7U-RbOKanYs&t=389s
// func (b brute)
