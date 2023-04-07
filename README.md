# load-balancer
set up load balancer for different countries and with backup

* There are 5 containers for go application 2 - UK, 1 - US, 1 - others, 1 - backup
* For set up load balancing between different countries i used `ngx_http_geoip2_module`
* `geoip.mmdb` it's a lite version and there aren't a lot of countries, i noticed only `US`
* `health_check` it's a part of nginx+ program, i couldn't figure out how can set up this module separately
* for case with `health_check` need to add `health_check interval=5 fails=2 passes=2;` i used healthcheck in `docker-compose.yml`
