package comment

import (
	"context"
	"errors"
	"log"
)

var (
	ErrorFetchingComment = errors.New("failed to fetch comment by id")
	ErrorNotImplemented  = errors.New("not implemented yet")
)

// Comment - a representation of the comment
// structure for our service
type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

// Store - this interface defines all of the methods
// that our service needs in order to operate
type Store interface {
	GetComment(context.Context, string) (Comment, error)
	CreateComment(context.Context, Comment) (Comment, error)
}

// Service - is the struct on which all our
// logic will be built on top of
type Service struct {
	Store Store
}

// NewService - returns a * to a new service
func NewService(store Store) *Service {
	return &Service{
		Store: store,
	}
}

// GetComment - get comment by id
func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	log.Printf("Retrieving a comment with the id %s\n", id)
	cmt, err := s.Store.GetComment(ctx, id)
	if err != nil {
		log.Println(err)
		return Comment{}, ErrorFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrorNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	return ErrorNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	newCmt, err := s.Store.CreateComment(ctx, cmt)
	if err != nil {
		return Comment{}, err
	}
	return newCmt, nil
}
