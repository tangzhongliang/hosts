package snsstruct

type ServiceMessage struct {
}
type EPMessage struct {
}
type ServiceMessageResponse struct {
	Ok         bool   `json:"ok"`
	ErrCode    int    `json:"err_code"`
	ErrMessage string `json:"err_message"`
}
