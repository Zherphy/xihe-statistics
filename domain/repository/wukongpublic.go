package repository

import "project/xihe-statistics/domain"

type WuKongPublic interface {
	Add(*domain.WuKongPublic) error
	Get() ([]domain.WuKongPublic, error)
}