package repository

import "project/xihe-statistics/domain"

type CloudUsers struct {
	Users []string
}

type CloudRecord interface {
	Add(*domain.Cloud) error
	Get() (int64, error)
	GetUsers() (CloudUsers, error)
}
