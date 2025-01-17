FROM openeuler/openeuler:23.03 as BUILDER
RUN dnf update -y && \
    dnf install -y golang && \
    go env -w GOPROXY=https://goproxy.cn,direct

ARG USER
ARG PASS
RUN echo "machine github.com login $USER password $PASS" > /root/.netrc
RUN go env -w GOPRIVATE=github.com/opensourceways

# build binary
COPY . /go/src/github.com/opensourceways/xihe-statistics
RUN cd /go/src/github.com/opensourceways/xihe-statistics && GO111MODULE=on CGO_ENABLED=0 go build -buildmode=pie --ldflags "-s -linkmode 'external' -extldflags '-Wl,-z,now'"

# copy binary config and utils
FROM openeuler/openeuler:22.03
RUN dnf -y update && \
    dnf in -y shadow tzdata && \
    groupadd -g 5000 mindspore && \
    useradd -u 5000 -g mindspore -s /bin/bash -m mindspore

USER mindspore
WORKDIR /opt/app/

COPY  --chown=mindspore --from=BUILDER /go/src/github.com/opensourceways/xihe-statistics /opt/app
RUN chmod 550 /opt/app/xihe-statistics

ENTRYPOINT ["/opt/app/xihe-statistics"]