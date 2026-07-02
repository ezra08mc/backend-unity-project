package controller

import (
    "net/http"
    "strconv"

    "github.com/ezra08mc/backend-unity-project/config/middleware"
    "github.com/ezra08mc/backend-unity-project/contract"
    "github.com/ezra08mc/backend-unity-project/dto"
    "github.com/gin-gonic/gin"
)

type TodoController struct {
    service contract.TodoService
}

func (t *TodoController) GetPrefix() string { return "/api/todos" }

func (t *TodoController) InitService(service *contract.Service) {
    t.service = service.Todo
}

func (t *TodoController) InitRoute(app *gin.RouterGroup) {
    app.POST("/", middleware.Auth(), t.Create)
    app.GET("/", middleware.Auth(), t.GetActive)
    app.GET("/trash", middleware.Auth(), t.GetTrash)
    app.GET("/:id", middleware.Auth(), t.GetByID)
    app.PUT("/:id", middleware.Auth(), t.Update)
    app.DELETE("/:id", middleware.Auth(), t.SoftDelete)
    app.PATCH("/:id/restore", middleware.Auth(), t.Restore)
    app.DELETE("/:id/permanent", middleware.Auth(), t.PermanentDelete)

    admin := app.Group("/admin", middleware.Auth())
    admin.GET("/", t.AdminGetActive)
    admin.GET("/trash", t.AdminGetTrash)
    admin.GET("/:id", t.AdminGetByID)
    admin.PUT("/:id", t.AdminUpdate)
    admin.PATCH("/:id/restore", t.AdminRestore)
    admin.DELETE("/:id/permanent", t.AdminPermanentDelete)
}

// @Summary Create todo
// @Tags Todo User
// @Security BearerAuth
// @Accept json
// @Produce json
// @Param payload body dto.TodoRequest true "Todo data"
// @Success 201 {object} dto.TodoResponse
// @Failure 400 {object} errs.ErrorData
// @Router /api/todos [post]
func (t *TodoController) Create(ctx *gin.Context) {
    userID, err := getUserID(ctx)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    var p dto.TodoRequest
    if err := ctx.ShouldBindJSON(&p); err != nil {
        ctx.JSON(http.StatusBadRequest, gin.H{
            "message": "Validasi gagal",
            "errors":  gin.H{"title": "Title wajib diisi dan minimal 3 karakter"},
        })
        return
    }
    resp, err := t.service.CreateTodo(userID, p)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusCreated, resp)
}

// @Summary Get active todos
// @Tags Todo User
// @Security BearerAuth
// @Success 200 {object} dto.TodoListResponse
// @Router /api/todos [get]
func (t *TodoController) GetActive(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    l, o := getPagination(ctx)
    resp, err := t.service.GetActiveByUserID(userID, l, o)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary Get trash
// @Tags Todo User
// @Security BearerAuth
// @Success 200 {object} dto.TodoListResponse
// @Router /api/todos/trash [get]
func (t *TodoController) GetTrash(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    l, o := getPagination(ctx)
    resp, err := t.service.GetTrashByUserID(userID, l, o)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary Get todo detail
// @Tags Todo User
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} dto.TodoResponse
// @Router /api/todos/{id} [get]
func (t *TodoController) GetByID(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    id, _ := strconv.Atoi(ctx.Param("id"))
    resp, err := t.service.GetByIDAndUserID(id, userID)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary Update todo
// @Tags Todo User
// @Security BearerAuth
// @Param id path int true "ID"
// @Param payload body dto.TodoRequest true "Data"
// @Success 200 {object} dto.TodoResponse
// @Router /api/todos/{id} [put]
func (t *TodoController) Update(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    id, _ := strconv.Atoi(ctx.Param("id"))
    var p dto.TodoRequest
    ctx.ShouldBindJSON(&p)
    resp, err := t.service.UpdateByUserID(id, userID, p)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary Soft delete
// @Tags Todo User
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/todos/{id} [delete]
func (t *TodoController) SoftDelete(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := t.service.SoftDeleteByUserID(id, userID); err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// @Summary Restore todo
// @Tags Todo User
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/todos/{id}/restore [patch]
func (t *TodoController) Restore(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := t.service.RestoreByUserID(id, userID); err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// @Summary Permanent delete
// @Tags Todo User
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/todos/{id}/permanent [delete]
func (t *TodoController) PermanentDelete(ctx *gin.Context) {
    userID, _ := getUserID(ctx)
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := t.service.PermanentDeleteByUserID(id, userID); err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// @Summary [Admin] Get all active
// @Tags Todo Admin
// @Security BearerAuth
// @Success 200 {object} dto.TodoListResponse
// @Router /api/todos/admin [get]
func (t *TodoController) AdminGetActive(ctx *gin.Context) {
    l, o := getPagination(ctx)
    resp, err := t.service.GetAllActive(l, o)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary [Admin] Get all trash
// @Tags Todo Admin
// @Security BearerAuth
// @Success 200 {object} dto.TodoListResponse
// @Router /api/todos/admin/trash [get]
func (t *TodoController) AdminGetTrash(ctx *gin.Context) {
    l, o := getPagination(ctx)
    resp, err := t.service.GetAllTrash(l, o)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary [Admin] Get detail
// @Tags Todo Admin
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} dto.TodoResponse
// @Router /api/todos/admin/{id} [get]
func (t *TodoController) AdminGetByID(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    resp, err := t.service.GetByID(id)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary [Admin] Update any
// @Tags Todo Admin
// @Security BearerAuth
// @Param id path int true "ID"
// @Param payload body dto.TodoRequest true "Data"
// @Success 200 {object} dto.TodoResponse
// @Router /api/todos/admin/{id} [put]
func (t *TodoController) AdminUpdate(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    var p dto.TodoRequest
    ctx.ShouldBindJSON(&p)
    resp, err := t.service.Update(id, p)
    if err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, resp)
}

// @Summary [Admin] Restore any
// @Tags Todo Admin
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/todos/admin/{id}/restore [patch]
func (t *TodoController) AdminRestore(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := t.service.Restore(id); err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"success": true})
}

// @Summary [Admin] Permanent delete any
// @Tags Todo Admin
// @Security BearerAuth
// @Param id path int true "ID"
// @Success 200 {object} map[string]interface{}
// @Router /api/todos/admin/{id}/permanent [delete]
func (t *TodoController) AdminPermanentDelete(ctx *gin.Context) {
    id, _ := strconv.Atoi(ctx.Param("id"))
    if err := t.service.PermanentDelete(id); err != nil {
        HandlerError(ctx, err)
        return
    }
    ctx.JSON(http.StatusOK, gin.H{"success": true})
}