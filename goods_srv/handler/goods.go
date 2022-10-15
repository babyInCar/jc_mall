package handler

import (
	"context"
	"fmt"
	"jc_mall/goods_src/global"
	"jc_mall/goods_srv/model"
	"jc_mall/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type GoodsServer struct {
	proto.UnimplementedGoodsServer
}

func ModelToResponse(goods model.Goods) proto.GoodsInfoResponse {
	return proto.GoodsInfoResponse{
		Id:              goods.ID,
		CategoryId:      goods.CategoryID,
		Name:            goods.Name,
		GoodsSn:         goods.GoodsSn,
		ClickNum:        goods.ClickNum,
		SoldNum:         goods.SoldNum,
		FavNum:          goods.FavNum,
		MarketPrice:     goods.MarketPrice,
		ShopPrice:       goods.ShopPrice,
		GoodsDesc:       goods.GoodsDesc,
		ShipFree:        goods.ShipFree,
		GoodsCoverImage: goods.GoodsCoverImage,
		IsNew:           goods.IsNew,
		IsHot:           goods.IsHot,
		OnSale:          goods.OnSale,
		DescImages:      goods.DescImages,
		Images:          goods.Images,
		Category: &proto.CategoryBriefInfoResponse{
			Id:   goods.Category.ID,
			Name: goods.Category.Name,
		},
		Brand: &proto.BrandInfoResponse{
			Id:   goods.Brands.ID,
			Name: goods.Brands.Name,
			Logo: goods.Brands.Logo,
		},
	}
}

func (s *GoodsServer) GoodsList(ctx context.Context, req *proto.GoodsFilterRequest) (*proto.GoodsListResponse, error) {
	goodsListResponse := &proto.GoodsListResponse
	var goods []model.Goods
	// queryMap := map[string]interface{}{}
	localDB := global.DB.Model(model.Goods{})
	if req.KeyWords != "" {
		localDB = localDB.Where("name LIKE ?", "%"+req.KeyWords+"%")
	}
	if req.IsHot {
		localDB = localDB.Where("is_hot=true")
	}
	if req.IsNew {
		localDB = localDB.Where("is_new=true")
	}
	if req.PriceMin {
		localDB = localDB.Where("shop_price>=?", req.PriceMin)
	}
	if req.PriceMax {
		localDB = localDB.Where("shop_price<=?", req.PriceMax)
	}
	if req.Brand > 0 {
		localDB = localDB.Where("brand_id=?", req.Brand)
	}
	//通过category去查询到商品信息
	if req.TopCategory > 0 {
		var category model.Category
		if result := global.DB.First(&category, req.TopCategory); result.RowAffected == 0 {
			return nil, status.Errorf(codes.NotFound, "商品分类不存在")
		}
		var subQuery string
		if category.Level == 1 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id in (select id from category where parent_category_id=category_id=%s)", req.TopCategory)
		} else if category.Level == 2 {
			subQuery = fmt.Sprintf("select id from category where parent_category_id=%s", req.TopCategory)
		} else if category.Level == 3 {
			subQuery = fmt.Sprintf("select id from goods where id=%s", req.TopCategory)
		}
		localDB = localDB.Where("category_id in (%s)", subQuery)
	}
	var count int64
	localDB.Count(&count)
	goodsListResponse.Total = int32(count)

	result := localDB.Preload("Category").Preload("Brands").Scopes(Paginate(req.Pages), int(req.PagePerNums)).Find(&goods)
	// localDB = localDB.Scopes().
	if result.Error != nil {
		return nil, result.Error
	}
	for _, good := range goods {
		goodsInfoResponse := ModelToResponse(good)
		goodsListResponse.Data = append(goodsListResponse.Data, &goodsInfoResponse)
	}
	return goodsListResponse, nil

}
