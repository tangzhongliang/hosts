package snsstruct

type SnsEpMessage struct {
}

type SnsEpMessageResponse struct {
}
type SnsEpEmailBindRequest struct {
	Email       string
	Account     string
	AccountType string
}

type LineAccessTokenResponse struct {
	Scope        string `json:"scope"`
	AccessToken  string `json:"access_token"`
	TokenType    string `json:"token_type"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
}

type LineProfile struct {
	UserID        string `json:"userId"`
	DisplayName   string `json:"displayName"`
	PictureURL    string `json:"pictureUrl"`
	StatusMessage string `json:"statusMessage"`
}
