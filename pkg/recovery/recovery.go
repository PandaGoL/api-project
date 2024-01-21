package recovery

// Recovery - интерфейс работы с паниками
type Recovery interface {
	// Do - запускает отлов паник
	Do()
}

// Panickier - интерфейс для установки состояния паники модуля
type Panickier interface {
	SetPanicState(bool)
}
