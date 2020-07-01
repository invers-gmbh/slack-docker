# slack-docker [![CircleCI](https://circleci.com/gh/int128/slack-docker.svg?style=shield)](https://circleci.com/gh/int128/slack-docker)

A Slack integration to notify [Docker events](https://docs.docker.com/engine/reference/commandline/events/).

<img width="596" alt="slack-docker-screenshot" src="https://user-images.githubusercontent.com/321266/47410763-c7682d80-d7a1-11e8-8f05-c80786152604.png">


## Getting Started

Setup [an Incoming WebHook](https://my.slack.com/services/new/incoming-webhook) on your Slack workspace and get the WebHook URL.

```
# Docker
docker run -d \
      -e webhook=https://hooks.slack.com/services/... \
      -h "$(hostname)" \
      -v /var/run/docker.sock:/var/run/docker.sock \
      inverscom/slack-docker

# Docker Compose
curl -O https://raw.githubusercontent.com/int128/slack-docker/master/docker-compose.yml
docker-compose up -d
```

## Configuration

It supports the following options and environment variables:

```
Application Options [env-variable]:
      --webhook [$webhook]                      = Slack Incoming WebHook URL
      --image-regexp [$image_regexp]            = Filter events by image name (default to all) 
      --container-regexp [$container_regexp]    = Filter events by container name (default to all) 
      --action-regexp [$action_regexp]          = Filter events by action (default to all) - possible values: https://docs.docker.com/engine/reference/commandline/events/
      --type-regexp [$type_regexp]              = Filter events by type (default to all) - possible values: https://docs.docker.com/engine/reference/commandline/events/

Help Options:
  -h, --help          Show this help message
```


### Filter events by image name

```sh
webhook=https://hooks.slack.com/services/... image_regexp='^alpine$' ./slack-docker
```


## Contribution

This is an open source software licensed under Apache-2.0.
Feel free to open issues or pull requests.
