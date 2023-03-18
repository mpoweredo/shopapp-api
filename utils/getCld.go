package utils

import (
	"errors"
	"github.com/cloudinary/cloudinary-go/v2"
	"os"
)

func GetCld() (*cloudinary.Cloudinary, error) {
	cloudinaryURL := os.Getenv("CLOUDINARY_URL")

	cld, err := cloudinary.NewFromURL(cloudinaryURL)

	if err != nil {
		return nil, errors.New("couldn't connect with cloudinary")
	}

	return cld, nil
}
