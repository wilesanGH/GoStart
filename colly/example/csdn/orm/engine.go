package orm
import(
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/go-xorm/core"
)

var CsdnEngine *xorm.Engine

func InitialOrmEngine(driverName string, dataSourceName string){
	var err error
	CsdnEngine,err = xorm.NewEngine(driverName,dataSourceName)
	CsdnEngine.ShowSQL(true)
	CsdnEngine.Logger().SetLevel(core.LOG_INFO)
	CsdnEngine.SetMapper(core.SameMapper{})
	CsdnEngine.SetColumnMapper(core.SnakeMapper{})
	if err != nil{
		panic(err)
	}
}