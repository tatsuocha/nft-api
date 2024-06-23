package wire

// MEMO: 再宣言エラーが発生してしまうため、wireで自動生成後は関数はコメントアウトしておく
//func InitializeVoiceHandler() (*handler.VoiceHandler, error) {
//	wire.Build(
//		repository.NewVoiceRepository,
//		usecase.NewVoiceUseCase,
//		handler.NewVoiceHandler,
//	)
//	return &handler.VoiceHandler{}, nil
//}
//
//func InitializeEthereumHandler() (*handler.EthereumHandler, error) {
//	wire.Build(
//		repository.NewEthereumRepository,
//		usecase.NewEthereumUseCase,
//		handler.NewEthereumHandler,
//	)
//	return &handler.EthereumHandler{}, nil
//}
