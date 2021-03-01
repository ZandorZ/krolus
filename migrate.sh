#!/bin/sh
rm ./db/* -Rf
cd ./migrations
rm migrations
go build
./migrations
cd ..