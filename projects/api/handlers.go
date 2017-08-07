package api

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

var r api.Registerer = api.GetRegisterer()

func Initialize() {
	r.GET("/api", "/projects", GetAll)
	r.GET("/api", "/projects/:name", Get)
	r.POST("/api", "/projects", Create)
	r.PUT("/api", "/projects/:name", Update)
	r.DELETE("/api", "/projects/:name", Delete)
}

func GetAll(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Projects!")
}

func Get(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Project!")
}

func Create(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Project created!")
}

func Update(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Project updated!")
}

func Delete(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "Project deleted!")
}
