package dal

import (
	"bocchi/rpc/interaction/dal/db"
)

func Init() {
	db.Init()
	//cache.Init()
}
