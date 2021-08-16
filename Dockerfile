FROM ttbb/base:go AS build
COPY . /opt/sh/compile
WORKDIR /opt/sh/compile/pkg
RUN go build -o mysql_mate .


FROM ttbb/mysql:nake

LABEL maintainer="shoothzj@gmail.com"

COPY docker-build /opt/sh/mysql/mate

COPY --from=build /opt/sh/compile/pkg/mysql_mate /opt/sh/mysql/mate/mysql_mate

CMD ["/usr/local/bin/dumb-init", "bash", "-vx", "/opt/sh/mysql/mate/scripts/start.sh"]