package photobase

import (
	"github.com/JulzDiverse/feedelphia/api"
)

type InMemoryPhotobase struct {
	photos []api.Photo
}

func NewInMemoryPhotobase() InMemoryPhotobase {
	return InMemoryPhotobase{
		photos: []api.Photo{},
	}
}

func (p *InMemoryPhotobase) List() ([]api.Photo, error) {
	return p.photos, nil
}

func (p *InMemoryPhotobase) Add(photo api.Photo) error {
	p.photos = append(p.photos, photo)
	return nil
}
