ARG GO_VERSION=1.20
ARG PORT=9999

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

WORKDIR /project/
COPY . /project
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/api ./app/entry/api/main.go


#FROM alpine:latest
#
#RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
#
#COPY --from=build /bin/api .
#COPY --from=build /go/src/api/.env .
#COPY --from=build /go/src/api/${FILEBASECONFIGNAME} .

EXPOSE ${PORT}
ENTRYPOINT [ "/bin/api" ]

