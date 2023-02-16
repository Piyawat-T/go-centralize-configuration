package mocks

import (
	"github.com/stretchr/testify/mock"
)

type Database struct {
	mock.Mock
}

func (mc *Database) CloseMySQLDatabase() {
	// Do nothing
}

func (mc *Database) Find(out interface{}, where ...interface{}) error {
	ret := mc.Called(out)
	var r0 error
	if ret.Get(0) != nil {
		r0 = ret.Get(0).(error)
	}
	return r0
}
