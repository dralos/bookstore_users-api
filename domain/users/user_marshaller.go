package users

type PublicUser struct {
	Id          int64  `json:"id"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

type PrivateUser struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
	Status      string `json:"status"`
}

func (users Users) Marshall(isPublic bool) []any {
	result := make([]any, len(users))
	for index, user := range users {
		result[index] = user.Marshall(isPublic)
	}
	return result
}

func (users *User) Marshall(isPublic bool) interface{} {
	if isPublic {
		return PublicUser{
			Id:          users.Id,
			DateCreated: users.DateCreated,
			Status:      users.Status,
		}
	}
	return PrivateUser{
		Id:          users.Id,
		FirstName:   users.FirstName,
		LastName:    users.LastName,
		Email:       users.Email,
		DateCreated: users.DateCreated,
		Status:      users.Status,
	}
}
