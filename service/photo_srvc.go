package service

import (
	"FinalProjectAssignment/model"
	"FinalProjectAssignment/repo"
	"context"
	"errors"
	"fmt"
)

type PhotoServiceInterface interface {
	PhotoRegister(ctx context.Context, photos *model.Photo) (*model.Photo, error)
	PhotoGet(ctx context.Context, id string) ([]*model.PhotoGet, error)
	PhotoUpdate(ctx context.Context, id string, photos *model.Photo) (*model.Photo, error)
	PhotoDelete(ctx context.Context, id string, photos *model.Photo) (*model.Photo, error)
}

type PhotoService struct {
	PhotoRepo repo.PhotoRepoInterface
}

func NewPhotoService(PhotoRepo repo.PhotoRepoInterface) PhotoServiceInterface {
	return &PhotoService{PhotoRepo: PhotoRepo}
}

func (p *PhotoService) PhotoRegister(ctx context.Context, photos *model.Photo) (*model.Photo, error) {
	title := photos.Title
	caption := photos.Caption
	url := photos.Photo_url
	fmt.Println("ini photo service : ", photos)
	if title == "" {
		return nil, errors.New("Please input photo title")
	}
	if caption == "" {
		return nil, errors.New("Please input photo title")
	}
	if url == "" {
		return nil, errors.New("Please input photo title")
	}

	photoRegister, err := p.PhotoRepo.PhotoRepoRegister(ctx, photos)
	fmt.Println("ini photo service : ", photoRegister)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", photoRegister)
	return photoRegister, nil
}

func (p *PhotoService) PhotoGet(ctx context.Context, id string) ([]*model.PhotoGet, error) {
	photo, err := p.PhotoRepo.PhotoRepoGet(ctx, id)
	if err != nil {
		return nil, err
	}
	fmt.Println("ini service photo", photo)
	return photo, nil
}

func (p *PhotoService) PhotoUpdate(ctx context.Context, id string, photos *model.Photo) (*model.Photo, error) {
	title := photos.Title
	caption := photos.Caption
	url := photos.Photo_url
	fmt.Println("ini photo service : ", photos)
	if title == "" {
		return nil, errors.New("Please input photo title")
	}
	if caption == "" {
		return nil, errors.New("Please input photo title")
	}
	if url == "" {
		return nil, errors.New("Please input photo title")
	}

	photoUpdate, err := p.PhotoRepo.PhotoRepoUpdate(ctx, id, photos)
	fmt.Println("ini photo service : ", photoUpdate)
	if err != nil {
		fmt.Println("Error While Register", err.Error())
		return nil, err
	}
	fmt.Println("service user:", photoUpdate)
	return photoUpdate, nil
}

func (p *PhotoService) PhotoDelete(ctx context.Context, id string, photos *model.Photo) (*model.Photo, error) {
	deleteUser, err := p.PhotoRepo.PhotoRepoDelete(ctx, id, photos)
	if err != nil {
		return nil, err
	}
	return deleteUser, nil
}
