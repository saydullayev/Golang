package models

type User struct {
	Name string
	PhoneNumber string

}

type UserGetListResponse struct {
	Count int
	Users []*User
}

type UserGEtListRequest struct {
	Offset int
	Limit int
}