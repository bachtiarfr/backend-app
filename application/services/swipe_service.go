package services

import (
	"backend-app/domain/entities"
	"backend-app/domain/repositories"
	"errors"
	"time"
)

type SwipeService interface {
	Swipe(userID uint, profileID uint, action string) error
	GetDailySwipes(userID uint) ([]entities.Swipe, error)
}

type swipeService struct {
	swipeRepo repositories.SwipeRepository
}

func NewSwipeService(swipeRepo repositories.SwipeRepository) SwipeService {
	return &swipeService{swipeRepo}
}

func (s *swipeService) Swipe(userID uint, profileID uint, action string) error {
	swipes, err := s.GetDailySwipes(userID)
	if err != nil {
		return err
	}

	if len(swipes) >= 10 {
		return errors.New("daily swipe limit reached")
	}

	for _, swipe := range swipes {
		if swipe.ProfileID == profileID {
			return errors.New("profile already swiped today")
		}
	}

	swipe := &entities.Swipe{
		UserID:    userID,
		ProfileID: profileID,
		Action:    action,
		CreatedAt: time.Now(),
	}

	return s.swipeRepo.CreateSwipe(swipe)
}

func (s *swipeService) GetDailySwipes(userID uint) ([]entities.Swipe, error) {
	swipes, err := s.swipeRepo.FindSwipesByUserID(userID)
	if err != nil {
		return nil, err
	}

	var dailySwipes []entities.Swipe
	for _, swipe := range swipes {
		if swipe.CreatedAt.After(time.Now().AddDate(0, 0, -1)) {
			dailySwipes = append(dailySwipes, swipe)
		}
	}

	return dailySwipes, nil
}
