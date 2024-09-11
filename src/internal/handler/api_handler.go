package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type SchedulerHandler struct {
}

func NewSchedulerHandler() *SchedulerHandler {
	return &SchedulerHandler{}
}

func (s SchedulerHandler) GetAllSubjects(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}

func (s SchedulerHandler) GetAllTeachers(ctx echo.Context) error {
	return echo.NewHTTPError(http.StatusNotImplemented, "Not Implemented")
}
