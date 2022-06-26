package service

import (
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/repo"
	"context"
	"errors"
	"fmt"
)

type CommentServiceInterface interface {
	CommentRegister(ctx context.Context, comments *model.Comment) (*model.Comment, error)
	CommentGet(ctx context.Context, id string) ([]*model.CommentGet, error)
	CommentUpdate(ctx context.Context, id string, comments *model.Comment) (*model.CommentShow, error)
	CommentDelete(ctx context.Context, id string, comments *model.Comment) (*model.Comment, error)
}

type CommentService struct {
	CommentRepo repo.CommentRepoInterface
}

func NewCommentService(CommentRepo repo.CommentRepoInterface) CommentServiceInterface {
	return &CommentService{CommentRepo: CommentRepo}
}

func (c *CommentService) CommentRegister(ctx context.Context, comments *model.Comment) (*model.Comment, error) {
	message := comments.Message
	fmt.Println("ini Comment service : ", comments)
	if message == "" {
		return nil, errors.New("Please input Comment Message")
	}
	CommentRegister, err := c.CommentRepo.CommentRepoRegister(ctx, comments)
	fmt.Println("ini Comment service : ", CommentRegister)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", CommentRegister)
	return CommentRegister, nil
}

func (c *CommentService) CommentGet(ctx context.Context, id string) ([]*model.CommentGet, error) {
	comment, err := c.CommentRepo.CommentRepoGet(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("ini service comment", comment)
	return comment, nil
}

func (c *CommentService) CommentUpdate(ctx context.Context, id string, comments *model.Comment) (*model.CommentShow, error) {
	message := comments.Message
	fmt.Println("ini Comment service : ", comments)
	if message == "" {
		return nil, errors.New("Please input Comment Message")
	}
	commentUpdate, err := c.CommentRepo.CommentRepoUpdate(ctx, id, comments)
	fmt.Println("ini photo service : ", commentUpdate)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", commentUpdate)
	return commentUpdate, nil
}

func (c *CommentService) CommentDelete(ctx context.Context, id string, comments *model.Comment) (*model.Comment, error) {
	deleteUser, err := c.CommentRepo.CommentRepoDelete(ctx, id, comments)
	if err != nil {
		return nil, err
	}
	return deleteUser, nil
}
