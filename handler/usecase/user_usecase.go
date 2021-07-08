package usecase

import (
	"encoding/json"
	"fasthttp_crud/handler"
	"fasthttp_crud/model"
	"fasthttp_crud/model/constant"
	"fasthttp_crud/util"
	"fasthttp_crud/util/log"
	"fmt"
	"github.com/go-pg/pg/v10"
	"net/http"
	"strconv"
	"strings"
)

type UserUsecase struct {
	iUserMapper handler.IUserMapper
	iUserRepo   handler.IUserRepo
}

func NewUserUsecase(iUserMapper handler.IUserMapper, iUserRepo handler.IUserRepo) handler.IUserUsecase {
	return &UserUsecase{iUserMapper, iUserRepo}
}

func (u *UserUsecase) GetAll(traceId, sLimit, sOffset string) (httpCode int, bodyResponse interface{}) {
	limit, _ := strconv.Atoi(sLimit)
	offset, _ := strconv.Atoi(sOffset)
	list, total, err := u.iUserRepo.GetAll(traceId, limit, offset)
	if err != nil {
		return http.StatusInternalServerError, util.Error(constant.ERROR_DB, err.Error())
	}
	listDTO := u.iUserMapper.ToListUserDTO(list)
	return http.StatusOK, util.SuccessWithData("success get data", listDTO, &total)
}

func (u *UserUsecase) GetUser(traceId, userId string) (httpCode int, bodyResponse interface{}) {
	user, err := u.iUserRepo.GetUserById(traceId, userId)
	if err == pg.ErrNoRows {
		return http.StatusOK, util.Error(constant.ERROR_DB, "data not found")
	} else if err != nil {
		return http.StatusInternalServerError, util.Error(constant.ERROR_DB, err.Error())
	}
	userDTO := u.iUserMapper.ToUserDTO(user)
	return http.StatusOK, util.SuccessWithData("success get data", userDTO, nil)
}

func (u *UserUsecase) AddNewUser(traceId string, reqBody []byte) (httpCode int, bodyResponse interface{}) {
	payload := model.CreateUserDTO{}
	err := json.Unmarshal(reqBody, &payload)
	if err != nil {
		log.Error(err, traceId)
		return http.StatusBadRequest, util.Error(constant.ERROR_REQUEST_BODY, "request body not well formated")
	} else if payload.Username == "" {
		log.Error(err, traceId)
		return http.StatusBadRequest, util.Error(constant.ERROR_REQUEST_BODY, "username cant be empty")
	}
	user := u.iUserMapper.ToUserOnCreateUser(payload)
	user.EncrpytPasswordWithHashHMAC_SHA256()
	inserted, err := u.iUserRepo.AddNewUser(traceId, user)
	if !inserted {
		return http.StatusConflict, util.Error(constant.ERROR_DB, "data already exist with username "+user.Username)
	} else if err != nil {
		return http.StatusInternalServerError, util.Error(constant.ERROR_DB, err.Error())
	}
	return http.StatusOK, util.Success("success add user data with id "+user.Id)
}

func (u *UserUsecase) UpdateUser(traceId, userId string, reqBody []byte) (httpCode int, bodyResponse interface{}) {
	oldUser, err :=  u.iUserRepo.GetUserById(traceId, userId)
	if err == pg.ErrNoRows {
		return http.StatusOK, util.Error(constant.ERROR_DB, fmt.Sprintf("data with id %v not found", userId))
	} else if err != nil {
		return http.StatusInternalServerError, util.Error(constant.ERROR_DB, err.Error())
	}
	payload := model.UpdateUserDTO{}
	err = json.Unmarshal(reqBody, &payload)
	if err != nil {
		log.Error(err, traceId)
		return http.StatusBadRequest, util.Error(constant.ERROR_REQUEST_BODY, "request body not well formated")
	} else if payload.Username == "" {
		log.Error(err, traceId)
		return http.StatusBadRequest, util.Error(constant.ERROR_REQUEST_BODY, "username cant be empty")
	}
	updatedUser := u.iUserMapper.ToUserOnUpdateUser(payload, oldUser)
	updatedUser.EncrpytPasswordWithHashHMAC_SHA256()
	err = u.iUserRepo.UpdateUserById(traceId, updatedUser)
	if err != nil {
		return http.StatusInternalServerError, util.Error(constant.ERROR_DB, err.Error())
	}
	return http.StatusOK, util.Success("success update user data with id "+updatedUser.Id)
}

func (u *UserUsecase) DeleteUser(traceId, userId string) (httpCode int, bodyResponse interface{}) {
	err := u.iUserRepo.DeleteUserById(traceId, userId)
	if err != nil {
		if strings.Contains(err.Error(), "data deleted fail, row affected") {
			return http.StatusBadRequest, util.Error(constant.ERROR_REQUEST_PATH, "user id not found")
		}
		return http.StatusInternalServerError, util.Error(constant.ERROR_DB, err.Error())
	}
	return http.StatusOK, util.Success("success delete user data with id "+userId)
}
