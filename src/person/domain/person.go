package domain

type Person struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Name   string `gorm:"type:varchar(100);not null" json:"name"`
	Age    int    `gorm:"not null" json:"age"`
	Gender bool   `gorm:"not null" json:"gender"`
}

func NewPerson(name string, age int, gender bool) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		Gender: gender,
	}
}
