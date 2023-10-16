package api

import (
	"context"
	"database/sql"
	"github.com/gin-gonic/gin"
	"github.com/linqcod/kanban-backend/cmd/api/middleware"
	"github.com/linqcod/kanban-backend/internal/handler"
	"github.com/linqcod/kanban-backend/internal/repository"
	"go.uber.org/zap"
)

func InitRouter(ctx context.Context, logger *zap.SugaredLogger, db *sql.DB) *gin.Engine {
	router := gin.Default()
	router.Use(middleware.CORSMiddleware())

	// init services, repos, handlers
	productRepo := repository.NewProductRepository(ctx, db)
	productHandler := handler.NewProductHandler(logger, productRepo)

	api := router.Group("/api/v1")
	{
		products := api.Group("/products")
		{
			products.GET("", productHandler.GetAllProducts)
		}
	}

	return router
}
