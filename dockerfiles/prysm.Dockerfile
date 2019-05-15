FROM ubuntu:18.04

ENV DEBIAN_FRONTEND noninteractive
RUN apt-get update &&\
	apt-get install -y apt-utils expect git git-extras software-properties-common \
	inetutils-tools wget ca-certificates curl build-essential libssl-dev golang-go \
  	pkg-config zip g++ zlib1g-dev unzip python tmux openssh-server iperf3

RUN wget https://github.com/bazelbuild/bazel/releases/download/0.25.1/bazel-0.25.1-installer-linux-x86_64.sh
RUN chmod +x bazel-0.25.1-installer-linux-x86_64.sh
RUN ./bazel-0.25.1-installer-linux-x86_64.sh

RUN git clone https://github.com/prysmaticlabs/prysm.git
WORKDIR /prysm/
RUN bazel build //beacon-chain:beacon-chain

FROM ubuntu:18.04
WORKDIR /
COPY --from=0 /prysm/bazel-bin/beacon-chain/linux_amd64_stripped/beacon-chain .

ENTRYPOINT ["/beacon-chain"]