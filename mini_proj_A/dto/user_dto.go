package dto

import (
	"net/mail"
	"strings"
	"time"

	"gopractice/mini_proj_A/models"
)

// CreateUserRequest defines the incoming payload to create a new user.
type CreateUserRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
	Role  string `json:"role"`
}

// Validate checks all fields in CreateUserRequest and returns field-level error messages.
func (req *CreateUserRequest) Validate() map[string]string {
	errs := make(map[string]string)

	// 1. Name validation
	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		errs["name"] = "name is required"
	} else if len(req.Name) < 2 || len(req.Name) > 50 {
		errs["name"] = "name must be between 2 and 50 characters"
	}

	// 2. Email validation
	req.Email = strings.TrimSpace(strings.ToLower(req.Email))
	if req.Email == "" {
		errs["email"] = "email is required"
	} else {
		_, err := mail.ParseAddress(req.Email)
		if err != nil || !strings.Contains(req.Email, ".") {
			errs["email"] = "invalid email address format"
		}
	}

	// 3. Age validation
	if req.Age <= 0 {
		errs["age"] = "age must be greater than 0"
	} else if req.Age > 120 {
		errs["age"] = "age must be 120 or below"
	}

	// 4. Role validation
	req.Role = strings.TrimSpace(strings.ToLower(req.Role))
	allowedRoles := map[string]bool{"admin": true, "user": true, "manager": true}
	if req.Role == "" {
		req.Role = "user" // Default value if not specified
	} else if !allowedRoles[req.Role] {
		errs["role"] = "role must be one of: 'admin', 'user', or 'manager'"
	}

	return errs
}

// UpdateUserRequest defines what fields can be updated.
type UpdateUserRequest struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Role string `json:"role"`
}

// Validate checks fields for UpdateUserRequest.
func (req *UpdateUserRequest) Validate() map[string]string {
	errs := make(map[string]string)

	req.Name = strings.TrimSpace(req.Name)
	if req.Name == "" {
		errs["name"] = "name cannot be empty"
	} else if len(req.Name) < 2 || len(req.Name) > 50 {
		errs["name"] = "name must be between 2 and 50 characters"
	}

	if req.Age <= 0 || req.Age > 120 {
		errs["age"] = "age must be between 1 and 120"
	}

	req.Role = strings.TrimSpace(strings.ToLower(req.Role))
	allowedRoles := map[string]bool{"admin": true, "user": true, "manager": true}
	if req.Role != "" && !allowedRoles[req.Role] {
		errs["role"] = "role must be one of: 'admin', 'user', or 'manager'"
	}

	return errs
}

// UserResponse defines the shape of user data exposed to the client.
type UserResponse struct {
	ID        int64     `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Age       int       `json:"age"`
	Role      string    `json:"role"`
	CreatedAt time.Time `json:"created_at"`
}

// ToUserResponse converts a database User model into an outgoing DTO.
func ToUserResponse(u *models.User) UserResponse {
	return UserResponse{
		ID:        u.ID,
		Name:      u.Name,
		Email:     u.Email,
		Age:       u.Age,
		Role:      u.Role,
		CreatedAt: u.CreatedAt,
	}
}

// ToUserResponses converts a slice of DB models to response DTOs.
func ToUserResponses(users []models.User) []UserResponse {
	res := make([]UserResponse, 0, len(users))
	for i := range users {
		res = append(res, ToUserResponse(&users[i]))
	}
	return res
}
