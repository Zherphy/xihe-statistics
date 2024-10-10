package app

import (
	"github.com/opensourceways/xihe-statistics/domain/repository"
)

type D1Service interface {
	Get() (D1DTO, error)
}

func NewD1Service(
	ub repository.UserWithBigModel,
	ur repository.UserWithRepo,
) D1Service {
	return d1Service{
		bms: bigModelRecordService{
			ub: ub,
		},
		urs: repoRecordService{
			ur: ur,
		},
	}
}

type d1Service struct {
	bms bigModelRecordService
	urs repoRecordService
}

func (s d1Service) Get() (dto D1DTO, err error) {
	bdto, err := s.bms.GetBigModelRecordAll()
	if err != nil {
		return
	}
	bigModelUsers := bdto.Users

	rdto, err := s.urs.Get()
	if err != nil {
		return
	}
	repoUsers := rdto.Users

	duCounts := len(bigModelUsers) + len(repoUsers)

	users := append(bigModelUsers, repoUsers...)
	users = RemoveRepeatedElement(users)

	dto = D1DTO{
		Counts:          len(users),
		DuplicateCounts: duCounts,
		Users:           users,
		UpdateAt:        getLocalTime(),
	}

	return
}

type D1DTO struct {
	Counts          int      `json:"counts"`
	DuplicateCounts int      `json:"duplicate_counts"`
	Users           []string `json:"users"`
	UpdateAt        string   `json:"update_at"`
}
