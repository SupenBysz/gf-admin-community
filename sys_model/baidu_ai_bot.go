package sys_model

type AiBotAppListRes struct {
	ID          string `json:"id" dc:"应用ID"`
	Name        string `json:"name" dc:"应用名称"`
	Description string `json:"description" dc:"应用简介"`
}
