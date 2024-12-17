# <h1 align="center"> Sentinel </h1>

<p align="center">
    <img src=".github/assets/README_COVER.JPEG" style="border-radius:5%" width="800" alt="">
</p>

<p align="center">
    Indexing, monitoring and alerting solution for Morpho vaults time-locked operations.
</p>

## Introduction

On a MetaMorpho Vault, there are several actions with a potential impact against the user's interest.
The current implementation includes a timelock before the actions can be performed, which can be used to alert the users.

Sentinel will monitor the vaults, look for any timelocked actions and send appropriate notifications to the users.

## Architecture

Coming soon.

## Getting Started

### Locally

You can run Sentinel locally by forking the network where the MetaMorpho Vault is deployed.

First, you need to fork the network using anvil:
```shell
anvil -f $RPC_URL --auto-impersonate
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
INFO[2024-12-17T18:24:55+01:00]
INFO[2024-12-17T18:24:55+01:00]      _____            __  _            __
INFO[2024-12-17T18:24:55+01:00]     / ___/___  ____  / /_(_)___  ___  / /
INFO[2024-12-17T18:24:55+01:00]     \__ \/ _ \/ __ \/ __/ / __ \/ _ \/ /
INFO[2024-12-17T18:24:55+01:00]    ___/ /  __/ / / / /_/ / / / /  __/ /
INFO[2024-12-17T18:24:55+01:00]   /____/\___/_/ /_/\__/_/_/ /_/\___/_/
INFO[2024-12-17T18:24:55+01:00]
INFO[2024-12-17T18:24:55+01:00]
DEBU[2024-12-17T18:24:55+01:00] All services initialized, starting sentinel..
INFO[2024-12-17T18:25:09+01:00] A new timelocked-action has been detected!    block_number=21423732 current_timelock=96h0m0s description="Timelock decreased to 95h59m59s" valid_at="2024-12-21 18:24:59 +0100 CET" vault=0x38989BBA00BDF8181F4082995b3DEAe96163aC5D
```

### Production

Coming soon.

## Contributing

We welcome all contributions, specially for new connectors to send notifications via new channels!

Check out [the dedicated file](./CONTRIBUTING.md) for more.

## Authors

This project is currently being maintained by the folks at [Quartz Technology](https://github.com/quartz-technology).

