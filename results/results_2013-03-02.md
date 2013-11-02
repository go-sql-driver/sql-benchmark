#SQL-Benchmark

A synthetic benchmark to compare the performance of various sql-drivers for Go's database/sql package

## Results
### Contributed Results
* [March 19, 2013](results-osx_2013-03-19.md) // OS X 10.8.3 / MySQL 5.6.10 by [Arne Hormann](https://github.com/arnehormann)
* [March 19, 2013](results-linux_2013-03-19.md) // Linux / MySQL 5.5.27 by [Arne Hormann](https://github.com/arnehormann)

### Older Results
* [February 26, 2013](results_2013-02-26.md) // Show [Diff to March 02, 2013](https://github.com/Go-SQL-Driver/SQL-Benchmark/commit/52078ce64a397d94b08bc66f5d7acedf13d3364e)

### Setup
* Intel Core i5-2500K (3.30 GHz), 8 GB RAM
* MySQL 5.1, Windows 7 x64
* Current [Go-MySQL-Driver](https://github.com/go-sql-driver/mysql) and [mymysql](https://github.com/ziutek/mymysql) versions as of March 02, 2013
* Java7 (JDBC) + MySQL Connector/J 5.1.23 for comparision

### Notes
* Please don't try to compare the benchmark sets (Query | QueryRow | Exec) to each other. They test different things. So you can't say QueryRow is faster than Query. In fact QueryRow is just a shortcut which uses Query internally. View the source to understand what we test here.
* The benchmarks are designed to minimize response latency from the server. We try to compare driver performance here and not to test the MySQL Server ;)
* The setup MySQL 5.1 + Windows isn't ideal. Maybe I'll setup a GNU/Linux + MariaDB 10.0 enviornment another time and run this again.
* This benchmark isn't concurrent while most production apps will be. You might get many more queries/second by running them concurrent. I'll add a test some time (feel free to contribute).
* The Go results tend to vary much more. Probably this is caused by the garbage collector.

### Go tip
* go version devel +646803253bb4 Tue Feb 26 09:51:33 2013
* reverted [changeset ddb9e6365e57](http://code.google.com/p/go/source/detail?r=ddb9e6365e570c2619d88427176a465e8b76b4aa) because of [Issue #4902](http://golang.org/issue/4902)

<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/go-sql-driver/mysql">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.23</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>4697 queries/second</td>
        <td>3111 queries/second</td>
        <td><b>5266</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td>5053 queries/second</td>
        <td>4156 queries/second</td>
        <td><b>6353</b> queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td><b>7868</b> queries/second</td>
        <td>6740 queries/second</td>
        <td> - </td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td><b>14712</b> queries/second</td>
        <td>7053 queries/second</td>
        <td>14063 queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td><b>16436</b> queries/second</td>
        <td>14452 queries/second</td>
        <td>14327 queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td><b>29127</b> queries/second</td>
        <td>27239 queries/second</td>
        <td>24231 queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td><b>30533</b> queries/second</td>
        <td>28784 queries/second</td>
        <td>25654 queries/second</td>
    </tr>
</table>

### Go1.0.3
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/go-sql-driver/mysql">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.23</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>3754 queries/second</td>
        <td>2888 queries/second</td>
        <td><b>5266</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td>5000 queries/second</td>
        <td>3724 queries/second</td>
        <td><b>6353</b> queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td><b>7855</b> queries/second</td>
        <td>6838 queries/second</td>
        <td> - </td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td>7981 queries/second</td>
        <td>6992 queries/second</td>
        <td><b>14063</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td><b>16490</b> queries/second</td>
        <td>13857 queries/second</td>
        <td>14327 queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td><b>29427</b> queries/second</td>
        <td>27939 queries/second</td>
        <td>24231 queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td><b>30834</b> queries/second</td>
        <td>29170 queries/second</td>
        <td>25654 queries/second</td>
    </tr>
</table>

### Original Logs
#### Go tip

```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.1291218s [ 4697 queries/second ]
PreparedQuery: 2.0041147s [ 4990 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 12.709727s [ 7868 queries/second ]
SimpleQueryRow: 6.7973888s [ 14712 queries/second ]
PreparedQueryRow: 6.2203558s [ 16076 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.4331964s [ 29127 queries/second ]
PreparedExec: 3.2751873s [ 30533 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.2221843s [ 3103 queries/second ]
PreparedQuery: 2.4981429s [ 4003 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 14.8358486s [ 6740 queries/second ]
SimpleQueryRow: 14.4258251s [ 6932 queries/second ]
PreparedQueryRow: 7.2694158s [ 13756 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.67121s [ 27239 queries/second ]
PreparedExec: 3.4741987s [ 28784 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.1351221s [ 4684 queries/second ]
PreparedQuery: 1.9791132s [ 5053 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 12.8407344s [ 7788 queries/second ]
SimpleQueryRow: 6.8563922s [ 14585 queries/second ]
PreparedQueryRow: 6.084348s [ 16436 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.4891995s [ 28660 queries/second ]
PreparedExec: 3.3341907s [ 29992 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.2141838s [ 3111 queries/second ]
PreparedQuery: 2.4061376s [ 4156 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 15.228871s [ 6566 queries/second ]
SimpleQueryRow: 14.178811s [ 7053 queries/second ]
PreparedQueryRow: 6.9193957s [ 14452 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.6822106s [ 27158 queries/second ]
PreparedExec: 3.5232015s [ 28383 queries/second ]


```


#### Go1.0.3

```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.692154s [ 3714 queries/second ]
PreparedQuery: 2.0001144s [ 5000 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 12.7307282s [ 7855 queries/second ]
SimpleQueryRow: 12.5297167s [ 7981 queries/second ]
PreparedQueryRow: 6.0643468s [ 16490 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.409195s [ 29332 queries/second ]
PreparedExec: 3.2571863s [ 30701 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.4911997s [ 2864 queries/second ]
PreparedQuery: 2.8191612s [ 3547 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 14.790846s [ 6761 queries/second ]
SimpleQueryRow: 14.3098185s [ 6988 queries/second ]
PreparedQueryRow: 7.2164128s [ 13857 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.6492087s [ 27403 queries/second ]
PreparedExec: 3.4471971s [ 29009 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.6641524s [ 3754 queries/second ]
PreparedQuery: 2.0211156s [ 4948 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 12.9847427s [ 7701 queries/second ]
SimpleQueryRow: 12.5747192s [ 7952 queries/second ]
PreparedQueryRow: 6.1143498s [ 16355 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.3981944s [ 29427 queries/second ]
PreparedExec: 3.2431855s [ 30834 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.4631981s [ 2888 queries/second ]
PreparedQuery: 2.6851535s [ 3724 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 14.6238365s [ 6838 queries/second ]
SimpleQueryRow: 14.302818s [ 6992 queries/second ]
PreparedQueryRow: 7.2354138s [ 13821 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 3.5792047s [ 27939 queries/second ]
PreparedExec: 3.4281961s [ 29170 queries/second ]


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
