package usecase

import "github.com/asatake/trophies-management/api/trophy/domain"

type ContentRepository interface {
	FindContentBy(int) (domain.Content, error)
	FindAchievementsBy(int) ([]domain.Achievement, error)
	FindChecksBy(int) (domain.Check, error)
}
