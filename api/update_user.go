package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/minhphong306/mindX/db/sqlc"
	"net/http"
)

type updateUserRequest struct {
	Id               int64  `json:"id" binding:"required"`
	Name             string `json:"name"  binding:"required"`
	PermanentAddress string `json:"permanent_address"  binding:"required"`
	CurrentAddress   string `json:"current_address" binding:"required"`
	CurrentStatus    int32  `json:"current_status"  binding:"required"`
}

func (server *Server) updateUser(ctx *gin.Context) {
	var req updateUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.UpdateUserParams{
		ID:               req.Id,
		Name:             req.Name,
		PermanentAddress: req.PermanentAddress,
		CurrentAddress:   req.CurrentAddress,
		CurrentStatus:    req.CurrentStatus,
	}

	user, err := server.store.UpdateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
