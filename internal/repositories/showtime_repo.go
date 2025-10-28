package repositories

import (
	"time"

	"github.com/BraveHeart-tex/Cinema-Core-Service/internal/models"
	"gorm.io/gorm"
)

type ShowtimeRepository struct {
	db *gorm.DB
}

func NewShowtimeRepository(db *gorm.DB) *ShowtimeRepository {
	return &ShowtimeRepository{db: db}
}

func (r *ShowtimeRepository) Create(showtime *models.Showtime) error {
	return r.db.Create(showtime).Error
}

func (r *ShowtimeRepository) ExistsOverlap(theaterID uint, start, end time.Time) (bool, error) {
	var count int64
	err := r.db.Model(&models.Showtime{}).
		Where("theater_id = ? AND start_time < ? AND end_time > ?", theaterID, end, start).
		Count(&count).Error
	return count > 0, err
}
