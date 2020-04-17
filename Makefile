.PHONY: help
help:
	@echo "make help    -- ヘルプを表示する"
	@echo "make deps    -- 依存ライブラリをダウンロードする"
	@echo "make run     -- 開発サーバーを起動する"
	@echo "make fmt     -- ソースコードを整形する"
	@echo "make lint    -- golint を実行する"

.PHONY: deps
deps:
	go get -d
	go mod tidy

.PHONY: run
run: deps
	bee run -downdoc=true -gendoc=true

.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: lint
lint:
	golint -set_exit_code ./...
