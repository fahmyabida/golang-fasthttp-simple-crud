package usecase

import (
	"encoding/base64"
	"fasthttp_crud/handler"
	"fasthttp_crud/util/log"
	"fmt"
	"net/http"
	"strings"
)

type LoginUsecase struct {
	iUserRepo handler.IUserRepo
}

func NewLoginUsecase(iUserRepo handler.IUserRepo) handler.ILoginUsecase {
	return &LoginUsecase{iUserRepo}
}

func (u *LoginUsecase) Login(traceId, authorization string) (httpCode int, desc string) {
	if authorization == "" {
		return http.StatusBadRequest, "header Authorization not found"
	}
	userPasswordBase64 := strings.SplitN(authorization, " ", 2)
	userPassword, err := base64.StdEncoding.DecodeString(userPasswordBase64[1])
	if err != nil {
		log.Error(err, traceId)
		return http.StatusInternalServerError, "decode base64 failed"
	}
	sliceUserPassword := strings.SplitN(string(userPassword), ":", 2)
	username := sliceUserPassword[0]
	password := sliceUserPassword[1]
	user, err := u.iUserRepo.GetUserByUsername(traceId, username)
	if err != nil {
		return http.StatusUnauthorized, "user not found"
	}
	if !user.IsValidPassword(password) {
		return http.StatusUnauthorized, "password is invalid"
	}
	return http.StatusOK, fmt.Sprintf("login successfull! welcome %v %v", user.FirstName, user.LastName)
}
