package repositories

import (
	"github.com/opensourceways/xihe-statistics/domain"
	"github.com/opensourceways/xihe-statistics/domain/repository"
)

type CloudRecordMapper interface {
	AddCloudRecord(*CloudRecordDO) error
	GetCloudRecordCount(cloudId string) (int64, error)
	GetUsers() (CloudUsersDO, error)
}

func NewCloudRecordRepository(mapper CloudRecordMapper) repository.CloudRecord {
	return &cloudRecord{mapper}
}

type cloudRecord struct {
	mapper CloudRecordMapper
}

func (impl *cloudRecord) Add(d *domain.Cloud) (err error) {
	do := new(CloudRecordDO)
	do.toCloudRecordDO(d)

	return impl.mapper.AddCloudRecord(do)
}

func (impl *cloudRecord) GetUsers() (repository.CloudUsers, error) {
	do, err := impl.mapper.GetUsers()
	if err != nil {
		return repository.CloudUsers{}, err
	}

	return repository.CloudUsers{
		Users: do.Users,
	}, nil
}

func (impl *cloudRecord) Get(cloudType domain.CloudType) (int64, error) {
	return impl.mapper.GetCloudRecordCount(cloudType.CloudId())
}

func (do *CloudRecordDO) toCloudRecordDO(d *domain.Cloud) {
	*do = CloudRecordDO{
		UserName: d.UserName.Account(),
		CloudId:  d.CloudId,
		CreateAt: d.CreateAt,
	}
}

type CloudRecordDO struct {
	UserName string
	CloudId  string
	CreateAt int64
}

type CloudUsersDO struct {
	Users []string
}
