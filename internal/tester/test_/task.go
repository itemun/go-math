package test

// Task is the task for the math testing.
type Task struct {
	question string
	answer   int
}

func (t Task) IsRightAnswer(answer int) bool {
	panic("implement me")
}

func (t Task) String() string {
	return t.question
}
