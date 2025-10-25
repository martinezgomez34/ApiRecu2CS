package application

import "api/src/person/domain"

type GetPersonByIDService struct {
    repo domain.PersonRepository
}

func NewGetPersonByIDService(repo domain.PersonRepository) *GetPersonByIDService {
    return &GetPersonByIDService{repo: repo}
}

func (s *GetPersonByIDService) GetByID(id uint) (*domain.Person, error) {
    return s.repo.GetByID(id)
}
