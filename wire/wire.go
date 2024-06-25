package wire

// MEMO: 再宣言エラーが発生してしまうため、wireで自動生成後はコメントアウトしておく
//func InitializeVoiceHandler() (*handler.VoiceHandler, error) {
//	wire.Build(
//		postgresql.NewCommonDao,
//		postgresql.NewVoiceDao,
//		repository.NewVoiceRepository,
//		usecase.NewVoiceUseCase,
//		handler.NewVoiceHandler,
//	)
//	return &handler.VoiceHandler{}, nil
//}
//
//func InitializeEthereumHandler() (*handler.EthereumHandler, error) {
//	wire.Build(
//		ethereum.NewEthereumDao,
//		repository.NewEthereumRepository,
//		usecase.NewEthereumUseCase,
//		handler.NewEthereumHandler,
//	)
//	return &handler.EthereumHandler{}, nil
//}
