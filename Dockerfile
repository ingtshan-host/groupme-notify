FROM debian:jessie-slim
COPY fs /

ARG GROUP
ARG MESSAGE

CMD ["/notify", "-groupId", "$GROUP", "-m", "$MESSAGE"]