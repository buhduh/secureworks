#!/bin/bash

#I don't care if this fails
#wrap it in a script
docker kill $1 > /dev/null 2>&1
exit 0
