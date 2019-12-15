# Nginx Load Balancer Demo

Docker compose with app backend and nginx as simple proxy/load balancer. See https://docs.nginx.com/nginx/admin-guide/load-balancer/http-load-balancer/#proxy_pass

Note that nginx will only resolve the DNS record for the backend at startup, so scaling the app later won't have any effect until nginx is restarted.

## Usage

Start app with 3 backends:

```sh
$ docker-compose up --scale app=3
```

Each request should return a different hostname:

```sh
$ curl -sS localhost:8080
2b513cebc169
```

Scale up and restart nginx:

```sh
$ docker-compose up --scale app=5
$ docker-compose restart nginx
```
