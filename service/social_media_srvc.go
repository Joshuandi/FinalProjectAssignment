package service

import (
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/repo"
	"context"
	"errors"
	"fmt"
)

type SocialMediaServiceInterface interface {
	SocialMediaRegister(ctx context.Context, sm *model.SocialMedia) (*model.SocialMedia, error)
	SocialMediaGet(ctx context.Context) ([]*model.SocialMediaShow, error)
	SocialMediaUpdate(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error)
	SocialMediaDelete(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error)
}

type SocialMediaService struct {
	SocialMediaRepo repo.SocialMediaRepoInterface
}

func NewSocialMediaService(SocialMediaRepo repo.SocialMediaRepoInterface) SocialMediaServiceInterface {
	return &SocialMediaService{SocialMediaRepo: SocialMediaRepo}
}

func (c *SocialMediaService) SocialMediaRegister(ctx context.Context, sm *model.SocialMedia) (*model.SocialMedia, error) {
	name := sm.Name
	socmed_url := sm.SocialMedia_url
	fmt.Println("ini SocialMedia service : ", sm)
	if name == "" {
		return nil, errors.New("Please input SocialMedia Name")
	}
	if socmed_url == "" {
		return nil, errors.New("Please input SocialMedia URL")
	}
	SocialMediaRegister, err := c.SocialMediaRepo.SocialMediaRepoRegister(ctx, sm)
	fmt.Println("ini SocialMedia service : ", SocialMediaRegister)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", SocialMediaRegister)
	return SocialMediaRegister, nil
}

func (c *SocialMediaService) SocialMediaGet(ctx context.Context) ([]*model.SocialMediaShow, error) {
	SocialMedia, err := c.SocialMediaRepo.SocialMediaRepoGet(ctx)
	if err != nil {
		return nil, err
	}
	fmt.Println("ini service SocialMedia", SocialMedia)
	return SocialMedia, nil
}

func (c *SocialMediaService) SocialMediaUpdate(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error) {
	name := sm.Name
	socmed_url := sm.SocialMedia_url
	fmt.Println("ini SocialMedia service : ", sm)
	if name == "" {
		return nil, errors.New("Please input SocialMedia Name")
	}
	if socmed_url == "" {
		return nil, errors.New("Please input SocialMedia URL")
	}
	SocialMediaUpdate, err := c.SocialMediaRepo.SocialMediaRepoUpdate(ctx, id, sm)
	fmt.Println("ini photo service : ", SocialMediaUpdate)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", SocialMediaUpdate)
	return SocialMediaUpdate, nil
}

func (c *SocialMediaService) SocialMediaDelete(ctx context.Context, id string, sm *model.SocialMedia) (*model.SocialMedia, error) {
	deleteUser, err := c.SocialMediaRepo.SocialMediaRepoDelete(ctx, id, sm)
	if err != nil {
		return nil, err
	}
	return deleteUser, nil
}
//