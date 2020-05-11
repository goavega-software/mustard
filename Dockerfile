FROM node:14-alpine AS node
LABEL maintainer="Goavega Open Source Team <opensource@goavega.com>"
ENV BUILD_DIR /home/

RUN set -ex && \
    apk add yarn
COPY ./ ${BUILD_DIR}
WORKDIR ${BUILD_DIR}client
RUN set -ex && \
    yarn install && \
    yarn prod

FROM golang:1.14-alpine AS go
ENV APP_DIR /var/www/
ENV BUILD_DIR /home/
RUN set -ex && \
	mkdir -p ${APP_DIR}js/ && \
	mkdir -p ${APP_DIR}css/ && \
	mkdir -p ${APP_DIR}views/ 
COPY --from=node /home/js/* ${APP_DIR}js/
COPY --from=node /home/css/* ${APP_DIR}css/
COPY --from=node /home/views/* ${APP_DIR}views/
COPY ./ ${BUILD_DIR}
WORKDIR ${BUILD_DIR}
RUN set -ex && \
    go build ./ && \
	cp ./mustard ${APP_DIR}

FROM alpine:latest
ENV APP_DIR /var/www/
ENV PORT 80

COPY --from=go ${APP_DIR} ${APP_DIR}
COPY ./dockersupport/start.sh /usr/local/bin/
WORKDIR ${APP_DIR}
RUN set -ex && \
    chmod u+x ./mustard && \
    chmod u+x /usr/local/bin/start.sh
EXPOSE ${PORT}
CMD ["start.sh"]