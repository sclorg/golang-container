Golang s2i Docker image
===================

This repository contains the source for building various versions of
the Go application as a reproducible s2i Docker image.
Users can choose between RHEL and CentOS based builder images.
The resulting image can be run using [Docker](http://docker.io) or using [source-to-image](https://github.com/openshift/source-to-image/).


Usage
---------------------

Simplest usage via s2i:

s2i build ./src centos/golang-1.8-centos7:latest test-app

where ./src directory contains git repository with golang application that has complete dependencies. You can't use incremental build in this case.

or

s2i build -e IMPORT_URL='github.com/cpuguy83/go-md2man' ./src centos/golang-1.8-centos7:latest test-app

where ./src directory contains git repository with golang application(in this case github.com/cpuguy83/go-md2man) that has complete or incomplete dependencies. You can use incremental build in this case.

Both will build test-app application image.
