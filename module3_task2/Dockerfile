FROM ubuntu:22.04
ENV PATH="/usr/local/go/bin:${PATH}"
RUN apt-get update && \
    apt-get -y install sudo curl
RUN curl -sL https://deb.nodesource.com/setup_14.x | bash - \
    && apt-get install -y --no-install-recommends \
    nodejs=14.* \
    make=4.* \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/*
RUN npm install -g npm@9.5.1
ARG GO_VERSION="1.20.4"
ARG GO_CHECKSUM="698ef3243972a51ddb4028e4a1ac63dc6d60821bf18e59a807e051fee0a385bd"
ENV CGO_ENABLED=0
ENV GO11MODULES=on
RUN curl --silent --show-error --location --output /tmp/go.tgz \
    "https://golang.org/dl/go${GO_VERSION}.linux-amd64.tar.gz" \
  && sha256sum /tmp/go.tgz | grep -q "${GO_CHECKSUM}" \
  && tar -C /usr/local -xzf /tmp/go.tgz \
  && go version | grep "${GO_VERSION}" \
  && rm -f /tmp/go.tgz
RUN useradd -m student
RUN usermod -aG sudo student
RUN echo 'student ALL=(ALL) NOPASSWD:ALL' >> /etc/sudoers
USER student
CMD ["/bin/bash"]
