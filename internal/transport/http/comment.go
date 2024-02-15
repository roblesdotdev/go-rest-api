package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/roblesdotdev/go-rest-api/internal/comment"
)

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.CreateComment(r.Context(), cmt)
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}
