package green_api

type APIRequest struct {
	IDInstance       string `json:"idInstance"`
	APITokenInstance string `json:"apiTokenInstance"`
	ChatID           string `json:"chatId"`
	Message          string `json:"message"`
	URLFile          string `json:"urlFile"`
	FileName         string `json:"fileName"`
	Caption          string `json:"caption"`
}

type APIResponse map[string]interface{}

type UseCaseGreenApi interface {
	GetSettings(idInstance, apiTokenInstance string) (APIResponse, error)
	GetStateInstance(idInstance, apiTokenInstance string) (APIResponse, error)
	SendMessage(idInstance, apiTokenInstance, chatID, message string) (APIResponse, error)
	SendFileByUrl(idInstance, apiTokenInstance, chatID, urlFile, fileName, caption string) (APIResponse, error)
}
