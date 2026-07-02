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

    rootGroup := app.Group("/")
    apiGroup := app.Group("/api")

    for _, c := range allController {
        c.InitService(service)
        
        var targetGroup *gin.RouterGroup
        if len(c.GetPrefix()) >= 4 && c.GetPrefix()[:4] == "/api" {
            subPath := c.GetPrefix()[4:] 
            targetGroup = apiGroup.Group(subPath)
        } else {
            targetGroup = rootGroup.Group(c.GetPrefix())
        }

        targetGroup.Use(middleware.CORSMiddleware())
        c.InitRoute(targetGroup)
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

func getUserID(ctx *gin.Context) (int, error) {
    rawID, exists := ctx.Get("user_id")
    if !exists {
        err := errs.Unauthorized("Unauthorized")
        HandlerError(ctx, err)
        return 0, err 
    }
    return rawID.(int), nil     
}

func getPagination(ctx *gin.Context) (int, int){
	limit, _ := strconv.Atoi(ctx.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(ctx.DefaultQuery("offset", "0"))
	return limit, offset
}