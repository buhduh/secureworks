#!/bin/bash

curl -X POST -d '{"username": "bob", "unix_timestamp": 3235, "event_uuid": "67ads7f", "ip_address": "234.234.234"}' http://localhost${PORT_STR}
