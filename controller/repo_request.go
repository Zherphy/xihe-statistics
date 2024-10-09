package controller

import (
	"github.com/opensourceways/xihe-statistics/app"
	"github.com/opensourceways/xihe-statistics/domain"
)

type AddRepoRecordRequest struct {
	UserName string `json:"username"`
	RepoName string `json:"repo_name"`
}

func (req *AddRepoRecordRequest) toCmd() (cmd app.RepoRecordAddCmd, err error) {
	username, err := domain.NewAccount(req.UserName)
	if err != nil {
		return
	}

	cmd.UserName = username
	cmd.RepoName = req.RepoName

	return
}
