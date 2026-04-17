package green_api

import (
	"green-api-web-ui/internal/infrastructure/http"
	"green-api-web-ui/internal/interfaces/green_api"
)

type UseCase struct {
	client *http.Client
}

func NewUseCase(client *http.Client) green_api.UseCaseGreenApi {
	return &UseCase{
		client: client,
	}
}

func (uc *UseCase) GetSettings(idInstance, apiTokenInstance string) (green_api.APIResponse, error) {
	return uc.client.GetSettings(idInstance, apiTokenInstance)
}

func (uc *UseCase) GetStateInstance(idInstance, apiTokenInstance string) (green_api.APIResponse, error) {
	return uc.client.GetStateInstance(idInstance, apiTokenInstance)
}

func (uc *UseCase) SendMessage(idInstance, apiTokenInstance, chatID, message string) (green_api.APIResponse, error) {
	return uc.client.SendMessage(idInstance, apiTokenInstance, chatID, message)
}

func (uc *UseCase) SendFileByUrl(idInstance, apiTokenInstance, chatID, urlFile, fileName, caption string) (green_api.APIResponse, error) {
	return uc.client.SendFileByUrl(idInstance, apiTokenInstance, chatID, urlFile, fileName, caption)
}
