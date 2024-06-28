package stress_test

import (
	cosmosMath "cosmossdk.io/math"
	alloraMath "github.com/allora-network/allora-chain/math"
	testCommon "github.com/allora-network/allora-chain/test/common"
	emissionstypes "github.com/allora-network/allora-chain/x/emissions/types"
	"github.com/stretchr/testify/require"
)

// Broadcast the tx to create a new topic
func createTopic(
	m testCommon.TestConfig,
	epochLength int64,
	creator NameAccountAndAddress,
) (topicId uint64) {
	createTopicRequest := &emissionstypes.MsgCreateNewTopic{
		Creator:         creator.aa.addr,
		Metadata:        "ETH 24h Prediction",
		LossLogic:       "bafybeid7mmrv5qr4w5un6c64a6kt2y4vce2vylsmfvnjt7z2wodngknway",
		LossMethod:      "loss-calculation-eth.wasm",
		InferenceLogic:  "bafybeigx43n7kho3gslauwtsenaxehki6ndjo3s63ahif3yc5pltno3pyq",
		InferenceMethod: "allora-inference-function.wasm",
		EpochLength:     epochLength,
		GroundTruthLag:  0,
		DefaultArg:      "ETH",
		PNorm:           alloraMath.NewDecFromInt64(3),
		AlphaRegret:     alloraMath.NewDecFromInt64(1),
		AllowNegative:   true,
		Tolerance:       alloraMath.MustNewDecFromString("0.01"),
	}

	txResp, err := m.Client.BroadcastTx(m.Ctx, creator.aa.acc, createTopicRequest)
	require.NoError(m.T, err)

	_, err = m.Client.WaitForTx(m.Ctx, txResp.TxHash)
	require.NoError(m.T, err)

	createTopicResponse := &emissionstypes.MsgCreateNewTopicResponse{}
	err = txResp.Decode(createTopicResponse)
	require.NoError(m.T, err)

	incrementCountTopics()

	m.T.Log(topicLog(createTopicResponse.TopicId, "creator", creator.name, "created topic"))

	return createTopicResponse.TopicId
}

// broadcast a tx to fund a topic
func fundTopic(
	m testCommon.TestConfig,
	topicId uint64,
	sender NameAccountAndAddress,
	amount int64,
) error {
	m.T.Log(topicLog(topicId, "funded topic with ", amount, "from", sender.name))
	txResp, err := m.Client.BroadcastTx(
		m.Ctx,
		sender.aa.acc,
		&emissionstypes.MsgFundTopic{
			Sender:  sender.aa.addr,
			TopicId: topicId,
			Amount:  cosmosMath.NewInt(amount),
		},
	)
	if err != nil {
		return err
	}
	_, err = m.Client.WaitForTx(m.Ctx, txResp.TxHash)
	if err != nil {
		return err
	}
	resp := &emissionstypes.MsgFundTopicResponse{}
	err = txResp.Decode(resp)
	if err != nil {
		return err
	}
	return nil
}
