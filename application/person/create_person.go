package personapp

import (
	domain "FindPeople/domain/person"
	"FindPeople/infrastructure/external"
	"FindPeople/infrastructure/persistence"
)

type CreatePersonInput struct {
	FullName string
}

func CreatePerson(in CreatePersonInput) (*domain.Person, error) {
	first := domain.New(in.FullName, "", "", 0).FirstName()

	age, err := external.FetchAge(first)
	if err != nil {
		return nil, err
	}
	gender, err := external.FetchGender(first)
	if err != nil {
		return nil, err
	}
	nat, err := external.FetchNationality(first)
	if err != nil {
		return nil, err
	}

	p := domain.New(in.FullName, gender, nat, age)
	if !p.IsValid() {
		return nil, err
	}
	return p, persistence.PersonRepo.Save(p)
}
