package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/gorilla/mux"
	"github.com/roblesdotdev/go-rest-api/internal/comment"
)

type Response struct {
	Message string
}

type PostCommentRequest struct {
	Slug   string `json:"slug" validate:"required"`
	Author string `json:"author" validate:"required"`
	Body   string `json:"body" validate:"required"`
}

func (pc *PostCommentRequest) toComment() comment.Comment {
	return comment.Comment{
		Slug:   pc.Slug,
		Author: pc.Author,
		Body:   pc.Body,
	}
}

func (h *Handler) PostComment(w http.ResponseWriter, r *http.Request) {
	var cmt PostCommentRequest
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	validate := validator.New()
	if err := validate.Struct(cmt); err != nil {
		http.Error(w, "not a valid comment", http.StatusBadRequest)
		return
	}

	postedComment, err := h.Service.CreateComment(r.Context(), cmt.toComment())
	if err != nil {
		log.Println(err)
		return
	}

	if err := json.NewEncoder(w).Encode(postedComment); err != nil {
		panic(err)
	}
}

func (h *Handler) GetComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	cmt, err := h.Service.GetComment(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err = json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) UpdateComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var cmt comment.Comment
	if err := json.NewDecoder(r.Body).Decode(&cmt); err != nil {
		return
	}

	cmt, err := h.Service.UpdateComment(r.Context(), id, cmt)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(cmt); err != nil {
		panic(err)
	}
}

func (h *Handler) DeleteComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := h.Service.DeleteComment(r.Context(), id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Deleted successfully"}); err != nil {
		panic(err)
	}
}
