# go-logger-benchmark


# Benchmark
利用go testing的benchmark功能，对列出的三方库进行压力测试，利用testing的Parallel进行并发测试，分别把cpu参数设置为1，2，4进行测试。

测试脚本： https://github.com/senses2008/go-logger-benchmark

zap的性能很突出，zerolog也符合“Zero Allocation JSON Logger”的定义。

## 写文本格式日志
|库名	| op time(cpu-1) |	allocs size(cpu-1) |	alloc times(cpu-1) |	op time(cpu-2) |	allocs size(cpu-2)|	alloc count(cpu-2)|	op time(cpu-4)|	allocs size(cpu-4)|	alloc count(cpu-4)|
|-- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
|logrus|10850 ns/op|	368 B/op|	11 allocs/op|	13538 ns/op |	368 B/op|	11 allocs/op|	13925 ns/op|	368 B/op|	11 allocs/op|
|zap	|912 ns/op|	0 B/op|	0 allocs/op |	510 ns/op|	0 B/op|	0 allocs/op|	268 ns/op|	0 B/op|	0 allocs/op|
|go-logging|	10143 ns/op|	920 B/op|	15 allocs/op|	13608 ns/op|	920 B/op|	15 allocs/op|	13342 ns/op|	920 B/op|	15 allocs/op|
|zerolog|	8294 ns/op|	0 B/op|	0 allocs/op|	10622 ns/op|	0 B/op|	0 allocs/op|	10414 ns/op|	0 B/op|	0 allocs/op|
|seelog|	10315 ns/op|	440 B/op|	11 allocs/op|	13015 ns/op|	440 B/op|	11 allocs/op|	12318 ns/op|	440 B/op|	11 allocs/op|
|log15	|12184 ns/op|	840 B/op|	14 allocs/op|	16081 ns/op|	840 B/op|	14 allocs/op|	15236 ns/op|	840 B/op|	14 allocs/op|
|go-kit	|7289 ns/op|	96 B/op|	2 allocs/op|	9025 ns/op|	96 B/op|	2 allocs/op|	9516 ns/op|	96 B/op	|2 allocs/op|

## 写Json格式日志
|库名	| op time(cpu-1) |	allocs size(cpu-1) |	alloc times(cpu-1) |	op time(cpu-2) |	allocs size(cpu-2)|	alloc count(cpu-2)|	op time(cpu-4)|	allocs size(cpu-4)|	alloc count(cpu-4)|
|-- | --- | --- | --- | --- | --- | --- | --- | --- | --- |
|logrus|	16302 ns/op|	2144 B/op|	32 allocs/op|	20125 ns/op|	2144 B/op|	32 allocs/op|	19650 ns/op|	2144 B/op|	32 allocs/op|
|zap|	1159 ns/op|	192 B/op|	1 allocs/op|	630 ns/op|	192 B/op|	1 allocs/op|	372 ns/op|	192 B/op	|1 allocs/op|
|zerolog	|9050 ns/op|	0 B/op	|0 allocs/op|	11362 ns/op|	0 B/op|	0 allocs/op|	12227 ns/op|	0 B/op|	0 allocs/op|
|log15	|19535 ns/op|	2240 B/op|	32 allocs/op|	23066 ns/op|	2240 B/op|	32 allocs/op|	23577 ns/op|	2240 B/op|	32 allocs/op|
|go-kit|	9221 ns/op|	298 B/op|	5 allocs/op|	12312 ns/op|	298 B/op|	24 allocs/op|	15428 ns/op	|298 B/op	|5 allocs/op|
