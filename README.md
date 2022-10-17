# http-rejector

This trivial HTTP service just returns a 403 rejection for all requests
except some health checks.

Listens on port 8000.  Available as the Docker image tyemcq/http-rejector.
Trivial to deploy to a Google Kubernetes Engine cluster using the included
[workload.yaml](/workload.yaml).

To be a health check, the HTTP method must be "GET" and either the User-Agent
header must have one of the specified prefixes or the URL path must exactly
match one of the specified values (including query parameters).  The list of
UA prefixes and the list of path values can be customized by the below
environment variables.  Empty values in each comma-separated list are ignored
so you can set `USERAGENT_PREFIX=,` to ignore the User-Agent header.

* HEALTH_PATH can be set to a comma-separated list of URL paths. The default
    value is an empty list.

* USERAGENT_PREFIX can be set to a comma-separated list of User-Agent header
    prefixes. The default value is "GoogleHC/".

Therefore, by default, any health checks issued by any Google Cloud Provider
load balancer will get an empty 200 success response.
