FROM scratch
COPY ./out/uuid-srv /
ENTRYPOINT ["/uuid-srv"]
