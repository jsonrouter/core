# RUNNING BENCHMARKS

You need to run relevant server(s) by cd /*platform* e.g cd /fasthttp
then running

go run serve.go  

-or-

go build && ./*platform*

and then in /benchmarks run 

go test -run=XXX -bench=.

 