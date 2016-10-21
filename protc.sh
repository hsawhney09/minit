#!/usr/bin/env bash

ROOTPATH=`dirname $0`
PROTO_PATH=${ROOTPATH}/../../../:${ROOTPATH}/:${ROOTPATH}/vendor/
protoc --proto_path="${PROTO_PATH}" --gogo_out=. ${ROOTPATH}/mesos/proto/*.proto
protoc --proto_path="${PROTO_PATH}" --gogo_out=. ${ROOTPATH}/mesos/proto/scheduler/*.proto
#protoc --proto_path="${PROTO_PATH}" --gogo_out=. ${ROOTPATH}/mesos/proto/messages/*.proto
#protoc --proto_path="${PROTO_PATH}" --gogo_out=. ./executor/*.proto