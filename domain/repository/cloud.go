package repository

import "github.com/opensourceways/xihe-statistics/domain"

type CloudUsers struct {
	Users []string
}

type CloudRecord interface {
	Add(*domain.Cloud) error
	Get(domain.CloudType) (int64, error)
	GetUsers() (CloudUsers, error)
}
