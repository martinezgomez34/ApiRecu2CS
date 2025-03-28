package infrastructure

import (
	"api/src/person/domain"
	"sync"
)

type InMemoryPersonRepository struct {
    persons []*domain.Person
    mutex   sync.Mutex
}

func NewInMemoryPersonRepository() *InMemoryPersonRepository {
    return &InMemoryPersonRepository{
        persons: make([]*domain.Person, 0),
    }
}

func (r *InMemoryPersonRepository) Save(person *domain.Person) error {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    r.persons = append(r.persons, person)
    return nil
}

func (r *InMemoryPersonRepository) CountByGender() (maleCount, femaleCount int, err error) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    for _, person := range r.persons {
        if person.Gender {
            maleCount++
        } else {
            femaleCount++
        }
    }
    return
}

func (r *InMemoryPersonRepository) GetAll() ([]*domain.Person, error) {
    r.mutex.Lock()
    defer r.mutex.Unlock()

    return r.persons, nil
}