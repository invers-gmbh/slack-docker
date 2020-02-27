# slack-docker [![CircleCI](https://circleci.com/gh/int128/slack-docker.svg?style=shield)](https://circleci.com/gh/int128/slack-docker)

A Slack integration to notify [Docker events](https://docs.docker.com/engine/reference/commandline/events/).

<img width="596" alt="slack-docker-screenshot" src="https://user-images.githubusercontent.com/321266/47410763-c7682d80-d7a1-11e8-8f05-c80786152604.png">


## Getting Started

Setup [an Incoming WebHook](https://my.slack.com/services/new/incoming-webhook) on your Slack workspace and get the WebHook URL.

Install slack-docker by brew tap or from the [releases](https://github.com/int128/slack-docker/releases).

```sh
brew tap int128/slack-docker
brew install slack-docker
```

Run the command with the WebHook URL.

```sh
slack-docker --webhook=https://hooks.slack.com/services/...
```

You can run on Docker or Docker Compose as well.

```
# Docker
docker run -d -e webhook=https://hooks.slack.com/services/... -h "$(hostname)" -v /var/run/docker.sock:/var/run/docker.sock inverscom/slack-docker

# Docker Compose
curl -O https://raw.githubusercontent.com/int128/slack-docker/master/docker-compose.yml
docker-compose up -d
```


## Configuration

It supports the following options and environment variables:

```
Application Options:
      --webhook=      Slack Incoming WebHook URL [$webhook]
      --image-regexp= Filter events by image name (default to all) [$image_regexp]
      --container-regexp= Filter events by container name (default to all) [$container_regexp]
      --action-regexp= Filter events by action (default to all) [$action_regexp]
      --type-regexp= Filter events by type (default to all) [$type_regexp]

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
