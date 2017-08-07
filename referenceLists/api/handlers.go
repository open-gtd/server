package api

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

var r api.Registerer = api.GetRegisterer()

func Initialize() {
	r.GET("/api", "/referenceLists", GetAll)
	r.GET("/api", "/referenceLists/:name", Get)
	r.POST("/api", "/referenceLists", Create)
	r.PUT("/api", "/referenceLists/:name", Update)
	r.DELETE("/api", "/referenceLists/:name", Delete)
}

func GetAll(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "ReferenceLists!")
}

func Get(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "ReferenceList!")
}

func Create(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "ReferenceList created!")
}

func Update(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "ReferenceList updated!")
}

func Delete(rq api.Request, rs api.Response) error {
	return rs.String(http.StatusOK, "ReferenceList deleted!")
}
