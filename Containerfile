ARG ALPINE_TAG=latest

FROM alpine:${ALPINE_TAG}

LABEL maintainer="Vinu Kochukuttan <vkochuku@redhat.com>" \
      description="A container image with various network tools pre-installed."

COPY server.go /

RUN apk -U upgrade \
    && apk add --no-cache \
        apache2-utils \
        bash \
        bash-completion \
        bind-tools \
        ca-certificates \
        coreutils \
        curl \
        drill \
        findutils \
        fping \
        git \
	go \
        gping \
        ipcalc \
        iperf \
        iproute2 \
        iputils \
        jq \
        libidn \
        macchanger \
        masscan \
        mtr \
        net-tools \
        netcat-openbsd \
        nftables \
        ngrep \
        nload \
        nmap \
        nmap-nping \
        nmap-scripts \
        oath-toolkit-oathtool \
        openssh-client \
        openssl \
        procps \
        socat \
        speedtest-cli \
        tcpdump \
        tcptraceroute \
        tshark \
        wget \
        python3 \
        py3-requests \
    && apk add --repository=https://dl-cdn.alpinelinux.org/alpine/edge/testing/ --no-cache \
        envsubst \
        grpcurl \
        hping3 \
    && apk add --repository=https://dl-cdn.alpinelinux.org/alpine/edge/community/ --no-cache \
        arp-scan \
        kcat \
    && rm -rf /var/cache/apk/* \
    && echo 'export PS1="[network-tools]\$ "' >> /root/.bash_profile

CMD ["/bin/bash", "--login"]
