package app

import "project/xihe-statistics/domain/repository"

type D2Service interface {
	Get() (D2DTO, error)
}

func NewD2Service(
	fileupload repository.FileUploadRecord,
	wukongpublic repository.WuKongPublic,
	userwithrepo repository.UserWithRepo,
) D2Service {
	return &d2Service{
		fileupload:   fileupload,
		wukongpublic: wukongpublic,
		userwithrepo: userwithrepo,
	}
}

type d2Service struct {
	fileupload   repository.FileUploadRecord
	wukongpublic repository.WuKongPublic
	userwithrepo repository.UserWithRepo
}

func (s *d2Service) Get() (dto D2DTO, err error) {
	// fileupload
	fu, err := s.fileupload.Get()
	if err != nil {
		return
	}
	user1 := fu.Users

	// wukongpublic
	wps, err := s.wukongpublic.Get()
	if err != nil {
		return
	}
	user2 := make([]string, len(wps))
	for i := range wps {
		user2[i] = wps[i].UserName.Account()
	}

	// user create repo
	rr, err := s.userwithrepo.Get()
	if err != nil {
		return
	}
	user3 := rr.Users

	// append
	users := append(user1, user2...)
	users = append(users, user3...)
	users = RemoveRepeatedElement(users)

	return D2DTO{
		Counts:   len(users),
		Users:    users,
		UpdateAt: getLocalTime(),
	}, nil
}

type D2DTO struct {
	Counts   int      `json:"counts"`
	Users    []string `json:"users"`
	UpdateAt string   `json:"update_at"`
}
