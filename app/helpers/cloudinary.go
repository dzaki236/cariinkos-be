package helpers

import (
	"context"
	"errors"
	"os"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type PhotoUploader interface {
	Upload(filePath string) (string, string, error)
	Delete(publicID string) error
}

type CloudinaryPhotoUploader struct {
	client *cloudinary.Cloudinary
	err    error
}

func NewCloudinaryPhotoUploader() *CloudinaryPhotoUploader {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")
	if cloudinaryURL == "" {
		return &CloudinaryPhotoUploader{err: errors.New("CLOUDINARY_URL is required")}
	}
	client, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return &CloudinaryPhotoUploader{err: err}
	}
	return &CloudinaryPhotoUploader{client: client}
}

func (u *CloudinaryPhotoUploader) Upload(filePath string) (string, string, error) {
	if u.err != nil {
		return "", "", u.err
	}
	result, err := u.client.Upload.Upload(context.Background(), filePath, uploader.UploadParams{ResourceType: "image"})
	if err != nil {
		return "", "", err
	}
	return result.SecureURL, result.PublicID, nil
}

func (u *CloudinaryPhotoUploader) Delete(publicID string) error {
	if u.err != nil {
		return u.err
	}
	_, err := u.client.Upload.Destroy(context.Background(), uploader.DestroyParams{PublicID: publicID, ResourceType: "image"})
	return err
}
