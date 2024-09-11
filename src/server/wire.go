//go:build wireinject
// +build wireinject

package server

import (
	"github.com/google/wire"
	"teacherScheduler/src/internal/handler"
)

func NewController() (*handler.SchedulerHandler, error) {
	wire.Build(handler.NewSchedulerHandler)
	return &handler.SchedulerHandler{}, nil
}
