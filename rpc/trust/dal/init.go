package dal

import (
	"bocchi/rpc/trust/dal/cache"
	"bocchi/rpc/trust/dal/db"
)

func Init() {
	db.Init()
	cache.Init()
}
