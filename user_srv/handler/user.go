package handler

import (
	"context"
	"fmt"
	"go/format"
	"jc_mall/user_srv/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServer struct{}

func ModelToRspone(user model.User) proto.UserInfoResponse {
	//在grpc的message中字段有默认值，你不能随便赋值nil进去，容易出错
	userInfoRsp := proto.UserInfoResponse{
		Id:       user.ID,
		PassWord: user.Password,
		NickName: user.NickName,
		Gender:   user.Gender,
		Role:     int32(user.Role),
		Mobile:   user.Mobile,
	}
	if user.Birthday != nil {
		userInfoRsp.BirthDay = uint64(user.Birthday.Unix())
	}
	return userInfoRsp
}

func Paginate(page, pageSize) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}

		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}

func (s *UserServer) GetUserList(ctx context.Context, req *proto.PageInfo) (*proto.UserListResponse, error) {
	//获取用户列表
	var users []model.User
	result := global.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	fmt.Println("用户列表")
	rsp := &proto.UserListResponse{}
	rsp.Total = int32(result.RowsAffected)

	global.DB.Scopes(Paginate(int(req.Pn), int(req.PSize))).Find(&users)

	for _, user := range users {
		userInfoRsp := ModelToRsponse(user)
		rsp.Data = append(rsp.Data, &userInfoRsp)
	}
	return rsp, nil
}

func (s *UserServer) GetUserByMobile(ctx context.Context, req *proto.MobileRequest) *proto.UserInfoResponse {
	var user model.User
	result := global.DB.Where(&model.User{Mobile: req.Mobile}).First(&user)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, format:"用户不存在")
	}
	
	if result.Error != nil{
		return nil, result.Error
	}

	userInfoRsp := ModelToRspone(result)
	return 	&userInfoRsp, nil
}

func (s *UserServer) GetUserById(ctx context.Context,req *proto.IdRequest) (*proto.UserInfoResponse, error){
	var user model.User
	result := global.DB.First(&user, req.Id)
	if result.RowsAffected == 0 {
		return nil, status.Errorf(codes.NotFound, format:"用户不存在")
	}
	
	if result.Error != nil{
		return nil, result.Error
	}

	userInfoRsp := ModelToRspone(result)
	return 	&userInfoRsp, nil
}
