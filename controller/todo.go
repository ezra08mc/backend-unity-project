package controller

import (
	"net/http"
	"strconv"

	"github.com/ezra08mc/backend-unity-project/config/middleware"
	"github.com/ezra08mc/backend-unity-project/config/pkg/errs"
	"github.com/ezra08mc/backend-unity-project/contract"
	"github.com/ezra08mc/backend-unity-project/dto"
	"github.com/gin-gonic/gin"
)

type TodoController struct {
	service contract.TodoService
}

func (t *TodoController) GetPrefix() string {
	return "/api/todos"
}

func (t *TodoController) InitService(service *contract.Service) {
	t.service = service.Todo
}

func (t *TodoController) InitRoute(app *gin.RouterGroup) {

	// User routes (Scoped access)
	app.POST("/", middleware.Auth(), t.create)
	app.GET("/", middleware.Auth(), t.getActive)
	app.GET("/trash", middleware.Auth(), t.getTrash)
	app.GET("/:id", middleware.Auth(), t.getByID)
	app.PUT("/:id", middleware.Auth(), t.update)
	app.DELETE("/:id", middleware.Auth(), t.softDelete)
	app.PATCH("/:id/restore", middleware.Auth(), t.restore)
	app.DELETE("/:id/permanent", middleware.Auth(), t.permanentDelete)


	// Admin routes (Global access)
	admin := app.Group("/admin", middleware.Auth())
	admin.GET("/", t.adminGetActive)
	admin.GET("/trash", t.adminGetTrash)
	admin.GET("/:id", t.adminGetByID)
	admin.PUT("/:id", t.adminUpdate)
	admin.PATCH("/:id/restore", t.adminRestore)
	admin.DELETE("/:id/permanent", t.adminPermanentDelete)
}


// User handlers (Scoped access)

// Create godoc
// @Summary      Create new todo
// @Description  Create a new todo item for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        payload  body      dto.TodoRequest  true  "Todo data"
// @Success      201      {object}  dto.TodoResponse
// @Failure      400      {object}  errs.MessageErrData
// @Router       /api/todos [post]
func (t *TodoController) create(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	var payload dto.TodoRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("Invalid request payload"))
		return
	}

	response, err := t.service.CreateTodo(userID, payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusCreated, response)
}

// GetActive godoc
// @Summary      Get all active todos
// @Description  Get all active (not deleted) todos for the authenticated user with pagination
// @Tags         Todo User
// @Security     BearerAuth
// @Produce      json
// @Param        limit   query     int  false  "Limit"
// @Param        offset  query     int  false  "Offset"
// @Success      200     {array}   dto.TodoResponse
// @Router       /api/todos [get]
func (t *TodoController) getActive(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	limit, offset := getPagination(ctx)
	response, err := t.service.GetActiveByUserID(userID, limit, offset)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetTrash godoc
// @Summary      Get trashed todos
// @Description  Get all soft-deleted todos for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Produce      json
// @Param        limit   query     int  false  "Limit"
// @Param        offset  query     int  false  "Offset"
// @Success      200     {array}   dto.TodoResponse
// @Router       /api/todos/trash [get]
func (t *TodoController) getTrash(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	limit, offset := getPagination(ctx)
	response, err := t.service.GetTrashByUserID(userID, limit, offset)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// GetByID godoc
// @Summary      Get todo detail
// @Description  Get specific todo details for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  dto.TodoResponse
// @Router       /api/todos/{id} [get]
func (t *TodoController) getByID(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	response, err := t.service.GetByIDAndUserID(id, userID)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// Update godoc
// @Summary      Update todo
// @Description  Update specific active todo for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int              true  "Todo ID"
// @Param        payload  body      dto.TodoRequest  true  "Todo data"
// @Success      200      {object}  dto.TodoResponse
// @Router       /api/todos/{id} [put]
func (t *TodoController) update(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	var payload dto.TodoRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("Invalid request payload"))
		return
	}

	response, err := t.service.UpdateByUserID(id, userID, payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// SoftDelete godoc
// @Summary      Soft delete todo
// @Description  Move a todo to trash (soft delete) for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/todos/{id} [delete]
func (t *TodoController) softDelete(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := t.service.SoftDeleteByUserID(id, userID); err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Todo moved to trash"})
}

// Restore godoc
// @Summary      Restore trashed todo
// @Description  Restore a soft-deleted todo back to active for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/todos/{id}/restore [patch]
func (t *TodoController) restore(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := t.service.RestoreByUserID(id, userID); err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Todo restored successfully"})
}

// PermanentDelete godoc
// @Summary      Permanent delete todo
// @Description  Force delete a todo permanently for the authenticated user
// @Tags         Todo User
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/todos/{id}/permanent [delete]
func (t *TodoController) permanentDelete(ctx *gin.Context) {
	userID, err := getUserID(ctx)
	if err != nil {
		HandlerError(ctx, err)
		return
	}

	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := t.service.PermanentDeleteByUserID(id, userID); err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Todo permanently deleted"})
}


// Admin handlers (Scoped access)

// AdminGetActive godoc
// @Summary      [Admin] Get all users' active todos
// @Description  Get all active todos globally (Admin only)
// @Tags         Todo Admin
// @Security     BearerAuth
// @Produce      json
// @Param        limit   query     int  false  "Limit"
// @Param        offset  query     int  false  "Offset"
// @Success      200     {array}   dto.TodoResponse
// @Router       /api/todos/admin [get]
func (t *TodoController) adminGetActive(ctx *gin.Context) {
	limit, offset := getPagination(ctx)
	response, err := t.service.GetAllActive(limit, offset)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// AdminGetTrash godoc
// @Summary      [Admin] Get all users' trashed todos
// @Description  Get all soft-deleted todos globally (Admin only)
// @Tags         Todo Admin
// @Security     BearerAuth
// @Produce      json
// @Param        limit   query     int  false  "Limit"
// @Param        offset  query     int  false  "Offset"
// @Success      200     {array}   dto.TodoResponse
// @Router       /api/todos/admin/trash [get]
func (t *TodoController) adminGetTrash(ctx *gin.Context) {
	limit, offset := getPagination(ctx)
	response, err := t.service.GetAllTrash(limit, offset)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// AdminGetByID godoc
// @Summary      [Admin] Get any todo detail
// @Description  Get any specific todo details globally (Admin only)
// @Tags         Todo Admin
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  dto.TodoResponse
// @Router       /api/todos/admin/{id} [get]
func (t *TodoController) adminGetByID(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	response, err := t.service.GetByID(id)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// AdminUpdate godoc
// @Summary      [Admin] Update any todo
// @Description  Update any user's todo (Admin only)
// @Tags         Todo Admin
// @Security     BearerAuth
// @Accept       json
// @Produce      json
// @Param        id       path      int              true  "Todo ID"
// @Param        payload  body      dto.TodoRequest  true  "Todo data"
// @Success      200      {object}  dto.TodoResponse
// @Router       /api/todos/admin/{id} [put]
func (t *TodoController) adminUpdate(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	var payload dto.TodoRequest
	if err := ctx.ShouldBindJSON(&payload); err != nil {
		HandlerError(ctx, errs.BadRequest("Invalid request payload"))
		return
	}

	response, err := t.service.Update(id, payload)
	if err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, response)
}

// AdminRestore godoc
// @Summary      [Admin] Restore any trashed todo
// @Description  Restore any soft-deleted todo globally (Admin only)
// @Tags         Todo Admin
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/todos/admin/{id}/restore [patch]
func (t *TodoController) adminRestore(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := t.service.Restore(id); err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Todo restored successfully by Admin"})
}

// AdminPermanentDelete godoc
// @Summary      [Admin] Permanent delete any todo
// @Description  Force delete any todo permanently (Admin only)
// @Tags         Todo Admin
// @Security     BearerAuth
// @Produce      json
// @Param        id   path      int  true  "Todo ID"
// @Success      200  {object}  map[string]interface{}
// @Router       /api/todos/admin/{id}/permanent [delete]
func (t *TodoController) adminPermanentDelete(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	if err := t.service.PermanentDelete(id); err != nil {
		HandlerError(ctx, err)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"success": true, "message": "Todo permanently deleted by Admin"})
}