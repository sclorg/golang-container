Golang container images
====================

This repository contains the source for building of
the Go application as a reproducible s2i container image.
Users can choose between RHEL, Fedora and CentOS based builder images.
The resulting image can be run using [podman](https://github.com/containers/libpod), [Docker](http://docker.io) or using [source-to-image](https://github.com/openshift/source-to-image/).
Also the whole pipeline from build to app deployment could be run on top of the [Openshift Origin](https://www.okd.io/) or [Red Hat's Openshift](https://www.openshift.com/).

For more information about contributing, see
[the Contribution Guidelines](https://github.com/sclorg/welcome/blob/master/contribution.md).
For more information about concepts used in these container images, see the
[Landing page](https://github.com/sclorg/welcome).


Versions
---------------
Golang versions currently provided are:
* [Golang 1.8 (go-toolset 7, where applicable)](1.8)
* [Golang 1.9](1.9)
* [Golang 1.10](1.10)
* [Golang 1.11](1.11)

RHEL versions currently supported are:
* RHEL7

CentOS versions currently supported are:
* CentOS7

Fedora versions currently supported are:
* Fedora 27
* Fedora 28
* Fedora 29


Usage
---------------------------------

For information about usage of Dockerfile for Golang Toolset 7 and Fedora golang,
see usage documentation for respective Go versions [Go1.8](1.8/README.md) [Go1.9](1.9/README.md) [Go1.10](1.10/README.md).

