package controller

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/macduyhai/go-rabbitMQ/logger"
	"github.com/macduyhai/go-rabbitMQ/producer/config"
	"github.com/macduyhai/go-rabbitMQ/producer/dtos"
	"github.com/macduyhai/go-rabbitMQ/producer/models"
	"github.com/macduyhai/go-rabbitMQ/producer/service"
	"github.com/macduyhai/go-rabbitMQ/producer/utils"
)

type Controller struct {
	userService service.RMQProducerService
}

func NewController(config *config.Config) Controller {
	rbmq := service.InitConnectionRB(config.RMQURL, config.UserQueue)
	rmqservice := service.GetRMQProducerService(rbmq)
	return Controller{userService: *rmqservice}
}

func (ctl *Controller) UserBuy(context *gin.Context) {
	var request dtos.BuyRequest

	err := context.ShouldBindJSON(&request)
	if err != nil {
		logger.LogError("Decode json create request error: " + err.Error())
		utils.ResponseError400(context, err.Error())
		return
	}
	// timeNow := utils.TimeIn("Asia/Ho_Chi_Minh")
	usbuy := models.UserBuy{
		Username: request.Username,
		Products: request.Products,
	}
	jsonuser, _ := json.Marshal(usbuy)
	err = ctl.userService.PublishMessage("application/json", jsonuser)
	logger.LogInfor(string(jsonuser))
	if err != nil {
		utils.ResponseError400(context, err.Error())
	} else {
		utils.ResponseSuccess200(context, string(jsonuser), "Buy success")
	}
}
func (ctl *Controller) Ping(ctx *gin.Context) {
	ctx.JSON(200, gin.H{"message": "Pong"})
}
