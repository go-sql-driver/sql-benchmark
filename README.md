#SQL-Benchmark

A synthetic benchmark to compare the performance of various sql-drivers for Go's database/sql package

## Results
* Intel Core i5-2500K (3.30 GHz), 8 GB RAM
* MySQL 5.1, Windows 7 x64
* Current [Go-MySQL-Driver](https://github.com/Go-SQL-Driver/MySQL) and [mymysql](https://github.com/ziutek/mymysql) versions as of February 25, 2013

### Overview
#### Go1.0.3
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/Go-SQL-Driver/MySQL">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.23</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>2789 queries/second</td>
        <td>1762 queries/second</td>
        <td><b>5266</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td>3279 queries/second</td>
        <td>2015 queries/second</td>
        <td><b>6353</b> queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td><b>7622</b> queries/second</td>
        <td>5446 queries/second</td>
        <td> - </td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td>7783 queries/second</td>
        <td>5746 queries/second</td>
        <td><b>14063</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td><b>15953</b> queries/second</td>
        <td>11139 queries/second</td>
        <td>14327 queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td><b>29332</b> queries/second</td>
        <td>10531 queries/second</td>
        <td>24231 queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td><b>30626</b> queries/second</td>
        <td>26873 queries/second</td>
        <td>25654 queries/second</td>
    </tr>
</table>

### Original Logs
#### Go1.0.3

```
D:\Development\Go\SQL-Benchmark>go build sqlBenchmark.go

D:\Development\Go\SQL-Benchmark>sqlBenchmark.exe
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.6292076s [ 2755 queries/second ]
PreparedQuery: 3.1391796s [ 3186 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 13.1207505s [ 7622 queries/second ]
SimpleQueryRow: 12.8477349s [ 7783 queries/second ]
PreparedQueryRow: 6.2683586s [ 15953 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.409195s [ 29332 queries/second ]
PreparedExec: 3.2651867s [ 30626 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 5.7593294s [ 1736 queries/second ]
PreparedQuery: 4.9632839s [ 2015 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 18.9370832s [ 5281 queries/second ]
SimpleQueryRow: 17.6690106s [ 5660 queries/second ]
PreparedQueryRow: 9.2945316s [ 10759 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 9.6345511s [ 10379 queries/second ]
PreparedExec: 3.7902168s [ 26384 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.585205s [ 2789 queries/second ]
PreparedQuery: 3.0501745s [ 3279 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 13.3307625s [ 7501 queries/second ]
SimpleQueryRow: 13.0667473s [ 7653 queries/second ]
PreparedQueryRow: 6.5433743s [ 15283 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.4541975s [ 28950 queries/second ]
PreparedExec: 3.3351908s [ 29983 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 5.6753246s [ 1762 queries/second ]
PreparedQuery: 4.9732845s [ 2011 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 18.3610502s [ 5446 queries/second ]
SimpleQueryRow: 17.4029953s [ 5746 queries/second ]
PreparedQueryRow: 8.9775135s [ 11139 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 9.4955431s [ 10531 queries/second ]
PreparedExec: 3.7212129s [ 26873 queries/second ]


D:\Development\Go\SQL-Benchmark>
```

### Java
Same machine, Java (JDK7 / 64 bit) + [MySQL Connector/J 5.1.23](http://dev.mysql.com/downloads/connector/j/)
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 1.899s [ 5266 queries/second ]
PreparedQuery: 1.574s [ 6353 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 7.111s [ 14063 queries/second ]
PreparedQueryRow: 6.98s [ 14327 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.127s [ 24231 queries/second ]
PreparedExec: 3.898s [ 25654 queries/second ]
```
