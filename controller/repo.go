package controller

import (
	"net/http"

	"github.com/opensourceways/xihe-statistics/app"
	"github.com/opensourceways/xihe-statistics/domain/repository"

	"github.com/gin-gonic/gin"
)

func AddRouterForRepoRecordController(
	rg *gin.RouterGroup,
	ur repository.UserWithRepo,
) {
	ctl := RepoRecordController{
		rs: app.NewRepoRecordService(ur),
	}

	rg.POST("/v1/d1/repo", ctl.AddRepoRecord)
	rg.GET("/v1/d1/repo", ctl.GetRepoRecord)

}

type RepoRecordController struct {
	baseController
	rs app.RepoRecordService
}

// @Summary Add
// @Description add user query repo record
// @Tags  D1
// @Param  body  body  AddRepoRecordRequest  true  "body of repo records"
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1/repo [post]
func (ctl *RepoRecordController) AddRepoRecord(ctx *gin.Context) {
	req := AddRepoRecordRequest{}
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, respBadRequestBody)
	}

	cmd, err := req.toCmd()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, newResponseCodeError(
			errorBadRequestParam, err,
		))
		return
	}

	err = ctl.rs.Add(&cmd)
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData("success"))
}

// @Summary Get
// @Description get user create repository record
// @Tags  D1
// @Accept json
// @Success 200 {object}
// @Produce json
// @Router /v1/d1/repo [get]
func (ctl *RepoRecordController) GetRepoRecord(ctx *gin.Context) {

	rrd, err := ctl.rs.Get()
	if err != nil {
		ctl.sendRespWithInternalError(ctx, newResponseError(err))

		return
	}

	ctx.JSON(http.StatusOK, newResponseData(rrd))
}
