package orm
import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var Engine *xorm.Engine

func InitialOrmEngine(driverName string, dataSourceName string){
	var err error
	Engine,err = xorm.NewEngine(driverName,dataSourceName)
	if err != nil{
		panic(err)
	}
}