# build stage
FROM node:lts-alpine as build-stage
WORKDIR /front
COPY front/package*.json ./
RUN npm install
COPY ./front .
RUN npm run build

FROM golang:1.12.5 as Builder
RUN mkdir /go/src/accounting_software
COPY ./server /go/src/accounting_software/server
WORKDIR /go/src/accounting_software/server/service
RUN go get
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /accounting_software /go/src/accounting_software/server/service

FROM alpine:latest
WORKDIR /
RUN mkdir /usr/share/accounting_software
RUN mkdir /templates
COPY --from=build-stage /front/dist/index.html ./templates
COPY --from=Builder /accounting_software .
COPY  --from=Builder /go/src/accounting_software/server/config.ini ./config.ini
RUN mkdir /migdir
COPY --from=Builder /go/src/accounting_software/server/migrations /migdir
RUN ls -la /
EXPOSE 80

#FROM alpine:latest
#ARG TYPE
#WORKDIR /
#RUN mkdir /usr/share/accounting_software
#COPY --from=build-stage /front/dist/static /usr/share/accounting_software/static
#RUN mkdir /templates
#COPY --from=build-stage /front/dist/index.html ./templates
#COPY --from=Builder /accounting_software .
#COPY  --from=Builder /go/src/accounting_software/server/config.${TYPE}.ini ./config.ini
#RUN mkdir /migdir
#COPY --from=Builder /go/src/accounting_software/server/migrations /migdir
#RUN ls -la /
#EXPOSE 80

#fileserver port
EXPOSE 55555
CMD ["/accounting_software"]
