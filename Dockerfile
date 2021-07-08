FROM golang:alpine as build-env
LABEL maintainer="Fahmy Abida <fahmyabida@gmail.com>"
ARG SERVICE_NAME=fasthttp_crud

RUN mkdir /builder
WORKDIR /builder

ADD . /builder/

RUN go build -mod=vendor -o ${SERVICE_NAME} -mod=vendor .

FROM alpine

COPY  --from=build-env /builder/${SERVICE_NAME}  ./${SERVICE_NAME}
EXPOSE 80

RUN mkdir -p logs

RUN apk add --no-cache tzdata
ENV TZ Asia/Jakarta

ENTRYPOINT ["./fasthttp_crud"]