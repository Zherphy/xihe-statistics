package pgsql

import (
	"context"

	"github.com/opensourceways/xihe-statistics/infrastructure/repositories"
)

func NewGitLabRecordMapper(table GitLabRecord) repositories.GitLabRecordMapper {
	return gitlabRecord{table}
}

type gitlabRecord struct {
	table GitLabRecord
}

func (m gitlabRecord) InsertCloneCount(
	do *repositories.CloneCountDO,
) (err error) {
	col, err := m.toGitLabCloneRecord(do)
	if err != nil {
		return
	}

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

func (m gitlabRecord) GetLast() (do repositories.CloneCountDO, err error) {
	col := &GitLabRecord{}

	f := func(ctx context.Context) error {
		return cli.getLast(
			ctx, m.table,
			col,
		)
	}

	if err = withContext(f); err != nil {
		return
	}

	do, err = m.toCloneCountDO(col)
	if err != nil {
		return
	}

	return
}

func (m gitlabRecord) toGitLabCloneRecord(
	do *repositories.CloneCountDO,
) (GitLabRecord, error) {
	return GitLabRecord{
		Counts:   do.Counts,
		CreateAt: do.CreateAt,
	}, nil
}

func (m gitlabRecord) toCloneCountDO(col *GitLabRecord) (repositories.CloneCountDO, error) {
	return repositories.CloneCountDO{
		Counts:   col.Counts,
		CreateAt: col.CreateAt,
	}, nil
}
