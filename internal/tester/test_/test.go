// Package test store domain model of application.
// Folder has name "test_" because it is not folder for test :)
package test

// Test is the list from some questions for testing mathematical.
type Test struct {
	tasks       []Task
	currentTask int
}

// NewTest returns a new Test.
func NewTest(tasks ...Task) *Test {
	panic("implement me")
}

func (t *Test) IsRightAnswer(answer int) bool {
	panic("implement me")
}

// Next returns true if the test is not over.
//
//	for test.Next() {
//		question := test.RequestQuestion()
//	}
func (t *Test) Next() bool {
	panic("implement me")
}

func (t *Test) RequestQuestion() string {
	panic("implement me")
}
