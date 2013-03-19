#!/bin/bash

# first download connector jar into test directory
# depends on go, java, awk, sed, tee, sort, tail, cat, echo, read and exit
# all except go and java should be present on anything Unixy (only tested on OS X)

# !!! this is pretty raw and doesn't do any checks
# !!! read and understand it before you run it

unify-format() {
	sed -n '/BENCHMARKING/{s/  *[^ ]* \(.*\) \[run \([12]\).*/BENCHMARK|\1|\2/p;};/.*:.*/{s/\([^:]*\):[^\[]*\[ \([0-9]*\).*/\1|\2/p;}'\
	 | awk 'BEGIN{FS="|";OFS="\t";BENCHMARKED="Java";BENCHRUN=1}{if(NF==3){BENCHMARKED=$2;BENCHRUN=$3}else{print BENCHMARKED, BENCHRUN, $1, $2}}'
}

benchmark-java() {
	CONNECTOR=$(ls mysql-connector-java-*-bin.jar|tail -1)
	java -cp .:$CONNECTOR SQLBenchmark
}

GO_VERSION=$(go version|sed 's/^[^0-9+]*\([+.0-9a-f]*\).*/\1/')

GO_LOG=go-$GO_VERSION.log
JAVA_LOG_1=java-run1.log
JAVA_LOG_2=java-run2.log

RESULT=all-go$GO_VERSION.tsv

TEMP_LOG=java+go.log.tsv

# nag the user with a dialog...
echo -n "Do you understand what this does? Ctrl-C if not, else Enter."; read || exit

# compile java
javac -cp $(ls mysql-connector-java-*-bin.jar|tail -1) SQLBenchmark.java

# empty temp log
> $TEMP_LOG

echo "Benchmarking Java, run 1 (raw in $JAVA_LOG_1)"
benchmark-java | tee $JAVA_LOG_1\
 | unify-format >> $TEMP_LOG

echo "Benchmarking Go (raw in $GO_LOG)"
go run sqlBenchmark.go | tee $GO_LOG\
 | unify-format >> $TEMP_LOG

echo "Benchmarking Java, run 2 (raw in $JAVA_LOG_2)"
benchmark-java | tee $JAVA_LOG_2\
 | unify-format >> $TEMP_LOG

echo "Sorting results (will be in $RESULT)"
sort -t "	" -k 3,3 -k 4,4n < $TEMP_LOG > $RESULT

# cleanup
rm $TEMP_LOG
