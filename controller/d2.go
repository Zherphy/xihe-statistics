package controller

import (
	"net/http"

	"github.com/opensourceways/xihe-statistics/app"
	"github.com/opensourceways/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForD2Controller(
	rg *gin.RouterGroup,
	fr repository.FileUploadRecord,
	wp repository.WuKongPublic,
	ur repository.UserWithRepo,
	cr repository.CloudRecord,
) {
	ctl := FileUploadRecordController{
		fs: app.NewD2Service(fr, wp, ur, cr),
	}

	rg.GET("/v1/d2", ctl.GetFileUploadRecord)

}

type FileUploadRecordController struct {
	baseController
	fs app.D2Service
}

// @Summary Get
// @Description get d2
// @Tags  D2
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d2 [get]
func (ctl *FileUploadRecordController) GetFileUploadRecord(ctx *gin.Context) {
	dto, err := ctl.fs.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(dto))
}
