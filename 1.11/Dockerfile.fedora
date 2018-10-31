FROM registry.fedoraproject.org/f29/s2i-base:latest

ENV NAME=golang \
    VERSION=1.11 \
    ARCH=x86_64

ENV SUMMARY="Platform for building and running Go $VERSION based applications" \
    DESCRIPTION="Go $VERSION available as container is a base platform for \
building and running various Go $VERSION applications and frameworks. \
Go is an easy to learn, powerful, statically typed language in the C/C++ \
tradition with garbage collection, concurrent programming support, and memory safety features."

LABEL summary="$SUMMARY" \
      description="$DESCRIPTION" \
      io.k8s.description="$DESCRIPTION" \
      io.k8s.display-name="Go $VERSION" \
      io.openshift.tags="builder,golang,golang18,rh-golang18,go" \
      com.redhat.component="$NAME" \
      name="$FGC/$NAME" \
      version="$VERSION" \
      architecture="$ARCH" \
      maintainer="Jakub Cajka <jcajka@fedoraproject.org>" \
      usage="docker run $FGC/$NAME"

RUN INSTALL_PKGS="golang" && \
    dnf install -y --setopt=tsflags=nodocs $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    dnf clean all -y

# Copy the S2I scripts from the specific language image to $STI_SCRIPTS_PATH.
COPY ./s2i/bin/ $STI_SCRIPTS_PATH
# Copy manpage
COPY ./root/help.1 /

RUN chown -R 1001:0 $STI_SCRIPTS_PATH && chown -R 1001:0 /opt/app-root

USER 1001

# Set the default CMD to print the usage of the language image.
CMD $STI_SCRIPTS_PATH/usage
