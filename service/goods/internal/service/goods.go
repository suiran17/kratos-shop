package service

import (
	v1 "goods/api/goods/v1"
	"goods/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
)

// GoodsService is a goods service.
type GoodsService struct {
	v1.UnimplementedGoodsServer
	cac *biz.CategoryUsecase
	// bc      *biz.BrandUsecase
	// gt      *biz.GoodsTypeUsecase
	// s       *biz.SpecificationUsecase
	// ga      *biz.GoodsAttrUsecase
	// g       *biz.GoodsUsecase
	// esGoods *biz.EsGoodsUsecase
	log *log.Helper
}

// NewGoodsService new a goods service.
func NewGoodsService(
// bc *biz.BrandUsecase,
	cac *biz.CategoryUsecase,
// gt *biz.GoodsTypeUsecase,
// s *biz.SpecificationUsecase,
// ga *biz.GoodsAttrUsecase,
// gc *biz.GoodsUsecase,
// esGoods *biz.EsGoodsUsecase,
	logger log.Logger,
) *GoodsService {
	return &GoodsService{
		// bc:      bc,
		cac: cac,
		// gt:      gt,
		// s:       s,
		// ga:      ga,
		// g:       gc,
		// esGoods: esGoods,
		log: log.NewHelper(logger),
	}
}
