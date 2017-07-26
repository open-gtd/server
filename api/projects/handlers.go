package projects

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

func RegisterHandlers(r api.Registerer) {
	r.GET("/projects", GetAll)
	r.GET("/projects/:name", Get)
	r.POST("/projects", Create)
	r.PUT("/projects/:name", Update)
	r.DELETE("/projects/:name", Delete)
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
