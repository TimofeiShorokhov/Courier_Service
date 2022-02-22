package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"stlab.itechart-group.com/go/food_delivery/courier_service/dao"
)


// @Summary  CreateDeliveryService
// @Tags DeliveryService
// @Description create a Delivery Service
// @ID CreateDeliveryService
// @Accept  json
// @Produce json
// @Param input body dao.DeliveryService true "Delivery Service"
// @Success 200 {object} dao.DeliveryService
// @Failure 400,404 {string} string
// @Failure 500 {string} string
// @Router /deliveryservice [post]
func (h *Handler) CreateDeliveryService(ctx *gin.Context) {
	var service dao.DeliveryService
	if err := ctx.ShouldBindJSON(&service); err != nil {
		log.Println(err)
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}
	newDelivServ, err:=h.services.CreateDeliveryService(service)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("Error: %s", err)})
		return
	}
	ctx.JSON(http.StatusOK, newDelivServ)
}
