package main

import (
	"GoStart/crawler/engine"
	"GoStart/crawler/parser/city"
	"GoStart/crawler/scheduler"
	"GoStart/crawler/persist"
)

func main() {
	//params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "cra", "cra", "202.120.222.140:80", "crawler")
	//orm.InitialOrmEngine("mysql",params)
	//cl.GetPageCount(nil)
	e := engine.ConcurrentEngine{
		Scheduler:&scheduler.QueuedScheduler{},
		WorkerCount:100,
		ItemChan: persist.ItemSaver(),
	}
/*	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun",
		//Url:"http://www.t66y.com/thread0806.php?fid=7",
		ParserFunc:city.ParseCityList,
		NeedVPN:false,
	})*/

	e.Run(engine.Request{
		Url:"http://www.zhenai.com/zhenghun/shanghai",
		//Url:"http://www.t66y.com/thread0806.php?fid=7",
		ParserFunc:city.ParseCity,
		NeedVPN:false,
	})
}




/*,
engine.Request{Url:"http://www.zhenai.com/zhenghun",
ParserFunc:city.ParseCityList,
NeedVPN:false,
}*/
