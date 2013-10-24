### Setup
* Intel Core i7 (2.20 GHz), 16 GB RAM
* OS X 10.8.3 (12D78); full disk encryption enabled
* MySQL 5.6.10 Source distribution (`brew install mysql` without configuration changes)
* Current [Go-MySQL-Driver](https://github.com/Go-SQL-Driver/MySQL) and [mymysql](https://github.com/ziutek/mymysql) versions as of March 19, 2013
* Java7 (JDBC) + MySQL Connector/J 5.1.24 for comparison

### Results for go 1.0.3
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/Go-SQL-Driver/MySQL">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.24</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>2873 queries/second</td>
        <td>2407 queries/second</td>
        <td><b>4088</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td><b>3843</b> queries/second</td>
        <td>3004 queries/second</td>
        <td>3369 queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td><b>4795</b> queries/second</td>
        <td>4623 queries/second</td>
        <td>- queries/second</td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td>4932 queries/second</td>
        <td>4800 queries/second</td>
        <td><b>8463</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td>8544 queries/second</td>
        <td><b>8714</b> queries/second</td>
        <td>5200 queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td>14332 queries/second</td>
        <td>14808 queries/second</td>
        <td><b>15620</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td>15118 queries/second</td>
        <td><b>15780</b> queries/second</td>
        <td>15454 queries/second</td>
    </tr>
</table>

### Results for go tip
<table>
    <tr>
        <th>Benchmark</th>
        <th><a href="https://github.com/Go-SQL-Driver/MySQL">Go-MySQL-Driver</a></th>
        <th><a href="https://github.com/ziutek/mymysql">mymysql godrv</a></th>
        <th><a href="http://dev.mysql.com/downloads/connector/j/">Java (JDBC) + MySQL Connector/J 5.1.24</a></th>
    </tr>
    <tr>
        <th>SimpleQuery</th>
        <td>3983 queries/second</td>
        <td>3303 queries/second</td>
        <td><b>4088</b> queries/second</td>
    </tr>
    <tr>
        <th>PreparedQuery</th>
        <td><b>4655</b> queries/second</td>
        <td>3758 queries/second</td>
        <td>3369 queries/second</td>
    </tr>
    <tr>
        <th>AutoQueryRow</th>
        <td>5755 queries/second</td>
        <td><b>8526</b>* queries/second</td>
        <td>- queries/second</td>
    </tr>
    <tr>
        <th>SimpleQueryRow</th>
        <td><b>8585</b> queries/second</td>
        <td>8571 queries/second</td>
        <td>8463 queries/second</td>
    </tr>
    <tr>
        <th>PreparedQueryRow</th>
        <td><b>9652</b> queries/second</td>
        <td>9166 queries/second</td>
        <td>5200 queries/second</td>
    </tr>
    <tr>
        <th>SimpleExec</th>
        <td><b>20784</b> queries/second</td>
        <td>20050 queries/second</td>
        <td>15620 queries/second</td>
    </tr>
    <tr>
        <th>PreparedExec</th>
        <td><b>23588</b> queries/second</td>
        <td>22753 queries/second</td>
        <td>15454 queries/second</td>
    </tr>
</table>

`* AutoQueryRow`: mymysql builds a single query string instead of using prepared statements

### Original Logs
#### Go 1.0.3
```
$ go version
go version go1.0.3
```
##### Run 1
```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.413593s [ 2929 queries/second ]
PreparedQuery: 2.593549s [ 3856 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 20.42619s [ 4896 queries/second ]
SimpleQueryRow: 19.975014s [ 5006 queries/second ]
PreparedQueryRow: 11.704284s [ 8544 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.977532s [ 14332 queries/second ]
PreparedExec: 6.614706s [ 15118 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 4.154453s [ 2407 queries/second ]
PreparedQuery: 3.328669s [ 3004 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 21.632837s [ 4623 queries/second ]
SimpleQueryRow: 20.809298s [ 4806 queries/second ]
PreparedQueryRow: 11.323829s [ 8831 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.736055s [ 14845 queries/second ]
PreparedExec: 6.290393s [ 15897 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.437407s [ 2909 queries/second ]
PreparedQuery: 2.562842s [ 3902 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 20.436384s [ 4893 queries/second ]
SimpleQueryRow: 20.08296s [ 4979 queries/second ]
PreparedQueryRow: 10.953022s [ 9130 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.738728s [ 14840 queries/second ]
PreparedExec: 6.263373s [ 15966 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 4.121502s [ 2426 queries/second ]
PreparedQuery: 3.289298s [ 3040 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 21.343546s [ 4685 queries/second ]
SimpleQueryRow: 20.832256s [ 4800 queries/second ]
PreparedQueryRow: 11.37943s [ 8788 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.673795s [ 14984 queries/second ]
PreparedExec: 6.299963s [ 15873 queries/second ]
```
##### Run 2
```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.480331s [ 2873 queries/second ]
PreparedQuery: 2.578976s [ 3878 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 20.853913s [ 4795 queries/second ]
SimpleQueryRow: 20.276176s [ 4932 queries/second ]
PreparedQueryRow: 10.959359s [ 9125 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.664772s [ 15004 queries/second ]
PreparedExec: 6.208022s [ 16108 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 4.125382s [ 2424 queries/second ]
PreparedQuery: 3.29141s [ 3038 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 21.116172s [ 4736 queries/second ]
SimpleQueryRow: 20.764883s [ 4816 queries/second ]
PreparedQueryRow: 11.40949s [ 8765 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.653944s [ 15029 queries/second ]
PreparedExec: 6.260656s [ 15973 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.450241s [ 2898 queries/second ]
PreparedQuery: 2.60198s [ 3843 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 20.501758s [ 4878 queries/second ]
SimpleQueryRow: 20.113424s [ 4972 queries/second ]
PreparedQueryRow: 10.811276s [ 9250 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.598204s [ 15156 queries/second ]
PreparedExec: 6.202888s [ 16122 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 4.111154s [ 2432 queries/second ]
PreparedQuery: 3.316876s [ 3015 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 21.104781s [ 4738 queries/second ]
SimpleQueryRow: 20.834429s [ 4800 queries/second ]
PreparedQueryRow: 11.475476s [ 8714 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.752991s [ 14808 queries/second ]
PreparedExec: 6.337248s [ 15780 queries/second ]
```

#### Go tip
```
$ go version
go version devel +786e094255c9 Tue Mar 19 07:08:26 2013 +0100 darwin/amd64
```

##### Run 1
```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.510925006s [ 3983 queries/second ]
PreparedQuery: 2.14814344s [ 4655 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.376778601s [ 5755 queries/second ]
SimpleQueryRow: 11.327457312s [ 8828 queries/second ]
PreparedQueryRow: 10.301616676s [ 9707 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.638879971s [ 21557 queries/second ]
PreparedExec: 4.178864187s [ 23930 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.020216718s [ 3311 queries/second ]
PreparedQuery: 2.660903352s [ 3758 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 11.575896249s [ 8639 queries/second ]
SimpleQueryRow: 11.574478322s [ 8640 queries/second ]
PreparedQueryRow: 10.824283133s [ 9238 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.867722732s [ 20543 queries/second ]
PreparedExec: 4.341147858s [ 23035 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.50137433s [ 3998 queries/second ]
PreparedQuery: 2.146637653s [ 4658 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.357114312s [ 5761 queries/second ]
SimpleQueryRow: 11.647742719s [ 8585 queries/second ]
PreparedQueryRow: 10.360052005s [ 9652 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.695827143s [ 21296 queries/second ]
PreparedExec: 4.177790515s [ 23936 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.991561306s [ 3343 queries/second ]
PreparedQuery: 2.643930236s [ 3782 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 11.573065959s [ 8641 queries/second ]
SimpleQueryRow: 11.557347126s [ 8653 queries/second ]
PreparedQueryRow: 10.780680308s [ 9276 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.955322367s [ 20180 queries/second ]
PreparedExec: 4.35608785s [ 22956 queries/second ]
```

##### Run 2
```
*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.492747789s [ 4012 queries/second ]
PreparedQuery: 2.140412159s [ 4672 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.099773537s [ 5848 queries/second ]
SimpleQueryRow: 11.079477674s [ 9026 queries/second ]
PreparedQueryRow: 10.330586657s [ 9680 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.615779532s [ 21665 queries/second ]
PreparedExec: 4.14832651s [ 24106 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 1]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.027637513s [ 3303 queries/second ]
PreparedQuery: 2.652219341s [ 3770 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 11.568385456s [ 8644 queries/second ]
SimpleQueryRow: 11.533978356s [ 8670 queries/second ]
PreparedQueryRow: 10.774763586s [ 9281 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.779220431s [ 20924 queries/second ]
PreparedExec: 4.321005243s [ 23143 queries/second ]


*************************************************************
   BENCHMARKING Go-MySQL-Driver [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.508798034s [ 3986 queries/second ]
PreparedQuery: 2.141992203s [ 4669 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 17.281003539s [ 5787 queries/second ]
SimpleQueryRow: 11.285806624s [ 8861 queries/second ]
PreparedQueryRow: 10.311076985s [ 9698 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.811444584s [ 20784 queries/second ]
PreparedExec: 4.239370382s [ 23588 queries/second ]


*************************************************************
   BENCHMARKING mymysql godrv [run 2]
*************************************************************

-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 3.006637533s [ 3326 queries/second ]
PreparedQuery: 2.654193854s [ 3768 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
AutoQueryRow: 11.72828808s [ 8526 queries/second ]
SimpleQueryRow: 11.667093527s [ 8571 queries/second ]
PreparedQueryRow: 10.90937567s [ 9166 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 4.987597721s [ 20050 queries/second ]
PreparedExec: 4.395009337s [ 22753 queries/second ]
```

### Java
```
$ java - version
java version "1.7.0_15"
Java(TM) SE Runtime Environment (build 1.7.0_15-b03)
Java HotSpot(TM) 64-Bit Server VM (build 23.7-b01, mixed mode)
```

##### Run 1
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.43s [ 4115 queries/second ]
PreparedQuery: 2.968s [ 3369 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 11.6s [ 8621 queries/second ]
PreparedQueryRow: 19.185s [ 5212 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.402s [ 15620 queries/second ]
PreparedExec: 6.471s [ 15454 queries/second ]
```

##### Run 2
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.446s [ 4088 queries/second ]
PreparedQuery: 2.914s [ 3432 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 11.816s [ 8463 queries/second ]
PreparedQueryRow: 19.201s [ 5208 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.362s [ 15718 queries/second ]
PreparedExec: 6.21s [ 16103 queries/second ]
```

##### Run 3
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.435s [ 4107 queries/second ]
PreparedQuery: 2.931s [ 3412 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 11.708s [ 8541 queries/second ]
PreparedQueryRow: 19.161s [ 5219 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.355s [ 15736 queries/second ]
PreparedExec: 6.235s [ 16038 queries/second ]
```

##### Run 4
```
-------------------------------------------------------------
   [10000 * Query 100 Rows]
-------------------------------------------------------------
SimpleQuery: 2.424s [ 4125 queries/second ]
PreparedQuery: 2.899s [ 3449 queries/second ]

-------------------------------------------------------------
   [100 * QueryRow] * 1000
-------------------------------------------------------------
SimpleQueryRow: 11.757s [ 8506 queries/second ]
PreparedQueryRow: 19.232s [ 5200 queries/second ]

-------------------------------------------------------------
   [100000 * Exec]
-------------------------------------------------------------
SimpleExec: 6.398s [ 15630 queries/second ]
PreparedExec: 6.378s [ 15679 queries/second ]
```
