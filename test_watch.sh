#!/usr/bin/env bash

if [ -z "$1" ]; then
  nix shell nixpkgs#watchexec -c watchexec --exts go go test ./... -v
else
  nix shell nixpkgs#watchexec -c watchexec --exts go go test -run "Test_all/day$1" -v
fi
