FROM nginx:latest

RUN apt-get update \
    && apt-get install -y nginx-extras \
    && rm -rf /var/lib/apt/lists/*

COPY geoip.mmdb /usr/share/GeoIP/geoip.mmdb

CMD ["nginx", "-g", "daemon off;"]