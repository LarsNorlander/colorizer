.PHONEY: test
test:
	go run main.go > out.txt
	code out.txt
