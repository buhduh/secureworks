# Secureworks Programming Assessment

## Instructions
$make

This should vendor, build, test, and deploy a docker container accessible at http://localhost:8080
The port is a variable that can be modified in Makefile
See scripts/test.sh for example curl requests

$make all 

Will show all available make endpoints.  For the purpose of this test no endpoints should be required.

$make test

To see what's available, much of it is cruft from development that I don't see the need to rip out.

## Dependencies

To aid in deployment the maxmind database has been previously downloaded and committed.

All other "dependencies" are pulled in dynamically.  That being said, here's a list of required runtime dependencies:

  * https://dev.maxmind.com/geoip/geoip2/geolite2/ # maxmind free database
  * github.com/mattn/go-sqlite3 # sqlite3 golang library
  * github.com/oschwald/maxminddb-golang # maxmind golang library

The maxmind golang library wasn't strictly necessary, parsing CSV was totally possible, it worked...

The following docker images were consumed during build:
  
  * keinos/sqlite3:latest as database
  * alpine:latest as maxmind
  * golang:1.11 as vendor
  * golang:1.11 as sourcer
  * ubuntu:devel as app

I would have preferred using alpine:latest but it requries disabling CGO to force correct building of the net package which breaks the maxmind(I think...) package.
Deploying against ubuntu:latest exposed an issue with glibc on Ubuntu 18.04, luckily Ubuntu 20.04 came with a compatible glibc binary.

All other golang libraries are from "core" golang.

Tested on Kubuntu 18.04 and macOS Sierra Version 10.12.6
