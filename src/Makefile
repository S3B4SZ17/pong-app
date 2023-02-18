.PHONY: build

build:
	@docker build --build-arg ARG_CONFIG_FILE="/pong-app/config.yaml" --no-cache -t pong-app:1.0.0 .

run:
	@echo "=========== App is startin up ===="
	@echo "CONFIG ==== 😁 Exporting environemnt variables"
	@echo "SUCCESS ===  ✔ Environment variables exported"
	@echo "INIT ======  ⚡ Running server"
	@go run main.go