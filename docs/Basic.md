**[NGINX's Official Docs](https://nginx.org/en/docs/) |
[Installation](https://docs.nginx.com/nginx/admin-guide/installing-nginx/installing-nginx-open-source/) |
[Introduction](https://nginx.org/en/docs/#introduction) |
[Pitfalls](https://www.nginx.com/resources/wiki/start/topics/tutorials/config_pitfalls/) |
[Location Tester](https://nginx.viraptor.info/) |
[Regex Guide](https://stackoverflow.com/a/59846239) |
[All NGINX Variables](http://nginx.org/en/docs/varindex.html) | [Nginx Conf Example](https://www.nginx.com/resources/wiki/start/topics/examples/fullexample2/)**

## CORE
**master**
- purpose: maintain and evaluate the configuration and worker.
- count: 1
- not configurable

**worker**
- purpose: execute the request.
- count: **YOUR_CPU_CORE** (default)
- `worker_processes auto;` will spawn as much as **N** core
- `worker_processes 5;`, spawn 5
- configurable at `/etc/nginx/nginx.conf` [[more]](http://nginx.org/en/docs/ngx_core_module.html#worker_processes)


## FLOW
always going first to `/etc/nginx/nginx.conf` (main configuration), and inside of it will include:
```
include /etc/nginx/conf.d/*.conf;
include /etc/nginx/sites-enabled/*;
```
that's why all files under `/etc/nginx/sites-enabled/*;` always run, not the `/etc/nginx/sites-available/*`. Usually the `sites-available` is used to store the configuration and `sites-enabled;` will create sibling to it.


## CMD
| cmd | Desc |
| --- | --- |
| service nginx status | check nginx status |
| service nginx reload / nginx -s reload | reload the conf |
| service nginx start | start the nginx |
| service nginx stop / nginx -s stop / nginx -s quit | stop the nginx |
| service nginx restart | restart the nginx |
| nginx -t | test nginx's conf |
| **nginx -t && nginx -s reload** | **reload the conf if pass the test** |


## SYNTAX
**simple**
```
proxy_cache_key $request_uri;
```
**block**
```
map "$slow:$abnormal" $logable {
    "1:1" 1;
    default 0;
}
```
**context**, braces in braces (events, http, server, and location)
```
server {
    location / {
        proxy_pass http://127.0.0.1:9000;
    }
}
```
**comment**
```
# Hello World
```