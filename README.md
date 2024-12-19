# <h1 align="center"> Sentinel </h1>

<p align="center">
    <img src=".github/assets/README_COVER.JPEG" style="border-radius:5%" width="800" alt="">
</p>

<p align="center">
    Indexing, monitoring and alerting solution for MetaMorpho vaults time-locked operations.
</p>

## Table of Contents

- [Introduction](#introduction)
- [Architecture](#architecture)
- [Getting Started](#getting-started)
  - [Locally - From Source](#locally---from-source)
  - [Locally - Using Docker](#locally---using-docker)
  - [Production](#production)
- [Contributing](#contributing)
- [Authors](#authors)

## Introduction

On a MetaMorpho Vault, there are several actions with a potential impact against the user's interest.
The current implementation includes a timelock before the actions can be performed, which can be used to alert the users.

Sentinel will monitor the vaults, look for any timelocked actions and send appropriate notifications to the users.

## Architecture

Coming soon.

## Getting Started

### Locally - From Source

You can run Sentinel locally by forking the network where the MetaMorpho Vault is deployed.

First, you need to fork the network using anvil:
```shell
anvil -f $RPC_URL --auto-impersonate --fork-block-number 21429070
```

Then, you can start Sentinel:
```shell
go run main.go -c etc/config.example.yaml
```

Finally, you need to impersonate the administrator of the vault to perform an action that is timelocked:
```shell
# Sends ETH to the account to impersonate
cast send 0x2566f66f68ed438726AD904524FB306A03FdB80B --value 1ether --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Perform the action - Here decreasing the timelock from 96h to 95h59m59s
cast send 0x38989BBA00BDF8181F4082995b3DEAe96163aC5D --from 0x2566f66f68ed438726AD904524FB306A03FdB80B "submitTimelock(uint256)" 345599 --unlocked
```

And you should see a log similar to the following appear in your terminal:
```shell
INFO[2024-12-19T19:51:54+01:00]
INFO[2024-12-19T19:51:54+01:00]      _____            __  _            __
INFO[2024-12-19T19:51:54+01:00]     / ___/___  ____  / /_(_)___  ___  / /
INFO[2024-12-19T19:51:54+01:00]     \__ \/ _ \/ __ \/ __/ / __ \/ _ \/ /
INFO[2024-12-19T19:51:54+01:00]    ___/ /  __/ / / / /_/ / / / /  __/ /
INFO[2024-12-19T19:51:54+01:00]   /____/\___/_/ /_/\__/_/_/ /_/\___/_/
INFO[2024-12-19T19:51:54+01:00]
INFO[2024-12-19T19:51:54+01:00]
INFO[2024-12-19T19:51:54+01:00] All services initialized, starting sentinel..
INFO[2024-12-19T19:51:54+01:00] Fetched new block from EL Client              block_number=21429072
INFO[2024-12-19T19:51:54+01:00] Capturing events logs emitted by contracts at block  block_number=21429072 contract_addresses="[0xA9c3D3a366466Fa809d1Ae982Fb2c46E5fC41101 0x38989BBA00BDF8181F4082995b3DEAe96163aC5D]"
INFO[2024-12-19T19:51:54+01:00] Captured events logs emitted at block         block_number=21429072 captured_events_logs=1
INFO[2024-12-19T19:51:54+01:00] Decoding captured event log emitted by contract  event_contract_address=0x38989BBA00BDF8181F4082995b3DEAe96163aC5D
INFO[2024-12-19T19:51:54+01:00] Decoded timelocked action                     kind="Timelock Decrease"
INFO[2024-12-19T19:51:54+01:00] A new timelocked-action has been detected!    block_number=21429072 chain=mainnet current_timelock=96h0m0s description="Timelock decreased to 95h59m59s" kind="Timelock Decrease" valid_at="2024-12-22 12:25:06 +0100 CET" vault=0x38989BBA00BDF8181F4082995b3DEAe96163aC5D
```

### Locally - Using Docker

You can run Sentinel in Docker by forking the network where the MetaMorpho Vault is deployed.

First, you need to fork the network using anvil:
```shell
anvil -f $RPC_URL --auto-impersonate --fork-block-number 21429070
```

Then, you can start Sentinel:
```shell
docker compose up
```

Finally, you need to impersonate the administrator of the vault to perform an action that is timelocked:
```shell
# Sends ETH to the account to impersonate
cast send 0x2566f66f68ed438726AD904524FB306A03FdB80B --value 1ether --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80

# Perform the action - Here decreasing the timelock from 96h to 95h59m59s
cast send 0x38989BBA00BDF8181F4082995b3DEAe96163aC5D --from 0x2566f66f68ed438726AD904524FB306A03FdB80B "submitTimelock(uint256)" 345599 --unlocked
```

And you should see a log similar to the following appear in your terminal:
```shell
sentinel-1  | time="2024-12-19T19:05:23Z" level=info
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="     _____            __  _            __"
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="    / ___/___  ____  / /_(_)___  ___  / /"
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="    \\__ \\/ _ \\/ __ \\/ __/ / __ \\/ _ \\/ / "
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="   ___/ /  __/ / / / /_/ / / / /  __/ /  "
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="  /____/\\___/_/ /_/\\__/_/_/ /_/\\___/_/   "
sentinel-1  | time="2024-12-19T19:05:23Z" level=info
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="                                         "
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="All services initialized, starting sentinel.."
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="Fetched new block from EL Client" block_number=21429072
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="Capturing events logs emitted by contracts at block" block_number=21429072 contract_addresses="[0xA9c3D3a366466Fa809d1Ae982Fb2c46E5fC41101 0x38989BBA00BDF8181F4082995b3DEAe96163aC5D]"
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="Captured events logs emitted at block" block_number=21429072 captured_events_logs=1
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="Decoding captured event log emitted by contract" event_contract_address=0x38989BBA00BDF8181F4082995b3DEAe96163aC5D
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="Decoded timelocked action" kind="Timelock Decrease"
sentinel-1  | time="2024-12-19T19:05:23Z" level=info msg="A new timelocked-action has been detected!" block_number=21429072 chain=mainnet current_timelock=96h0m0s description="Timelock decreased to 95h59m59s" kind="Timelock Decrease" valid_at="2024-12-22 11:19:34 +0000 UTC" vault=0x38989BBA00BDF8181F4082995b3DEAe96163aC5D
```

Note that by default, the docker-compose file uses the `etc/config.example.docker.yaml`.

### Production

Coming soon.

## Contributing

We welcome all contributions, specially for new connectors to send notifications via new channels!

Check out [the dedicated file](./CONTRIBUTING.md) for more.

## Authors

This project is currently being maintained by the folks at [Quartz Technology](https://github.com/quartz-technology).

