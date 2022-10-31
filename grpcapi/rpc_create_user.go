package grpcapi

import (
	"context"
	"e-wallet-app/modules/user/userbiz"
	"e-wallet-app/modules/user/userenum"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userrepo"
	"e-wallet-app/modules/user/userstore"
	"e-wallet-app/protogen"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Server) CreateUser(ctx context.Context, protoReq *protogen.CreateUserRequest) (*protogen.CreateUserResponse, error) {
	var req usermodel.CreateUserRequest
	err := mapstructure.Decode(protoReq, &req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	req.FillDefault()
	req.Role = int32(userenum.RoleUser)
	store := userstore.NewSqlStore(s.appCtx.GetMainDBConnection())
	repo := userrepo.NewUserRepo(store)
	biz := userbiz.NewUserBiz(repo)

	user, err := biz.Create(ctx, &req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var res protogen.CreateUserResponse
	err = mapstructure.Decode(user, &res)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	res.AppRole = userenum.Role(user.Role).String()
	res.CreatedAt = timestamppb.New(user.CreatedAt)
	res.UpdatedAt = timestamppb.New(user.UpdatedAt)

	return &res, nil
}
