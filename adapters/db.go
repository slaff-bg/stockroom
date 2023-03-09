package adapters

type DBRepo struct {
	DB dbGorm // Should be changed only here ...
}

type DBS interface {
	DBConn() *DBRepo
}

func DBConn(db DBS) *DBRepo {
	return db.DBConn()
}
