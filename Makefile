TAGS = 

ifeq ($(TAG),)
	TAGS=v0.0.1
else
	TAGS=${TAG}
endif

.PHONY: clean static binary docker image archive export

default: clean binary

image: clean binary docker

binary:
	echo "Building sentinel_tunnel binary..."
	CGO_ENABLED=0 go build -v -ldflags "-X sentinel_tunnel/config.VERSION=${TAGS} -s -w"

docker:
	docker build -t k8s-test/sentinel_tunnel:${TAGS} .
	rm -rf sentinel_tunnel

archive:
	docker save `docker images -a --format {{.Repository}}:{{.Tag}} | grep "k8s-test/sentinel_tunnel"` | gzip > sentinel_tunnel-${TAGS}.tgz

export: image archive

clean:
	rm -rf sentinel_tunnel *.tgz
	docker images -a --format {{.Repository}}:{{.Tag}} | grep "k8s-test/sentinel_tunnel" | xargs -i docker rmi {} >/dev/null 2>&1
	docker images -a --format {{.Repository}}:{{.ID}} | grep "none" | awk -F":" '{print $$2}' | xargs -i docker rmi {} >/dev/null 2>&1 || true
