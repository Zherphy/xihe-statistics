package repository

import (
	"github.com/opensourceways/xihe-statistics/domain"
)

type RegisterCounts struct {
	Counts int64
}

type RegisterRecord interface {
	Add(*domain.RegisterRecord) error
	Get() (RegisterCounts, error)
}
