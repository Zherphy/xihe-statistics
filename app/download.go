package app

import (
	"errors"

	"github.com/opensourceways/xihe-statistics/domain"
	"github.com/opensourceways/xihe-statistics/domain/repository"
)

type DownloadRecordService interface {
	Add(*DownloadRecordAddCmd) error
	Get() (DownloadRecordDTO, error)
}

type downloadRecordService struct {
	dr repository.DownloadRecord
	gs repository.Gitlab
}

func NewDownloadRecordService(
	dr repository.DownloadRecord,
	gs repository.Gitlab,
) DownloadRecordService {
	return downloadRecordService{
		dr: dr,
		gs: gs,
	}
}

func (s downloadRecordService) Add(cmd *DownloadRecordAddCmd) (err error) {
	download := new(domain.DownloadRecord)
	cmd.toDownloadRecord(download)

	return s.dr.Add(download)
}

func (s downloadRecordService) Get() (dto DownloadRecordDTO, err error) {
	// git clone
	cc, err := s.gs.Get()
	if err != nil {
		return
	}

	// download
	counts, err := s.dr.Get()
	if err != nil {
		return
	}

	// clone plus download
	dto = DownloadRecordDTO{
		Counts:   counts + cc.Counts,
		UpdateAt: getLocalTime(),
	}

	return
}

func (cmd *DownloadRecordAddCmd) toDownloadRecord(
	d *domain.DownloadRecord,
) {
	*d = domain.DownloadRecord{
		UserName:     cmd.UserName,
		DownloadPath: cmd.DownloadPath,
		CreateAt:     cmd.CreateAt,
	}
}

func (cmd DownloadRecordAddCmd) Validate() error {
	b := cmd.UserName == nil ||
		cmd.DownloadPath == "" ||
		cmd.CreateAt == 0

	if b {
		return errors.New("invalid cmd of add download record")
	}

	return nil
}

type DownloadRecordDTO struct {
	Counts   int64  `json:"counts"`
	UpdateAt string `json:"update_at"`
}

type DownloadRecordAddCmd struct {
	domain.DownloadRecord
}
