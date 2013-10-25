package main

import (
	"database/sql"
	"sync"
	"sync/atomic"
)

func bmSimpleExec(db *sql.DB, n int) error {
	for i := 0; i < n; i++ {
		if _, err := db.Exec("DO 1"); err != nil {
			return err
		}
	}
	return nil
}

func bmPreparedExec(db *sql.DB, n int) error {
	stmt, err := db.Prepare("DO 1")
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		if _, err := stmt.Exec(); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func bmSimpleQueryRow(db *sql.DB, n int) error {
	var num int

	for i := 0; i < n; i++ {
		if err := db.QueryRow("SELECT 1").Scan(&num); err != nil {
			return err
		}
	}
	return nil
}

func bmPreparedQueryRow(db *sql.DB, n int) error {
	var num int

	stmt, err := db.Prepare("SELECT 1")
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		if err := stmt.QueryRow().Scan(&num); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func bmPreparedQueryRowParam(db *sql.DB, n int) error {
	var num int

	stmt, err := db.Prepare("SELECT ?")
	if err != nil {
		return err
	}

	for i := 0; i < n; i++ {
		if err := stmt.QueryRow(i).Scan(&num); err != nil {
			return err
		}
	}

	return stmt.Close()
}

func bmEchoMixed5(db *sql.DB, n int) error {
	stmt, err := db.Prepare("SELECT ?, ?, ?, ?, ?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Some random data with different types
	type entry struct {
		id    int64
		name  string
		ratio float64
		other interface{}
		hire  bool
	}

	in := entry{
		id:    42,
		name:  "Gopher",
		ratio: 1.618,
		other: nil,
		hire:  true,
	}

	var out entry

	for i := 0; i < n; i++ {
		if err := stmt.QueryRow(
			in.id,
			in.name,
			in.ratio,
			in.other,
			in.hire,
		).Scan(
			&out.id,
			&out.name,
			&out.ratio,
			&out.other,
			&out.hire,
		); err != nil {
			return err
		}
	}
	return nil
}

func bmSelectLargeString(db *sql.DB, n int) error {
	var str string
	for i := 0; i < n; i++ {
		if err := db.QueryRow("SELECT REPEAT('A', 10000)").Scan(&str); err != nil {
			return err
		}
	}
	return nil
}

func bmSelectPreparedLargeString(db *sql.DB, n int) error {
	stmt, err := db.Prepare("SELECT REPEAT('A', 10000)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var str string
	for i := 0; i < n; i++ {
		if err := stmt.QueryRow().Scan(&str); err != nil {
			return err
		}
	}
	return nil
}

func bmSelectLargeBytes(db *sql.DB, n int) error {
	var raw []byte
	for i := 0; i < n; i++ {
		if err := db.QueryRow("SELECT REPEAT('A', 10000)").Scan(&raw); err != nil {
			return err
		}
	}
	return nil
}

func bmSelectPreparedLargeBytes(db *sql.DB, n int) error {
	stmt, err := db.Prepare("SELECT REPEAT('A', 10000)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var raw []byte
	for i := 0; i < n; i++ {
		if err := stmt.QueryRow().Scan(&raw); err != nil {
			return err
		}
	}
	return nil
}

func bmSelectLargeRaw(db *sql.DB, n int) error {
	var raw sql.RawBytes
	for i := 0; i < n; i++ {
		rows, err := db.Query("SELECT REPEAT('A', 10000)")
		if err != nil {
			return err
		}

		if !rows.Next() {
			return sql.ErrNoRows
		}

		if err = rows.Scan(&raw); err != nil {
			return err
		}

		if err = rows.Close(); err != nil {
			return err
		}
	}
	return nil
}

func bmSelectPreparedLargeRaw(db *sql.DB, n int) error {
	stmt, err := db.Prepare("SELECT REPEAT('A', 10000)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	var raw sql.RawBytes
	for i := 0; i < n; i++ {
		rows, err := stmt.Query()
		if err != nil {
			return err
		}

		if !rows.Next() {
			return sql.ErrNoRows
		}

		if err = rows.Scan(&raw); err != nil {
			return err
		}

		if err = rows.Close(); err != nil {
			return err
		}
	}
	return nil
}

func runPreparedExecConcurrent(db *sql.DB, n int, co int) error {
	stmt, err := db.Prepare("DO 1")
	if err != nil {
		return err
	}

	remain := int64(n)
	var wg sync.WaitGroup
	wg.Add(co)

	for i := 0; i < co; i++ {
		go func() {
			for {
				if atomic.AddInt64(&remain, -1) < 0 {
					wg.Done()
					return
				}

				if _, err1 := stmt.Exec(); err1 != nil {
					wg.Done()
					err = err1
					return
				}
			}
		}()
	}
	wg.Wait()
	stmt.Close()
	return err
}

func bmPreparedExecConcurrent1(db *sql.DB, n int) error {
	return runPreparedExecConcurrent(db, n, 1)
}

func bmPreparedExecConcurrent2(db *sql.DB, n int) error {
	return runPreparedExecConcurrent(db, n, 2)
}

func bmPreparedExecConcurrent4(db *sql.DB, n int) error {
	return runPreparedExecConcurrent(db, n, 4)
}

func bmPreparedExecConcurrent8(db *sql.DB, n int) error {
	return runPreparedExecConcurrent(db, n, 8)
}

func bmPreparedExecConcurrent16(db *sql.DB, n int) error {
	return runPreparedExecConcurrent(db, n, 16)
}

func runPreparedQueryConcurrent(db *sql.DB, n int, co int) error {
	stmt, err := db.Prepare("SELECT ?, \"foobar\"")
	if err != nil {
		return err
	}

	remain := int64(n)
	var wg sync.WaitGroup
	wg.Add(co)

	for i := 0; i < co; i++ {
		go func() {
			var id int
			var str string
			for {
				if atomic.AddInt64(&remain, -1) < 0 {
					wg.Done()
					return
				}

				if err1 := stmt.QueryRow(i).Scan(&id, &str); err1 != nil {
					wg.Done()
					err = err1
					return
				}
			}
		}()
	}
	wg.Wait()
	stmt.Close()
	return err
}

func bmPreparedQueryConcurrent1(db *sql.DB, n int) error {
	return runPreparedQueryConcurrent(db, n, 1)
}

func bmPreparedQueryConcurrent2(db *sql.DB, n int) error {
	return runPreparedQueryConcurrent(db, n, 2)
}

func bmPreparedQueryConcurrent4(db *sql.DB, n int) error {
	return runPreparedQueryConcurrent(db, n, 4)
}

func bmPreparedQueryConcurrent8(db *sql.DB, n int) error {
	return runPreparedQueryConcurrent(db, n, 8)
}

func bmPreparedQueryConcurrent16(db *sql.DB, n int) error {
	return runPreparedQueryConcurrent(db, n, 16)
}
