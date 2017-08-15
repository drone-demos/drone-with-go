# drone-with-go [![Build Status](https://ci.cloudlockng.com/api/badges/josebarn/drone-with-go/status.svg)](https://ci.cloudlockng.com/josebarn/drone-with-go)

An example of how to test Go code with Drone.

Inspect the .drone.yml file for details on the build steps for this simple go project.

Notice the .drone.sig is used to ensure secrets are not exposed; and the pipeline is split into multiple phases.
