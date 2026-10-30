package config

import "cariinkos/app/facades"

func init() {
	config := facades.Config()
	config.Add("cloudinary", map[string]any{
		"url": config.Env("CLOUDINARY_URL", ""),
	})
}
