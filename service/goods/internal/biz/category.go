package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
)

type Category struct {
	ID               int32
	Name             string
	ParentCategoryID int32
	SubCategory      []*Category
	Level            int32
	IsTab            bool
	Sort             int32
}

type CategoryInfo struct {
	ID             int32
	Name           string
	ParentCategory int32
	Level          int32
	IsTab          bool
	Sort           int32
}

type CategoryRepo interface {
	AddCategory(context.Context, *CategoryInfo) (*CategoryInfo, error)
	UpdateCategory(context.Context, *CategoryInfo) error
	Category(context.Context) ([]*Category, error)
	GetCategoryByID(ctx context.Context, id int32) (*CategoryInfo, error)
	SubCategory(context.Context, CategoryInfo) ([]*CategoryInfo, error)
	DeleteCategory(context.Context, int32) error
	GetCategoryAll(context.Context, int32, int32) ([]interface{}, error)
}

type CategoryUsecase struct {
	repo CategoryRepo
	log  *log.Helper
}

func NewCategoryUsecase(repo CategoryRepo, logger log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (c *CategoryUsecase) CreateCategory(ctx context.Context, r *CategoryInfo) (*CategoryInfo, error) {
	cateInfo, err := c.repo.AddCategory(ctx, r)
	if err != nil {
		return nil, err
	}
	return cateInfo, nil
}
