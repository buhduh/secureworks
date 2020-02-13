######maxmind######
FROM alpine:latest as maxmind
ARG HOST_IP_DB=${HOST_IP_DB}
ARG CONT_IP_DB=${CONT_IP_DB}
COPY ${HOST_IP_DB} ${CONT_IP_DB}

######vendor######
#FROM golang:1.11 as vendor

#ARG VENDOR_DIR=${VENDOR_DIR}
#ARG APP_SRC=${APP_SRC}

#COPY ${APP_SRC} ${VENDOR_DIR}/src
#WORKDIR ${VENDOR_DIR}/src
#RUN go mod init secureworks
#RUN go mod vendor

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

ENV GEN_GO        ${GEN_GO}
ENV SQL_DB        ${SQL_DB}
ENV CONT_IP_DB    ${CONT_IP_DB}
ENV CREATE_DB_SQL ${CREATE_DB_SQL}

COPY ${APP_SRC} /go/src/
COPY scripts/generate.sh ${CONT_GEN_SCRIPT}
COPY scripts/createdb.sql ${CREATE_DB_SQL}

RUN chmod +x ${CONT_GEN_SCRIPT}
RUN go generate /go/src/secureworks/main.go

######tester######
FROM golang:1.11 as tester

ARG CONT_IP_DB=${CONT_IP_DB}
ARG CREATE_DB_SQL=${CREATE_DB_SQL}

COPY --from=sourcer /go/src /go/src
COPY --from=maxmind ${CONT_IP_DB} ${CONT_IP_DB}
COPY --from=sourcer ${CREATE_DB_SQL} ${CREATE_DB_SQL}
RUN go test -v secureworks/...

######builder######
#FROM golang:1.11 as builder

#ARG CONT_IP_DB=${CONT_IP_DB}
#ARG CONT_BIN=${CONT_BIN}

#COPY --from=sourcer ${CONT_IP_DB} ${CONT_IP_DB}
#COPY --from=sourcer /go/src /go/src

#RUN go build -o ${CONT_BIN} secureworks
