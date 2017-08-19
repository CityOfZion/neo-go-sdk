BRANCH = "master"
VERSION = $(shell cat ./VERSION)

check-version:
	@echo "=> Checking if VERSION exists as Git tag..."
	(! git rev-list ${VERSION})

push-tag:
	git checkout ${BRANCH}
	git pull origin ${BRANCH}
	git tag ${VERSION}
	git push origin ${BRANCH} --tags

test:
	@go test $(shell glide nv) -cover

vet:
	@go vet $(shell glide nv)