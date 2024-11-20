package certgen

import (
	"certgen/internal/reqstate"
	"net/http"
)

func httpError(w http.ResponseWriter, req *http.Request, statusCode int, err error) {
	reqState := reqstate.Get(req)
	reqState.Logger.WarnContext(req.Context(), err.Error())
	http.Error(w, http.StatusText(statusCode), statusCode)
}

func httpBadRequestError(w http.ResponseWriter, req *http.Request, err error) {
	httpError(w, req, http.StatusBadRequest, err)
}

func httpInternalServerError(w http.ResponseWriter, req *http.Request, err error) {
	httpError(w, req, http.StatusInternalServerError, err)
}
