package helpers

import (
	"context"
	"errors"
	"fmt"
	"os"

	"cariinkos/app/facades"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type PhotoUploader interface {
	Upload(ctx context.Context, filePath string) (string, string, error)
	Delete(ctx context.Context, publicID string) error
}

type CloudinaryPhotoUploader struct {
	client *cloudinary.Cloudinary
	err    error
}

func NewCloudinaryPhotoUploader() *CloudinaryPhotoUploader {
	cloudinaryURL := facades.Config().GetString("cloudinary.url")
	if cloudinaryURL == "" {
		return &CloudinaryPhotoUploader{err: errors.New("CLOUDINARY_URL is required")}
	}
	client, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return &CloudinaryPhotoUploader{err: err}
	}
	return &CloudinaryPhotoUploader{client: client}
}

func (u *CloudinaryPhotoUploader) Upload(ctx context.Context, filePath string) (string, string, error) {
	if u.err != nil {
		return "", "", u.err
	}
	file, err := os.Open(filePath)
	if err != nil {
		return "", "", fmt.Errorf("open file for cloudinary upload: %w", err)
	}
	defer file.Close()
	result, err := u.client.Upload.Upload(ctx, file, uploader.UploadParams{
		ResourceType: "image",
		Folder:       "cariinkos/photos",
	})
	if err != nil {
		return "", "", err
	}
	if result == nil {
		return "", "", errors.New("cloudinary upload returned an empty response")
	}
	if result.Error.Message != "" {
		return "", "", fmt.Errorf("cloudinary rejected upload: %s", result.Error.Message)
	}
	imageURL := result.SecureURL
	if imageURL == "" {
		imageURL = result.URL
	}
	if imageURL == "" || result.PublicID == "" {
		return "", "", errors.New("cloudinary upload response is missing secure_url/url or public_id")
	}
	return imageURL, result.PublicID, nil
}

func (u *CloudinaryPhotoUploader) Delete(ctx context.Context, publicID string) error {
	if u.err != nil {
		return u.err
	}
	_, err := u.client.Upload.Destroy(ctx, uploader.DestroyParams{PublicID: publicID, ResourceType: "image"})
	return err
}
