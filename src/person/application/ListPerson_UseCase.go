package application

import "api/src/person/domain"

type ListPersonService struct {
    repo domain.PersonRepository
}

func NewListPersonService(repo domain.PersonRepository) *ListPersonService {
    return &ListPersonService{repo: repo}
}

func (s *ListPersonService) GetAllPersons() ([]*domain.Person, error) {
    return s.repo.GetAll()
}