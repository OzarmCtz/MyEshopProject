package global

import (
	"database/sql"

	adm "github.com/OzarmCtz/e_shop_backend_v1/appli/datasources/mysql"
)

var (
	MysqlUserDbPool   = make(map[int32]*sql.DB)
	QueriesUserDbPool = make(map[int32]*adm.Queries)
)
