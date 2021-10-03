package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/minhphong306/mindX/db/sqlc"
	"github.com/spf13/cast"
	"net/http"
)

type createUserRequest struct {
	Name             string `json:"name"  binding:"required"`
	PermanentAddress string `json:"permanent_address"  binding:"required"`
	CurrentAddress   string `json:"current_address" binding:"required"`
	CurrentStatus    int32  `json:"current_status"  binding:"required"`
}

type userResponse struct {
	ID               int64  `json:"id"`
	Name             string `json:"name"`
	PermanentAddress string `json:"permanent_address"`
	CurrentAddress   string `json:"current_address"`
	CurrentStatus    int32  `json:"current_status"`
}

func newUserResponse(user db.User) userResponse {
	return userResponse{
		ID:               user.ID,
		Name:             cast.ToString(user.Name),
		PermanentAddress: cast.ToString(user.PermanentAddress),
		CurrentAddress:   cast.ToString(user.CurrentAddress),
		CurrentStatus:    cast.ToInt32(user.CurrentStatus),
	}
}

func (server *Server) createUser(ctx *gin.Context) {
	var req createUserRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateUserParams{
		Name:             req.Name,
		PermanentAddress: req.PermanentAddress,
		CurrentAddress:   req.CurrentAddress,
		CurrentStatus:    req.CurrentStatus,
	}

	user, err := server.store.CreateUser(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	rsp := newUserResponse(user)
	ctx.JSON(http.StatusOK, rsp)
}
