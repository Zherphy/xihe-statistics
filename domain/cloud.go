package domain

import "errors"

const (
	cloudTypeCpu = "cpu"
	cloudTypeNpu = "npu"
	cloudIdCpu   = "cpu_001"
	cloudIdNpu   = "ascend_001"
)

type Cloud struct {
	UserName Account
	CloudId  string
	CreateAt int64
}

type CloudType interface {
	CloudId() string
}

type cloudType string

func (t cloudType) CloudId() string {
	if t == cloudTypeCpu {
		return cloudIdCpu
	}

	return cloudIdNpu
}

func NewCloudType(v string) (CloudType, error) {
	switch v {
	case cloudTypeCpu, cloudTypeNpu:
		return cloudType(v), nil
	}

	return nil, errors.New("not supported type")
}
