package services

import (
	"context"
	"errors"

	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"cariinkos/app/facades"

	"github.com/google/uuid"
)

type KosPhotoService struct {
	repo     repositoryContracts.KosPhotoRepositoryContract
	kosRepo  repositoryContracts.KosRepositoryContract
	uploader helpers.PhotoUploader
}

func NewKosPhotoService(repo repositoryContracts.KosPhotoRepositoryContract, kosRepo repositoryContracts.KosRepositoryContract, uploader helpers.PhotoUploader) *KosPhotoService {
	return &KosPhotoService{repo: repo, kosRepo: kosRepo, uploader: uploader}
}

func (s *KosPhotoService) GetAllByKosID(kosID string) ([]models.KosPhoto, error) {
	return s.repo.GetAllByKosID(kosID)
}

func (s *KosPhotoService) Create(ctx context.Context, userID, kosID string, request dto.CreateKosPhotoRequest) (models.KosPhoto, error) {
	if err := CheckKosOwner(userID, kosID, s.kosRepo); err != nil {
		return models.KosPhoto{}, err
	}
	if request.File == nil {
		return models.KosPhoto{}, helpers.ErrPhotoRequired
	}
	tempPath, err := request.File.Store("temp/cloudinary")
	if err != nil {
		return models.KosPhoto{}, err
	}
	defer func() { _ = facades.Storage().Delete(tempPath) }()
	imageURL, publicID, err := s.uploader.Upload(ctx, facades.Storage().Path(tempPath))
	if err != nil {
		return models.KosPhoto{}, err
	}
	if imageURL == "" || publicID == "" {
		return models.KosPhoto{}, errors.New("photo upload returned an empty URL or public ID")
	}
	photo := models.KosPhoto{ID: uuid.NewString(), KosID: kosID, ImageURL: imageURL, PublicID: &publicID}
	if err := s.repo.Create(&photo); err != nil {
		_ = s.uploader.Delete(ctx, publicID)
		return models.KosPhoto{}, err
	}
	return photo, nil
}

func (s *KosPhotoService) Update(ctx context.Context, userID, kosID, photoID string, request dto.UpdateKosPhotoRequest) (models.KosPhoto, error) {
	if err := CheckKosOwner(userID, kosID, s.kosRepo); err != nil {
		return models.KosPhoto{}, err
	}
	photo, err := s.repo.GetByID(kosID, photoID)
	if err != nil {
		return models.KosPhoto{}, err
	}
	if request.File == nil {
		return models.KosPhoto{}, helpers.ErrPhotoRequired
	}
	tempPath, err := request.File.Store("temp/cloudinary")
	if err != nil {
		return models.KosPhoto{}, err
	}
	defer func() { _ = facades.Storage().Delete(tempPath) }()
	imageURL, publicID, err := s.uploader.Upload(ctx, facades.Storage().Path(tempPath))
	if err != nil {
		return models.KosPhoto{}, err
	}
	if imageURL == "" || publicID == "" {
		return models.KosPhoto{}, errors.New("photo upload returned an empty URL or public ID")
	}
	oldPublicID := photo.PublicID
	photo.ImageURL = imageURL
	photo.PublicID = &publicID
	if err := s.repo.Update(&photo); err != nil {
		_ = s.uploader.Delete(ctx, publicID)
		return models.KosPhoto{}, err
	}
	if oldPublicID != nil {
		_ = s.uploader.Delete(ctx, *oldPublicID)
	}
	return photo, nil
}

func (s *KosPhotoService) Delete(ctx context.Context, userID, kosID, photoID string) error {
	if err := CheckKosOwner(userID, kosID, s.kosRepo); err != nil {
		return err
	}
	photo, err := s.repo.GetByID(kosID, photoID)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(kosID, photoID); err != nil {
		return err
	}
	if photo.PublicID != nil {
		return s.uploader.Delete(ctx, *photo.PublicID)
	}
	return nil
}

var _ serviceContracts.KosPhotoService = (*KosPhotoService)(nil)
