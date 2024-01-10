package storage

type DatabaseItf interface {
	Connect() error
	Disconnect() error
}
