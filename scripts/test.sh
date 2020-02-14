#!/bin/bash

curl -X POST -d '{"username": "bob", "unix_timestamp": 3235, "event_uuid": "67ads7f", "ip_address": "91.207.175.104"}' http://localhost:8080
curl -X POST -d '{"username": "bob", "unix_timestamp": 3236, "event_uuid": "67ads7g", "ip_address": "91.207.175.104"}' http://localhost:8080
curl -X POST -d '{"username": "bob", "unix_timestamp": 3238, "event_uuid": "67ads7h", "ip_address": "91.207.175.104"}' http://localhost:8080
curl -X POST -d '{"username": "bob", "unix_timestamp": 3230, "event_uuid": "67ads7i", "ip_address": "91.207.175.104"}' http://localhost:8080
curl -X POST -d '{"username": "bob", "unix_timestamp": 3233, "event_uuid": "67ads7j", "ip_address": "91.207.175.104"}' http://localhost:8080
