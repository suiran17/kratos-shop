package service

import (
	"context"
	v1 "goods/api/goods/v1"
	"goods/internal/domain"
)

func (g *GoodsService) CreateGoodsType(ctx context.Context, r *v1.GoodsTypeRequest) (*v1.GoodsTypeResponse, error) {
	id, err := g.gt.GoosTypeCreate(ctx, &domain.GoodsType{
		Name:      r.Name,
		TypeCode:  r.TypeCode,
		NameAlias: r.NameAlias,
		IsVirtual: r.IsVirtual,
		Desc:      r.Desc,
		Sort:      r.Sort,
	})
	if err != nil {
		return nil, err
	}
	return &v1.GoodsTypeResponse{
		Id: id,
	}, nil
}
