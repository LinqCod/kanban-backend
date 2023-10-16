package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/linqcod/kanban-backend/internal/common/errorTypes"
	"github.com/linqcod/kanban-backend/internal/handler/dto"
	"github.com/linqcod/kanban-backend/internal/model"
	"go.uber.org/zap"
	"net/http"
)

type ProductRepository interface {
	GetAllProducts() ([]*model.Product, error)
}

type ProductHandler struct {
	logger *zap.SugaredLogger
	repo   ProductRepository
}

func NewProductHandler(logger *zap.SugaredLogger, repo ProductRepository) *ProductHandler {
	return &ProductHandler{
		logger: logger,
		repo:   repo,
	}
}

func (h ProductHandler) GetAllProducts(c *gin.Context) {
	products, err := h.repo.GetAllProducts()
	if err != nil {
		h.logger.Errorf("error while getting products: %v", err)
		c.JSON(http.StatusInternalServerError, dto.ErrorDTO{
			Error: errorTypes.ErrDBDataReception.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, products)
}
