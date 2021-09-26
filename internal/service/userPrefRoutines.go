package service

import (
	"github.com/monkeydioude/drannoc/internal/entity"
	"github.com/monkeydioude/drannoc/internal/repository"
)

func FetchUserPreferences(
	repo *repository.UserPreferences,
	userID string,
) (*entity.UserPreferences, error) {
	userPref := entity.NewUserPreferences()
	_, err := repo.Load(userPref, userID)
	if err != nil {
		return nil, err
	}

	return userPref, err
}
