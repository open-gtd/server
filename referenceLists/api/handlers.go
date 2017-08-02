package api

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

func RegisterHandlers(r api.Registerer) {
	r.GET("/referenceLists", GetAll)
	r.GET("/referenceLists/:name", Get)
	r.POST("/referenceLists", Create)
	r.PUT("/referenceLists/:name", Update)
	r.DELETE("/referenceLists/:name", Delete)
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
