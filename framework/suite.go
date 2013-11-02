package framework

import (
	"database/sql"
	"fmt"
	"runtime"
	"sort"
	"time"
)

type driver struct {
	name string
	db   *sql.DB
}

type Result struct {
	Err      error
	Queries  int
	Duration time.Duration
	Allocs   uint64
	Bytes    uint64
}

func (res *Result) QueriesPerSecond() float64 {
	return float64(res.Queries) / res.Duration.Seconds()
}

func (res *Result) AllocsPerQuery() int {
	return int(res.Allocs) / res.Queries
}

func (res *Result) BytesPerQuery() int {
	return int(res.Bytes) / res.Queries
}

var memStats runtime.MemStats

type benchmark struct {
	name string
	n    int
	bm   func(*sql.DB, int) error
}

func (b *benchmark) run(db *sql.DB) Result {
	runtime.GC()

	runtime.ReadMemStats(&memStats)
	var (
		startMallocs    = memStats.Mallocs
		startTotalAlloc = memStats.TotalAlloc
		startTime       = time.Now()
	)

	err := b.bm(db, b.n)

	endTime := time.Now()
	runtime.ReadMemStats(&memStats)

	return Result{
		Err:      err,
		Queries:  b.n,
		Duration: endTime.Sub(startTime),
		Allocs:   memStats.Mallocs - startMallocs,
		Bytes:    memStats.TotalAlloc - startTotalAlloc,
	}
}

type BenchmarkSuite struct {
	drivers     []driver
	benchmarks  []benchmark
	WarmUp      func(*sql.DB) error
	Repetitions int
	PrintStats  bool
}

func (bs *BenchmarkSuite) AddDriver(name, drv, dsn string) error {
	db, err := sql.Open(drv, dsn)
	if err != nil {
		return fmt.Errorf("Error registering driver '%s': %s", name, err.Error())
	}

	if err = db.Ping(); err != nil {
		return fmt.Errorf("Error on driver '%s': %s", name, err.Error())
	}

	bs.drivers = append(bs.drivers, driver{
		name: name,
		db:   db,
	})
	return nil
}

func (bs *BenchmarkSuite) AddBenchmark(name string, n int, bm func(*sql.DB, int) error) {
	bs.benchmarks = append(bs.benchmarks, benchmark{
		name: name,
		n:    n,
		bm:   bm,
	})
}

func (bs *BenchmarkSuite) Run() {
	startTime := time.Now()

	if len(bs.drivers) < 1 {
		fmt.Println("No drivers registered to run benchmarks with!")
		return
	}

	if len(bs.benchmarks) < 1 {
		fmt.Println("No benchmark functions registered!")
		return
	}

	if bs.WarmUp != nil {
		for _, driver := range bs.drivers {
			fmt.Println("Warming up " + driver.name + "...")
			if err := bs.WarmUp(driver.db); err != nil {
				fmt.Println(err.Error())
				return
			}
		}
		fmt.Println()
	}

	fmt.Println("Run Benchmarks...")
	fmt.Println()

	var qps []float64
	if bs.Repetitions > 1 && bs.PrintStats {
		qps = make([]float64, bs.Repetitions)
	} else {
	     bs.PrintStats = false
	}

	for _, benchmark := range bs.benchmarks {
		fmt.Println(benchmark.name, benchmark.n, "iterations")
		for _, driver := range bs.drivers {
			fmt.Println(driver.name)
			for i := 0; i < bs.Repetitions; i++ {
				res := benchmark.run(driver.db)
				if res.Err != nil {
					fmt.Println(res.Err.Error())
				} else {
					fmt.Println(
						" " +
						res.Duration.String(), "\t   ",
						int(res.QueriesPerSecond()+0.5), "queries/sec\t   ",
						res.AllocsPerQuery(), "allocs/query\t   ",
						res.BytesPerQuery(), "B/query",
					)
					if bs.Repetitions > 1 {
						qps[i] = res.QueriesPerSecond()
					}
				}
			}
			
			if bs.PrintStats {
		          var totalQps float64
		          for i := range qps {
		               totalQps += qps[i]
		          }
		          
		          sort.Float64s(qps)
		          
		          fmt.Println(
		               " -- " +
		               "avg", int(totalQps/float64(len(qps)) +0.5), "qps;  "+
		               "median", int(qps[len(qps)/2]+0.5), "qps",
		          )
		     }
		}
		
		fmt.Println()
	}
	endTime := time.Now()
	fmt.Println("Finished... Total running time:", endTime.Sub(startTime).String())
}
