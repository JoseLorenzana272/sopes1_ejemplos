FROM grafana/grafana:11.0.0

RUN grafana-cli plugins install redis-datasource
VOLUME /var/lib/grafana
EXPOSE 3000
