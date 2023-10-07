# ==================================================================================== #
# DEVELOPMENT
# ==================================================================================== #

## environment variables
export UNDERSHORTS_REDIS_URL=redis://:PASSWORD@localhost:6379

## redis: launch a redis dev server inside a docker container
.PHONY: redis
redis:
	docker run -d -p 6379:6379 --name undershorts-redis redis sh -c "redis-server --requirepass PASSWORD"

## build: build the application
.PHONY: build
build:
	go build -o=./tmp/bin/undershorts ./cmd/undershorts/main.go

## run: run the application
.PHONY: run
run: build
	./tmp/bin/undershorts

## run/live: run the application with reloading on file changes
.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "./tmp/bin/undershorts" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

