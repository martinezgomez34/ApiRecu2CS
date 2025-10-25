package domain

type PersonRepository interface {
    Save(person *Person) error
    CountByGender() (maleCount, femaleCount int, err error)
    GetAll() ([]*Person, error)
    GetByID(id uint) (*Person, error)
}