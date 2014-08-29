#!/bin/sh

gox -os="darwin linux windows" -arch="386 amd64" -output="pkg/{{.Dir}}_{{.OS}}_{{.Arch}}"
ghr -u ryotarai $1 pkg/
