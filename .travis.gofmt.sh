#!/bin/bash

#auto format code
if [ -n "$(gofmt -1 .)" ]; then
  gofmt -d
  echo "Code formatting complete"
  exit 1
fi