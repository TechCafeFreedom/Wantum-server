package constants

type key string

const (
	// ユーザのUIDをcontextに格納するためのKey
	AuthCtxKey key = "AUTHED_UID"
	// DB DEFAULT USER
	DbDefaultUser string = "root"
	// DB DEFAULT PASSWORD
	DbDefaultPassword string = "password"
	// DB DEFAULT HOST
	DbDefaultHost string = "localhost"
	// DB DEFAULT PORT
	DbDefaultPort string = "3306"
	// DB DEFAULT NAME
	DbDefaultName string = "wantum"
)
