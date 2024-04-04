package dal

import (
	"bocchi/rpc/user/dal/cache"
	"bocchi/rpc/user/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
