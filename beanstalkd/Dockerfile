FROM alpine:3.19.1
MAINTAINER jon brookes <jon@ajbc.co>

## Add Tini and Beanstalkd
RUN apk add --update --no-cache tini beanstalkd && rm -fr /var/cache/apk/*

## Ensure /opt/beanstalkd/data directory exists and is writable by nobody
RUN mkdir -p /opt/beanstalkd/data && chown nobody:nobody /opt/beanstalkd/data

ENTRYPOINT ["/sbin/tini", "--"]

## Expose the port
EXPOSE 11300

CMD ["beanstalkd", "-p", "11300", "-u", "nobody", "-b", "/opt/beanstalkd/data"]
