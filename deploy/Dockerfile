FROM alpine:3.19
WORKDIR /opt/img-cloud-update
ENV LANG=en_US.utf8
COPY deploy/entrypoint.sh /opt/img-cloud-update/entrypoint.sh
COPY server/server /opt/img-cloud-update/server/server
COPY server/config.yaml /opt/img-cloud-update/server/config.yaml
COPY server/data  /opt/img-cloud-update/server/data/
COPY web/dist/ /opt/img-cloud-update/web/
COPY deploy/nginx.conf /etc/nginx/http.d/nginx.conf
RUN set -ex \
    && apk add --no-cache nginx musl-locales bash \
    && echo "LANG=en_US.utf8" > /etc/locale.conf \
    && chmod +x ./entrypoint.sh 
EXPOSE 80
ENTRYPOINT ["/bin/bash", "./entrypoint.sh"]