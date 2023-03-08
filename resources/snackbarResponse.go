package resources

const (
	ERROR   string = "ERROR"
	SUCCESS        = "SUCCESS"
	WARNING        = "WARNING"
	INFO           = "INFO"
)

type SnackbarResponse struct {
	Message string `json:"message"`
	Type    string `json:"type"`

	Description string `json:"description"`
}
