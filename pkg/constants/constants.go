package constants

type key string

const (
	AuthCtxKey key = "AUTHED_UID" // ユーザのUIDをcontextに格納するためのKey
)
