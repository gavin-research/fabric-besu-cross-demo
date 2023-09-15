# Fabric Besu Cross Demo

We're going to set up a functional stage for Fabric vs. Besu, and then proceed to perform a similar implementation of Besu vs. Besu.

Below are the code snippets to kickstart the development:

```
$ git clone git@github.com:gavin-research/fabric-besu-cross-demo.git
$ cd fabric-besu-cross-demo
$ git switch besu-besu-cross-demo 
$ cd demo
$ make build
$ make network (fail... WIP Issue1)
```

WIP Issue 1
```
Creating besu_besu-node4_1                ... done
make -C ../../../contracts/erc721 truffle-migrate-demo
make[2]: Entering directory '/home/iadmin/wip/fabric-besu-cross-demo/contracts/erc721'
npm run compile

> fabric-besu-cross-demo-erc721@1.0.0 compile
> truffle compile

Error: EACCES: permission denied, open '/root/.config/truffle/config.json'
You don't have access to this file.
```

WIP Issue 2
```
make -C chains/fabric network
```

```
### Creating channel erc20-channel
2023-09-15 02:09:54.744 CEST [channelCmd] InitCmdFactory -> INFO 001 Endorser and orderer connections initialized
Error: got unexpected status: BAD_REQUEST -- error applying config update to existing channel 'erc20-channel': error authorizing update: error validating ReadSet: proposed update requires that key [Group]  /Channel/Application be at version 0, but it is currently at version 1
make[1]: *** [Makefile:62: channel] Error 1
make[1]: Leaving directory '/home/iadmin/fabric-besu-cross-demo/demo/chains/fabric'
make: *** [Makefile:14: network] Error 2
make: Leaving directory '/home/iadmin/fabric-besu-cross-demo/demo/chains/fabric'
iadmin@gavindev02:~/fabric-besu-cross-demo/demo$ 
```

# Linked Repositories


## Detailed Explanation

IBC-Token related commands

```
Usage:
  erc20cli erc20 [flags]
  erc20cli erc20 [command]

Available Commands:
  mint        Mint token
  approve     Approve token
  allowance   Get allowance
  balanceOf   Get balance
  totalSupply Get totalSupply

Flags:
  -h, --help   help for erc20

Global Flags:
      --home string   set home directory (default "/home/iadmin/.erc20cli")
```
```
Usage:
  erc721cli erc721 [command]

Available Commands:
  address     Get Contract address
  approve     Approve token
  balanceOf   Get balance
  getApproved Get owner for token
  mint        Mint token
  ownerOf     Get owner

Flags:
  -h, --help   help for erc721

Global Flags:
      --config string   config file (default is $HOME/.fabric-besu-cross-demo.yaml)
```


# Besu Besu Cross Demo

