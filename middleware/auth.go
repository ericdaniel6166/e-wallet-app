package middleware

import (
	"e-wallet-app/common"
	"e-wallet-app/component"
	"e-wallet-app/component/token"
	"e-wallet-app/modules/user/userenum"
	"e-wallet-app/modules/user/usermodel"
	"e-wallet-app/modules/user/userstore"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

const (
	authHeaderKey  = "Authorization"
	authTypeBearer = "Bearer"
)

func ErrWrongAuthHeader(err error) *common.AppError {
	return common.NewCustomError(
		err,
		fmt.Sprintf("wrong authentication header"),
		fmt.Sprintf("ErrWrongAuthHeader"),
	)
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ")
	//"Authorization" : "Bearer {token}"

	if parts[0] != authTypeBearer || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", ErrWrongAuthHeader(nil)
	}

	return parts[1], nil
}

// RequiredAuth
// 1. Get token from header
// 2. Validate token and parse to payload
// 3. From the token payload, we use username to find from DB
func RequiredAuth(appCtx component.AppContext, role userenum.Role) gin.HandlerFunc {
	//tokenMaker, err := token.NewJWTMaker(appCtx.TokenSymmetricKey())
	tokenMaker, err := token.NewPasetoMaker(appCtx.TokenSymmetricKey())
	if err != nil {
		panic(common.ErrUnauthorized(err))
	}

	return func(ctx *gin.Context) {

		tok, err := extractTokenFromHeaderString(ctx.GetHeader(authHeaderKey))

		if err != nil {
			panic(common.ErrUnauthorized(err))
		}

		payload, err := tokenMaker.VerifyToken(tok)
		if err != nil {
			panic(common.ErrUnauthorized(err))
		}

		store := userstore.NewSqlStore(appCtx.GetMainDBConnection())
		user, err := store.GetUserByUsername(ctx.Request.Context(), payload.Username)
		//101&001 =001 == 001

		if !(user.GetRole()&role == role) {
			panic(common.ErrNoPermission(common.NoPermission))
		}

		if user.Status == false {
			panic(common.ErrNoPermission(common.ErrEntityBlocked(usermodel.EntityName, common.RecordIsBlocked)))
		}

		ctx.Set(common.CurrentUser, user)
		ctx.Next()
	}
}
