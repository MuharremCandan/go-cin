package cloud

import "github.com/cloudinary/cloudinary-go/v2"

// NewCloud creates a new Cloudinary client.
func NewCloud(cloudinaryURL string) (*cloudinary.Cloudinary, error) {
	cld, err := cloudinary.NewFromURL(cloudinaryURL)
	if err != nil {
		return nil, err
	}
	return cld, nil
}
