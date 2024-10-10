#!/bin/bash

CURDIR=$(cd $(dirname $0); pwd)

SVC_NAME=credit-server
BinaryName=credit-server

echo "$CURDIR/bin/${BinaryName} -conf=?"
exec $CURDIR/bin/${BinaryName}