[![CircleCI](https://circleci.com/gh/ethereum/eth2-client-tests.svg?style=svg)](https://circleci.com/gh/ethereum/eth2-client-tests)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

# Ethereum 2.0 tests

This repository contains a series of sanity tests to be built against Ethereum 2.0 clients.

## `tester`

The tester program connects to an existing Genesis server and deploys, maintains and tests a testnet.

See https://www.github.com/Whiteblock/genesis to set up a testnet server.

### Check genesis is available locally

`tester genesis up`

### Create a testnet

`tester genesis tesnet --blockchain [prysm|artemis|lighthouse]`

Optionally, you can indicate a file to store the testnet ID.

`tester genesis tesnet --blockchain [prysm|artemis|lighthouse] -f out`

### Check all nodes in the testnet can serve traffic on a port

`tester network --testnet <testnetId>`

# Chains supported

| Name     | Image                           |
|----------|---------------------------------|
|Prysm     |dockerfiles/prysm.Dockerfile     |
|Artemis   |dockerfiles/artemis.Dockerfile   |
|Lighthouse|dockerfiles/lighthouse.Dockerfile|
|Lodestar  |dockerfiles/lodestar.Dockerfile  |

# Contributing

See the TODO file for possible contributions.