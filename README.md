This is a cli tool to manage docker image tag management.
Lets automate the to replace point-and-click workflow on https://hub.docker.com


## Issue:
As a best practice each docker image has to be used with an exact tag like `sequenceiq/cloudbreak:0.4.1`. 
To achive it right now the following steps has to be done manually:

- maintain 2 automated build on Docker hub:
  - `latest` image tag pointing to master branch on github
  - `x.y.z` image tag pointing to x.y.z git tagged
- once the x.y.z tag built once, it has to be deleted. Please note that now we refer to delete the automated build, not the tagged image.

We want to make it automated via jenkins/CircleCI

## Work In Progress

The only official [lib](https://github.com/CenturyLinkLabs/docker-reg-client) for the Docker registry,
is able to do read-only operations, like getting image metadat, or listing the tags of an image/repo,
however i can’t get write permission ...


```
[REQ] ===>  PUT /v1/repositories/sequenceiq/cloudbreak/ HTTP/1.1
Host: index.docker.io
User-Agent: Go 1.1 package http
Content-Length: 3
Authorization: Basic ZGVhcl9oYWNrZXI6aWtub3diYXNlNjRpc3Vuc2FmZTsp
X-Docker-Token: true
Accept-Encoding: gzip

[RESP] ===>  HTTP/1.1 403 FORBIDDEN
Connection: close
Transfer-Encoding: chunked
Content-Type: application/json
Date: Thu, 09 Apr 2015 09:20:30 GMT
Server: nginx/1.6.2
Set-Cookie: csrftoken=53wN5L9B8xx4dfX6RZqe6QIWcp9WQRzv; expires=Thu, 07-Apr-2016 09:20:30 GMT; httponly; Max-Age=31449600; Path=/; secure
Set-Cookie: sessionid=lmo7ovbxqquj2zlo1w7kevyjgurw0bgq; expires=Sun, 12-Apr-2015 20:40:30 GMT; httponly; Max-Age=300000; Path=/; secure
Strict-Transport-Security: max-age=31536000
Vary: Cookie
X-Frame-Options: SAMEORIGIN

```
