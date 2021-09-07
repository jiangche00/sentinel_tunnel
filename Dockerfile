FROM debian:buster-slim
ADD sentinel_tunnel /opt/sentinel_tunnel
RUN chmod +x /opt/sentinel_tunnel
ENV PATH=$PATH:/opt
WORKDIR /opt
ENTRYPOINT [ "/opt/sentinel_tunnel" ]
