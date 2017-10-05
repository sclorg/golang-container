FROM centos:centos7
MAINTAINER Tom Stellard <tstellar@redhat.com>

ENV SUMMARY="Red Hat Go Toolset 7 container image" \
    DESCRIPTION="Platform for building applications using Red Hat Go Toolset 7"

LABEL com.redhat.component="go-toolset-7-docker" \
      name="devtools/go-toolset-7-rhel7" \
      version="7" \
      summary="$SUMMARY" \
      description="$DESCRIPTION" \
      io.k8s.description="$DESCRIPTION" \
      io.k8s.display-name="Go Toolset 7"

RUN yum install -y yum-utils && \
    yum install -y centos-release-scl && \
    yum-config-manager --enable centos-sclo-rh-testing && \
    INSTALL_PKGS="go-toolset-7" && \
    yum install -y --setopt=tsflags=nodocs $INSTALL_PKGS && \
    rpm -V $INSTALL_PKGS && \
    yum clean all -y

COPY ./root /

ENV HOME=/opt/app-root/src \
    PATH=/opt/app-root/src/bin:/opt/app-root/bin:/opt/rh/go-toolset-7/root/usr/bin/:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

RUN mkdir -p ${HOME} && \
    groupadd -r default -f -g 1001 && \
    useradd -u 1001 -r -g default -d ${HOME} -s /sbin/nologin \
    -c "Default Application User" default && \
    chown -R 1001:1001 /opt/app-root /usr/bin/container-entrypoint /usr/bin/usage && \
    chmod u+x /usr/bin/usage

USER 1001

WORKDIR ${HOME}

# Enable the SCL for all bash scripts.
ENV BASH_ENV=/opt/app-root/etc/scl_enable \
    ENV=/opt/app-root/etc/scl_enable \
    PROMPT_COMMAND=". /opt/app-root/etc/scl_enable"

ENTRYPOINT ["container-entrypoint"]
CMD ["usage"]
