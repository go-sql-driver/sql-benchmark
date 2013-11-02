package main

import (
	"./framework"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/ziutek/mymysql/godrv"
)

func warmup(db *sql.DB) error {
	db.SetMaxIdleConns(16)

	for i := 0; i < 100000; i++ {
		rows, err := db.Query("SELECT \"Hello Gophers!\"")
		if err != nil {
			return err
		}

		if err = rows.Close(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var err error

	bs := framework.BenchmarkSuite{
		WarmUp:      warmup,
		Repetitions: 3,
		PrintStats:  true,
	}

	if err = bs.AddDriver("mymysql godrv", "mymysql", "/root/root"); err != nil {
		fmt.Println(err)
		return
	}
	if err = bs.AddDriver("Go-MySQL-Driver", "mysql", "root:root@/"); err != nil {
		fmt.Println(err)
		return
	}

	bs.AddBenchmark("SimpleExec", 250000, bmSimpleExec)
	bs.AddBenchmark("PreparedExec", 250000, bmPreparedExec)
	bs.AddBenchmark("SimpleQueryRow", 250000, bmSimpleQueryRow)
	bs.AddBenchmark("PreparedQueryRow", 250000, bmPreparedQueryRow)
	bs.AddBenchmark("PreparedQueryRowParam", 250000, bmPreparedQueryRowParam)
	bs.AddBenchmark("EchoMixed5", 250000, bmEchoMixed5)
	bs.AddBenchmark("SelectLargeString", 50000, bmSelectLargeString)
	bs.AddBenchmark("SelectPreparedLargeString", 50000, bmSelectPreparedLargeString)
	bs.AddBenchmark("SelectLargeBytes", 50000, bmSelectLargeBytes)
	bs.AddBenchmark("SelectPreparedLargeBytes", 50000, bmSelectPreparedLargeBytes)
	bs.AddBenchmark("SelectLargeRaw", 50000, bmSelectLargeRaw)
	bs.AddBenchmark("SelectPreparedLargeRaw", 50000, bmSelectPreparedLargeRaw)
	bs.AddBenchmark("PreparedExecConcurrent1", 500000, bmPreparedExecConcurrent1)
	bs.AddBenchmark("PreparedExecConcurrent2", 500000, bmPreparedExecConcurrent2)
	bs.AddBenchmark("PreparedExecConcurrent4", 500000, bmPreparedExecConcurrent4)
	bs.AddBenchmark("PreparedExecConcurrent8", 500000, bmPreparedExecConcurrent8)
	bs.AddBenchmark("PreparedExecConcurrent16", 500000, bmPreparedExecConcurrent16)
	bs.AddBenchmark("PreparedQueryConcurrent1", 500000, bmPreparedQueryConcurrent1)
	bs.AddBenchmark("PreparedQueryConcurrent2", 500000, bmPreparedQueryConcurrent2)
	bs.AddBenchmark("PreparedQueryConcurrent4", 500000, bmPreparedQueryConcurrent4)
	bs.AddBenchmark("PreparedQueryConcurrent8", 500000, bmPreparedQueryConcurrent8)
	bs.AddBenchmark("PreparedQueryConcurrent16", 500000, bmPreparedQueryConcurrent16)

	bs.Run()
}
