package pgsql

import (
	"context"
	"project/xihe-statistics/infrastructure/repositories"
)

func NewCloudRecordMapper(table CloudRecord) repositories.CloudRecordMapper {
	return &cloudRecord{table}
}

type cloudRecord struct {
	table CloudRecord
}

func (m *cloudRecord) AddCloudRecord(do *repositories.CloudRecordDO) (err error) {
	col := toCloudRecordCol(do)

	f := func(ctx context.Context) error {
		return cli.create(
			ctx, m.table,
			col,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m *cloudRecord) GetCloudRecordCount(cloudId string) (counts int64, err error) {
	f := func(ctx context.Context) error {
		return cli.whereDistinctCount(
			ctx, m.table,
			"1 = ? AND cloud_id = ?",
			1, cloudId,
			"username", &counts,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	return
}

func (m *cloudRecord) GetUsers() (
	do repositories.CloudUsersDO,
	err error,
) {

	var users []interface{}

	f := func(ctx context.Context) error {
		return cli.distinct(
			ctx, m.table,
			"username", &users,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	u, err := toArryString(users)
	if err != nil {
		return
	}
	do = repositories.CloudUsersDO{
		Users: u,
	}

	return
}

func toCloudRecordCol(do *repositories.CloudRecordDO) CloudRecord {

	return CloudRecord{
		UserName: do.UserName,
		CloudId:  do.CloudId,
		CreateAt: do.CreateAt,
	}
}
