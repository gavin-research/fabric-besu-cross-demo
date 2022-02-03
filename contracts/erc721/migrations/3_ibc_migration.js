const IBCHost = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHost");
const IBCClient = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCClient");
const IBCConnection = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCConnection");
const IBCChannel = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCChannel");
const IBCHandler = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCHandler");
const IBCMsgs = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCMsgs");
const IBCIdentifier = artifacts.require("@hyperledger-labs/yui-ibc-solidity/IBCIdentifier");
const CrossSimpleModule = artifacts.require("@datachainlab/cross-solidity/CrossSimpleModule");
const MultisigClient = artifacts.require("@datachainlab/ibc-proxy-solidity/MultisigClient");
const ProxyMultisigClient = artifacts.require("@datachainlab/ibc-proxy-solidity/ProxyMultisigClient");
const MockClient = artifacts.require("@hyperledger-labs/yui-ibc-solidity/MockClient");
const ERC721ContractModule = artifacts.require("ERC721ContractModule");

const PortCross = "cross"
const MockClientType = "mock-client"
const EthMultisigClientType = "ethmultisig-client"
const ProxyMultisigClientType = "proxyclient"

const deployCore = async (deployer) => {
  await deployer.deploy(IBCIdentifier);
  await deployer.link(IBCIdentifier, [
    IBCHost, IBCHandler,
    MultisigClient, ProxyMultisigClient
  ]);

  await deployer.deploy(IBCMsgs);
  await deployer.link(IBCMsgs, [
    IBCClient,
    IBCConnection,
    IBCChannel,
    IBCHandler,
  ]);

  await deployer.deploy(IBCClient);
  await deployer.link(IBCClient, [IBCHandler, IBCConnection, IBCChannel]);

  await deployer.deploy(IBCConnection);
  await deployer.link(IBCConnection, [IBCHandler, IBCChannel]);

  await deployer.deploy(IBCChannel);
  await deployer.link(IBCChannel, [IBCHandler]);

  await deployer.deploy(MockClient);
  await deployer.deploy(MultisigClient);
  await deployer.deploy(ProxyMultisigClient);

  await deployer.deploy(IBCHost);
  await deployer.deploy(IBCHandler, IBCHost.address);
};

const deployApp = async (deployer) => {
    await deployer.deploy(ERC721ContractModule);
    await deployer.deploy(CrossSimpleModule, IBCHost.address, IBCHandler.address, ERC721ContractModule.address, true);
};

const init = async (deployer) => {
  const ibcHost = await IBCHost.deployed();
  const ibcHandler = await IBCHandler.deployed();
  const erc20ContractModule = await ERC721ContractModule.deployed();

  for(const promise of [
    () => ibcHost.setIBCModule(IBCHandler.address),
    () => ibcHandler.bindPort(PortCross, CrossSimpleModule.address),
    () => ibcHandler.registerClient(MockClientType, MockClient.address),
    () => ibcHandler.registerClient(EthMultisigClientType, MultisigClient.address),
    () => ibcHandler.registerClient(ProxyMultisigClientType, ProxyMultisigClient.address),
    () => erc20ContractModule.setCaller(CrossSimpleModule.address)
  ]) {
    const result = await promise();
    console.log(result);
    if(!result.receipt.status) {
      throw new Error(`transaction failed to execute. ${result.tx}`);
    }
  }
}

module.exports = async function(deployer) {
  await deployCore(deployer);
  await deployApp(deployer);
  await init(deployer);
};
