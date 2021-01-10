FROM debian:jessie-slim
# need certs to trust groupme
RUN apt-get update && apt-get install -y ca-certificates

COPY fs /

CMD ["/notify"]