.PHONY: deps install local run build

deps:
	@printf "Installing all dependencies\n\n"
	go get -u -v github.com/githubnemo/CompileDaemon
	@printf "Installing glide dependencies\n\n"
	glide install

install:
	@printf "Installing API\n\n"
	go install -v github.com/disiqueira/frango/src/search/cmd/search

local:
	@printf "Installing Pre-commit\n\n"
	brew install pre-commit

	@printf "Installing Pre-commit hooks\n\n"
	pre-commit install

run:
	@printf "Running API\n\n"
	docker-compose --project-name "frango" up -d

build:
	@printf "Running API\n\n"
	docker-compose build
