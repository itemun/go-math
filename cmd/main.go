package main

import (
	"go.uber.org/zap"

	"math_tester/internal/tester"
	"math_tester/internal/tester/test_/repository/memory"
	"math_tester/pkg/process"
)

func main() {
	telegramBot := tester.NewTelegramBot(
		tester.NewService(
			memory.NewTestRepository(),
		),
	)

	go func() {
		if err := telegramBot.Run(); err != nil {
			process.Terminate()
		}
	}()

	process.WaitForTermination()
	if err := telegramBot.Close(); err != nil {
		zap.L().Error("failed to close telegram bot", zap.Error(err))
	}
}
