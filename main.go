package main

import (
	"fmt"

	"github.com/mahesadhanaa/crud-golang/config"
	"github.com/mahesadhanaa/crud-golang/src/modules/profile/model"
	"github.com/mahesadhanaa/crud-golang/src/modules/profile/repository"
)

func main() {

	db, err := config.GetPostgresDB()

	if err != nil {
		fmt.Println(err)
	}

	mhs := model.NewProfile()
	mhs.ID = "M1"
	mhs.FirstName = "Mahesa"
	mhs.LastName = "Dhana"
	mhs.Email = "mhs@gmail.com"
	mhs.Password = "123456"

	profileRepositoryPostgres := repository.NewProfileRepositoryPostgres(db)

	profiles, err := getProfiles(profileRepositoryPostgres)

	if err != nil {
		fmt.Println(err)
	}

	for _, v := range profiles {
		fmt.Println(v)
	}
}

func saveProfile(p *model.Profile, repo repository.ProfileRepository) error {

	err := repo.Save(p)

	if err != nil {
		return err
	}

	return nil
}

func updateProfile(p *model.Profile, repo repository.ProfileRepository) error {

	err := repo.Update(p.ID, p)

	if err != nil {
		return err
	}
	return nil
}

func deleteProfile(id string, repo repository.ProfileRepository) error {

	err := repo.Delete(id)

	if err != nil {
		return err
	}

	return nil
}

func getProfile(id string, repo repository.ProfileRepository) (*model.Profile, error) {
	profile, err := repo.FindByID(id)

	if err != nil {
		return nil, err
	}

	return profile, nil
}

func getProfiles(repo repository.ProfileRepository) (model.Profiles, error) {
	profiles, err := repo.FindAll()

	if err != nil {
		return nil, err
	}

	return profiles, nil
}
