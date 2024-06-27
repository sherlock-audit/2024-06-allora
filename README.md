
# Allora contest details

- Join [Sherlock Discord](https://discord.gg/MABEWyASkp)
- Submit findings using the issue page in your private contest repo (label issues as med or high)
- [Read for more details](https://docs.sherlock.xyz/audits/watsons)

# Q&A

# Audit scope


[allora-chain @ ccad6d27e55b27a7ec3b2aebd7e55f1bc26798ed](https://github.com/allora-network/allora-chain/tree/ccad6d27e55b27a7ec3b2aebd7e55f1bc26798ed)
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

[allora-inference-base @ 54844054aebc65e4d8039bb4603e6f60269f8f3b](https://github.com/allora-network/allora-inference-base/tree/54844054aebc65e4d8039bb4603e6f60269f8f3b)
- [allora-inference-base/cmd/keys/main.go](allora-inference-base/cmd/keys/main.go)
- [allora-inference-base/cmd/node/addresses.go](allora-inference-base/cmd/node/addresses.go)
- [allora-inference-base/cmd/node/appchain.go](allora-inference-base/cmd/node/appchain.go)
- [allora-inference-base/cmd/node/execute.go](allora-inference-base/cmd/node/execute.go)
- [allora-inference-base/cmd/node/flags.go](allora-inference-base/cmd/node/flags.go)
- [allora-inference-base/cmd/node/main.go](allora-inference-base/cmd/node/main.go)
- [allora-inference-base/cmd/node/nooplogger.go](allora-inference-base/cmd/node/nooplogger.go)
- [allora-inference-base/cmd/node/role.go](allora-inference-base/cmd/node/role.go)
- [allora-inference-base/cmd/node/types.go](allora-inference-base/cmd/node/types.go)




[allora-chain @ ccad6d27e55b27a7ec3b2aebd7e55f1bc26798ed](https://github.com/allora-network/allora-chain/tree/ccad6d27e55b27a7ec3b2aebd7e55f1bc26798ed)
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


