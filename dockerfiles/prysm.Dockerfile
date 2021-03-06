FROM ubuntu:18.04

RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y apt-utils expect git git-extras software-properties-common \
	inetutils-tools wget ca-certificates curl build-essential libssl-dev golang-go \
  	pkg-config zip g++ zlib1g-dev unzip python tmux openssh-server iperf3

RUN wget https://github.com/bazelbuild/bazel/releases/download/0.25.1/bazel-0.25.1-installer-linux-x86_64.sh
RUN chmod +x bazel-0.25.1-installer-linux-x86_64.sh
RUN ./bazel-0.25.1-installer-linux-x86_64.sh

RUN git clone https://github.com/prysmaticlabs/prysm.git
WORKDIR /prysm/
RUN bazel build //beacon-chain:beacon-chain
RUN bazel build //validator:validator

FROM ubuntu:18.04
WORKDIR /
COPY --from=0 /prysm/bazel-bin/beacon-chain/linux_amd64_stripped/beacon-chain .
COPY --from=0 /prysm/bazel-bin/validator/linux_amd64_pure_stripped/validator .

RUN apt-get update && \
	DEBIAN_FRONTEND=noninteractive apt-get install -y apt-utils expect git git-extras software-properties-common \
	inetutils-tools wget ca-certificates curl build-essential libssl-dev golang-go \
  	pkg-config zip g++ zlib1g-dev unzip python tmux openssh-server iperf3 lsof

ENTRYPOINT ["/beacon-chain"]