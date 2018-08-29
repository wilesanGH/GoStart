package main

import (
	"GoStart/crawler/orm"
	"fmt"
	"GoStart/crawler/parser/cl"
)



func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "P@ssw0rd", "202.120.222.140:80", "crawler")
	orm.InitialOrmEngine("mysql",params)
	cl.GetPageCount(nil)
	/*engine.Run(engine.Request{
		//Url:"http://www.zhenai.com/zhenghun",
		Url:"http://www.t66y.com/thread0806.php?fid=7",
		ParserFunc:cl.GetPageCount,
		NeedVPN:true,
	})*/
}




/*,
engine.Request{Url:"http://www.zhenai.com/zhenghun",
ParserFunc:city.ParseCityList,
NeedVPN:false,
}*/
