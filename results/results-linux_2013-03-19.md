### Setup
* Intel Xeon CPU 3.20GHz, 4 GB RAM
* Linux 64 Bit
* MySQL 5.5.27
* Current [Go-MySQL-Driver](https://github.com/Go-SQL-Driver/MySQL) and [mymysql](https://github.com/ziutek/mymysql) versions as of March 19, 2013
* Java7 (JDBC) + MySQL Connector/J 5.1.24 for comparison

### Results for go 1.0.3, Java run 1 and 2
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/Go-SQL-Driver/MySQL">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.24</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>2037 queries/second</td>
        <td>1631 queries/second</td>
        <td><b>3323</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td>2928 queries/second</td>
        <td>2246 queries/second</td>
        <td><b>5195</b> queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td><b>3645</b> queries/second</td>
        <td>3323 queries/second</td>
        <td>-</td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td>3835 queries/second</td>
        <td>3401 queries/second</td>
        <td><b>11071</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td>8803 queries/second</td>
        <td>8045 queries/second</td>
        <td><b>11453</b> queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td>10027 queries/second</td>
        <td>9474 queries/second</td>
        <td><b>11455</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td>10278 queries/second</td>
        <td>9831 queries/second</td>
        <td><b>11832</b> queries/second</td>
    </tr>
</table>

### Results for go tip, Java run 3 and 4
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/Go-SQL-Driver/MySQL">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.24</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td><b>3743</b> queries/second</td>
        <td>2683 queries/second</td>
        <td>3301 queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td>4101 queries/second</td>
        <td>3018 queries/second</td>
        <td><b>5198</b> queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td>4372 queries/second</td>
        <td><b>9678</b>* queries/second</td>
        <td>-</td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td><b>11217</b> queries/second</td>
        <td>9738 queries/second</td>
        <td>11086 queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td>9849 queries/second</td>
        <td>8784 queries/second</td>
        <td><b>11391</b> queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td><b>12907</b> queries/second</td>
        <td>12410 queries/second</td>
        <td>11429 queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td><b>13222</b> queries/second</td>
        <td>12696 queries/second</td>
        <td>11748 queries/second</td>
   </tr>
</table>

`* AutoQueryRow`: mymysql builds a single query string instead of using prepared statements

### Original Logs

#### Go 1.0.3
```
$ go version
go version go1.0.3
```

```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 4.90898s [ 2037 queries/second ]
PreparedQuery: 3.415631s [ 2928 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 25.992804s [ 3847 queries/second ]
SimpleQueryRow: 26.018075s [ 3843 queries/second ]
PreparedQueryRow: 11.359661s [ 8803 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 9.937693s [ 10063 queries/second ]
PreparedExec: 9.673266s [ 10338 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 6.090524s [ 1642 queries/second ]
PreparedQuery: 4.447945s [ 2248 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 30.092542s [ 3323 queries/second ]
SimpleQueryRow: 29.399428s [ 3401 queries/second ]
PreparedQueryRow: 12.43072s [ 8045 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 10.555367s [ 9474 queries/second ]
PreparedExec: 10.172413s [ 9831 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 4.721605s [ 2118 queries/second ]
PreparedQuery: 3.251801s [ 3075 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 27.436285s [ 3645 queries/second ]
SimpleQueryRow: 26.077592s [ 3835 queries/second ]
PreparedQueryRow: 11.229181s [ 8905 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 9.97262s [ 10027 queries/second ]
PreparedExec: 9.729652s [ 10278 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 6.130481s [ 1631 queries/second ]
PreparedQuery: 4.451507s [ 2246 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 28.611437s [ 3495 queries/second ]
SimpleQueryRow: 28.730313s [ 3481 queries/second ]
PreparedQueryRow: 12.427371s [ 8047 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 10.319815s [ 9690 queries/second ]
PreparedExec: 10.147441s [ 9855 queries/second ]
```


#### Go tip
```
$ go version
go version devel +786e094255c9 Tue Mar 19 07:08:26 2013 +0100 linux/amd64
```

```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.660608864s [ 3759 queries/second ]
PreparedQuery: 2.416900888s [ 4138 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 22.660756965s [ 4413 queries/second ]
SimpleQueryRow: 8.815792086s [ 11343 queries/second ]
PreparedQueryRow: 10.028589997s [ 9971 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 7.676686034s [ 13026 queries/second ]
PreparedExec: 7.508063974s [ 13319 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.724476606s [ 2685 queries/second ]
PreparedQuery: 3.301542098s [ 3029 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 10.320880617s [ 9689 queries/second ]
SimpleQueryRow: 10.252676766s [ 9754 queries/second ]
PreparedQueryRow: 11.309210151s [ 8842 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 8.017343577s [ 12473 queries/second ]
PreparedExec: 7.853554703s [ 12733 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.671647612s [ 3743 queries/second ]
PreparedQuery: 2.43836213s [ 4101 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 22.872479121s [ 4372 queries/second ]
SimpleQueryRow: 8.914766029s [ 11217 queries/second ]
PreparedQueryRow: 10.153325208s [ 9849 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 7.747658594s [ 12907 queries/second ]
PreparedExec: 7.563278955s [ 13222 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.726763041s [ 2683 queries/second ]
PreparedQuery: 3.312916071s [ 3018 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 10.33306149s [ 9678 queries/second ]
SimpleQueryRow: 10.269378078s [ 9738 queries/second ]
PreparedQueryRow: 11.38467273s [ 8784 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 8.058210282s [ 12410 queries/second ]
PreparedExec: 7.876202872s [ 12696 queries/second ]
```

### Java
```
java version "1.7.0_05"
Java(TM) SE Runtime Environment (build 1.7.0_05-b06)
Java HotSpot(TM) 64-Bit Server VM (build 23.1-b03, mixed mode)
```
#### Run 1
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.009s [ 3323 queries/second ]
PreparedQuery: 1.921s [ 5206 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 9.033s [ 11071 queries/second ]
PreparedQueryRow: 8.731s [ 11453 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 8.66s [ 11547 queries/second ]
PreparedExec: 8.452s [ 11832 queries/second ]
```

#### Run 2
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.004s [ 3329 queries/second ]
PreparedQuery: 1.925s [ 5195 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 9.001s [ 11110 queries/second ]
PreparedQueryRow: 8.729s [ 11456 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 8.73s [ 11455 queries/second ]
PreparedExec: 8.427s [ 11867 queries/second ]
```

#### Run 3
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.029s [ 3301 queries/second ]
PreparedQuery: 1.921s [ 5206 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 9.005s [ 11105 queries/second ]
PreparedQueryRow: 8.731s [ 11453 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 8.744s [ 11436 queries/second ]
PreparedExec: 8.475s [ 11799 queries/second ]
```

#### Run 4
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.016s [ 3316 queries/second ]
PreparedQuery: 1.924s [ 5198 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 9.02s [ 11086 queries/second ]
PreparedQueryRow: 8.779s [ 11391 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 8.75s [ 11429 queries/second ]
PreparedExec: 8.512s [ 11748 queries/second ]
```
