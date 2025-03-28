package domain

type Person struct {
    Name   string
    Age    int
    Gender bool // true: male, false: female
}

func NewPerson(name string, age int, gender bool) *Person {
    return &Person{
        Name:   name,
        Age:    age,
        Gender: gender,
    }
}