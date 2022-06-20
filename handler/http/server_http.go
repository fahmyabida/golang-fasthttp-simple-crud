package http

import (
	"encoding/json"
	"fasthttp_crud/handler"
	"fasthttp_crud/util"
	"fasthttp_crud/util/log"
	"fmt"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

type ServerHttp struct {
	router        *router.Router
	iLoginUsecase handler.ILoginUsecase
	iUserUsecase  handler.IUserUsecase
}

func NewServerHttp(iLoginUsecase handler.ILoginUsecase, iUserUsecase handler.IUserUsecase) handler.IServerHttp {
	router := router.New()
	return &ServerHttp{router, iLoginUsecase, iUserUsecase}
}

func (h *ServerHttp) Routing() {
	h.router.GET("/", func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.SetContentType("application/json")
		ctx.Response.SetStatusCode(200)
		hMap := make(map[string]interface{})
		hMap["message"] = "pong"
		asByteJSON, _ := json.Marshal(hMap)
		ctx.Response.SetBody(asByteJSON)
		fmt.Fprint(ctx)
	})
	h.router.POST("/login", h.Login)
	h.router.GET("/users", h.GetAllUser)
	h.router.GET("/user/{id}", h.GetUser)
	h.router.POST("/user", h.CreateUser)
	h.router.PUT("/user/{id}", h.UpdateUser)
	h.router.DELETE("/user/{id}", h.RemoveUser)
	portHttp := 80
	fmt.Printf("Http listen on port %v \n", portHttp)
	fmt.Println("Ready to serve!")
	fasthttp.ListenAndServe(fmt.Sprintf(":%v", portHttp), h.router.Handler)
}

func (h *ServerHttp) Login(ctx *fasthttp.RequestCtx) {
	traceId := util.CreateTraceID()
	authorization := string(ctx.Request.Header.Peek("Authorization"))
	log.Message(
		traceId,
		"IN",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		"Header Authorization : "+string(authorization))
	httCode, description := h.iLoginUsecase.Login(traceId, authorization)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(httCode)
	hMap := make(map[string]interface{})
	hMap["message"] = description
	asByteJSON, _ := json.Marshal(hMap)
	ctx.Response.SetBody(asByteJSON)
	fmt.Fprint(ctx)
	log.Message(
		traceId,
		"OUT",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Response.Body()))
}

func (h *ServerHttp) GetAllUser(ctx *fasthttp.RequestCtx) {
	traceId := util.CreateTraceID()
	limit := string(ctx.QueryArgs().Peek("limit"))
	offset := string(ctx.QueryArgs().Peek("offset"))
	log.Message(
		traceId,
		"IN",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Request.Body()))
	httCode, bodyResponse := h.iUserUsecase.GetAll(traceId, limit, offset)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(httCode)
	asByteJSON, _ := json.Marshal(bodyResponse)
	ctx.Response.SetBody(asByteJSON)
	fmt.Fprint(ctx)
	log.Message(
		traceId,
		"OUT",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Response.Body()))
}

func (h *ServerHttp) GetUser(ctx *fasthttp.RequestCtx) {
	traceId := util.CreateTraceID()
	userId := fmt.Sprintf("%v", ctx.UserValue("id"))
	log.Message(
		traceId,
		"IN",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Request.Body()))
	httCode, bodyResponse := h.iUserUsecase.GetUser(traceId, userId)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(httCode)
	asByteJSON, _ := json.Marshal(bodyResponse)
	ctx.Response.SetBody(asByteJSON)
	fmt.Fprint(ctx)
	log.Message(
		traceId,
		"OUT",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Response.Body()))
}

func (h *ServerHttp) CreateUser(ctx *fasthttp.RequestCtx) {
	traceId := util.CreateTraceID()
	reqBody := ctx.Request.Body()
	log.Message(
		traceId,
		"IN",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Request.Body()))
	httCode, bodyResponse := h.iUserUsecase.AddNewUser(traceId, reqBody)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(httCode)
	asByteJSON, _ := json.Marshal(bodyResponse)
	ctx.Response.SetBody(asByteJSON)
	fmt.Fprint(ctx)
	log.Message(
		traceId,
		"OUT",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Response.Body()))
}

func (h *ServerHttp) UpdateUser(ctx *fasthttp.RequestCtx) {
	traceId := util.CreateTraceID()
	reqBody := ctx.Request.Body()
	userId := fmt.Sprintf("%v", ctx.UserValue("id"))
	log.Message(
		traceId,
		"IN",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Request.Body()))
	httCode, bodyResponse := h.iUserUsecase.UpdateUser(traceId, userId, reqBody)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(httCode)
	asByteJSON, _ := json.Marshal(bodyResponse)
	ctx.Response.SetBody(asByteJSON)
	fmt.Fprint(ctx)
	log.Message(
		traceId,
		"OUT",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Response.Body()))
}

func (h *ServerHttp) RemoveUser(ctx *fasthttp.RequestCtx) {
	traceId := util.CreateTraceID()
	userId := fmt.Sprintf("%v", ctx.UserValue("id"))
	log.Message(
		traceId,
		"IN",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Request.Body()))
	httCode, bodyResponse := h.iUserUsecase.DeleteUser(traceId, userId)
	ctx.Response.Header.SetContentType("application/json")
	ctx.Response.SetStatusCode(httCode)
	asByteJSON, _ := json.Marshal(bodyResponse)
	ctx.Response.SetBody(asByteJSON)
	fmt.Fprint(ctx)
	log.Message(
		traceId,
		"OUT",
		"URL",
		string(ctx.Method()),
		string(ctx.Request.RequestURI()),
		string(ctx.Response.Body()))
}
