FROM golang:1.18.2 as builder

WORKDIR /go/src/github.com/meshplus/pier-client-ethereum/
ARG http_proxy=""
ARG https_proxy=""
ENV PATH=$PATH:/go/bin
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/lib
COPY . /go/src/github.com/meshplus/pier-client-ethereum/

RUN go env -w GOPROXY=https://goproxy.cn,direct \
    && go env -w GOFLAGS="--mod=mod" \
    && version=$(bash /go/src/github.com/meshplus/pier-client-ethereum/scripts/version.sh) \
    && echo $version \
    && cd /go/src/github.com/meshplus \
    && git clone -b $version https://github.com/meshplus/pier.git \
    && cd /go/src/github.com/meshplus/pier \
    && go install github.com/gobuffalo/packr/v2/packr2@v2.8.3 \
    && make install \
    && cp ./build/wasm/lib/linux-amd64/libwasmer.so /lib \
    && cd /go/src/github.com/meshplus/pier-client-ethereum/ \
    && make eth \
    && pier init relay \
    && mkdir /root/.pier/ethereum /root/.pier/plugins \
    && cp /go/src/github.com/meshplus/pier-client-ethereum/build/eth-client /root/.pier/plugins/appchain_plugin \
    && cp -r /go/src/github.com/meshplus/pier-client-ethereum/config/* /root/.pier/ethereum

FROM frolvlad/alpine-glibc:glibc-2.32

COPY --from=0 /go/bin/pier /usr/local/bin/pier
COPY --from=0 /root/.pier /root/.pier
COPY --from=0 /lib/libwasmer.so /lib/libwasmer.so
ENV LD_LIBRARY_PATH=$LD_LIBRARY_PATH:/lib

EXPOSE 44544 44555
ENTRYPOINT ["/bin/sh", "-c","pier start"]