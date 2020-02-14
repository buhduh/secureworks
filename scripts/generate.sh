#!/bin/bash
cat << EOF > $GEN_GO
/*
File automatically created with go generate, DO NOT MODIFY
*/
package constants

const (
  IP_DB string = "$CONT_IP_DB"
  SQL_DB       = "$SQL_DB"
  PORT         = "$PORT"
)

EOF
