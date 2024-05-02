# protoから自動生成する
.PHONY: buf-gen
bufgen:
	buf generate

# bufのlintを実行する
.PHONY: buf-lint
buflint:
	buf lint

# buf.gen.ymlの変更をlockファイルに反映する
.PHONY: buf-update
buf-update:
	buf mod update

# bufのキャッシュをクリアする
.PHONY: buf-clear
buf-clear:
	buf mod clear-cache


# 現在は、go-migrateを使っているので、このコマンド(gormのAuto migrate)は不要
# //DBのマイグレーションをする
# .PHONY: migrate
# migrate:
# 	go run cmd/migration/main.go
