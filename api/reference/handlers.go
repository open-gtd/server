package reference

import (
	"net/http"

	"github.com/open-gtd/server/api"
)

func RegisterHandlers(r api.Registerer) {
	r.GET("/reference", GetAll)
	r.GET("/reference/:name", Get)
	r.POST("/reference", Create)
	r.PUT("/reference/:name", Update)
	r.DELETE("/reference/:name", Delete)
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
