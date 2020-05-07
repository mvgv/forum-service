package validation

/*ApplicationError formata a mensagem de erro retornada pela api*/
type ApplicationError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"status_code"`
}
