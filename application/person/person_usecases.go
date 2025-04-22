package personapp

import (
	"FindPeople/infrastructure/persistence"
	"FindPeople/models"
)

func GetAllPeople() ([]models.Person, error) {
	return persistence.PersonRepo.FindAll()
}

func FindByLastName(last string) ([]models.Person, error) {
	return persistence.PersonRepo.FindByLastName(last)
}

func UpdatePerson(id uint, fullName string) (*models.Person, error) {
	return persistence.PersonRepo.UpdateName(id, fullName)
}

func LinkFriends(userID, friendID uint) error {
	return persistence.PersonRepo.LinkFriends(userID, friendID)
}

func GetFriends(userID uint) ([]models.Person, error) {
	return persistence.PersonRepo.GetFriends(userID)
}
