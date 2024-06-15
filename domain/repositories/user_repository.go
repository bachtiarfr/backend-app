package repositories

import (
	"backend-app/domain/entities"
	"backend-app/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user *entities.User) error
	FindUserByEmail(email string) (*entities.User, error)
	UpgradeToPremium(userID uint, packageType string) error
}

type userRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db:  db,
		log: logging.InitLogger(),
	}
}

func (r *userRepository) CreateUser(user *entities.User) error {
	if err := r.db.Create(user).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"module": "repository",
			"error":  err.Error(),
		}).Error("Error creating user")
		return err
	}
	return nil
}

func (r *userRepository) FindUserByEmail(email string) (*entities.User, error) {
	var user entities.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"module": "repository",
			"error":  err.Error(),
		}).Error("Error finding user by email")
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) UpgradeToPremium(userID uint, packageType string) error {
	var user entities.User
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}
	user.Premium = true
	return r.db.Save(&user).Error
}
