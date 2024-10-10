package message

import "github.com/opensourceways/xihe-statistics/domain"

type BigModelRecordHandler interface {
	AddBigModelRecord(*domain.UserWithBigModel) error
}

type RepoRecordHandler interface {
	AddRepoRecord(*domain.UserWithRepo) error
}

type RegisterRecordHandler interface {
	AddRegisterRecord(*domain.RegisterRecord) error
}

type FileUploadRecordHandler interface {
	AddUploadFileRecord(*domain.FileUploadRecord) error
}

type DownloadRecordHandler interface {
	AddDownloadRecord(*domain.DownloadRecord) error
}

type TrainRecordHandler interface {
	AddTrainRecord(*domain.TrainRecord) error
}

type CloudRecordHandler interface {
	AddCloudRecord(d *domain.Cloud) error
}
