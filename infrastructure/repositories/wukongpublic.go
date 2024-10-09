package repositories

import (
	"github.com/opensourceways/xihe-statistics/domain"
	"github.com/opensourceways/xihe-statistics/domain/repository"
)

type WuKongPublicMapper interface {
	Get() ([]WuKongPublicRecordsDO, error)
}

func NewWuKongPublicRepository(mapper WuKongPublicMapper) repository.WuKongPublic {
	return &wukongPublic{
		mapper: mapper,
	}
}

type wukongPublic struct {
	mapper WuKongPublicMapper
}

func (impl *wukongPublic) Add(wp *domain.WuKongPublic) error {
	return nil
}

func (impl *wukongPublic) Get() (wps []domain.WuKongPublic, err error) {
	dos, err := impl.mapper.Get()
	if err != nil {
		return
	}

	wps = make([]domain.WuKongPublic, len(dos))
	for i := range dos {
		if wps[i], err = dos[i].toWuKongPublic(); err != nil {
			return
		}
	}

	return
}

func (r WuKongPublicRecordsDO) toWuKongPublic() (wp domain.WuKongPublic, err error) {
	if wp.UserName, err = domain.NewAccount(r.UserName); err != nil {
		return
	}

	wp.CreateAt = r.CreateAt

	return
}

type WuKongPublicRecordsDO struct {
	UserName string
	CreateAt int64
}
