package repository

import "github.com/opensourceways/xihe-statistics/domain"

type WuKongPublic interface {
	Add(*domain.WuKongPublic) error
	Get() ([]domain.WuKongPublic, error)
}
