FROM golang:1.14.2 as builder

RUN mkdir -p /go/src/github.com/meshplus/pier-client-ethereum
WORKDIR /go/src/github.com/meshplus/pier-client-ethereum

# Cache dependencies
COPY . .

RUN mv build/pier ../pier

# Build real binaries
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go get -u github.com/gobuffalo/packr/packr
RUN cd ../pier && make install

RUN cd ../pier-client-ethereum && \
    make eth && \
    cp build/eth-client.so /go/bin/eth-client.so

# Final image
FROM frolvlad/alpine-glibc

WORKDIR /root

# Copy over binaries from the builder
COPY --from=builder /go/bin/pier /usr/local/bin
COPY ./build/pier/build/libwasmer.so /lib
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/lib

RUN ["pier", "init"]

RUN mkdir -p /root/.pier/plugins
COPY --from=builder /go/bin/*.so /root/.pier/plugins/
COPY config/validating.wasm /root/.pier/validating.wasm
COPY scripts/docker_entrypoint.sh /root/docker_entrypoint.sh
RUN chmod +x /root/docker_entrypoint.sh

COPY config /root/.pier/ether
COPY config/pier.toml /root/.pier/pier.toml

ENV APPCHAIN_NAME=ether

EXPOSE 44555 44544

ENTRYPOINT ["/root/docker_entrypoint.sh", "$APPCHAIN_NAME"]