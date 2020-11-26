A short demo showing that GoLang channels are non-deterministic when writing. But, a `close(chan)` is in terms of all readers receiving the close "signal". 

```shell
$ go run test.go 
Configuring A() and B() go-routines as stop channel readers...
Sending stop <- struct{}{}...
B

Configuring A() and B() again...
Calling close(stop)...
A
A
B


$ go run test.go 
Configuring A() and B() go-routines as stop channel readers...
Sending stop <- struct{}{}...
A

Configuring A() and B() again...
Calling close(stop)...
B
A
B
```

References:
- <https://golang.org/ref/spec#Select_statements>
- <https://stackoverflow.com/questions/31539804/multiple-receivers-on-a-single-channel-who-gets-the-data>
