// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package server

import (
	"teacherScheduler/src/internal/handler"
)

// Injectors from wire.go:

func NewController() (*handler.SchedulerHandler, error) {
	schedulerHandler := handler.NewSchedulerHandler()
	return schedulerHandler, nil
}
