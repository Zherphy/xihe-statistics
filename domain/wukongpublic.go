package domain

type WuKongPublic struct {
	UserName Account `json:"username"`
	CreateAt int64   `json:"create_at"`
}
