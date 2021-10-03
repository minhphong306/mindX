package api

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/minhphong306/mindX/db/sqlc"
	"github.com/minhphong306/mindX/token"
	"github.com/minhphong306/mindX/util"
)

// Server serves HTTP requests for our mindX service.
type Server struct {
	config     util.Config
	store      db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and set up routing.
func NewServer(config util.Config, store db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	server.setupRouter()
	return server, nil
}

func (server *Server) setupRouter() {
	router := gin.Default()

	authRoutes := router.Group("/admin").Use(authMiddleware(server.tokenMaker))
	authRoutes.GET("/get-users", server.getListUser)
	authRoutes.POST("/create-user", server.createUser)
	authRoutes.PUT("/update-user", server.updateUser)

	//authRoutes.PUT("/update-user", server.createUser)
	//authRoutes.PUT("/get-user", server.createUser)
	//authRoutes.PUT("/delete-user", server.createUser)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
