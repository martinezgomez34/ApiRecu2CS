package infrastructure

import (
	"api/src/person/domain"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
	"os"
)

type PostgresPersonRepository struct {
	db *gorm.DB
}

func NewPostgresPersonRepository() (*PostgresPersonRepository, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("No se pudo cargar .env")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	port := os.Getenv("DB_PORT")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&domain.Person{})
	return &PostgresPersonRepository{db: db}, nil
}

func (r *PostgresPersonRepository) Save(person *domain.Person) error {
	return r.db.Create(person).Error
}

func (r *PostgresPersonRepository) CountByGender() (maleCount, femaleCount int, err error) {
	var males, females int64
	if err = r.db.Model(&domain.Person{}).Where("gender = ?", true).Count(&males).Error; err != nil {
		return
	}
	if err = r.db.Model(&domain.Person{}).Where("gender = ?", false).Count(&females).Error; err != nil {
		return
	}
	return int(males), int(females), nil
}

func (r *PostgresPersonRepository) GetAll() ([]*domain.Person, error) {
	var persons []*domain.Person
	err := r.db.Find(&persons).Error
	return persons, err
}