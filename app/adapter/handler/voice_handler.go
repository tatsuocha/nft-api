package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nft-api/app/adapter/request"
	"nft-api/app/usecase"
)

type VoiceHandler struct {
	usecase usecase.VoiceUsecase
}

func NewVoiceHandler(voiceUsecase usecase.VoiceUsecase) *VoiceHandler {
	return &VoiceHandler{usecase: voiceUsecase}
}

func (voiceHandler *VoiceHandler) RegisterVoiceRoutes(router *gin.Engine) {
	router.GET("/voice", voiceHandler.get)
	router.POST("/voice", voiceHandler.create)
}

func (voiceHandler *VoiceHandler) get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"response": voiceHandler.usecase.Get()})
}

// MEMO: リクエストサンプル（PowerShell
// Invoke-WebRequest -Uri http://localhost:8080/voice -Method Post -ContentType "application/json" -Body '{"name": "example voice", "content": "example voice"}'
func (voiceHandler *VoiceHandler) create(context *gin.Context) {
	var voiceCreateRequest request.VoiceCreateRequest
	if err := context.ShouldBindJSON(&voiceCreateRequest); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voiceHandler.usecase.Create(voiceCreateRequest)
	context.JSON(http.StatusCreated, gin.H{"response": "Voice created successfully"})
}
