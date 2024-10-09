package repository

import (
	"github.com/opensourceways/xihe-statistics/domain"
)

type RepoRecords struct {
	Users  []string
	Counts int
}

type UserWithRepo interface {
	Add(*domain.UserWithRepo) error
	Get() (RepoRecords, error)
}
