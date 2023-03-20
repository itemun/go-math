package tester

import (
	"context"
)

type MathTester interface {
	StartTest(ctx context.Context, userID string, number int) error
	RequestQuestion(ctx context.Context, userID string) (string, error)
	HandleAnswer(ctx context.Context, userID string, answer string) error
	StopTest(ctx context.Context, userID string)
}

// Bot is the implementation of the telegram bot for the math testing service.
type Bot struct {
	tester MathTester
}

func NewTelegramBot(tester MathTester) *Bot {
	panic("implement me")
}

// Run starts the bot. It blocks until the bot is stopped.
// For stopping the bot, call the Close method.
// Run returns an error if the bot is stopped by an error.
func (b *Bot) Run() error {
	// There is we in the infinite loop receiving messages from the telegram application
	// and process it, using the handleMessage method.
	panic("not impl")
}

func (b *Bot) Close() error {
	return nil
}

// handleMessage processes the message from the telegram application.
func (b *Bot) handleMessage( /* message tg.Message */ ) {
	switch {
	// If the message is a command, we run start or stop method.
	// If the message is a text, we run handleAnswer method.
	// If answer right we request next question, and send it.
	}
}
