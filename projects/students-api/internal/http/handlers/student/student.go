package student

import (
	"encoding/json"
	"errors"
	"io"
	"log/slog"
	"net/http"

	"github.com/students-api/internal/types"
	"github.com/students-api/internal/utils/response"
)

func New() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var student types.Student
		json.NewDecoder(r.Body).Decode(&student)
		err := json.NewDecoder(r.Body).Decode(&student)
		if errors.Is(err, io.EOF) {
			response.WriteJson(w, http.StatusBadRequest, err.Error())
			return
		}
		slog.Info("creating student")
		response.WriteJson(w, http.StatusCreated, map[string]string{"success": "OK"})
	}
}
