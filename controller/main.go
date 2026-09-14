package controller

import (
	// "app/calc"
	"github.com/go-faker/faker/v4"
	"app/models"

	
)


func GenerateUser(count int) []models.User {
	var users []models.User
	for count >= 0 {
		users = append(users, models.User{
			Name: faker.Name(),
			PhoneNumber: faker.Phonenumber(),
		})
		count --
	}
	return users
}