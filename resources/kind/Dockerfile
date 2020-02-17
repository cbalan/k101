FROM alpine:3.9 AS build
RUN wget -O /kind https://github.com/kubernetes-sigs/kind/releases/download/v0.7.0/kind-linux-amd64
RUN chmod +x /kind

FROM docker:18.09.6 as docker

FROM alpine:3.9
COPY --from=build /kind /bin/kind
COPY --from=docker /usr/local/bin/docker /bin/docker
WORKDIR /
ENTRYPOINT ["/bin/kind"]
