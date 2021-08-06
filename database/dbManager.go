package dbManager

type dbinfo struct {
	host     string
	user     string
	password string
	dbname   string
}

func createDbStruct(host string, user string, password string, dbname string) dbinfo {
	return dbinfo{host, user, password, dbname}
}
