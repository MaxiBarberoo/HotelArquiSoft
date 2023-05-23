package dto

type AdminDto struct {
	Name      string `json:"name"`
	LastName  string `json:"last_name"`
	UserEmail string `json:"admin_email"`
	Password  string `json:"password"`
	Id        int    `json:"id"`
}

type AdminsDto []AdminDto
