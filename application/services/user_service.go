package services

import (
	"backend-app/application/dto"
	"backend-app/domain/entities"
	"backend-app/domain/repositories"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UserService interface {
	SignUp(userDto dto.UserDTO) error
	Login(userDto dto.UserDTO) (string, error)
	PurchasePremium(userID uint, packageType string) error
}

type userService struct {
	userRepo repositories.UserRepository
}

func NewUserService(userRepo repositories.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) SignUp(userDto dto.UserDTO) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDto.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user := &entities.User{
		Username: userDto.Username,
		Email:    userDto.Email,
		Password: string(hashedPassword),
	}

	return s.userRepo.CreateUser(user)
}

func (s *userService) Login(userDto dto.UserDTO) (string, error) {
	user, err := s.userRepo.FindUserByEmail(userDto.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userDto.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": user.ID,
		"exp":    time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// PurchasePremium allows a user to purchase a premium package
func (s *userService) PurchasePremium(userID uint, packageType string) error {
	if packageType == "no_swipe_quota" {
		return s.userRepo.UpgradeToPremium(userID, packageType)
	}
	return errors.New("invalid package type")
}
