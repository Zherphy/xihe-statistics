package controller

import (
	"net/http"

	"github.com/opensourceways/xihe-statistics/app"
	"github.com/opensourceways/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForCloudRecordController(
	rg *gin.RouterGroup,
	cr repository.CloudRecord,
) {
	ctl := CloudRecordController{
		cs: app.NewCloudRecodeService(cr),
	}

	rg.GET("/v1/cloud/:type", ctl.Get)
}

type CloudRecordController struct {
	baseController

	cs app.CloudRecordService
}

// @Summary Get
// @Description get cloud record
// @Tags  cloud
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/cloud/{type} [get]
func (ctl *CloudRecordController) Get(ctx *gin.Context) {
	cmd, err := app.ToCountCloudCmd(ctx.Param("type"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponseCodeError(errorBadRequestParam, err))

		return
	}

	dto, err := ctl.cs.Get(cmd)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
