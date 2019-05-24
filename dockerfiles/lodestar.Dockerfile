FROM ubuntu:18.04

# Install dependencies (namely nodejs)
RUN env DEBIAN_FRONTEND=noninteractive \
    apt-get update && \
	apt-get install -y apt-utils build-essential curl git && \
    curl -sL https://deb.nodesource.com/setup_10.x | bash - && \
    apt-get install -y nodejs && \
    curl -o- -L https://yarnpkg.com/install.sh | bash

RUN env DEBIAN_FRONTEND=noninteractive \
    apt-get install -y apt-utils expect git git-extras software-properties-common \
    	inetutils-tools wget ca-certificates curl build-essential libssl-dev \
      	pkg-config zip zlib1g-dev unzip python tmux openssh-server iperf3 lsof

# Install lodestar
RUN git clone https://github.com/ChainSafe/lodestar.git
RUN cd /lodestar && $HOME/.yarn/bin/yarn install --pure-lockfile
RUN cd /lodestar && $HOME/.yarn/bin/yarn run build



ENTRYPOINT ["/bin/bash"]
