package repository

import (
	"github.com/mahesadhanaa/crud-golang/src/modules/profile/model"
)

// profile repository
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(*model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (model.Profile, error)
}
