package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"gopractice/mini_proj_A/dto"
	"gopractice/mini_proj_A/models"
	"gopractice/mini_proj_A/repository"
)

type UserHandler struct {
	repo repository.UserRepository
}

func NewUserHandler(repo repository.UserRepository) *UserHandler {
	return &UserHandler{repo: repo}
}

// CreateUser handles POST /users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// 1. Decode JSON Request Body
	var req dto.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.SendError(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error(), nil)
		return
	}

	// 2. Validate Request DTO fields
	if fieldErrors := req.Validate(); len(fieldErrors) > 0 {
		dto.SendError(w, http.StatusBadRequest, "Validation failed", fieldErrors)
		return
	}

	// 3. Business logic check: duplicate email
	exists, err := h.repo.ExistsByEmail(r.Context(), req.Email)
	if err != nil {
		dto.SendError(w, http.StatusInternalServerError, "Failed to verify email availability", nil)
		return
	}
	if exists {
		dto.SendError(w, http.StatusConflict, "Email already registered", map[string]string{
			"email": "a user with this email address already exists",
		})
		return
	}

	// 4. Map DTO to DB Model
	newUser := &models.User{
		Name:  req.Name,
		Email: req.Email,
		Age:   req.Age,
		Role:  req.Role,
	}

	// 5. Save to database
	_, err = h.repo.Create(r.Context(), newUser)
	if err != nil {
		dto.SendError(w, http.StatusInternalServerError, "Failed to save user: "+err.Error(), nil)
		return
	}

	// 6. Return Response DTO (201 Created)
	dto.SendSuccess(w, http.StatusCreated, "User created successfully", dto.ToUserResponse(newUser))
}

// GetUserByID handles GET /users/{id}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Go 1.22+ standard library extracts path variables via r.PathValue
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		dto.SendError(w, http.StatusBadRequest, "Invalid user ID, must be a positive integer", nil)
		return
	}

	user, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			dto.SendError(w, http.StatusNotFound, "User not found", nil)
			return
		}
		dto.SendError(w, http.StatusInternalServerError, "Failed to fetch user", nil)
		return
	}

	dto.SendSuccess(w, http.StatusOK, "User retrieved successfully", dto.ToUserResponse(user))
}

// GetAllUsers handles GET /users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.repo.GetAll(r.Context())
	if err != nil {
		dto.SendError(w, http.StatusInternalServerError, "Failed to fetch users", nil)
		return
	}

	dto.SendSuccess(w, http.StatusOK, "Users retrieved successfully", dto.ToUserResponses(users))
}

// UpdateUser handles PUT /users/{id}
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		dto.SendError(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	var req dto.UpdateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		dto.SendError(w, http.StatusBadRequest, "Invalid JSON payload", nil)
		return
	}

	if fieldErrors := req.Validate(); len(fieldErrors) > 0 {
		dto.SendError(w, http.StatusBadRequest, "Validation failed", fieldErrors)
		return
	}

	// Fetch existing user
	existing, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			dto.SendError(w, http.StatusNotFound, "User not found", nil)
			return
		}
		dto.SendError(w, http.StatusInternalServerError, "Failed to fetch user", nil)
		return
	}

	// Update fields
	existing.Name = req.Name
	existing.Age = req.Age
	if req.Role != "" {
		existing.Role = req.Role
	}

	if err := h.repo.Update(r.Context(), existing); err != nil {
		dto.SendError(w, http.StatusInternalServerError, "Failed to update user", nil)
		return
	}

	dto.SendSuccess(w, http.StatusOK, "User updated successfully", dto.ToUserResponse(existing))
}

// DeleteUser handles DELETE /users/{id}
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil || id <= 0 {
		dto.SendError(w, http.StatusBadRequest, "Invalid user ID", nil)
		return
	}

	if err := h.repo.Delete(r.Context(), id); err != nil {
		if errors.Is(err, repository.ErrNotFound) {
			dto.SendError(w, http.StatusNotFound, "User not found", nil)
			return
		}
		dto.SendError(w, http.StatusInternalServerError, "Failed to delete user", nil)
		return
	}

	dto.SendSuccess(w, http.StatusOK, "User deleted successfully", map[string]int64{"deleted_id": id})
}
