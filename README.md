
# Allora contest details

- Join [Sherlock Discord](https://discord.gg/MABEWyASkp)
- Submit findings using the issue page in your private contest repo (label issues as med or high)
- [Read for more details](https://docs.sherlock.xyz/audits/watsons)

# Q&A

### Q: On what chains are the smart contracts going to be deployed?
Our primary "smart contracts" consist of business logic (staking and math) in a cosmos-sdk based module. They run on our own cosmos chain.
___

### Q: If you are integrating tokens, are you allowing only whitelisted tokens to work with the codebase or any complying with the standard? Are they assumed to have certain properties, e.g. be non-reentrant? Are there any types of [weird tokens](https://github.com/d-xo/weird-erc20) you want to integrate?
We will in theory support token transfers of any cosmos-based IBC token. However, our system only interacts with the native currency of our chain, and we do not actually do any interactions with other types of tokens sent to our chain.
___

### Q: Are there any limitations on values set by admins (or other roles) in the codebase, including restrictions on array lengths?
No.
Any limitations on values passed to the chain by protocol-specific actors are explicitly caught via validation. There is no special knowledge of unwritten limitations expected of any protocol-specific actor.
___

### Q: Are there any limitations on values set by admins (or other roles) in protocols you integrate with, including restrictions on array lengths?
No.
There are no protocols we integrate with beyond IBC. We do not enforce any atypical constraints on our IBC code.

___

### Q: For permissioned functions, please list all checks and requirements that will be made before calling the function.
The permissioned functions are:
* AddToWhitelistAdmin
* RemoveFromWhitelistAdmin
* UpdateParams of which there are two: one in the `emissions` module and the other in the `mint` module

For each, the caller must simply have been added as a "whitelist admin" in the `emissions` module. This means that the emissions "tx function" `AddToWhitelistAdmin()` was called with an address of the admin, that `RemoveFromWhitelistAdmin()` was not successfully called with their address as an argument since last calling `AddToWhitelistAdmin()`, and that the address is a valid Bech32 string.

___

### Q: Is the codebase expected to comply with any EIPs? Can there be/are there any deviations from the specification?
No.
___

### Q: Are there any off-chain mechanisms or off-chain procedures for the protocol (keeper bots, arbitrage bots, etc.)?
Yes. These were also submitted as part of the audit. These are all facets of the Blockless network, including a head node that validators are expected to run in parallel to their node, and "Blockless worker nodes" that are ran by reputers and workers to submit losses or inferences and forecasts, respectively, when requested by the chain.

Each validator sets their env var to be a url that connects to their own Blockless head node. This is a simple server that initiates requests to all workers subscribed to a topic. We use a feature called "subgroups" where reputers and workers of each topics can subscribed to head nodes, specify what role they are to play (reputer or worker), specify the topic they are to play that role in (an integer). When a chain request is sent out, it includes data that is needed by reputers and workers to do their job, such as the nonce that determines which worker payloads should be judged by reputers.

After the head node receives a request, it conducts roll call, where workers that are part of that subgroup signal that they are ready to respond. The responsive workers then share their results with each other in a consensus mechanism between each other, signs their respective data, then a selected leader commits the batched data to the chain.

You can learn more about this and Blockless in general from their docs:
https://blockless.network/docs/protocol/core-concepts

Furthermore, delegators must claim their own rewards as they are awarded by calling `RewardDelegateStake()`, like a keeper of sorts.

___

### Q: Are there any hardcoded values that you intend to change before (some) deployments?
No.
___

### Q: If the codebase is to be deployed on an L2, what should be the behavior of the protocol in case of sequencer issues (if applicable)? Should Sherlock assume that the Sequencer won't misbehave, including going offline?
L2 not relevant. We're an L1.
___

### Q: Should potential issues, like broken assumptions about function behavior, be reported if they could pose risks in future integrations, even if they might not be an issue in the context of the scope? If yes, can you elaborate on properties/invariants that should hold?
Yes absolutely, here are the invariants/properties that should hold. 
Topics
* The size of ActiveTopics should always be less than nextTopicId
* The size of churnableTopics should always be less than or equal to the size of ActiveTopics
* The size of rewardableTopics should always be less than or equal to ChurnableTopics
* The size of churnableTopics should always be less than or equal to params.MaxTopicsPerBlock
Allora Staking (reputers, as opposed to cosmos validator staking):
* The balance of the totalStake in the emissions keeper should always equal the cosmos-sdk /x/bank balance of the "allorastaking" bank account
* the sum of all topicStake in every topic should always equal the totalStake
* the sum of all stakeReputerAuthority for a given topic should always equal the topic stake for that topic
* the sum of all stakeReputerAuthority should equal the total stake
* the stakeSumFromDelegator should always equal the sum of delegatedStakes[topicid, delegator, all reputers for that delegator])
* the stakeFromDelegatorsUponReputer should always equal the sum of delegatedStakes[topicid, all delegators for that reputer, reputer]
* the sum of all (delegateRewardsPerShare * delegated stake - reward debt) = the balance of the /x/bank AlloraPendingRewardForDelegatorAccountName module account
* the state of the stakeRemovalsByBlock and stakeRemovalsByActor should always be identical in terms of the number of keys in the map, and the content of those keys
* the sum of all stakeRemovals should always be less than or equal to the totalStake
Mint module invariants
* The annual percentage yield of a staker should not exceed 12%
___

### Q: Please discuss any design choices you made.
Delegation is modeled after the Sushi masterchef implementation, where the provisioning of rewards for delegators is maintained by themselves by calling `RewardDelegateStake()`. 

We implement a stake withdrawal delay on all actors with stake. This is a lengthy withdrawal period of time to ensure actors are committed to participating on the chain.

When stake is withdrawn, it still has the effect of being active so the consequences of placing stake there are fully felt. You can cancel a stake withdrawal at any time.

Like Tendermint, when your stake withdrawal period ends, the money is automatically sent to your wallet.
Similarly, when workers are rewarded, funds are automatically sent to their wallets.

We chose a "pull" pattern over "push" (from perspective of a worker) where the chain initiates Blockless requests to solicit losses and inferences. This offers us an efficiency gain by batching responses from actors via Blockless, so a fraction of the inbound requests are needed. The use of Blockless in this pattern also lowered the burden placed on nodes to connect to the Allora network, as Blockless abstracts much of the configuration necessary to connect to networks away, allowing workers and reputers to focus more on what they're best at.

We chose a "consumer-driven" pricing model for inferences because we expect that actors can drop out if they feel the price for their data are not worth their time. This plus off-chain communication mechanisms (e.g. public forums) are a fine, uncomplicated way to achieve price discovery. In this model, consumers pay the network for information from a particular topic, and those funds, as well as the topic's "weight", gradually drops per an exponential drip at the end of each topic's epoch. This ensures that the amount of funds a topic has left to spend is related to its "weight", so when funds are paid, the topic loses some probability of being churnable among the set of all topics.

We choose to implement a "competition" among topics by weight, where a topic with higher weight (one that has brought more value to ALLO via staked ALLO and/or consumer fees) is preferred to be churned over others. "Churning topics" means completing the process of soliciting inferences, judging them, then rewarding all involved. We only allow a fixed number of topics to be churned per block, which in theory drives competition among topics.

We use nonces to track requests based on the block that the request was sent. So if a worker request was sent at a particular height X, that is later associated to the returned inferences when they are returned. With this nonce, reputers can find said nonces then grade them.

We elected to combine much of the code into a monolithic module `emissions` purely for developer expediency.
___

### Q: Please list any known issues and explicitly state the acceptable risks for each known issue.
All admins are super admins in that they can add and remove each other. All risks associated with that are acceptable risks.

Admins can set bad parameter values. This is an acceptable risk.

___

### Q: We will report issues where the core protocol functionality is inaccessible for at least 7 days. Would you like to override this value?
Yes. 12 hours, please.
___

### Q: Please provide links to previous audits (if any).
We have a PDF of a previous, closed audit report for a slightly earlier version of the protocol we can share upon request.
___

### Q: Please list any relevant protocol resources.
Whitepaper:
https://whitepaper.assets.allora.network/whitepaper.pdf

Docs:
https://docs.allora.network/

Website:
https://www.allora.network/
___

### Q: Additional audit information.
The correct commit to use will be available on Friday morning for both repos requested. This may be the same as the commits presented at time of writing, but this is not guaranteed.

We ask that auditors look for both typical security vulnerabilities, but also departures of the codebase from the intentions defined in the whitepaper. This includes but is not limited to functions not correctly implemented or tied together as defined in the whitepaper. This requires an understanding of the math of the whitepaper, which we urge auditors to develop.
___



# Audit scope


[allora-inference-base @ 9630a7d691d48a8b0fbdda34dc1c13c3188b0706](https://github.com/allora-network/allora-inference-base/tree/9630a7d691d48a8b0fbdda34dc1c13c3188b0706)
- [allora-inference-base/cmd/keys/main.go](allora-inference-base/cmd/keys/main.go)
- [allora-inference-base/cmd/node/addresses.go](allora-inference-base/cmd/node/addresses.go)
- [allora-inference-base/cmd/node/appchain.go](allora-inference-base/cmd/node/appchain.go)
- [allora-inference-base/cmd/node/execute.go](allora-inference-base/cmd/node/execute.go)
- [allora-inference-base/cmd/node/flags.go](allora-inference-base/cmd/node/flags.go)
- [allora-inference-base/cmd/node/main.go](allora-inference-base/cmd/node/main.go)
- [allora-inference-base/cmd/node/nooplogger.go](allora-inference-base/cmd/node/nooplogger.go)
- [allora-inference-base/cmd/node/role.go](allora-inference-base/cmd/node/role.go)
- [allora-inference-base/cmd/node/types.go](allora-inference-base/cmd/node/types.go)

[allora-chain @ 3a97afe7af027c96749fac7c4327ae85359a61c8](https://github.com/allora-network/allora-chain/tree/3a97afe7af027c96749fac7c4327ae85359a61c8)
- [allora-chain/app/api.go](allora-chain/app/api.go)
- [allora-chain/app/app.go](allora-chain/app/app.go)
- [allora-chain/app/app.yaml](allora-chain/app/app.yaml)
- [allora-chain/app/export.go](allora-chain/app/export.go)
- [allora-chain/app/ibc.go](allora-chain/app/ibc.go)
- [allora-chain/app/params/config.go](allora-chain/app/params/config.go)
- [allora-chain/app/params/encoding.go](allora-chain/app/params/encoding.go)
- [allora-chain/app/topics_handler.go](allora-chain/app/topics_handler.go)
- [allora-chain/app/upgrades.go](allora-chain/app/upgrades.go)
- [allora-chain/app/upgrades/vintegration/upgrades.go](allora-chain/app/upgrades/vintegration/upgrades.go)
- [allora-chain/cmd/allorad/cmd/commands.go](allora-chain/cmd/allorad/cmd/commands.go)
- [allora-chain/cmd/allorad/cmd/root.go](allora-chain/cmd/allorad/cmd/root.go)
- [allora-chain/cmd/allorad/main.go](allora-chain/cmd/allorad/main.go)
- [allora-chain/math/collections.go](allora-chain/math/collections.go)
- [allora-chain/math/dec.go](allora-chain/math/dec.go)
- [allora-chain/math/utils.go](allora-chain/math/utils.go)
- [allora-chain/x/emissions/keeper/expected_keepers.go](allora-chain/x/emissions/keeper/expected_keepers.go)
- [allora-chain/x/emissions/keeper/genesis.go](allora-chain/x/emissions/keeper/genesis.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/common.go](allora-chain/x/emissions/keeper/inference_synthesis/common.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/network_inference_builder.go](allora-chain/x/emissions/keeper/inference_synthesis/network_inference_builder.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/network_inferences.go](allora-chain/x/emissions/keeper/inference_synthesis/network_inferences.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/network_losses.go](allora-chain/x/emissions/keeper/inference_synthesis/network_losses.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/network_regrets.go](allora-chain/x/emissions/keeper/inference_synthesis/network_regrets.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/nonce_mgmt.go](allora-chain/x/emissions/keeper/inference_synthesis/nonce_mgmt.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_bootstrap.go](allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_bootstrap.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_factory.go](allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_factory.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_forecast_implied.go](allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_forecast_implied.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_weight.go](allora-chain/x/emissions/keeper/inference_synthesis/synth_palette_weight.go)
- [allora-chain/x/emissions/keeper/inference_synthesis/types.go](allora-chain/x/emissions/keeper/inference_synthesis/types.go)
- [allora-chain/x/emissions/keeper/keeper.go](allora-chain/x/emissions/keeper/keeper.go)
- [allora-chain/x/emissions/keeper/migrator.go](allora-chain/x/emissions/keeper/migrator.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server.go](allora-chain/x/emissions/keeper/msgserver/msg_server.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_demand.go](allora-chain/x/emissions/keeper/msgserver/msg_server_demand.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_losses.go](allora-chain/x/emissions/keeper/msgserver/msg_server_losses.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_params.go](allora-chain/x/emissions/keeper/msgserver/msg_server_params.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_registrations.go](allora-chain/x/emissions/keeper/msgserver/msg_server_registrations.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_stake.go](allora-chain/x/emissions/keeper/msgserver/msg_server_stake.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_topics.go](allora-chain/x/emissions/keeper/msgserver/msg_server_topics.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_util_sort.go](allora-chain/x/emissions/keeper/msgserver/msg_server_util_sort.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_util_topic_activation.go](allora-chain/x/emissions/keeper/msgserver/msg_server_util_topic_activation.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_whitelist.go](allora-chain/x/emissions/keeper/msgserver/msg_server_whitelist.go)
- [allora-chain/x/emissions/keeper/msgserver/msg_server_worker_payload.go](allora-chain/x/emissions/keeper/msgserver/msg_server_worker_payload.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server.go](allora-chain/x/emissions/keeper/queryserver/query_server.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_forecasts.go](allora-chain/x/emissions/keeper/queryserver/query_server_forecasts.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_inferences.go](allora-chain/x/emissions/keeper/queryserver/query_server_inferences.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_losses.go](allora-chain/x/emissions/keeper/queryserver/query_server_losses.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_params.go](allora-chain/x/emissions/keeper/queryserver/query_server_params.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_registrations.go](allora-chain/x/emissions/keeper/queryserver/query_server_registrations.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_stake.go](allora-chain/x/emissions/keeper/queryserver/query_server_stake.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_topics.go](allora-chain/x/emissions/keeper/queryserver/query_server_topics.go)
- [allora-chain/x/emissions/keeper/queryserver/query_server_whitelist.go](allora-chain/x/emissions/keeper/queryserver/query_server_whitelist.go)
- [allora-chain/x/emissions/keeper/topic_weight.go](allora-chain/x/emissions/keeper/topic_weight.go)
- [allora-chain/x/emissions/migrations/v2/migrations.go](allora-chain/x/emissions/migrations/v2/migrations.go)
- [allora-chain/x/emissions/module/abci.go](allora-chain/x/emissions/module/abci.go)
- [allora-chain/x/emissions/module/autocli.go](allora-chain/x/emissions/module/autocli.go)
- [allora-chain/x/emissions/module/depinject.go](allora-chain/x/emissions/module/depinject.go)
- [allora-chain/x/emissions/module/module.go](allora-chain/x/emissions/module/module.go)
- [allora-chain/x/emissions/module/rewards/reputer_rewards.go](allora-chain/x/emissions/module/rewards/reputer_rewards.go)
- [allora-chain/x/emissions/module/rewards/rewards.go](allora-chain/x/emissions/module/rewards/rewards.go)
- [allora-chain/x/emissions/module/rewards/rewards_internal.go](allora-chain/x/emissions/module/rewards/rewards_internal.go)
- [allora-chain/x/emissions/module/rewards/scores.go](allora-chain/x/emissions/module/rewards/scores.go)
- [allora-chain/x/emissions/module/rewards/topic_rewards.go](allora-chain/x/emissions/module/rewards/topic_rewards.go)
- [allora-chain/x/emissions/module/rewards/topic_skimming.go](allora-chain/x/emissions/module/rewards/topic_skimming.go)
- [allora-chain/x/emissions/module/rewards/worker_rewards.go](allora-chain/x/emissions/module/rewards/worker_rewards.go)
- [allora-chain/x/emissions/types/codec.go](allora-chain/x/emissions/types/codec.go)
- [allora-chain/x/emissions/types/errors.go](allora-chain/x/emissions/types/errors.go)
- [allora-chain/x/emissions/types/events.go](allora-chain/x/emissions/types/events.go)
- [allora-chain/x/emissions/types/forecast.go](allora-chain/x/emissions/types/forecast.go)
- [allora-chain/x/emissions/types/genesis.go](allora-chain/x/emissions/types/genesis.go)
- [allora-chain/x/emissions/types/inference.go](allora-chain/x/emissions/types/inference.go)
- [allora-chain/x/emissions/types/keys.go](allora-chain/x/emissions/types/keys.go)
- [allora-chain/x/emissions/types/msg_create_topic.go](allora-chain/x/emissions/types/msg_create_topic.go)
- [allora-chain/x/emissions/types/msg_insert_bulk_reputer_payload.go](allora-chain/x/emissions/types/msg_insert_bulk_reputer_payload.go)
- [allora-chain/x/emissions/types/msg_insert_bulk_worker_payload.go](allora-chain/x/emissions/types/msg_insert_bulk_worker_payload.go)
- [allora-chain/x/emissions/types/msg_register.go](allora-chain/x/emissions/types/msg_register.go)
- [allora-chain/x/emissions/types/params.go](allora-chain/x/emissions/types/params.go)
- [allora-chain/x/emissions/types/reputer_value_bundle.go](allora-chain/x/emissions/types/reputer_value_bundle.go)
- [allora-chain/x/emissions/types/rewards.go](allora-chain/x/emissions/types/rewards.go)
- [allora-chain/x/emissions/types/types.go](allora-chain/x/emissions/types/types.go)
- [allora-chain/x/emissions/types/value_bundle_builder.go](allora-chain/x/emissions/types/value_bundle_builder.go)
- [allora-chain/x/emissions/types/withheld_worker_attributed_value.go](allora-chain/x/emissions/types/withheld_worker_attributed_value.go)
- [allora-chain/x/emissions/types/worker_attributed_value.go](allora-chain/x/emissions/types/worker_attributed_value.go)
- [allora-chain/x/emissions/types/worker_data_bundle.go](allora-chain/x/emissions/types/worker_data_bundle.go)
- [allora-chain/x/ibc/gmp/ibc_middleware.go](allora-chain/x/ibc/gmp/ibc_middleware.go)
- [allora-chain/x/ibc/gmp/types.go](allora-chain/x/ibc/gmp/types.go)
- [allora-chain/x/mint/keeper/emissions.go](allora-chain/x/mint/keeper/emissions.go)
- [allora-chain/x/mint/keeper/genesis.go](allora-chain/x/mint/keeper/genesis.go)
- [allora-chain/x/mint/keeper/keeper.go](allora-chain/x/mint/keeper/keeper.go)
- [allora-chain/x/mint/keeper/migrator.go](allora-chain/x/mint/keeper/migrator.go)
- [allora-chain/x/mint/keeper/msg_server.go](allora-chain/x/mint/keeper/msg_server.go)
- [allora-chain/x/mint/keeper/query_server.go](allora-chain/x/mint/keeper/query_server.go)
- [allora-chain/x/mint/module/abci.go](allora-chain/x/mint/module/abci.go)
- [allora-chain/x/mint/module/autocli.go](allora-chain/x/mint/module/autocli.go)
- [allora-chain/x/mint/module/module.go](allora-chain/x/mint/module/module.go)
- [allora-chain/x/mint/types/codec.go](allora-chain/x/mint/types/codec.go)
- [allora-chain/x/mint/types/errors.go](allora-chain/x/mint/types/errors.go)
- [allora-chain/x/mint/types/events.go](allora-chain/x/mint/types/events.go)
- [allora-chain/x/mint/types/expected_keepers.go](allora-chain/x/mint/types/expected_keepers.go)
- [allora-chain/x/mint/types/genesis.go](allora-chain/x/mint/types/genesis.go)
- [allora-chain/x/mint/types/keys.go](allora-chain/x/mint/types/keys.go)
- [allora-chain/x/mint/types/params.go](allora-chain/x/mint/types/params.go)




[allora-inference-base @ 9630a7d691d48a8b0fbdda34dc1c13c3188b0706](https://github.com/allora-network/allora-inference-base/tree/9630a7d691d48a8b0fbdda34dc1c13c3188b0706)
- [allora-inference-base/cmd/keys/main.go](allora-inference-base/cmd/keys/main.go)
- [allora-inference-base/cmd/node/addresses.go](allora-inference-base/cmd/node/addresses.go)
- [allora-inference-base/cmd/node/appchain.go](allora-inference-base/cmd/node/appchain.go)
- [allora-inference-base/cmd/node/execute.go](allora-inference-base/cmd/node/execute.go)
- [allora-inference-base/cmd/node/flags.go](allora-inference-base/cmd/node/flags.go)
- [allora-inference-base/cmd/node/main.go](allora-inference-base/cmd/node/main.go)
- [allora-inference-base/cmd/node/nooplogger.go](allora-inference-base/cmd/node/nooplogger.go)
- [allora-inference-base/cmd/node/role.go](allora-inference-base/cmd/node/role.go)
- [allora-inference-base/cmd/node/types.go](allora-inference-base/cmd/node/types.go)

