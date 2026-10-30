package seeders

import (
	"cariinkos/app/facades"
	"cariinkos/app/models"

	"github.com/google/uuid"
)

type UserSeeder struct {
}

// Signature The name and signature of the seeder.
func (s *UserSeeder) Signature() string {
	return "UserSeeder"
}

// Run executes the seeder logic.
func (s *UserSeeder) Run() error {
	password, err := facades.Hash().Make("password1234")
	if err != nil {
		return err
	}

	users := []models.User{
		{ID: uuid.NewString(), Name: "Budi Santoso", Email: "budi@example.com", Password: password, Role: "owner"},
		{ID: uuid.NewString(), Name: "Siti Aminah", Email: "siti@example.com", Password: password, Role: "owner"},
		{ID: uuid.NewString(), Name: "Andi Wijaya", Email: "andi@example.com", Password: password, Role: "owner"},
		{ID: uuid.NewString(), Name: "Dewi Lestari", Email: "dewi@example.com", Password: password, Role: "owner"},
		{ID: uuid.NewString(), Name: "Rizky Pratama", Email: "rizky@example.com", Password: password, Role: "owner"},
	}

	for _, user := range users {
		if err := facades.Orm().Query().Create(&user); err != nil {
			return err
		}
	}

	return nil
}
