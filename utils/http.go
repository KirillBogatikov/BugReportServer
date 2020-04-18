package utils

import "net/http"

func BadRequest(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusBadRequest)
	writer.Write([]byte("Bad request"))
}

func Forbidden(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusForbidden)
	writer.Write([]byte("Access denied"))
}

func Ok(writer http.ResponseWriter) {
	writer.WriteHeader(http.StatusOK)
	writer.Write([]byte("Report process started. Thanks for your activity"))
}
