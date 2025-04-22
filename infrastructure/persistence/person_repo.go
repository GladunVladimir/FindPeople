package persistence

import (
	"FindPeople/database"
	domain "FindPeople/domain/person"
	"FindPeople/models"

	"gorm.io/gorm"
)

type Repository struct{}

var PersonRepo = Repository{}

// --- сохранение ---
func (Repository) Save(p *domain.Person) error {
	return database.DB.Create(toModel(p)).Error
}

// --- выборки ---
func (Repository) FindAll() ([]models.Person, error) {
	var out []models.Person
	return out, database.DB.Find(&out).Error
}

func (Repository) FindByLastName(last string) ([]models.Person, error) {
	var out []models.Person
	return out, database.DB.
		Where("full_name ILIKE ?", "%"+last+"%").
		Find(&out).Error
}

// --- обновление ---
func (Repository) UpdateName(id uint, fullName string) (*models.Person, error) {
	var m models.Person
	if err := database.DB.First(&m, id).Error; err != nil {
		return nil, err
	}
	m.FullName = fullName
	return &m, database.DB.Save(&m).Error
}

// --- дружба ---
func (Repository) LinkFriends(userID, friendID uint) error {
	var u, f models.Person
	if err := database.DB.First(&u, userID).Error; err != nil {
		return err
	}
	if err := database.DB.First(&f, friendID).Error; err != nil {
		return err
	}
	return database.DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&u).Association("Friends").Append(&f); err != nil {
			return err
		}
		return tx.Model(&f).Association("Friends").Append(&u)
	})
}

func (Repository) GetFriends(userID uint) ([]models.Person, error) {
	var p models.Person
	if err := database.DB.Preload("Friends").First(&p, userID).Error; err != nil {
		return nil, err
	}

	out := make([]models.Person, 0, len(p.Friends))
	for _, f := range p.Friends {
		out = append(out, *f)
	}
	return out, nil
}

// --- helper ---
func toModel(p *domain.Person) *models.Person {
	return &models.Person{
		FullName:    p.FullName,
		Gender:      p.Gender,
		Nationality: p.Nationality,
		Age:         p.Age,
	}
}
