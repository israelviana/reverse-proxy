package ports

type IRepo interface {
	StartConnection() error
	GetAllBlockedIPs() ([]string, error)
}
