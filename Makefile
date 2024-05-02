# protoから自動生成する
.PHONY: bufgen
bufgen:
	buf generate


# 現在は、go-migrateを使っているので、このコマンド(gormのAuto migrate)は不要
# //DBのマイグレーションをする
# .PHONY: migrate
# migrate:
# 	go run cmd/migration/main.go
