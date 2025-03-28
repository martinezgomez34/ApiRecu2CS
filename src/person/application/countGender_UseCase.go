package application

import "api/src/person/domain"

type CountByGenderService struct {
	repo domain.PersonRepository
}

func NewStatsService(repo domain.PersonRepository) *CountByGenderService {
	return &CountByGenderService{repo: repo}
}

func (s *CountByGenderService) GetGenderStats() (maleCount, femaleCount int, err error) {
	return s.repo.CountByGender()
}