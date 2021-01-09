FROM debian:jessie-slim
# need certs to trust groupme
RUN apt-get update && apt-get install -y ca-certificates

COPY fs /

ARG GROUP
ARG MESSAGE

CMD ["/notify", "-groupId", "$GROUP", "-m", "$MESSAGE"]