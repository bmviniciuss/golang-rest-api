package comment

import (
	"context"
	"errors"
	"fmt"
)

var (
	ErrFetchingComment = errors.New("failed to fetch comment")
	ErrNotImplemented  = errors.New("not implemented")
	ErrCreatingComment = errors.New("failed to create comment")
)

type Comment struct {
	ID     string
	Slug   string
	Body   string
	Author string
}

type Repository interface {
	GetComment(ctx context.Context, id string) (Comment, error)
	CreateComment(ctx context.Context, cmt *Comment) error
	DeleteComment(ctx context.Context, uuid string) error
	UpdateComment(ctx context.Context, cmt Comment) (Comment, error)
}

type Service struct {
	repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{
		repository: repository,
	}
}

func (s *Service) GetComment(ctx context.Context, id string) (Comment, error) {
	fmt.Println("GetComment: ", id)

	cmt, err := s.repository.GetComment(ctx, id)
	if err != nil {
		fmt.Println("Error: ", err)
		return Comment{}, ErrFetchingComment
	}
	return cmt, nil
}

func (s *Service) UpdateComment(ctx context.Context, cmt Comment) error {
	return ErrNotImplemented
}

func (s *Service) DeleteComment(ctx context.Context, id string) error {
	fmt.Println("DeleteComment")
	err := s.repository.DeleteComment(ctx, id)

	if err != nil {
		fmt.Println("Error: ", err)
		return err
	}

	return ErrNotImplemented
}

func (s *Service) CreateComment(ctx context.Context, cmt Comment) (Comment, error) {
	fmt.Println("Creating comment...")
	err := s.repository.CreateComment(ctx, &cmt)

	if err != nil {
		fmt.Println("Error: ", err)
		return Comment{}, ErrCreatingComment
	}
	fmt.Println("Comment created")
	return cmt, nil
}
