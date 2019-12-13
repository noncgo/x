#!/bin/sh -e
go build -o _example.app/Contents/MacOS/app github.com/tie/dyld/cmd/app
open _example.app
