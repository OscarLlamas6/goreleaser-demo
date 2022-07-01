FROM scratch
COPY goreleaser-demo /usr/bin/goreleaser-demo
ENTRYPOINT ["/usr/bin/goreleaser-demo"]