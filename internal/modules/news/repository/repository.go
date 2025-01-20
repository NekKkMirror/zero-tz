package repository

import (
	"context"
	"fmt"
	"github.com/NekKkMirror/zero-tz/internal/modules/news/domain"
	"github.com/NekKkMirror/zero-tz/internal/modules/news/domain/model"

	reform "gopkg.in/reform.v1"
)

type Repository interface {
	UpdateNewsPartial(ctx context.Context, input domain.UpdateNewsRequest, newsID int64) error
	ListNews(ctx context.Context, page, limit int) ([]domain.NewsDetail, error)
}

type repository struct {
	db *reform.DB
}

func NewRepository(db *reform.DB) Repository {
	return &repository{db: db}
}

func (r *repository) UpdateNewsPartial(ctx context.Context, input domain.UpdateNewsRequest, newsID int64) error {
	var existing domain.News
	if err := r.db.FindByPrimaryKeyTo(&existing, newsID); err != nil {
		return fmt.Errorf("news item not found, id=%d: %w", newsID, err)
	}

	if input.Title != nil {
		existing.Title = *input.Title
	}
	if input.Content != nil {
		existing.Content = *input.Content
	}

	if err := r.db.Update(&existing); err != nil {
		return fmt.Errorf("failed to update news: %w", err)
	}

	if input.Categories != nil {
		_, err := r.db.Exec("DELETE FROM NewsCategories WHERE NewsId = ?", newsID)
		if err != nil {
			return fmt.Errorf("failed to delete old categories: %w", err)
		}

		for _, catID := range *input.Categories {
			cat := domain.NewsCategory{NewsID: newsID, CategoryID: catID}
			if err := r.db.Insert(&cat); err != nil {
				return fmt.Errorf("failed to insert category link: %w", err)
			}
		}
	}

	return nil
}

func (r *repository) ListNews(ctx context.Context, page, limit int) ([]domain.NewsDetail, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	offset := (page - 1) * limit

	var newsList []domain.News
	query := fmt.Sprintf(`SELECT * FROM News ORDER BY Id LIMIT %d OFFSET %d`, limit, offset)
	if err := r.db.SelectAllTo(&newsList, query); err != nil {
		return nil, err
	}

	results := make([]domain.NewsDetail, 0, len(newsList))
	for _, n := range newsList {
		var cats []domain.NewsCategory
		catQ := "SELECT * FROM NewsCategories WHERE NewsId = ?"
		err := r.db.SelectAllTo(&cats, catQ, n.ID)
		if err != nil {
			return nil, err
		}

		categoryIDs := make([]int64, 0, len(cats))
		for _, c := range cats {
			categoryIDs = append(categoryIDs, c.CategoryID)
		}

		results = append(results, domain.NewsDetail{
			ID:         n.ID,
			Title:      n.Title,
			Content:    n.Content,
			Categories: categoryIDs,
		})
	}

	return results, nil
}
