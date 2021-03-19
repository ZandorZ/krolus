#!/bin/sh
rm ~/.krolus/dev/* -Rf
cd ./migrations
rm migrations
go build
./migrations
cd ..