package handler

type ILoginUsecase interface {
	Login(traceId, authorization string) (httpCode int, desc string)
}

type IUserUsecase interface {
	GetAll(traceId, limit, offset string) (httpCode int, bodyResponse interface{})
	GetUser(traceId, userId string) (httpCode int, bodyResponse interface{})
	AddNewUser(traceId string, reqBody []byte) (httpCode int, bodyResponse interface{})
	UpdateUser(traceId, userId string, reqBody []byte) (httpCode int, bodyResponse interface{})
	DeleteUser(traceId, userId string) (httpCode int, bodyResponse interface{})
}
