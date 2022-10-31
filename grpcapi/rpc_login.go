package grpcapi

import (
	"context"
	"e-wallet-app/component/token"
	"e-wallet-app/modules/session/sessionstore"
	"e-wallet-app/modules/user/userbiz"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userrepo"
	"e-wallet-app/modules/user/userstore"
	"e-wallet-app/protogen"
	"github.com/mitchellh/mapstructure"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/peer"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s Server) Login(ctx context.Context, protoReq *protogen.LoginRequest) (*protogen.LoginResponse, error) {
	var req usermodel.LoginRequest
	err := mapstructure.Decode(protoReq, &req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	req.FillDefault()

	peerInfo, ok := peer.FromContext(ctx)
	if !ok {
		return nil, status.Errorf(codes.Internal, "cannot get peer information in ctx")
	}

	req.ClientIP = peerInfo.Addr.String()
	mD, _ := metadata.FromIncomingContext(ctx)
	req.UserAgent = mD["user-agent"][0]

	//tokenMaker, err := token.NewJWTMaker(s.appCtx.TokenSymmetricKey())
	tokenMaker, err := token.NewPasetoMaker(s.appCtx.TokenSymmetricKey())

	userStore := userstore.NewSqlStore(s.appCtx.GetMainDBConnection())
	sessionStore := sessionstore.NewSqlStore(s.appCtx.GetMainDBConnection())
	repo := userrepo.NewFullUserRepo(userStore, sessionStore)

	duration := s.appCtx.RefreshTokenDuration()
	biz := userbiz.NewFullUserBiz(repo, s.appCtx, tokenMaker, s.appCtx.AccessTokenDuration(), &duration)

	res, err := biz.Login(ctx, &req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	var protoRes protogen.LoginResponse
	err = mapstructure.Decode(res, &protoRes)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	protoRes.IssuedAt = timestamppb.New(res.IssuedAt)
	protoRes.ExpiredAt = timestamppb.New(res.ExpiredAt)
	return &protoRes, nil
}
