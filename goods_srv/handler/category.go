package handler

import (
	"context"
	"encoding/json"
	"jc_mall/goods_srv/global"
	"jc_mall/goods_srv/model"
	"jc_mall/goods_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GoodsServer struct{}

func (s *GoodsServer) GetAllCategorysList(context.Context, *emptypb.Empty) (*proto.CategoryListResponse, error) {
	var categorys []model.Category
	global.DB.Where(&model.Category{Level: 1}).Preload("SubCategory.SubCategory").Find(&categorys)
	b, _ := json.Marshal(&categorys)
	return &proto.CategoryListResponse{JsonData: string(b)}, nil
	// global.DB.Find(&categorys)
}

func (s *GoodsServer) GetSubCategory(ctx context.Context, req *proto.CategoryListRequest) (*proto.SubCategoryListResponse, error) {

	categoryListResponse := proto.SubCategoryListResponse{}
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品分类信息不存在")
	}

	var SubCategorys model.Category
	var subCategoryResponse []*proto.CategoryInfoResponse
	categoryListResponse.Info = &proto.CategoryInfoResponse{
		Id:             category.ID,
		Name:           category.Name,
		Level:          category.Level,
		IsTab:          category.IsTab,
		ParentCategory: category.ParentCategoryID,
	}

	global.DB.Where(&model.Category{ParentCategoryId: req.Id}).Find(&SubCategorys)
	for _, subCategory := range SubCategorys {
		subCategoryResponse = append(subCategoryResponse, &proto.CategoryInfoResponse{
			Id:             subCategory.Id,
			Name:           subCategory.Name,
			Level:          subCategory.Level,
			IsTab:          subCategory.IsTab,
			ParentCategory: subCategory.ParentCategoryID,
		})
	}

	categoryListResponse.SubCategorys = subCategoryResponse
	return &categoryListResponse, nil
}

func (s *GoodsServer) CreateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*proto.CategoryInfoResponse, error) {

	category := model.Category{}

	category.Name = req.Name
	category.Level = req.Level
	if req.Level != 1 {
		category.ParentCategoryID = req.ParentCategory
	}
	category.IsTab = req.IsTab

	global.DB.Save(&category)
	return &proto.CategoryInfoRequest{Id: int32(category.ID)}, nil
}

func (s *GoodsServer) UpdateCategory(ctx context.Context, req *proto.CategoryInfoRequest) (*emptypb.Empty, error) {
	var category model.Category
	if result := global.DB.First(&category, req.Id); result.RowAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品类目不存在")
	}
	if req.Name != "" {
		category.Name = req.Name
	}
	if req.ParentCategory != 0 {
		category.ParentCategoryId = req.ParentCategory
	}
	if req.Level != 0 {
		category.Level = req.Level
	}
	if req.IsTab {
		category.IsTab = true
	}
	global.DB.Save(&category)
	return &emptypb.Empty{}, nil
}

func (s *GoodsServer) DeleteCategory(ctx context.Context, req *proto.DeleteCategoryRequest) (*emptypb.Empty, error) {
	if result := global.DB.Delete(model.Category{}, req.Id); result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, "商品类目不存在")
	}
	return &emptypb.Empty{}, nil
}
