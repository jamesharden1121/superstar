package conf

const DriverName = "mysql"
const Url = "root:root@/superstar?charset=utf8&parseTime=True&loc=Local"

type DbConf struct {
	Host   string
	Port   int
	User   string
	Pwd    string
	DbName string
}

var MasterDbConfig DbConf = DbConf{
	"127.0.0.1",
	3306,
	"root",
	"root",
	"superstar",
}

var SlaveDbConfig DbConf = DbConf{
	"127.0.0.1",
	3306,
	"root",
	"root",
	"superstar",
}
