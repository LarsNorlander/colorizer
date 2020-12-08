.PHONEY: rbt
rbt:
	go run cmd/rbt/main.go > out.rbt.txt
	code out.rbt.txt

.PHONEY: mapper
mapper:
	go run cmd/mapper/main.go > out.mapper.txt
	code out.mapper.txt
