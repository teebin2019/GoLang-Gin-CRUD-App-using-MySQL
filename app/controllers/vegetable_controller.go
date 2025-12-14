package controllers

import (
	"net/http"

	"GoCRUDApplicationMySQL/app/models"
	"GoCRUDApplicationMySQL/app/repositories"

	"github.com/gin-gonic/gin"
)

type VegetableController struct {
	repo repositories.VegetableRepository
}

func NewVegetableController(repo repositories.VegetableRepository) *VegetableController {
	return &VegetableController{
		repo: repo,
	}
}

func (c *VegetableController) GetAllVegetables(ctx *gin.Context) {
	vegetables, err := c.repo.GetAllVegetables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch vegetables"})
		return
	}
	ctx.JSON(http.StatusOK, vegetables)
}

func (c *VegetableController) CreateVegetable(ctx *gin.Context) {
	var vegetable models.Vegetable
	if err := ctx.ShouldBindJSON(&vegetable); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := c.repo.CreateVegetable(&vegetable); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create vegetable"})
		return
	}

	ctx.JSON(http.StatusOK, vegetable)
}

// Implement other CRUD operations like GetVegetable, UpdateVegetable, and DeleteVegetable
