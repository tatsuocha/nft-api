package wire

// MEMO: 再宣言エラーが発生してしまうため、wireで自動生成後は当該関数はコメントアウトしておく
//func InitializeWire() (*adapter.VoiceHandler, error) {
//	wire.Build(
//		infrastructure.NewVoiceRepository,
//		usecase.NewVoiceUsecase,
//		adapter.NewVoiceHandler,
//	)
//	return &adapter.VoiceHandler{}, nil
//}
