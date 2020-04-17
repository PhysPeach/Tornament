help:
	@echo "make help    -- ヘルプを表示する"
	@echo "make fmt     -- ソースコードを整形する"
	@echo "make run     -- 開発サーバーを起動する"

fmt:
	go fmt .

run:
	bee run

.PHONY: help fmt run
