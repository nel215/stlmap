test:
	go test ./... -v -race
bench:
	go build stlmap.go stlmap_test.go && go test ./... -cpuprofile cpu.out -benchtime 3s -bench .
pprof:
	go tool pprof stlmap.test cpu.out
