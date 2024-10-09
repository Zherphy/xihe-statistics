package controller

import (
	"github.com/opensourceways/xihe-statistics/app"
	"github.com/opensourceways/xihe-statistics/domain"
)

type BigModelQueryWithTypeAndTimeRequest struct {
	BigModel  string `json:"bigmodel"`
	StartTime string `json:"start_time"`
	EndTime   string `json:"end_time"`
}

func (req BigModelQueryWithTypeAndTimeRequest) toCmd() (
	cmd app.BigModelCountIncreaseCmd,
	err error,
) {
	if cmd.BigModel, err = domain.NewBigModel(req.BigModel); err != nil {
		return
	}

	cmd.StartTime = req.StartTime
	cmd.EndTime = req.EndTime

	return
}
