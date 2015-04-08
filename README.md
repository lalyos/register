This is VIP cli tool to manage docker image tag management.
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



