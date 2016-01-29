# drone-with-go [![Build Status](http://beta.drone.io/api/badges/drone-demos/drone-with-go/status.svg)](http://beta.drone.io/drone-demos/drone-with-go)

An example of how to test Go code.

# Basic Testing
To run basic CI tests use the following in your `.drone.yml` file.

```yaml
build:
  image: golang:1.5.3
  commands:
    - go test ./...
```

# Advanced Testing
Using environment variables to configure Go testing is easy.
Configure environment variables by setting the `build` section's `environment`.

`.drone.yml`:
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

Below is a more advanced `.drone.yml` for notification integrations like HipChat.

```yaml

build:
  image: golang:1.5.3
  environment:
    - GO15VENDOREXPERIMENT=1
    - GOOS=linux
    - GOARCH=amd64
    - CGO_ENABLED=0
    - COVERALLS_TOKEN=$$COVERALLS_TOKEN
  commands:
    - go test -cover ./...

notify:
  hipchat:
    from: Your_Project
    notify: true
    room_id_or_name: Your_Room
    auth_token: $$HIPCHAT_DRONE_TOKEN
```
