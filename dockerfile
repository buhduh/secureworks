######database######
FROM keinos/sqlite3:latest as database

ARG SQL_DB=${SQL_DB}
ARG CREATE_DB_SQL=${CREATE_DB_SQL}

COPY scripts/createdb.sql ${CREATE_DB_SQL}

RUN sqlite3 ${SQL_DB} < ${CREATE_DB_SQL}

######maxmind######
FROM alpine:latest as maxmind
ARG HOST_IP_DB=${HOST_IP_DB}
ARG CONT_IP_DB=${CONT_IP_DB}
COPY ${HOST_IP_DB} ${CONT_IP_DB}

######vendor######
FROM golang:1.11 as vendor

ARG VENDOR_DIR=${VENDOR_DIR}
ARG APP_SRC=${APP_SRC}

COPY ${APP_SRC} ${VENDOR_DIR}/src
WORKDIR ${VENDOR_DIR}/src
RUN go mod init secureworks
RUN go mod vendor

######sourcer######
FROM golang:1.11 as sourcer

ARG CONT_BIN=${CONT_BIN}
ARG GEN_GO=${GEN_GO}
ARG SQL_DB=${SQL_DB}
ARG CONT_IP_DB=${CONT_IP_DB}
ARG CONT_GEN_SCRIPT=${CONT_GEN_SCRIPT}
ARG APP_SRC=${APP_SRC}
ARG VENDOR_DIR=${VENDOR_DIR}
ARG CREATE_DB_SQL=${CREATE_DB_SQL}
ARG PORT=${PORT}

ENV GEN_GO        ${GEN_GO}
ENV CONT_IP_DB    ${CONT_IP_DB}

COPY --from=vendor ${VENDOR_DIR}/src /go/src
COPY scripts/generate.sh ${CONT_GEN_SCRIPT}

RUN chmod +x ${CONT_GEN_SCRIPT}
RUN go generate /go/src/secureworks/main.go

######tester######
FROM golang:1.11 as tester

ARG CONT_IP_DB=${CONT_IP_DB}
ARG SQL_DB=${SQL_DB}

COPY --from=sourcer /go/src /go/src
COPY --from=maxmind ${CONT_IP_DB} ${CONT_IP_DB}
COPY --from=database ${SQL_DB} ${SQL_DB}

RUN go test -v secureworks/...

######builder######
FROM golang:1.11 as builder

ARG CONT_IP_DB=${CONT_IP_DB}
ARG SQL_DB=${SQL_DB}
ARG CONT_BIN=${CONT_BIN}

COPY --from=sourcer /go/src /go/src

RUN go build -o ${CONT_BIN} secureworks

######app######
FROM ubuntu:devel as app

ARG CONT_IP_DB=${CONT_IP_DB}
ARG SQL_DB=${SQL_DB}
ARG CONT_BIN=${CONT_BIN}
ARG PORT=${PORT}

ENV CONT_BIN ${CONT_BIN}

COPY --from=maxmind ${CONT_IP_DB} ${CONT_IP_DB}
COPY --from=database ${SQL_DB} ${SQL_DB}
COPY --from=builder ${CONT_BIN} ${CONT_BIN}

EXPOSE ${PORT}

ENTRYPOINT ${CONT_BIN}
