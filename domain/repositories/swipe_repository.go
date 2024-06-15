package repositories

import (
	"backend-app/domain/entities"
	"backend-app/logging"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type SwipeRepository interface {
	CreateSwipe(swipe *entities.Swipe) error
	FindSwipesByUserID(userID uint) ([]entities.Swipe, error)
}

type swipeRepository struct {
	db  *gorm.DB
	log *logrus.Logger
}

func NewSwipeRepository(db *gorm.DB) SwipeRepository {
	return &swipeRepository{
		db:  db,
		log: logging.InitLogger(),
	}
}

func (r *swipeRepository) CreateSwipe(swipe *entities.Swipe) error {
	if err := r.db.Create(swipe).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"module": "repository",
			"error":  err.Error(),
		}).Error("Error creating swipe")
		return err
	}
	return nil
}

func (r *swipeRepository) FindSwipesByUserID(userID uint) ([]entities.Swipe, error) {
	var swipes []entities.Swipe
	if err := r.db.Where("user_id = ?", userID).Find(&swipes).Error; err != nil {
		r.log.WithFields(logrus.Fields{
			"module": "repository",
			"error":  err.Error(),
		}).Error("Error finding swipes by user ID")
		return nil, err
	}
	return swipes, nil
}
