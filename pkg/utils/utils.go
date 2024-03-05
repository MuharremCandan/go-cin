package utils

import (
	"context"
	"mime/multipart"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

func UploadToCloudinary(file multipart.File, filePath string, cld *cloudinary.Cloudinary) (string, error) {
	uploadParams := uploader.UploadParams{
		PublicID: filePath,
	}
	result, err := cld.Upload.Upload(context.Background(), file, uploadParams)
	if err != nil {
		return "", err
	}
	imageUrl := result.SecureURL
	return imageUrl, nil
}
