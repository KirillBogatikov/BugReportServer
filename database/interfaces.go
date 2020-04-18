package database

import "server/models"

type ReviewerRepository interface {
	Add(reviewer *models.Reviewer) error
	Remove(reviewer models.Reviewer) error
	Check(reviewer models.Reviewer) error
	All(mapper models.InfoMapper) ([]models.Reviewer, error)
	Close() error
}
