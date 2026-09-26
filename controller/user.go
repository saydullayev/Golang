package controller

import (
	// "app/calc"
	"app/models"
	

	"log"

	"github.com/go-faker/faker/v4"
)


func (c *Controller) UserGenerate(count int) []*models.User {
	var users []*models.User
	for count > 0 {
		users = append(users, &models.User{
			Name: faker.Name(),
			PhoneNumber: faker.Phonenumber(),
		})
		count --
	}
	return users
}

func (c *Controller) UserGetList(req *models.UserGEtListRequest) *models.UserGetListResponse{

	log.Printf("UserGetList req : %+v\n", req)

	var (
		offset = c.Cfg.DefaultOffset
		limit = c.Cfg.DefaultLimit
	response = &models.UserGetListResponse{}
	)

	if req.Offset > 0 {
		offset = req.Offset
	}

	if req.Limit > 0 {
		limit = req.Limit
	}

	

	response.Count = len(c.Users)
	 if len(c.Users) < limit + offset {
		response.Users = c.Users[offset:]
	 } else {
		response.Users = c.Users[offset: offset + limit]

	 }



	return response
	
}
