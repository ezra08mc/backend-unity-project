package service

import (
	"github.com/ezra08mc/backend-unity-project/config/pkg/errs"
	"github.com/ezra08mc/backend-unity-project/config/pkg/token"
	"github.com/ezra08mc/backend-unity-project/contract"
	"github.com/ezra08mc/backend-unity-project/database"
	"github.com/ezra08mc/backend-unity-project/dto"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type authService struct {
	authRepo contract.AuthRepository
}

func ImplAuthService(authRepo contract.AuthRepository) contract.AuthService {
	return &authService{authRepo: authRepo}
}

func (s *authService) Register(req dto.RegisterRequest) (*dto.RegisterResponse, error) {
	// Check if email already exists
	_, err := s.authRepo.FindByEmail(req.Email)
	if err == nil {

		return nil, errs.BadRequest("email already registered")
	}

	if err != gorm.ErrRecordNotFound {
		return nil, errs.InternalServerError("failed to check email availability")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errs.InternalServerError("failed to process password")
	}

	user := &database.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     "user",
	}

	if err := s.authRepo.CreateUser(user); err != nil {
		return nil, errs.InternalServerError("failed to create user account")
	}

	return &dto.RegisterResponse{
		Success: true,
		Message: "Registration successful",
		Data: dto.UserData{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}

func (s *authService) Login(req dto.LoginRequest) (*dto.LoginResponse, error) {
	user, err := s.authRepo.FindByEmail(req.Email)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.Unauthorized("invalid email or password")
		}
		return nil, errs.InternalServerError("failed to authenticate user")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errs.Unauthorized("invalid email or password")
	}

	tokenString, err := token.GenerateToken(&token.UserAuthToken{
		ID:       user.ID,
		Email:    user.Email,
		Username: user.Name,
		Role:     user.Role,
	})
	if err != nil {
		return nil, errs.InternalServerError("failed to generate authentication token")
	}

	return &dto.LoginResponse{
		Success: true,
		Message: "Login successful",
		Data: dto.LoginData{
			Token: tokenString,
			User: dto.UserData{
				ID:    user.ID,
				Name:  user.Name,
				Email: user.Email,
				Role:  user.Role,
			},
		},
	}, nil
}

func (s *authService) GetProfile(userID int) (*dto.ProfileResponse, error) {
	user, err := s.authRepo.FindByID(userID)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errs.NotFound("user not found")
		}
		return nil, errs.InternalServerError("failed to fetch user profile")
	}

	return &dto.ProfileResponse{
		Success: true,
		Data: dto.UserData{
			ID:    user.ID,
			Name:  user.Name,
			Email: user.Email,
			Role:  user.Role,
		},
	}, nil
}
