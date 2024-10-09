package controller

import (
	"github.com/opensourceways/xihe-statistics/app"
	"github.com/opensourceways/xihe-statistics/domain"
)

type MediaRequest struct {
	Name string `json:"name"`
}

func (req MediaRequest) toCmd() (cmd app.AddMediaCmd, err error) {
	if cmd.Name, err = domain.NewMeidaName(req.Name); err != nil {
		return
	}

	cmd.CreateAt = app.GetUnixLocalTime()

	return
}
