package api

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

var r api.Registerer = api.GetRegisterer()

func Initialize() {
	r.GET("/api", "/tasks", GetAll)
	r.GET("/api", "/tasks/:name", Get)
	r.POST("/api", "/tasks", Create)
	r.PUT("/api", "/tasks/:name", Update)
	r.DELETE("/api", "/tasks/:name", Delete)
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
