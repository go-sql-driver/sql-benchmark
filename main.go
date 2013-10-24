package main

import (
	"./framework"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func main() {
	var err error

	bs := framework.BenchmarkSuite{}

	if err = bs.AddDriver("mymysql godrv", "mymysql", "gotest/root/root"); err != nil {
		fmt.Println(err)
		return
	}
	if err = bs.AddDriver("Go-MySQL-Driver", "mysql", "root:root@/gotest"); err != nil {
		fmt.Println(err)
		return
	}

	bs.AddBenchmark("SimpleExec", 500000, bmSimpleExec)
	bs.AddBenchmark("PreparedExec", 500000, bmPreparedExec)
	bs.AddBenchmark("SimpleQueryRow", 500000, bmSimpleQueryRow)
	bs.AddBenchmark("PreparedQueryRow", 500000, bmPreparedQueryRow)
	bs.AddBenchmark("PreparedQueryRowParam", 500000, bmPreparedQueryRowParam)
	bs.AddBenchmark("EchoMixed5", 500000, bmEchoMixed5)
	bs.AddBenchmark("SelectLargeString", 100000, bmSelectLargeString)
	bs.AddBenchmark("SelectPreparedLargeString", 100000, bmSelectPreparedLargeString)
	bs.AddBenchmark("SelectLargeBytes", 100000, bmSelectLargeBytes)
	bs.AddBenchmark("SelectPreparedLargeBytes", 100000, bmSelectPreparedLargeBytes)
	bs.AddBenchmark("SelectLargeRaw", 100000, bmSelectLargeRaw)
	bs.AddBenchmark("SelectPreparedLargeRaw", 100000, bmSelectPreparedLargeRaw)

	bs.Run()
}
