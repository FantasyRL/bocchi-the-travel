package constants

import "time"

const (
	// service name
	APIServiceName  = "api"
	UserServiceName = "user"

	// db table name
	UserTableName = "user"

	// limit
	MaxConnections     = 1000
	MaxQPS             = 100
	MaxRequestBodySize = 114514 * 1024 * 1024
	MaxIdleConns       = 20
	MaxGoroutines      = 10
	MaxOpenConns       = 100
	ConnMaxLifetime    = 10 * time.Second

	// RPC
	MuxConnection  = 1
	RPCTimeout     = 3 * time.Second
	ConnectTimeout = 50 * time.Millisecond
)

const (
	// page
	PageNum  = 1
	PageSize = 10
)
