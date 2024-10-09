package app

import (
	"errors"

	"github.com/opensourceways/xihe-statistics/domain"
)

type CloudRecordCmd struct {
	User     domain.Account
	CloudId  string
	CreateAt int64
}

func (cmd CloudRecordCmd) Validate() error {
	b := cmd.User == nil ||
		cmd.CloudId == "" ||
		cmd.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add cloud record")
	}

	return nil
}

type CloudRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}

func (dto *CloudRecordDTO) toCloudRecordDTO(counts int64, update string) {
	*dto = CloudRecordDTO{
		Counts:   counts,
		UpdateAt: update,
	}
}

type CountCloudCmd struct {
	CloudType domain.CloudType
}

func ToCountCloudCmd(cloudType string) (*CountCloudCmd, error) {
	t, err := domain.NewCloudType(cloudType)
	if err != nil {
		return nil, err
	}

	return &CountCloudCmd{
		CloudType: t,
	}, nil
}
