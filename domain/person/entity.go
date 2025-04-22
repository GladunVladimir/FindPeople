package person

import "strings"

type Person struct {
	FullName    string
	Gender      string
	Nationality string
	Age         int
}

func New(fullName, gender, nationality string, age int) *Person {
	return &Person{
		FullName:    fullName,
		Gender:      gender,
		Nationality: nationality,
		Age:         age,
	}
}

func (p *Person) FirstName() string {
	parts := strings.Split(p.FullName, " ")
	if len(parts) > 0 {
		return parts[0]
	}
	return p.FullName
}

func (p *Person) IsValid() bool {
	return p.FullName != "" && p.Age >= 0
}
