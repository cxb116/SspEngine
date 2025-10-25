package interfaces

type ITranProtocol interface {
	DoRequest(payload any) ([]byte, error)
}
