package repository

import (
	"github.com/mahesadhanaa/crud-golang/src/modules/profile/model"
)

//ProfileRepository
type ProfileRepository interface {
	Save(*model.Profile) error
	Update(string, *model.Profile) error
	Delete(string) error
	FindByID(string) (*model.Profile, error)
	FindAll() (model.Profiles, error)
}
