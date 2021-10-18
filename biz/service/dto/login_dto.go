package dto

type UserInfoDTO struct {
	SessionId  string `json:"session_id,omitempty"`
	CoreUserId int64  `json:"core_user_id,omitempty"`

	UserName  string `json:"user_name,omitempty"`
	AvatarUrl string `json:"avatar_url,omitempty"`
	Slogan    string `json:"slogan,omitempty"`
}

type AnswerInfoDTO struct {
	UserInfoDTO
	ProductDescription string `json:"product_description"`
	RealPrize          int64  `json:"real_prize"`
}

type GetCaptureDTO struct {
	Token      string `json:"token,omitempty"`
	CaptureUrl string `json:"capture_url,omitempty"`
}

type AnswerListDTO struct {
	Lists      []AnswerInfoDTO `json:"lists"`
	Pagination Pagination      `json:"pagination"`
}

type Pagination struct {
	Page    int64 `json:"page,omitempty"`
	Limit   int64 `json:"limit,omitempty"`
	Total   int64 `json:"total,omitempty"`
	HasMore bool  `json:"has_more,omitempty"`
}
