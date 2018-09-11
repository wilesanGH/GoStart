package main

import (
	"GoStart/colly/example/csdn"
	"fmt"
	"GoStart/colly/example/csdn/orm"
)



func main() {
	params := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true", "root", "root", "127.0.0.1:3306", "crawler")
	orm.InitialOrmEngine("mysql",params)


	//example.Usst()
	//example.Cryptocoins()
	csdn.GetCSDNBlog2()
	//example.GetStockListA("e:\\seea.csv")
	//example.GetStockListB("e:\\seeb.csv")

}
