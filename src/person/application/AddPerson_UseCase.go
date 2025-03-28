package application

import "api/src/person/domain"

type AddPersonService struct {
    repo domain.PersonRepository
}

func NewAddPersonService(repo domain.PersonRepository) *AddPersonService {
    return &AddPersonService{repo: repo}
}

func (s *AddPersonService) AddPerson(name string, age int, gender bool) error {
    person := domain.NewPerson(name, age, gender)
    return s.repo.Save(person)
}
