package services

import (
	"cariinkos/app/dto"
	"cariinkos/app/helpers"
	"cariinkos/app/models"
	repositoryContracts "cariinkos/app/repositories/contracts"
	serviceContracts "cariinkos/app/services/contracts"

	"github.com/google/uuid"
)

type KosPhotoService struct {
	repo     repositoryContracts.KosPhotoRepositoryContract
	uploader helpers.PhotoUploader
}

func NewKosPhotoService(repo repositoryContracts.KosPhotoRepositoryContract, uploader helpers.PhotoUploader) *KosPhotoService {
	return &KosPhotoService{repo: repo, uploader: uploader}
}

func (s *KosPhotoService) GetAllByKosID(kosID string) ([]models.KosPhoto, error) {
	return s.repo.GetAllByKosID(kosID)
}

func (s *KosPhotoService) Create(kosID string, request dto.CreateKosPhotoRequest) (models.KosPhoto, error) {
	if request.File == nil {
		return models.KosPhoto{}, helpers.ErrPhotoRequired
	}
	imageURL, publicID, err := s.uploader.Upload(request.File.File())
	if err != nil {
		return models.KosPhoto{}, err
	}
	photo := models.KosPhoto{ID: uuid.NewString(), KosID: kosID, ImageURL: imageURL, PublicID: &publicID}
	if err := s.repo.Create(&photo); err != nil {
		_ = s.uploader.Delete(publicID)
		return models.KosPhoto{}, err
	}
	return photo, nil
}

func (s *KosPhotoService) Update(kosID, photoID string, request dto.UpdateKosPhotoRequest) (models.KosPhoto, error) {
	photo, err := s.repo.GetByID(kosID, photoID)
	if err != nil {
		return models.KosPhoto{}, err
	}
	if request.File == nil {
		return models.KosPhoto{}, helpers.ErrPhotoRequired
	}
	imageURL, publicID, err := s.uploader.Upload(request.File.File())
	if err != nil {
		return models.KosPhoto{}, err
	}
	oldPublicID := photo.PublicID
	photo.ImageURL = imageURL
	photo.PublicID = &publicID
	if err := s.repo.Update(&photo); err != nil {
		_ = s.uploader.Delete(publicID)
		return models.KosPhoto{}, err
	}
	if oldPublicID != nil {
		_ = s.uploader.Delete(*oldPublicID)
	}
	return photo, nil
}

func (s *KosPhotoService) Delete(kosID, photoID string) error {
	photo, err := s.repo.GetByID(kosID, photoID)
	if err != nil {
		return err
	}
	if err := s.repo.Delete(kosID, photoID); err != nil {
		return err
	}
	if photo.PublicID != nil {
		return s.uploader.Delete(*photo.PublicID)
	}
	return nil
}

var _ serviceContracts.KosPhotoService = (*KosPhotoService)(nil)
