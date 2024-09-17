default: testv1 testv2

testv1: buildv1
	cd v1; go test ./...

buildv1: gogetv1
	cd v1; go build ./...

gogetv1:
	cd v1; go get ./...

testv2: buildv2
	cd v2; go test ./...

buildv2: gogetv2
	cd v2; go build ./...

gogetv2:
	cd v2; go get ./...