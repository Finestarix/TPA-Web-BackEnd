package middleware

import (
	"github.com/gorilla/mux"
	"net/http"
)

var APIKey string

func SetCORS(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {

		vars := mux.Vars(request)
		APIKey = vars["key"]

		writer.Header().Set("Access-Control-Allow-Origin", "*")
		writer.Header().Set("Access-Control-Allow-Methods", "POST")
		writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, Accept-Encoding")

		handler.ServeHTTP(writer, request)

	})
}
