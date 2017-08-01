package api

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

func RegisterHandlers(r api.RestRegisterer) {
	r.GET("/tasks", GetAll)
	r.GET("/tasks/:name", Get)
	r.POST("/tasks", Create)
	r.PUT("/tasks/:name", Update)
	r.DELETE("/tasks/:name", Delete)
}

func GetAll(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Tasks!")
}

func Get(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Task!")
}

func Create(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Task created!")
}

func Update(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Task updated!")
}

func Delete(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Task deleted!")
}
