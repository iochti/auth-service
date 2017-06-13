package models

// GhubUser type represents a Oauth fetched Ghub User
type GhubUser struct {
	ID        string `json:"id"`
	Email     string `json:"email"`
	Login     string `json:"login"`
	AvatarURL string `json:"avatar_url"`
	Name      string `json:"name"`
}
