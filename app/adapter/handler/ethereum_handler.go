package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"nft-api/app/domain/usecase"
)

type EthereumHandler struct {
	useCase usecase.EthereumUseCase
}

func NewEthereumHandler(ethereumUseCase usecase.EthereumUseCase) *EthereumHandler {
	return &EthereumHandler{useCase: ethereumUseCase}
}

func (ethereumHandler EthereumHandler) RegisterEthereumRoutes(router *gin.Engine) {
	router.GET("/ethereum", ethereumHandler.get)
}

func (ethereumHandler EthereumHandler) get(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{"response": ethereumHandler.useCase.Get()})
}
