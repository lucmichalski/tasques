## K8S

This module allows you to run a Tasques server for playing around.

* `make install-eck` to install the [ECK](https://github.com/elastic/cloud-on-k8s) k8s operator
* `make deploy` spins up a working env
* `make teardown` tears ... it down..

### Configuration

Tasques can be configured via a config file or via environment variable; see the `tasques.yml` k8s definition file
for examples.