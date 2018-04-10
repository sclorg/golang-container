Golang s2i Docker image
===================

This repository contains the source for building various versions of
the Go application as a reproducible s2i Docker image.
Users can choose between RHEL and CentOS based builder images.
The resulting image can be run using [Docker](http://docker.io) or using [source-to-image](https://github.com/openshift/source-to-image/).


Usage
---------------------

Simplest usage via s2i:

s2i build ./src centos/go-toolset-7-centos7:latest test-app

where ./src directory contains git repository with golang application that has complete dependencies. You can't use incremental build in this case.

or

s2i build -e IMPORT_URL='github.com/cpuguy83/go-md2man' ./src centos/go-toolset-7-centos7:latest test-app

where ./src directory contains git repository with golang application(in this case github.com/cpuguy83/go-md2man) that has complete or incomplete dependencies. You can use incremental build in this case.

Both will build test-app application image.

Environment variables
---------------------

To set these environment variables, you can place them as a key value pair into a `.s2i/environment`
file inside your source code repository or specified via s2i invocation.


* **IMPORT_URL**

    Used to specify the golang application import URL (i.e. usually something like github.com/someorg/somerepo), that is build. Necessary for the incremental build to function.

* **INSTALL_URL**

    Used to specify the golang application import URL of the main package (i.e. usually something like github.com/someorg/somerepo/subfolder). Necessary if the main package is not in the root folder of the repository.
