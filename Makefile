
analyze-heap-escape:
	go build -gcflags="-m" ./... 2>&1 | grep 'heap'
