package dto

import "github.com/goravel/framework/contracts/filesystem"

type CreateKosPhotoRequest struct {
	File filesystem.File `json:"-"`
}

type UpdateKosPhotoRequest struct {
	File filesystem.File `json:"-"`
}
