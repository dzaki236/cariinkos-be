package services

import (
	"errors"
	"testing"

	"cariinkos/app/dto"
	"cariinkos/app/models"

	"github.com/goravel/framework/contracts/filesystem"
	"github.com/stretchr/testify/require"
)

type kosPhotoRepositoryFake struct {
	photo       models.KosPhoto
	createErr   error
	getErr      error
	updateErr   error
	deleteErr   error
	createCalls int
}

func (r *kosPhotoRepositoryFake) GetAllByKosID(string) ([]models.KosPhoto, error) { return nil, nil }
func (r *kosPhotoRepositoryFake) GetByID(string, string) (models.KosPhoto, error) {
	return r.photo, r.getErr
}
func (r *kosPhotoRepositoryFake) Create(photo *models.KosPhoto) error {
	r.createCalls++
	r.photo = *photo
	return r.createErr
}
func (r *kosPhotoRepositoryFake) Update(photo *models.KosPhoto) error {
	r.photo = *photo
	return r.updateErr
}
func (r *kosPhotoRepositoryFake) Delete(string, string) error { return r.deleteErr }

type photoUploaderFake struct {
	url        string
	publicID   string
	uploadErr  error
	deleteErr  error
	deletedIDs []string
}

func (u *photoUploaderFake) Upload(_ string) (string, string, error) {
	return u.url, u.publicID, u.uploadErr
}
func (u *photoUploaderFake) Delete(publicID string) error {
	u.deletedIDs = append(u.deletedIDs, publicID)
	return u.deleteErr
}

type filesystemFileContract filesystem.File

type photoFileFake struct{ filesystemFileContract }

func (photoFileFake) File() string { return "photo.jpg" }

func TestKosPhotoServiceCreateUploadsAndPersistsCloudinaryMetadata(t *testing.T) {
	repository := &kosPhotoRepositoryFake{}
	uploader := &photoUploaderFake{url: "https://cdn.example/photo.jpg", publicID: "kos/photo"}
	service := NewKosPhotoService(repository, uploader)

	photo, err := service.Create("kos-id", dto.CreateKosPhotoRequest{File: photoFileFake{}})

	require.NoError(t, err)
	require.Equal(t, "https://cdn.example/photo.jpg", photo.ImageURL)
	require.NotNil(t, photo.PublicID)
	require.Equal(t, "kos/photo", *photo.PublicID)
	require.Equal(t, photo, repository.photo)
	require.Equal(t, 1, repository.createCalls)
}

func TestKosPhotoServiceCreateDeletesCloudinaryAssetWhenPersistenceFails(t *testing.T) {
	repository := &kosPhotoRepositoryFake{createErr: errors.New("database unavailable")}
	uploader := &photoUploaderFake{url: "https://cdn.example/photo.jpg", publicID: "kos/photo"}
	service := NewKosPhotoService(repository, uploader)

	_, err := service.Create("kos-id", dto.CreateKosPhotoRequest{File: photoFileFake{}})

	require.Error(t, err)
	require.Equal(t, []string{"kos/photo"}, uploader.deletedIDs)
}

func TestKosPhotoServiceDoesNotDeleteAssetWhenRepositoryRecordIsMissing(t *testing.T) {
	repository := &kosPhotoRepositoryFake{getErr: errors.New("not found")}
	uploader := &photoUploaderFake{}
	service := NewKosPhotoService(repository, uploader)

	err := service.Delete("kos-id", "photo-id")

	require.Error(t, err)
	require.Empty(t, uploader.deletedIDs)
}

func TestKosPhotoServiceDeleteRemovesCloudinaryAssetAfterRepositoryDelete(t *testing.T) {
	publicID := "kos/photo"
	repository := &kosPhotoRepositoryFake{photo: models.KosPhoto{PublicID: &publicID}}
	uploader := &photoUploaderFake{}
	service := NewKosPhotoService(repository, uploader)

	err := service.Delete("kos-id", "photo-id")

	require.NoError(t, err)
	require.Equal(t, []string{publicID}, uploader.deletedIDs)
}

func TestKosPhotoServiceUpdateReplacesCloudinaryAsset(t *testing.T) {
	oldPublicID := "kos/old-photo"
	repository := &kosPhotoRepositoryFake{photo: models.KosPhoto{ID: "photo-id", KosID: "kos-id", PublicID: &oldPublicID}}
	uploader := &photoUploaderFake{url: "https://cdn.example/new-photo.jpg", publicID: "kos/new-photo"}
	service := NewKosPhotoService(repository, uploader)

	photo, err := service.Update("kos-id", "photo-id", dto.UpdateKosPhotoRequest{File: photoFileFake{}})

	require.NoError(t, err)
	require.Equal(t, "https://cdn.example/new-photo.jpg", photo.ImageURL)
	require.Equal(t, []string{oldPublicID}, uploader.deletedIDs)
	require.Equal(t, "kos/new-photo", *repository.photo.PublicID)
}
