.PHONEY: test
test:
	go install
	colorizer > test.txt
	code test.txt
