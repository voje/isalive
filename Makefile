install-templ:
	go install github.com/a-h/templ/cmd/templ@latest

gen-templates:
	templ generate

air: gen-templates
	ISALIVE_SITES="https://www.google.com,http://localhost:8081,http://wikipedia.org" \
	air

build-docker-image:
	./scripts/build_docker_images.sh
