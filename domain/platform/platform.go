package platform

type ProjectId struct {
	Id int `json:"id"`
}

type CloneTotal struct {
	Total int64 `json:"total"`
}

func (r *CloneTotal) CleanData() {
	if r.Total > 500 {
		r.Total = 500
	}
}

type PlatForm interface {
	GetProjectId(int) ([]ProjectId, error)
	GetCloneTotal(int) (CloneTotal, error)
}

func (r *ProjectId) IsAbnormal() bool {
	for _, v := range [...]int{2469, 2599, 2598, 2597, 3084, 3459, 3407, 3534, 3528} {
		if v == r.Id {
			return true
		}
	}

	return false
}
