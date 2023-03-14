package adapters

type DBRepo struct {
	DB dbGorm // Should be changed only here ...
}

func (dbr *DBRepo) Connect() {
	dbr.DB.INST = GetGormConn()
}
