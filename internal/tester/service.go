// Package tester contains the business logic of the application.
//
// Tester is the telegram bot for testing mathematical.
// This bot generate the test if you write /start command.
//
// Then it send you the question, and you must answer it.
// If you answer right, it sends you the next question.
// If you answer wrong, it writes you that you answer wrong,
// send you the right answer and write he next question.
// When tasks are over, it writes you questions, that you answer wrong.
// In the end it writes you the result of the test.
//
// If you want to stop the test, you can write /stop command.
package tester

import (
	"context"

	test "math_tester/internal/tester/test_"
)

type TestRepository interface {
	FindTestByUserID(userID string) (*test.Test, error)
	SaveTest(userID string, test test.Test) error
	DeleteTest(userID string) error
}

type Service struct {
	tests TestRepository
}

func NewService(tests TestRepository) *Service {
	panic("implement me")
}

// StartTest starts the test for the current user, with tasks for the given number.
func (m Service) StartTest(
	ctx context.Context, userID string, number int,
) error {
	panic("implement me")
}

// RequestQuestion returns the question for the user.
// If the test is over, it returns an empty string and error ErrTestIsOver.
func (m Service) RequestQuestion(
	ctx context.Context, userID string,
) (string, error) {
	panic("implement me")
}

func (m Service) HandleAnswer(
	ctx context.Context, userID string, answer string,
) error {
	panic("implement me")
}

func (m Service) StopTest(ctx context.Context, userID string) {
	panic("implement me")
}
