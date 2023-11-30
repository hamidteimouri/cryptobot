package postgres

import (
	"context"
	"errors"
	"gorm.io/gorm"
)

type Postgres struct {
	db *gorm.DB
}

func NewPostgres(db *gorm.DB) *Postgres {
	pds := &Postgres{db: db}
	return pds
}

func (p *Postgres) GetStateOfUser(ctx context.Context, telegramId string) (*CommandEntity, error) {

	entity := CommandEntity{}
	result := p.db.WithContext(ctx).Model(&entity).
		Where("telegram_id = ?", telegramId).
		First(&entity)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, result.Error
	}

	return &entity, nil
}
