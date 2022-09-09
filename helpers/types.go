package helpers

type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Error   Error       `json:"error"`
}
type Error struct {
	ErrorCode    int    `json:"error_code,omitempty"`
	ErrorSubCode int    `json:"error_sub_code,omitempty"`
	ErrorMessage string `json:"error_message,omitempty"`
	ErrprSubCode int    `json:"error_subcode,omitempty"`
}
