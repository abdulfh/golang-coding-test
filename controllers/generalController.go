package controllers

import (
	"idn/services"
	"idn/utils/constants"
	"idn/utils/response"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
)

type generalController struct {
	services services.GeneralServices
}

func InitGeneralController(services services.GeneralServices) *generalController {
	return &generalController{services}
}

func (controller *generalController) FizzBuzz(context *gin.Context) {
	response := response.ResponseApi{}

	result, err := controller.services.FizzBuzz()
	if err != nil {
		response.ResponseCode = constants.ERROR_RC511
		response.ResponseDesc = constants.ERROR_RM511
		response.Data = result
		context.JSON(http.StatusOK, response)
		return
	}

	response.ResponseCode = constants.ERROR_RC200
	response.ResponseDesc = constants.ERROR_RM200
	response.Data = result
	context.JSON(http.StatusOK, response)
}

func (controller *generalController) Multiple(context *gin.Context) {
	response := response.ResponseApi{}

	result, err := controller.services.Multiple()
	if err != nil {
		response.ResponseCode = constants.ERROR_RC511
		response.ResponseDesc = constants.ERROR_RM511
		response.Data = result
		context.JSON(http.StatusOK, response)
		return
	}

	response.ResponseCode = constants.ERROR_RC200
	response.ResponseDesc = constants.ERROR_RM200
	response.Data = result
	context.JSON(http.StatusOK, response)
}

func (controller *generalController) MarkPaid(context *gin.Context) {
	response := response.ResponseApi{}
	billAmount := context.Param("bill")

	result, err := controller.services.MarkPaid(billAmount)
	if err != nil {
		glog.Error(err.Error())
		response.ResponseCode = constants.ERROR_RC410
		response.ResponseDesc = constants.ERROR_RM410
		response.Data = result
		context.JSON(http.StatusOK, response)
		return
	}

	response.ResponseCode = constants.ERROR_RC200
	response.ResponseDesc = constants.ERROR_RM200
	response.Data = result
	context.JSON(http.StatusOK, response)
}
