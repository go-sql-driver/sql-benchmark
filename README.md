#SQL-Benchmark

A synthetic benchmark to compare the performance of various sql-drivers for Go's database/sql package

## Results (2013-11-02)
### Older Results
* [March 02, 2013](results_2013-03-02.md)
* [February 26, 2013](results_2013-02-26.md)

### Contributed Results
* [March 19, 2013](results-osx_2013-03-19.md) // OS X 10.8.3 / MySQL 5.6.10 by [Arne Hormann](https://github.com/arnehormann)
* [March 19, 2013](results-linux_2013-03-19.md) // Linux / MySQL 5.5.27 by [Arne Hormann](https://github.com/arnehormann)


### Setup
* Intel Core i5-2500K (4x 3.30 GHz), 8 GB RAM
* MySQL 5.5.34
* Ubuntu 13.10 x64 (Linux)
* Go 1.2rc2 linux/amd64
* [Go-MySQL-Driver v1.1 and v1.0.2](https://github.com/go-sql-driver/mysql) and [MyMySQL v1.5.3](https://github.com/ziutek/mymysql)

#### MySQL
* fresh & clean install via `apt-get install mysql-server`
* server version: 5.5.34-0ubuntu0.13.10.1
* connected via Unix Domain Socket (`/var/run/mysqld/mysqld.sock`)
* additional config:

```
thread_cache_size = 16
query_cache_limit = 2M
query_cache_size = 128M
```

#### Linux
* CPU frequency governor set to 'performance'

### Notes
* We did a complete rewrite of the benchmark suite for this round. The results are not comparable to previous rounds.
* The results should now be much more accurate and vary less.
* The benchmark suite now includes concurrent tests. These are probably the most interesting tests.
* The memory footprint (per query) is now also measured.
* Before each test the Garbage Collector is manually run to eliminate influence of garbage from previous tests.
* The benchmarks are designed to minimize response latency from the server. We try to compare driver performance here and not to benchmark the MySQL Server ;)
* The memory footprint includes allocations by the database/sql package. In fact, for Go-MySQL-Driver 1.1 this is the large part, as you can see [in this memory profile](http://files.julienschmidt.com/public/go/sql-benchmark/mysql.PreparedQueryConcurrent4.mem.svg). The `mysql-rows.(*mysqlRows).Columns` block will hopefully eliminated after the Go 1.2 release [with this patch](https://codereview.appspot.com/17580043/).


```
$ go run main.go benchmarks.go
Warming up mymysql godrv...
Warming up Go-MySQL-Driver 1.0...
Warming up Go-MySQL-Driver 1.1...

Run Benchmarks...

SimpleExec 250000 iterations
mymysql godrv
 3.844984027s 	    65020 queries/sec	    7 allocs/query	    351 B/query
 3.832080666s 	    65239 queries/sec	    7 allocs/query	    351 B/query
 3.817775996s 	    65483 queries/sec	    7 allocs/query	    351 B/query
 -- avg 65247 qps;  median 65239 qps
Go-MySQL-Driver 1.0
 3.602866751s 	    69389 queries/sec	    4 allocs/query	    113 B/query
 3.599880038s 	    69447 queries/sec	    4 allocs/query	    113 B/query
 3.591707362s 	    69605 queries/sec	    4 allocs/query	    113 B/query
 -- avg 69480 qps;  median 69447 qps
Go-MySQL-Driver 1.1
 3.569295906s 	    70042 queries/sec	    3 allocs/query	    97 B/query
 3.578053109s 	    69870 queries/sec	    3 allocs/query	    97 B/query
 3.572950976s 	    69970 queries/sec	    3 allocs/query	    97 B/query
 -- avg 69961 qps;  median 69970 qps

PreparedExec 250000 iterations
mymysql godrv
 3.68239709s 	    67891 queries/sec	    8 allocs/query	    392 B/query
 3.685228641s 	    67838 queries/sec	    8 allocs/query	    392 B/query
 3.690438516s 	    67743 queries/sec	    8 allocs/query	    392 B/query
 -- avg 67824 qps;  median 67838 qps
Go-MySQL-Driver 1.0
 3.46375733s 	    72176 queries/sec	    6 allocs/query	    167 B/query
 3.463701957s 	    72177 queries/sec	    6 allocs/query	    167 B/query
 3.465739159s 	    72135 queries/sec	    6 allocs/query	    167 B/query
 -- avg 72163 qps;  median 72176 qps
Go-MySQL-Driver 1.1
 3.418759203s 	    73126 queries/sec	    5 allocs/query	    149 B/query
 3.418744609s 	    73126 queries/sec	    5 allocs/query	    149 B/query
 3.428505491s 	    72918 queries/sec	    5 allocs/query	    149 B/query
 -- avg 73057 qps;  median 73126 qps

SimpleQueryRow 250000 iterations
mymysql godrv
 5.802717679s 	    43083 queries/sec	    26 allocs/query	    1177 B/query
 5.803354897s 	    43079 queries/sec	    26 allocs/query	    1177 B/query
 5.800961587s 	    43096 queries/sec	    26 allocs/query	    1177 B/query
 -- avg 43086 qps;  median 43083 qps
Go-MySQL-Driver 1.0
 5.063697563s 	    49371 queries/sec	    13 allocs/query	    395 B/query
 5.056738087s 	    49439 queries/sec	    13 allocs/query	    395 B/query
 5.068211838s 	    49327 queries/sec	    13 allocs/query	    395 B/query
 -- avg 49379 qps;  median 49371 qps
Go-MySQL-Driver 1.1
 4.985966338s 	    50141 queries/sec	    12 allocs/query	    363 B/query
 4.942562121s 	    50581 queries/sec	    12 allocs/query	    363 B/query
 4.935405828s 	    50654 queries/sec	    12 allocs/query	    363 B/query
 -- avg 50458 qps;  median 50581 qps

PreparedQueryRow 250000 iterations
mymysql godrv
 6.245600097s 	    40028 queries/sec	    29 allocs/query	    1227 B/query
 6.256448195s 	    39959 queries/sec	    29 allocs/query	    1226 B/query
 6.238429435s 	    40074 queries/sec	    29 allocs/query	    1227 B/query
 -- avg 40020 qps;  median 40028 qps
Go-MySQL-Driver 1.0
 5.351447088s 	    46716 queries/sec	    17 allocs/query	    457 B/query
 5.340188319s 	    46815 queries/sec	    17 allocs/query	    457 B/query
 5.341302407s 	    46805 queries/sec	    17 allocs/query	    457 B/query
 -- avg 46779 qps;  median 46805 qps
Go-MySQL-Driver 1.1
 5.119802775s 	    48830 queries/sec	    14 allocs/query	    382 B/query
 5.100032785s 	    49019 queries/sec	    14 allocs/query	    382 B/query
 5.098174033s 	    49037 queries/sec	    14 allocs/query	    382 B/query
 -- avg 48962 qps;  median 49019 qps

PreparedQueryRowParam 250000 iterations
mymysql godrv
 6.878633952s 	    36344 queries/sec	    32 allocs/query	    1263 B/query
 6.9165302s 	    36145 queries/sec	    32 allocs/query	    1263 B/query
 6.888402455s 	    36293 queries/sec	    32 allocs/query	    1263 B/query
 -- avg 36261 qps;  median 36293 qps
Go-MySQL-Driver 1.0
 5.598532572s 	    44655 queries/sec	    21 allocs/query	    540 B/query
 5.591264361s 	    44713 queries/sec	    21 allocs/query	    540 B/query
 5.680375421s 	    44011 queries/sec	    21 allocs/query	    540 B/query
 -- avg 44459 qps;  median 44655 qps
Go-MySQL-Driver 1.1
 5.455894144s 	    45822 queries/sec	    15 allocs/query	    400 B/query
 5.461734165s 	    45773 queries/sec	    15 allocs/query	    400 B/query
 5.470208259s 	    45702 queries/sec	    15 allocs/query	    400 B/query
 -- avg 45766 qps;  median 45773 qps

EchoMixed5 250000 iterations
mymysql godrv
 8.773040358s 	    28496 queries/sec	    58 allocs/query	    2406 B/query
 8.596984382s 	    29080 queries/sec	    58 allocs/query	    2406 B/query
 8.657759663s 	    28876 queries/sec	    58 allocs/query	    2406 B/query
 -- avg 28817 qps;  median 28876 qps
Go-MySQL-Driver 1.0
 6.823789206s 	    36637 queries/sec	    32 allocs/query	    1093 B/query
 6.829033082s 	    36608 queries/sec	    32 allocs/query	    1092 B/query
 6.825584568s 	    36627 queries/sec	    32 allocs/query	    1092 B/query
 -- avg 36624 qps;  median 36627 qps
Go-MySQL-Driver 1.1
 6.212412594s 	    40242 queries/sec	    17 allocs/query	    644 B/query
 6.201953199s 	    40310 queries/sec	    17 allocs/query	    644 B/query
 6.227740014s 	    40143 queries/sec	    17 allocs/query	    644 B/query
 -- avg 40232 qps;  median 40242 qps

SelectLargeString 50000 iterations
mymysql godrv
 3.710482335s 	    13475 queries/sec	    26 allocs/query	    21688 B/query
 3.72813078s 	    13412 queries/sec	    26 allocs/query	    21688 B/query
 3.71555593s 	    13457 queries/sec	    26 allocs/query	    21688 B/query
 -- avg 13448 qps;  median 13457 qps
Go-MySQL-Driver 1.0
 3.382473045s 	    14782 queries/sec	    13 allocs/query	    10656 B/query
 3.38769337s 	    14759 queries/sec	    13 allocs/query	    10656 B/query
 3.366275121s 	    14853 queries/sec	    13 allocs/query	    10656 B/query
 -- avg 14798 qps;  median 14782 qps
Go-MySQL-Driver 1.1
 3.273157066s 	    15276 queries/sec	    12 allocs/query	    10608 B/query
 3.279537597s 	    15246 queries/sec	    12 allocs/query	    10608 B/query
 3.282647996s 	    15232 queries/sec	    12 allocs/query	    10608 B/query
 -- avg 15251 qps;  median 15246 qps

SelectPreparedLargeString 50000 iterations
mymysql godrv
 3.729359051s 	    13407 queries/sec	    31 allocs/query	    21752 B/query
 3.709282036s 	    13480 queries/sec	    31 allocs/query	    21752 B/query
 3.718004082s 	    13448 queries/sec	    31 allocs/query	    21752 B/query
 -- avg 13445 qps;  median 13448 qps
Go-MySQL-Driver 1.0
 3.30230472s 	    15141 queries/sec	    18 allocs/query	    10728 B/query
 3.287424383s 	    15209 queries/sec	    18 allocs/query	    10728 B/query
 3.290013005s 	    15198 queries/sec	    18 allocs/query	    10728 B/query
 -- avg 15183 qps;  median 15198 qps
Go-MySQL-Driver 1.1
 3.24534887s 	    15407 queries/sec	    15 allocs/query	    10632 B/query
 3.243725426s 	    15414 queries/sec	    15 allocs/query	    10632 B/query
 3.242123356s 	    15422 queries/sec	    15 allocs/query	    10632 B/query
 -- avg 15414 qps;  median 15414 qps

SelectLargeBytes 50000 iterations
mymysql godrv
 3.758672532s 	    13303 queries/sec	    26 allocs/query	    21688 B/query
 3.736869351s 	    13380 queries/sec	    26 allocs/query	    21688 B/query
 3.824912123s 	    13072 queries/sec	    26 allocs/query	    21688 B/query
 -- avg 13252 qps;  median 13303 qps
Go-MySQL-Driver 1.0
 3.288067334s 	    15207 queries/sec	    13 allocs/query	    10656 B/query
 3.289624359s 	    15199 queries/sec	    13 allocs/query	    10656 B/query
 3.28935065s 	    15201 queries/sec	    13 allocs/query	    10656 B/query
 -- avg 15202 qps;  median 15201 qps
Go-MySQL-Driver 1.1
 3.268010834s 	    15300 queries/sec	    12 allocs/query	    10608 B/query
 3.266943391s 	    15305 queries/sec	    12 allocs/query	    10608 B/query
 3.272072635s 	    15281 queries/sec	    12 allocs/query	    10608 B/query
 -- avg 15295 qps;  median 15300 qps

SelectPreparedLargeBytes 50000 iterations
mymysql godrv
 3.704054898s 	    13499 queries/sec	    31 allocs/query	    21752 B/query
 3.707255824s 	    13487 queries/sec	    31 allocs/query	    21752 B/query
 3.700564381s 	    13511 queries/sec	    31 allocs/query	    21752 B/query
 -- avg 13499 qps;  median 13499 qps
Go-MySQL-Driver 1.0
 3.293667657s 	    15181 queries/sec	    18 allocs/query	    10728 B/query
 3.281278474s 	    15238 queries/sec	    18 allocs/query	    10728 B/query
 3.284310328s 	    15224 queries/sec	    18 allocs/query	    10728 B/query
 -- avg 15214 qps;  median 15224 qps
Go-MySQL-Driver 1.1
 3.242004446s 	    15423 queries/sec	    15 allocs/query	    10632 B/query
 3.250628059s 	    15382 queries/sec	    15 allocs/query	    10632 B/query
 3.241874523s 	    15423 queries/sec	    15 allocs/query	    10632 B/query
 -- avg 15409 qps;  median 15423 qps

SelectLargeRaw 50000 iterations
mymysql godrv
 3.495386785s 	    14305 queries/sec	    24 allocs/query	    11419 B/query
 3.487362519s 	    14337 queries/sec	    24 allocs/query	    11419 B/query
 3.487310296s 	    14338 queries/sec	    24 allocs/query	    11419 B/query
 -- avg 14327 qps;  median 14337 qps
Go-MySQL-Driver 1.0
 3.085953618s 	    16202 queries/sec	    11 allocs/query	    397 B/query
 3.07629808s 	    16253 queries/sec	    11 allocs/query	    397 B/query
 3.078047696s 	    16244 queries/sec	    11 allocs/query	    397 B/query
 -- avg 16233 qps;  median 16244 qps
Go-MySQL-Driver 1.1
 3.056899867s 	    16356 queries/sec	    10 allocs/query	    348 B/query
 3.059649175s 	    16342 queries/sec	    10 allocs/query	    348 B/query
 3.059414835s 	    16343 queries/sec	    10 allocs/query	    348 B/query
 -- avg 16347 qps;  median 16343 qps

SelectPreparedLargeRaw 50000 iterations
mymysql godrv
 3.485924755s 	    14343 queries/sec	    29 allocs/query	    11483 B/query
 3.486648758s 	    14340 queries/sec	    29 allocs/query	    11483 B/query
 3.493145476s 	    14314 queries/sec	    29 allocs/query	    11483 B/query
 -- avg 14333 qps;  median 14340 qps
Go-MySQL-Driver 1.0
 3.079810302s 	    16235 queries/sec	    16 allocs/query	    473 B/query
 3.075175601s 	    16259 queries/sec	    16 allocs/query	    472 B/query
 3.080132071s 	    16233 queries/sec	    16 allocs/query	    473 B/query
 -- avg 16242 qps;  median 16235 qps
Go-MySQL-Driver 1.1
 3.034226559s 	    16479 queries/sec	    13 allocs/query	    374 B/query
 3.035156325s 	    16474 queries/sec	    13 allocs/query	    373 B/query
 3.035375902s 	    16472 queries/sec	    13 allocs/query	    373 B/query
 -- avg 16475 qps;  median 16474 qps

PreparedExecConcurrent1 500000 iterations
mymysql godrv
 7.412726791s 	    67452 queries/sec	    8 allocs/query	    392 B/query
 7.595283865s 	    65830 queries/sec	    8 allocs/query	    392 B/query
 7.582002016s 	    65946 queries/sec	    8 allocs/query	    392 B/query
 -- avg 66409 qps;  median 65946 qps
Go-MySQL-Driver 1.0
 6.921190296s 	    72242 queries/sec	    6 allocs/query	    167 B/query
 6.898444843s 	    72480 queries/sec	    6 allocs/query	    166 B/query
 6.899012692s 	    72474 queries/sec	    6 allocs/query	    166 B/query
 -- avg 72399 qps;  median 72474 qps
Go-MySQL-Driver 1.1
 6.842107045s 	    73077 queries/sec	    5 allocs/query	    149 B/query
 6.825865048s 	    73251 queries/sec	    5 allocs/query	    149 B/query
 6.831175178s 	    73194 queries/sec	    5 allocs/query	    149 B/query
 -- avg 73174 qps;  median 73194 qps

PreparedExecConcurrent2 500000 iterations
mymysql godrv
 3.44565992s 	    145110 queries/sec	    8 allocs/query	    392 B/query
 3.505009461s 	    142653 queries/sec	    8 allocs/query	    392 B/query
 3.446869762s 	    145059 queries/sec	    8 allocs/query	    392 B/query
 -- avg 144274 qps;  median 145059 qps
Go-MySQL-Driver 1.0
 2.90463481s 	    172139 queries/sec	    6 allocs/query	    167 B/query
 2.899396938s 	    172450 queries/sec	    6 allocs/query	    166 B/query
 2.90158597s 	    172320 queries/sec	    6 allocs/query	    166 B/query
 -- avg 172303 qps;  median 172320 qps
Go-MySQL-Driver 1.1
 2.819298666s 	    177349 queries/sec	    5 allocs/query	    149 B/query
 2.828106403s 	    176797 queries/sec	    5 allocs/query	    149 B/query
 2.82974316s 	    176694 queries/sec	    5 allocs/query	    149 B/query
 -- avg 176947 qps;  median 176797 qps

PreparedExecConcurrent4 500000 iterations
mymysql godrv
 3.42759382s 	    145875 queries/sec	    8 allocs/query	    392 B/query
 3.434183403s 	    145595 queries/sec	    8 allocs/query	    392 B/query
 3.422888817s 	    146075 queries/sec	    8 allocs/query	    392 B/query
 -- avg 145848 qps;  median 145875 qps
Go-MySQL-Driver 1.0
 2.955758093s 	    169161 queries/sec	    6 allocs/query	    167 B/query
 2.956442816s 	    169122 queries/sec	    6 allocs/query	    167 B/query
 2.956962873s 	    169092 queries/sec	    6 allocs/query	    167 B/query
 -- avg 169125 qps;  median 169122 qps
Go-MySQL-Driver 1.1
 2.902781851s 	    172249 queries/sec	    5 allocs/query	    149 B/query
 2.893675059s 	    172791 queries/sec	    5 allocs/query	    149 B/query
 2.881353501s 	    173530 queries/sec	    5 allocs/query	    149 B/query
 -- avg 172856 qps;  median 172791 qps

PreparedExecConcurrent8 500000 iterations
mymysql godrv
 3.500235814s 	    142848 queries/sec	    8 allocs/query	    392 B/query
 3.507999161s 	    142531 queries/sec	    8 allocs/query	    392 B/query
 3.497887883s 	    142943 queries/sec	    8 allocs/query	    392 B/query
 -- avg 142774 qps;  median 142848 qps
Go-MySQL-Driver 1.0
 3.026177405s 	    165225 queries/sec	    6 allocs/query	    167 B/query
 3.026883385s 	    165186 queries/sec	    6 allocs/query	    167 B/query
 3.020821566s 	    165518 queries/sec	    6 allocs/query	    167 B/query
 -- avg 165310 qps;  median 165225 qps
Go-MySQL-Driver 1.1
 2.954903265s 	    169210 queries/sec	    5 allocs/query	    150 B/query
 2.965571144s 	    168602 queries/sec	    5 allocs/query	    150 B/query
 2.96929521s 	    168390 queries/sec	    5 allocs/query	    150 B/query
 -- avg 168734 qps;  median 168602 qps

PreparedExecConcurrent16 500000 iterations
mymysql godrv
 3.722864558s 	    134305 queries/sec	    8 allocs/query	    393 B/query
 3.701155651s 	    135093 queries/sec	    8 allocs/query	    393 B/query
 3.708896696s 	    134811 queries/sec	    8 allocs/query	    393 B/query
 -- avg 134736 qps;  median 134811 qps
Go-MySQL-Driver 1.0
 3.211510336s 	    155690 queries/sec	    6 allocs/query	    167 B/query
 3.204158577s 	    156047 queries/sec	    6 allocs/query	    167 B/query
 3.189459753s 	    156766 queries/sec	    6 allocs/query	    167 B/query
 -- avg 156168 qps;  median 156047 qps
Go-MySQL-Driver 1.1
 3.128369303s 	    159828 queries/sec	    5 allocs/query	    150 B/query
 3.11130211s 	    160704 queries/sec	    5 allocs/query	    150 B/query
 3.139291582s 	    159272 queries/sec	    5 allocs/query	    150 B/query
 -- avg 159935 qps;  median 159828 qps

PreparedQueryConcurrent1 500000 iterations
mymysql godrv
 14.533103278s 	    34404 queries/sec	    39 allocs/query	    1555 B/query
 14.340936029s 	    34865 queries/sec	    39 allocs/query	    1555 B/query
 14.507614095s 	    34465 queries/sec	    39 allocs/query	    1555 B/query
 -- avg 34578 qps;  median 34465 qps
Go-MySQL-Driver 1.0
 11.62010633s 	    43029 queries/sec	    24 allocs/query	    645 B/query
 11.615208271s 	    43047 queries/sec	    24 allocs/query	    645 B/query
 11.656044242s 	    42896 queries/sec	    24 allocs/query	    645 B/query
 -- avg 42991 qps;  median 43029 qps
Go-MySQL-Driver 1.1
 10.998830961s 	    45459 queries/sec	    17 allocs/query	    476 B/query
 11.014426678s 	    45395 queries/sec	    17 allocs/query	    476 B/query
 10.979119615s 	    45541 queries/sec	    17 allocs/query	    476 B/query
 -- avg 45465 qps;  median 45459 qps

PreparedQueryConcurrent2 500000 iterations
mymysql godrv
 7.831652593s 	    63843 queries/sec	    39 allocs/query	    1553 B/query
 7.842288669s 	    63757 queries/sec	    39 allocs/query	    1553 B/query
 7.797480045s 	    64123 queries/sec	    39 allocs/query	    1553 B/query
 -- avg 63908 qps;  median 63843 qps
Go-MySQL-Driver 1.0
 5.278819724s 	    94718 queries/sec	    24 allocs/query	    645 B/query
 5.291427353s 	    94492 queries/sec	    24 allocs/query	    645 B/query
 5.278954458s 	    94716 queries/sec	    24 allocs/query	    645 B/query
 -- avg 94642 qps;  median 94716 qps
Go-MySQL-Driver 1.1
 4.830366589s 	    103512 queries/sec	    17 allocs/query	    476 B/query
 4.831139401s 	    103495 queries/sec	    17 allocs/query	    476 B/query
 4.834301429s 	    103428 queries/sec	    17 allocs/query	    476 B/query
 -- avg 103478 qps;  median 103495 qps

PreparedQueryConcurrent4 500000 iterations
mymysql godrv
 7.951862714s 	    62878 queries/sec	    39 allocs/query	    1552 B/query
 7.935456591s 	    63008 queries/sec	    39 allocs/query	    1553 B/query
 8.00026767s 	    62498 queries/sec	    39 allocs/query	    1553 B/query
 -- avg 62795 qps;  median 62878 qps
Go-MySQL-Driver 1.0
 5.378643813s 	    92960 queries/sec	    24 allocs/query	    644 B/query
 5.332978832s 	    93756 queries/sec	    24 allocs/query	    644 B/query
 5.360460591s 	    93276 queries/sec	    24 allocs/query	    644 B/query
 -- avg 93331 qps;  median 93276 qps
Go-MySQL-Driver 1.1
 4.926703513s 	    101488 queries/sec	    17 allocs/query	    476 B/query
 4.937515591s 	    101266 queries/sec	    17 allocs/query	    476 B/query
 4.915720468s 	    101714 queries/sec	    17 allocs/query	    476 B/query
 -- avg 101489 qps;  median 101488 qps

PreparedQueryConcurrent8 500000 iterations
mymysql godrv
 7.971771249s 	    62721 queries/sec	    39 allocs/query	    1554 B/query
 8.008622996s 	    62433 queries/sec	    39 allocs/query	    1554 B/query
 8.016234752s 	    62373 queries/sec	    39 allocs/query	    1553 B/query
 -- avg 62509 qps;  median 62433 qps
Go-MySQL-Driver 1.0
 5.409094898s 	    92437 queries/sec	    24 allocs/query	    643 B/query
 5.407859645s 	    92458 queries/sec	    24 allocs/query	    643 B/query
 5.427255858s 	    92128 queries/sec	    24 allocs/query	    643 B/query
 -- avg 92341 qps;  median 92437 qps
Go-MySQL-Driver 1.1
 4.9584675s 	    100838 queries/sec	    17 allocs/query	    476 B/query
 4.977183317s 	    100458 queries/sec	    17 allocs/query	    475 B/query
 4.957529888s 	    100857 queries/sec	    17 allocs/query	    476 B/query
 -- avg 100718 qps;  median 100838 qps

PreparedQueryConcurrent16 500000 iterations
mymysql godrv
 8.11024536s 	    61650 queries/sec	    39 allocs/query	    1551 B/query
 8.04783136s 	    62129 queries/sec	    39 allocs/query	    1551 B/query
 8.105189303s 	    61689 queries/sec	    39 allocs/query	    1551 B/query
 -- avg 61823 qps;  median 61689 qps
Go-MySQL-Driver 1.0
 5.643842747s 	    88592 queries/sec	    24 allocs/query	    643 B/query
 5.629226543s 	    88822 queries/sec	    24 allocs/query	    642 B/query
 5.642114369s 	    88619 queries/sec	    24 allocs/query	    643 B/query
 -- avg 88678 qps;  median 88619 qps
Go-MySQL-Driver 1.1
 5.166585448s 	    96776 queries/sec	    17 allocs/query	    475 B/query
 5.138424004s 	    97306 queries/sec	    17 allocs/query	    475 B/query
 5.156050123s 	    96973 queries/sec	    17 allocs/query	    475 B/query
 -- avg 97018 qps;  median 96973 qps

Finished... Total running time: 16m20.053874375s

```
