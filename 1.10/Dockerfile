FROM centos/s2i-base-centos7

ENV NAME=golang \
    VERSION=1.10

ENV SUMMARY="Platform for building and running Go $VERSION based applications" \
    DESCRIPTION="Go $VERSION available as docker container is a base platform for \
building and running various Go $VERSION applications and frameworks. \
Go is an easy to learn, powerful, statically typed language in the C/C++ \
tradition with garbage collection, concurrent programming support, and memory safety features."

LABEL summary="$SUMMARY" \
      description="$DESCRIPTION" \
      io.k8s.description="$DESCRIPTION" \
      io.k8s.display-name="Go $VERSION" \
      io.openshift.tags="builder,golang,golang18,rh-golang18,go" \
      com.redhat.component="go-toolset-7" \
      name="centos/go-toolset-7-centos7" \
      version="1" \
      maintainer="Jakub Čajka <jcajka@redhat.com>" \
      usage="docker run centos/go-toolset-7-centos7"

RUN yum install -y centos-release-scl-rh && \
    yum-config-manager --enable centos-sclo-rh-testing && \
    INSTALL_PKGS="go-toolset-7" && \
    yum install -y --setopt=tsflags=nodocs $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    yum clean all -y

# Copy the S2I scripts from the specific language image to $STI_SCRIPTS_PATH.
COPY ./s2i/bin/ $STI_SCRIPTS_PATH

COPY ./root/ /

RUN chown -R 1001:0 $STI_SCRIPTS_PATH && chown -R 1001:0 $APP_ROOT

USER 1001

# Set the default CMD to print the usage of the language image.
CMD $STI_SCRIPTS_PATH/usage
