package controller

import (
	"net/http"
)

func HandleErrorOfSelect(w http.ResponseWriter, err error) bool {
	data := false
	switch err {
	case sql.ErrNoRows:
		data = true
	case nil:
		data = false
	default:
		data = true
	}

	return data
}