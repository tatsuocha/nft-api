package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nft-api/app/adapter/request"
	"nft-api/app/domain/usecase"
)

type VoiceHandler struct {
	useCase usecase.VoiceUseCase
}

func NewVoiceHandler(voiceUseCase usecase.VoiceUseCase) *VoiceHandler {
	return &VoiceHandler{useCase: voiceUseCase}
}

func (voiceHandler VoiceHandler) RegisterVoiceRoutes(router *gin.Engine) {
	router.GET("/voice", voiceHandler.get)
	router.POST("/voice", voiceHandler.create)
}

func (voiceHandler VoiceHandler) get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"response": voiceHandler.useCase.Get()})
}

// MEMO: リクエストサンプル（PowerShell)
// Invoke-WebRequest -Uri http://localhost:8080/voice -Method Post -ContentType "application/json" -Body '{"name": "example voice name", "content": "example voice"}'
func (voiceHandler VoiceHandler) create(context *gin.Context) {
	var voiceCreateRequest request.VoiceCreateRequest
	if err := context.ShouldBindJSON(&voiceCreateRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voiceHandler.useCase.Create(voiceCreateRequest.Name, voiceCreateRequest.Content)
	context.JSON(http.StatusCreated, gin.H{"response": "Voice created successfully"})
}
