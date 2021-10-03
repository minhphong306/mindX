package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/minhphong306/mindX/db/sqlc"
	"net/http"
)

type createLocationHistoryRequest struct {
	UserId      int64  `json:"user_id"  binding:"required"`
	Type        int32  `json:"type"  binding:"required"`
	LocationId  int32  `json:"location_id"`
	ManualInput string `json:"manual_input"  binding:"required"`
}

func (server *Server) createLocationHistory(ctx *gin.Context) {
	var req createLocationHistoryRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.CreateLocationHistoryParams{
		UserID:      req.UserId,
		Type:        req.Type,
		LocationID:  req.LocationId,
		ManualInput: req.ManualInput,
	}

	history, err := server.store.CreateLocationHistory(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, history)
}
