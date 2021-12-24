#!/bin/bash
go install github.com/cespare/reflex@latest
~/go/bin/reflex -r '\.go' -s -- sh -c "go run server.go"