# go build
FROM golang:1.26.4 as build
RUN mkdir -p /app/building
WORKDIR /app/building
ADD . /app/building
ENV GOPROXY https://goproxy.cn
RUN make build

# copy & run
FROM alpine:3.24.1
COPY --from=build /app/building/dist/bin/risk_engine /app/bin/
COPY --from=build /app/building/dist/conf/config.yaml /app/conf/
COPY --from=build /app/building/dist/demo /app/demo
EXPOSE 8889
WORKDIR /app/
CMD ["bin/risk_engine", "-c", "conf/config.yaml"]
