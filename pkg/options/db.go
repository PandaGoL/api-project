package options

type DB struct {
	Username       string
	Password       string
	Host           string
	DBName         string
	Port           string
	MaxConnections int32
}
