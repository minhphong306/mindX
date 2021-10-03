package api

import (
	"github.com/gin-gonic/gin"
	db "github.com/minhphong306/mindX/db/sqlc"
	"github.com/spf13/cast"
	"net/http"
)

func (server *Server) getListLocation(ctx *gin.Context) {
	q := ctx.Request.URL.Query()

	limit := cast.ToInt32(q.Get("limit"))
	if limit <= 0 {
		limit = 100
	}

	page := cast.ToInt32(q.Get("page"))
	if page <= 0 {
		page = 1
	}

	req := db.ListLocationHistoriesParams{
		Limit:  limit,
		Offset: (page - 1) * limit,
	}

	users, err := server.store.ListLocationHistories(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, users)
}
