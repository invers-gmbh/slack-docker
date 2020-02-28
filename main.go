package main

import (
	"context"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/int128/slack"
	"github.com/invers-gmbh/slack-docker/formatter"
	"github.com/jessevdk/go-flags"
	"io"
	"log"
	"regexp"
)

type options struct {
	Webhook             string `long:"webhook" env:"webhook" description:"Slack Incoming WebHook URL" required:"1"`
	DockerImageRegexp   string `long:"image-regexp" env:"image_regexp" description:"Filter events by image name (default to all)"`
	TypeRegexp          string `long:"type-regexp" env:"type_regexp" description:"Filter events by type (default to all)"`
	ContainerNameRegexp string `long:"container-name-regexp" env:"container_name_regexp" description:"Filter events by container name (default to all)"`
	ActionRegexp        string `long:"action-regexp" env:"action_regexp" description:"Filter events by action (default to all)"`
	TitleLink           string `long:"title-link" env:"title_link" description:"Link that should be used in titles of slack attachments"`
}

func (o *options) Run(ctx context.Context) error {
	var eventFilter formatter.EventFilter
	if o.DockerImageRegexp != "" {
		r, err := regexp.Compile(o.DockerImageRegexp)
		if err != nil {
			return fmt.Errorf("invalid image-regexp: %s", err)
		}
		eventFilter.ImageRegexp = r
	}
	if o.TypeRegexp != "" {
		r, err := regexp.Compile(o.TypeRegexp)
		if err != nil {
			return fmt.Errorf("invalid type-regexp: %s", err)
		}
		eventFilter.TypeRegexp = r
	}
	if o.ContainerNameRegexp != "" {
		r, err := regexp.Compile(o.ContainerNameRegexp)
		if err != nil {
			return fmt.Errorf("invalid container-regexp: %s", err)
		}
		eventFilter.ContainerNameRegexp = r
	}
	if o.DockerImageRegexp != "" {
		r, err := regexp.Compile(o.DockerImageRegexp)
		if err != nil {
			return fmt.Errorf("invalid image-regexp: %s", err)
		}
		eventFilter.ImageRegexp = r
	}
	if o.ActionRegexp != "" {
		r, err := regexp.Compile(o.ActionRegexp)
		if err != nil {
			return fmt.Errorf("invalid action-regexp: %s", err)
		}
		eventFilter.ActionRegexp = r
	}
	docker, err := client.NewEnvClient()
	if err != nil {
		return fmt.Errorf("could not create a Docker client: %s", err)
	}
	if err := o.showStartup(ctx, docker, eventFilter, o.TitleLink); err != nil {
		return err
	}
	if err := o.showEvents(ctx, docker, eventFilter, o.TitleLink); err != nil {
		return err
	}
	return nil
}

func (o *options) showStartup(ctx context.Context, docker *client.Client, filter formatter.EventFilter, titleLink string) error {
	v, err := docker.ServerVersion(ctx)
	if err != nil {
		return fmt.Errorf("could not get version from the Docker server: %s", err)
	}
	log.Printf("Connected to Docker server: %+v", v)
	if err := slack.Send(o.Webhook, formatter.Startup(v, filter, titleLink)); err != nil {
		return fmt.Errorf("could not send a message to Slack: %s", err)
	}
	return nil
}

func (o *options) showEvents(ctx context.Context, docker *client.Client, filter formatter.EventFilter, titleLink string) error {
	msgCh, errCh := docker.Events(ctx, types.EventsOptions{})
	for {
		select {
		case msg := <-msgCh:
			log.Printf("Event %+v", msg)
			m := formatter.Event(msg, filter, titleLink)
			if m != nil {
				if err := slack.Send(o.Webhook, m); err != nil {
					log.Printf("Error while sending a message to Slack: %s", err)
				}
			}
		case err := <-errCh:
			if err == io.EOF {
				break
			}
			log.Printf("Error while receiving events from Docker server: %s", err)
			if err := slack.Send(o.Webhook, formatter.Error(err)); err != nil {
				log.Printf("Error while sending a message to Slack: %s", err)
			}
		case <-ctx.Done():
			break
		}
	}
}

func main() {
	var o options
	args, err := flags.NewParser(&o, flags.HelpFlag).Parse()
	if err != nil {
		log.Fatal(err)
	}
	if len(args) > 0 {
		log.Fatalf("Too many arguments")
	}
	if err := o.Run(context.Background()); err != nil {
		log.Fatalf("Error: %s", err)
	}
}
