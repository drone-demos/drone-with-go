# drone-with-go [![Build Status](http://beta.drone.io/api/badges/drone-demos/drone-with-go/status.svg)](http://beta.drone.io/drone-demos/drone-with-go) [![Build Status](https://aircover.co/badges/drone-demos/drone-with-go/coverage.svg)](https://aircover.co/drone-demos/drone-with-go)

An example of how to test Go code with Drone.

# Basic Testing
To run basic CI tests use the following in your `.drone.yml` file.

```yaml
build:
  image: golang:1.5.3
  commands:
    - go test ./...
```

In this config `image: golang:1.5.3` references the official Golang Docker image hosted on the Docker Hub registry at https://hub.docker.com/r/_/golang/.
Go tests are execute with the `go test ./...` command.

# Advanced Testing

## Environment Variables
Using environment variables to configure Go testing is easy.
Configure environment variables by setting the `build` section's `environment`.

```yaml
build:
  image: golang:1.5.3
  environment:
    - GO15VENDOREXPERIMENT=1
    - GOOS=linux
    - GOARCH=amd64
    - CGO_ENABLED=0
  commands:
    - go test ./...
```

## Coverage
Drone tests work best with the coverage plugin and the [aircover.co](https://aircover.co) service.
We only want to send a coverage report when all tests pass, so the coverage plugin uses `publish`.
Also, we should specify a particular branch so that coverage reports are consistent.

```yaml
build:
  image: golang:1.5.3
  environment:
    - GO15VENDOREXPERIMENT=1
    - GOOS=linux
    - GOARCH=amd64
    - CGO_ENABLED=0
  commands:
    - go test -cover -coverprofile coverage.out

publish:
  coverage:
    when:
      branch: master
```

## Plugins
Notification plugins use `notify` for integrations like HipChat.
You can find a list of plugins at [readme.drone.io/plugins](http://readme.drone.io/plugins/).

```yaml
build:
  image: golang:1.5.3
  environment:
    - GO15VENDOREXPERIMENT=1
    - GOOS=linux
    - GOARCH=amd64
    - CGO_ENABLED=0
  commands:
    - go test -cover -coverprofile coverage.out

publish:
  coverage:
    when:
      branch: master

notify:
  hipchat:
    from: Your_Project
    notify: true
    room_id_or_name: Your_Room
    auth_token: $$HIPCHAT_DRONE_TOKEN
```
