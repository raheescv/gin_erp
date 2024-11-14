package resources

import "product-store/app/models"

type Resource struct {
	ID    uint   `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func UserResource(user models.User) Resource {
	return Resource{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}
