# Include common Makefile code.
BASE_IMAGE_NAME = golang
VERSIONS = 1.8 1.9 1.10 1.11 1.12
OPENSHIFT_NAMESPACES =

# HACK:  Ensure that 'git pull' for old clones doesn't cause confusion.
# New clones should use '--recursive'.
.PHONY: $(shell test -f common/common.mk || echo >&2 'Please do "git submodule update --init" first.')

include common/common.mk
