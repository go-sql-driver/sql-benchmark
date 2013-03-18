### Setup
* Intel Core i7 (2.20 GHz), 16 GB RAM, 512 GB Samsung 840 Pro SSD
* OS X 10.8.3 (12D78), full disk encryption enabled
* MySQL 5.6.10 Source distribution (`brew install mysql` without configuration changes)
* Current [Go-MySQL-Driver](https://github.com/Go-SQL-Driver/MySQL) and [mymysql](https://github.com/ziutek/mymysql) versions as of March 18, 2013
* Java7 (JDBC) + MySQL Connector/J 5.1.24 for comparison

### Results for go tip
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/Go-SQL-Driver/MySQL">Go-MySQL-Driver</a> (worst of full run)</th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.24</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>4016 queries/second</td>
        <td>2823 queries/second</td>
        <td><b>4087</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td><b>4692</b> queries/second</td>
        <td>3633 queries/second</td>
        <td>3426 queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td><b>5739</b> queries/second</td>
        <td> - </td>
        <td> - </td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td><b>8620</b> queries/second</td>
        <td> - </td>
        <td>8315 queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td><b>9608</b> queries/second</td>
        <td> - </td>
        <td>5135 queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td><b>18990</b> queries/second</td>
        <td> - </td>
        <td>14879 queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td><b>23416</b> queries/second</td>
        <td> - </td>
        <td>15571 queries/second</td>
    </tr>
</table>


### Original Logs
#### Go tip
```
$ go version
go version devel +43eb97ed849a Mon Mar 18 12:18:49 2013 +1100 darwin/amd64
```

```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.480917311s [ 4031 queries/second ]
PreparedQuery: 2.132368083s [ 4690 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.293874865s [ 5782 queries/second ]
SimpleQueryRow: 11.534512883s [ 8670 queries/second ]
PreparedQueryRow: 10.473929602s [ 9548 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 5.3381216s [ 18733 queries/second ]
PreparedExec: 4.381224318s [ 22825 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.542225167s [ 2823 queries/second ]
PreparedQuery: 2.752905968s [ 3633 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
panic: Received #1461 error from MySQL server: "Can't create more than max_prepared_stmt_count statements (current value: 16382)"
```

The cause was in [sqlBenchmark.go:82](https://github.com/Go-SQL-Driver/SQL-Benchmark/blob/c84d727235f54a38e3ea3b9d94b75e6eeb51967a/sqlBenchmark.go#L82), reported as [issue 65](https://github.com/ziutek/mymysql/issues/65). After removing the mymysql runs:


```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.479931922s [ 4032 queries/second ]
PreparedQuery: 2.131371178s [ 4692 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.37663407s [ 5755 queries/second ]
SimpleQueryRow: 11.600404342s [ 8620 queries/second ]
PreparedQueryRow: 10.40767065s [ 9608 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 5.265846228s [ 18990 queries/second ]
PreparedExec: 4.246508685s [ 23549 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.489874616s [ 4016 queries/second ]
PreparedQuery: 2.13081187s [ 4693 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.425357853s [ 5739 queries/second ]
SimpleQueryRow: 11.428583065s [ 8750 queries/second ]
PreparedQueryRow: 10.351344836s [ 9661 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 5.257325283s [ 19021 queries/second ]
PreparedExec: 4.270571415s [ 23416 queries/second ]
```

### Java
```
$ java - version
java version "1.7.0_15"
Java(TM) SE Runtime Environment (build 1.7.0_15-b03)
Java HotSpot(TM) 64-Bit Server VM (build 23.7-b01, mixed mode)
```

```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.447s [ 4087 queries/second ]
PreparedQuery: 2.919s [ 3426 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 12.026s [ 8315 queries/second ]
PreparedQueryRow: 19.474s [ 5135 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.721s [ 14879 queries/second ]
PreparedExec: 6.422s [ 15571 queries/second ]
```
