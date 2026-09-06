package student

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/students-api/internal/types"
	"github.com/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		slog.Info("creating student")

		var student types.Student

		// Decode request body
		err := json.NewDecoder(r.Body).Decode(&student)

		if errors.Is(err, io.EOF) {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GenralError(fmt.Errorf("request body is empty")),
			)
			return
		}

		if err != nil {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GenralError(err),
			)
			return
		}

		// Request validation
		err = validator.New().Struct(student)

		if err != nil {
			response.WriteJson(
				w,
				http.StatusBadRequest,
				response.GenralError(err),
			)
			return
		}

		response.WriteJson(
			w,
			http.StatusCreated,
			map[string]string{"success": "OK"},
		)
	}
}
