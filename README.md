# ✨ So you want to run an audit

This `README.md` contains a set of checklists for our audit collaboration.

Your audit will use two repos: 
- **an _audit_ repo** (this one), which is used for scoping your audit and for providing information to wardens
- **a _findings_ repo**, where issues are submitted (shared with you after the audit) 

Ultimately, when we launch the audit, this repo will be made public and will contain the smart contracts to be reviewed and all the information needed for audit participants. The findings repo will be made public after the audit report is published and your team has mitigated the identified issues.

Some of the checklists in this doc are for **C4 (🐺)** and some of them are for **you as the audit sponsor (⭐️)**.

---

# Audit setup

## 🐺 C4: Set up repos
- [ ] Create a new private repo named `YYYY-MM-sponsorname` using this repo as a template.
- [ ] Rename this repo to reflect audit date (if applicable)
- [ ] Rename auditt H1 below
- [ ] Update pot sizes
- [ ] Fill in start and end times in audit bullets below
- [ ] Add link to submission form in audit details below
- [ ] Add the information from the scoping form to the "Scoping Details" section at the bottom of this readme.
- [ ] Add matching info to the Code4rena site
- [ ] Add sponsor to this private repo with 'maintain' level access.
- [ ] Send the sponsor contact the url for this repo to follow the instructions below and add contracts here. 
- [ ] Delete this checklist.

# Repo setup

## ⭐️ Sponsor: Add code to this repo

- [ ] Create a PR to this repo with the below changes:
- [ ] Provide a self-contained repository with working commands that will build (at least) all in-scope contracts, and commands that will run tests producing gas reports for the relevant contracts.
- [ ] Make sure your code is thoroughly commented using the [NatSpec format](https://docs.soliditylang.org/en/v0.5.10/natspec-format.html#natspec-format).
- [ ] Please have final versions of contracts and documentation added/updated in this repo **no less than 24 hours prior to audit start time.**
- [ ] Be prepared for a 🚨code freeze🚨 for the duration of the audit — important because it establishes a level playing field. We want to ensure everyone's looking at the same code, no matter when they look during the audit. (Note: this includes your own repo, since a PR can leak alpha to our wardens!)


---

## ⭐️ Sponsor: Edit this README

Under "SPONSORS ADD INFO HERE" heading below, include the following:

- [ ] Modify the bottom of this `README.md` file to describe how your code is supposed to work with links to any relevent documentation and any other criteria/details that the C4 Wardens should keep in mind when reviewing. ([Here's a well-constructed example.](https://github.com/code-423n4/2022-08-foundation#readme))
  - [ ] When linking, please provide all links as full absolute links versus relative links
  - [ ] All information should be provided in markdown format (HTML does not render on Code4rena.com)
- [ ] Under the "Scope" heading, provide the name of each contract and:
  - [ ] source lines of code (excluding blank lines and comments) in each
  - [ ] external contracts called in each
  - [ ] libraries used in each
- [ ] Describe any novel or unique curve logic or mathematical models implemented in the contracts
- [ ] Does the token conform to the ERC-20 standard? In what specific ways does it differ?
- [ ] Describe anything else that adds any special logic that makes your approach unique
- [ ] Identify any areas of specific concern in reviewing the code
- [ ] Review the Gas award pool amount. This can be adjusted up or down, based on your preference - just flag it for Code4rena staff so we can update the pool totals across all comms channels. 
- [ ] Optional / nice to have: pre-record a high-level overview of your protocol (not just specific smart contract functions). This saves wardens a lot of time wading through documentation.
- [ ] See also: [this checklist in Notion](https://code4rena.notion.site/Key-info-for-Code4rena-sponsors-f60764c4c4574bbf8e7a6dbd72cc49b4#0cafa01e6201462e9f78677a39e09746)
- [ ] Delete this checklist and all text above the line below when you're ready.

---

# Canto audit details
- Total Prize Pool: $36,500 USDC 
  - HM awards: $27,000 USDC 
  - Analysis awards: XXX XXX USDC (Notion: Analysis pool)
  - QA awards: $3,000 USDC 
  - Bot Race awards: XXX XXX USDC (Notion: Bot Race pool)
  - Gas awards: XXX XXX USDC (Notion: Gas pool)
  - Judge awards: XXX XXX USDC (Notion: Judge Fee)
  - Lookout awards: XXX XXX USDC (Notion: Sum of Pre-sort fee + Pre-sort early bonus)
  - Scout awards: $500 USDC 
- Join [C4 Discord](https://discord.gg/code4rena) to register
- Submit findings [using the C4 form](https://code4rena.com/contests/2023-06-canto/submit)
- [Read our guidelines for more details](https://docs.code4rena.com/roles/wardens)
- Starts June 20, 2023 20:00 UTC 
- Ends June 23, 2023 20:00 UTC 

## Automated Findings / Publicly Known Issues

Automated findings output for the audit can be found [here](add link to report) within 24 hours of audit opening.

*Note for C4 wardens: Anything included in the automated findings output is considered a publicly known issue and is ineligible for awards.*

[ ⭐️ SPONSORS ADD INFO HERE ]

# Overview of Onboarding Middleware

The x/onboarding module implements IBC middleware for users who don’t have Canto tokens for initial gas spending. Additionally, this module will also convert IBC assets to their ERC20 equivalent if the mapping is already registered.

## Background Information
Canto is a layer-1 EVM built using Cosmos SDK. As such, IBC assets do not exist on the EVM unless an IBC <-> ERC20 mapping is registered. Once the mapping is created, users are able to convert between ERC20s (to use on the EVM) and IBC assets (to IBC elsewhere).  

# Abstract

We introduce a new module called `onboarding` to help users outside of Canto onboard seamlessly.

The module automatically swaps a portion of the assets being transferred to Canto network via IBC transfer for Canto without the need for a manual process, and converts the remaining assets to ERC20 tokens on Canto.

# Specification

When users transfer assets to the Canto network through Gravity Bridge, the IBC transfer automatically triggers swap and conversion to Canto via IBC middleware. These actions are triggered only when transferred through a whitelisted channel.

## User Flow

- User transfers assets to the Canto network through Gravity Bridge
- Check recipient address's Canto balance
- If the balance is less than the minimum threshold (e.g. `4canto`), swap the assets to Canto using the coinswap module
- If the remaining assets are registered as a ERC20, convert them to ERC20.

`Example:`
- User bridges 100 USDC to Canto using IBC transfer
- A small portion of the USDC will be swapped so the user receives 4 Canto tokens
- The remaining USDC will converted from an IBC asset to ERC20 representation automatically
- Now users have Canto token to spend on gas fees and do not have to manually convert their USDC IBC asset to an ERC20


## Module Parameters

```go
message Params {
  bool enable_onboarding = 1;

  string auto_swap_threshold = 3 [
    (gogoproto.customtype) = "github.com/cosmos/cosmos-sdk/types.Int",
    (gogoproto.nullable) = false
  ];

  repeated string whitelisted_channels = 4;
}
```

- **`EnableOnboarding`:** If this value is `false`, then it will disable the auto swap and convert. (default value: `true`)
- **`AutoSwapThreshold`:** the threshold of the amount of canto to be swapped. When the balance of canto is less than the threshold, the auto swap will be triggered. (default value: `4 canto`)
- **`WhitelistedChannels`:** The list of channels that will be whitelisted for the auto swap and convert. When the channel is not in the list, the auto swap and convert will not be triggered. (default value: `["channel-0"]` which is a channel for Gravity Bridge)

## OnRecvPacket

The `onboarding` module implements an IBC Middleware in order to swap and convert IBC transferred assets to Canto and ERC20 tokens with `Keeper.OnRecvPacket` callback.

1. A user performs an IBC transfer to the Canto network. This is done using a `FungibleTokenPacket` IBC packet.
2. Check that the onboarding conditions are met and skip to the next middleware if any condition is not satisfied:
   1. onboarding is enabled globally
   2. channel is authorized
   3. the recipient account is not a module account
3. Check the recipient's Canto balance and if the balance is less than the `AutoSwapThreshold`, swap the assets to Canto. Amount of the swapped Canto is always equal to the `AutoSwapThreshold` and the price is determined by the liquidity pool.
4. Check if the transferred asset is registered in the `x/erc20` module as a ERC20 token pair and the token pair is enabled. If so, convert the remaining assets to ERC20 tokens.

```go
func (k Keeper) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	ack exported.Acknowledgement,
) exported.Acknowledgement {
	// It always returns original ACK if the packet was a ICS20 transfer.
	// Which means even if the swap or conversion fails, it does not revert IBC transfer.
	// The asset transferred to the Canto network will still remain in the Canto network.
	params := k.GetParams(ctx)
	if !params.EnableOnboarding {
		return ack
	}
	// check source channel is in the whitelist channels
	var found bool
	for _, s := range params.WhitelistedChannels {
		if s == packet.DestinationChannel {
			found = true
		}
	}
	if !found {
		return ack
	}
	// Get recipient addresses in `canto1` and the original bech32 format
	_, recipient, senderBech32, recipientBech32, err := ibc.GetTransferSenderRecipient(packet)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}
	account := k.accountKeeper.GetAccount(ctx, recipient)
	// onboarding is not supported for module accounts
	if _, isModuleAccount := account.(authtypes.ModuleAccountI); isModuleAccount {
		return ack
	}
	standardDenom := k.coinswapKeeper.GetStandardDenom(ctx)
	var data transfertypes.FungibleTokenPacketData
	if err = transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		// NOTE: shouldn't happen as the packet has already
		// been decoded on ICS20 transfer logic
		err = errorsmod.Wrapf(types.ErrInvalidType, "cannot unmarshal ICS-20 transfer packet data")
		return channeltypes.NewErrorAcknowledgement(err.Error())
	}

	// parse the transferred denom
	transferredCoin := ibc.GetReceivedCoin(
		packet.SourcePort, packet.SourceChannel,
		packet.DestinationPort, packet.DestinationChannel,
		data.Denom, data.Amount,
	)
	autoSwapThreshold := k.GetParams(ctx).AutoSwapThreshold
	swapCoins := sdk.NewCoin(standardDenom, autoSwapThreshold)
	standardCoinBalance := k.bankKeeper.SpendableCoins(ctx, recipient).AmountOf(standardDenom)
	swappedAmount := sdk.ZeroInt()
	// Swap for users who have less canto than the autoSwapThreshold.
	if standardCoinBalance.LT(autoSwapThreshold) {
		swappedAmount, err = k.coinswapKeeper.TradeInputForExactOutput(ctx, coinswaptypes.Input{Coin: transferredCoin, Address: recipient.String()}, coinswaptypes.Output{Coin: swapCoins, Address: recipient.String()})
		if err != nil {
			// no-op: proceed with the remaining logic regardless of 
      // whether the swap is successful or not.
		}		
	}

	// convert coins to ERC20 token if denom is registered in erc20 module.
	pairID := k.erc20Keeper.GetTokenPairID(ctx, transferredCoin.Denom)
	if len(pairID) == 0 {
		// short-circuit: if the denom is not registered, conversion will fail
		// so we can continue with the rest of the stack
		return ack
	}
	pair, _ := k.erc20Keeper.GetTokenPair(ctx, pairID)
	if !pair.Enabled {
		// no-op: continue with the rest of the stack without conversion
		return ack
	}
	convertCoin := sdk.NewCoin(transferredCoin.Denom, transferredCoin.Amount.Sub(swappedAmount))
	// Build MsgConvertCoin, from recipient to recipient since IBC transfer already occurred
	convertMsg := erc20types.NewMsgConvertCoin(convertCoin, common.BytesToAddress(recipient.Bytes()), recipient)

	// NOTE: we don't use ValidateBasic the msg since we've already validated
	// the ICS20 packet data
	// Use MsgConvertCoin to convert the Cosmos Coin to an ERC20
	if _, err = k.erc20Keeper.ConvertCoin(sdk.WrapSDKContext(ctx), convertMsg); err != nil {
		return ack
	}

	// return original acknowledgement
	return ack
}
```

It is possible that the IBC transaction fails in any point of the stack execution and in that case the onboarding will not be triggered by the transaction, as it will rollback to the previous state. However, the onboarding process is non-atomic, meaning that even if the swap or conversion fails, it does not revert IBC transfer and the asset transferred to the Canto network will still remain in the Canto network.

## Swap

For swap, we use a forked version of IRISNET's [Coinswap module v1.6](https://github.com/irisnet/irismod/tree/v1.6.0/modules/coinswap), which includes some modifications.

IRISNET's Coinswap module uses an AMM-based swap. This means that onboarding swaps will be handled by AMM also. However, there are some modifications:

- Only token pairs on the whitelist can be created as a pool.
  - Pool creation fails if the token pair is not on the whitelist.
  - Initial whitelist: `Canto/USDC.grv`, `Canto/USDT.grv`, `Canto/ETH.grv`
- There is a limit on the number of Canto tokens for each pool.
  - Deposits will fail if the amount of Canto for the pool exceeds 10,000 Canto.
- Double swaps are disabled.

For risk management purposes, a swap will fail if the input coin amount exceeds a pre-defined limit (10 USDC, 10 USDT, 0.01 ETH) or if the swap amount limit is not defined.

```go
// Only those IBC denom tokens are allowed to convert to Canto.
const (
	UsdcIBCDenom = "ibc/17CD484EE7D9723B847D95015FA3EBD1572FD13BC84FB838F55B18A57450F25B"
	UsdtIBCDenom = "ibc/4F6A2DEFEA52CD8D90966ADCB2BD0593D3993AB0DF7F6AEB3EFD6167D79237B0"
	EthIBCDenom  = "ibc/DC186CA7A8C009B43774EBDC825C935CABA9743504CE6037507E6E5CCE12858A"
)

// Default parameters
var (
	DefaultFee                    = sdk.NewDecWithPrec(0, 0)
	DefaultPoolCreationFee        = sdk.NewInt64Coin(sdk.DefaultBondDenom, 0)
	DefaultTaxRate                = sdk.NewDecWithPrec(0, 0)
	// Limit the number of canto tokens for each pool.
	DefaultMaxStandardCoinPerPool = sdk.NewIntWithDecimal(10000, 18)
	// Pre-defined limits of swap amount for risk management purposes.
	DefaultMaxSwapAmount          = sdk.NewCoins(
		sdk.NewCoin(UsdcIBCDenom, sdk.NewIntWithDecimal(10, 6)),
		sdk.NewCoin(UsdtIBCDenom, sdk.NewIntWithDecimal(10, 6)),
		sdk.NewCoin(EthIBCDenom, sdk.NewIntWithDecimal(1, 17)),
	)
)
```

## Middleware ordering

The IBC middleware adds custom logic between the core IBC and the underlying application. Middlewares are implemented as stacks so that applications can define multiple layers of custom behavior. Function calls from the IBC core to the application travel from the top-level middleware to the bottom middleware, and then to the application.

For Canto the middleware stack ordering is defined as follows (from top to bottom):

1. IBC Transfer
2. Recovery Middleware
3. Onboarding Middleware

```go
// app.go
// create IBC module from bottom to top of stack
var transferStack porttypes.IBCModule

transferStack = transfer.NewIBCModule(app.TransferKeeper)
transferStack = recovery.NewIBCMiddleware(*app.RecoveryKeeper, transferStack)
transferStack = onboarding.NewIBCMiddleware(*app.OnboardingKeeper, transferStack)
```

Each module implements their own custom logic in the packet callback `OnRecvPacket`. When a packet arrives from the IBC core, the IBC transfer will be executed first, followed by an attempted recovery, and finally the onboarding will be executed

# Scope

| Contract | SLOC | Purpose | Modules used |  
| ----------- | ----------- | ----------- | ----------- |
| [x/onboarding](add-link) | --- | Contains logic for onboarding middleware | [`x/coinswap`](add-link) |
| [x/onboarding/ibc_callbacks.go](add-link) | 55 | Contains core logic | [`x/coinswap`](add-link) |
| [x/coinswap/pool.go](add-link) | 85 | Contains logic for dex pools |  |
| [x/coinswap/swap.go](add-link) | 102 | Contains logic for dex swaps |  |

## Out of scope

*All other modules are out of scope*

## Scoping Details 
```
- If you have a public code repo, please share it here:  https://github.com/Canto-Network/Canto
- How many contracts are in scope?:   0
- Total SLoC for these contracts?:  349
- How many external imports are there?: 3  
- How many separate interfaces and struct definitions are there for the contracts within scope?:  0
- Does most of your code generally use composition or inheritance?:   Composition
- How many external calls?:   0
- What is the overall line coverage percentage provided by your tests?:  100 for business logic
- Is there a need to understand a separate part of the codebase / get context in order to audit this part of the protocol?: true  
- Please describe required context:   This audit is for a Cosmos SDK module. Therefore, understanding of how Cosmos SDK blockchains work is required.
- Does it use an oracle?:  No
- Does the token conform to the ERC20 standard?:  true
- Are there any novel or unique curve logic or mathematical models?: No
- Does it use a timelock function?:  
- Is it an NFT?: 
- Does it have an AMM?: true   
- Is it a fork of a popular project?:   false
- Does it use rollups?:   
- Is it multi-chain?:  
- Does it use a side-chain?: false
- Describe any specific areas you would like addressed.: Our module involves IBC transfers and an AMM. We would like wardens to focus on these aspects.
- Is this fresh code or have they been audited, if so by who?: Fresh code
```

# Tests
### Testing Onboarding Middelware
```bash
cd Canto/x/onboarding
go test ./...
```
### Testing Coinswap Module
```bash
cd Canto/x/coinswap
go test ./...
```