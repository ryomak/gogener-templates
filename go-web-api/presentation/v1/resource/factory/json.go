package factory

import (
	"encoding/json"
	"net/http"

	"[[.ModName]]/pkg/merr"
	"[[.ModName]]/presentation/v1/resource"
)

func NoContent(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNoContent)
}

func JSON(w http.ResponseWriter, res interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

func ErrorJSON(w http.ResponseWriter, err error, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	var errRes *resource.Error
	switch {
	case merr.IsMyError(err):
		errRes = &resource.Error{Message: err.Error(), Code: err.(*merr.Error).Code}
	case merr.IsErrNoRows(err):
		errRes = &resource.Error{Message: err.Error(), Code: merr.CodeDBNotFound}
	case merr.IsErrValidator(err):
		errRes = &resource.Error{Message: err.Error(), Code: merr.CodeValidator}
	default:
		errRes = &resource.Error{Message: err.Error(), Code: merr.CodeUnexpected}
	}
	json.NewEncoder(w).Encode(errRes)
}
