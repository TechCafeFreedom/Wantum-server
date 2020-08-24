package constants

type key string

const (
	// ユーザのUIDをcontextに格納するためのKey
	AuthCtxKey key = "AUTHED_UID"
	// MYSQL DEFAULT USER
	MysqlDefaultUser string = "root"
	// MYSQL DEFAULT PASSWORD
	MysqlDefaultPassword string = "password"
	// MYSQL DEFAULT HOST
	MysqlDefaultHost string = "localhost"
	// MYSQL DEFAULT PORT
	MysqlDefaultPort string = "3306"
	// MYSQL DEFAULT NAME
	MysqlDefaultName string = "wantum"
)
