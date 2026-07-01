package controller

import (
	"errors"
	"log"
	"net/http"
	"strconv"

	"github.com/ezra08mc/backend-unity-project/config/middleware"
	"github.com/ezra08mc/backend-unity-project/config/pkg/errs"
	"github.com/ezra08mc/backend-unity-project/contract"
	"github.com/gin-gonic/gin"
)

type Controller interface {
	GetPrefix() string
	InitService(service *contract.Service)
	InitRoute(app *gin.RouterGroup)
}

func New(app *gin.Engine, service *contract.Service) {
	allController := []Controller{
		&AuthController{},
		&TodoController{},
	}

	for _, c := range allController {
		c.InitService(service)
		group := app.Group(c.GetPrefix())
		group.Use(middleware.CORSMiddleware())
		c.InitRoute(group)
		log.Printf("initiate route %s\n", c.GetPrefix())
	}
}

func HandlerError(ctx *gin.Context, err error) {
	var messageErr errs.MessageError
	if errors.As(err, &messageErr) {
		ctx.JSON(messageErr.Status(), messageErr)
		return
	}
	_ = ctx.Error(err).SetType(gin.ErrorTypePrivate)
	ctx.JSON(http.StatusInternalServerError, errs.InternalServerError("Internal Server Error"))
}

func getUserID(ctx *gin.Context) (int, error){
	rawID, exists := ctx.Get("user_id")
	if !exists {
		HandlerError(ctx, errs.Unauthorized("Unauthorized"))
	}
	return rawID.(int), nil		
}

func getPagination(ctx *gin.Context) (int, int){
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	return limit, offset
}