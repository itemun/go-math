// Package memory store repository for tests.
package memory

import (
	test "math_tester/internal/tester/test_"
)

type TestRepository struct {
	// this map contains list of tests for each user
	tests map[string]test.Test
}

func NewTestRepository() *TestRepository {
	panic("implement me")
}

func (t *TestRepository) FindTestByUserID(userID string) (*test.Test, error) {
	panic("implement me")
}

func (t *TestRepository) SaveTest(userID string, test test.Test) error {
	panic("implement me")
}

func (t *TestRepository) DeleteTest(userID string) error {
	panic("implement me")
}
