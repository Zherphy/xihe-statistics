package pgsql

import (
	"context"

	"github.com/opensourceways/xihe-statistics/infrastructure/repositories"
)

const fieldUserName = "username"

func NewWuKongPublicMapper(table WuKongPublic) repositories.WuKongPublicMapper {
	return &wukongPublic{table}
}

type wukongPublic struct {
	table WuKongPublic
}

func (impl *wukongPublic) Get() (dos []repositories.WuKongPublicRecordsDO, err error) {
	var records []interface{}

	f := func(ctx context.Context) error {
		err := cli.distinct(
			ctx, impl.table,
			fieldUserName, &records,
		)

		return err
	}

	if err = withContext(f); err != nil {
		return
	}

	users, err := toArryString(records)
	if err != nil {
		return
	}

	dos = make([]repositories.WuKongPublicRecordsDO, len(users))
	for i := range users {
		dos[i].UserName = users[i]
	}

	return
}
