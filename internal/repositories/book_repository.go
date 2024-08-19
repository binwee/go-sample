package repositories

import (
	"context"

	"github.com/binwee/go-sample/internal/models"
	"gorm.io/gorm"
)

type BookmarkRepository interface {
	GetAll(ctx context.Context) ([]models.Book, error)
	GetByID(ctx context.Context, id int) (*models.Book, error)
	Create(ctx context.Context, b models.Book) (*models.Book, error)
	Update(ctx context.Context, b models.Book) error
	Delete(ctx context.Context, id int) error
}

type bookmarkRepo struct {
	db *gorm.DB
}

func NewBookmarkRepository(db *gorm.DB) BookmarkRepository {
	return bookmarkRepo{
		db: db,
	}
}

func (r bookmarkRepo) GetAll(ctx context.Context) ([]models.Book, error) {
	return nil, nil
}

func (r bookmarkRepo) GetByID(ctx context.Context, id int) (*models.Book, error) {
	return nil, nil
}

func (r bookmarkRepo) Create(ctx context.Context, b models.Book) (*models.Book, error) {
	return nil, nil
}

func (r bookmarkRepo) Update(ctx context.Context, b models.Book) error {
	return nil
}

func (r bookmarkRepo) Delete(ctx context.Context, id int) error {
	return nil
}
