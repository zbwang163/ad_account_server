package dto

type UserInfoDTO struct {
	SessionId  string `json:"session_id,omitempty"`
	CoreUserId int64  `json:"core_user_id,omitempty"`

	UserName  string `json:"user_name"`
	AvatarUrl string `json:"avatar_url"`
	Slogan    string `json:"slogan"`
}

type GetCaptureDTO struct {
	Token      string `json:"token"`
	CaptureUrl string `json:"capture_url"`
}
