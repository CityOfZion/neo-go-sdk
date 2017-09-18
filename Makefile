BRANCH = "master"
VERSION = $(shell cat ./VERSION)

build-cli:
	@go build .

check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

push-tag:
ifndef GITHUB_TOKEN
	$(error GITHUB_TOKEN needs to be set to run push-tag)
endif
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags
	curl -sL https://git.io/goreleaser | GITHUB_TOKEN=${GITHUB_TOKEN} bash

test:
	@go test $(shell glide nv) -cover

vet:
	@go vet $(shell glide nv)