package handlers_adpters

import (
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"

	handlers "github.com/alaurentinoofficial/gocore/handlers"
	validators "github.com/alaurentinoofficial/gocore/validations"
	"github.com/gorilla/mux"
)

func UrlVars(request *http.Request) map[string]string {
	return mux.Vars(request)
}

func HttpHandler[T any, U any](next handlers.HandlerFunc[T, U], parseBody bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			request *T
			err     error
		)

		if parseBody {
			request, err = ParseBody[T](r.Body)
			if err != nil {
				slog.Error("Error parsing the body")
				Error(w, err)
				return
			}
		}

		response, err := next(context.Background(), *request)
		if err != nil {
			slog.Error("Error from application: " + err.Error())
			Error(w, err)
			return
		}

		Ok(w, response)
	}
}

func ParseBody[T any](stream io.ReadCloser) (*T, error) {
	var payload T
	err := json.NewDecoder(stream).Decode(&payload)

	return &payload, err
}

func Ok(w http.ResponseWriter, obj any) {
	w.Header().Set("Content-Type", "text/json")
	_ = json.NewEncoder(w).Encode(obj)
}

type ErrorResponse struct {
	Errors any `json:"errors"`
}

func Error(w http.ResponseWriter, err error) {
	w.Header().Set("Content-Type", "text/json")

	switch err {
	case validators.ErrNotFound:
		http.Error(w, "Not found", http.StatusNotFound)
		return
	case validators.ErrForbidden:
		http.Error(w, "Forbidden", http.StatusForbidden)
		return
	case validators.ErrMethodNotAllowed:
		http.Error(w, "Method not allowed", http.StatusForbidden)
		return
	case validators.ErrBadRequest:
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	slog.Error(err.Error())
	w.WriteHeader(http.StatusInternalServerError)
	_ = json.NewEncoder(w).Encode(ErrorResponse{Errors: err})
}
