// +build debug 2k

package build

import (
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/lotus/chain/actors/policy"
)

const BootstrappersFile = ""
const GenesisFile = ""

// const UpgradeBreezeHeight = -1
// const BreezeGasTampingDuration = -1

// const UpgradeSmokeHeight = -1
// const UpgradeIgnitionHeight = -2
// const UpgradeRefuelHeight = -3
// const UpgradeTapeHeight = -4

// const UpgradeActorsV2Height = -1
// const UpgradeLiftoffHeight = -5

// const UpgradeKumquatHeight = 0
// const UpgradeCalicoHeight = 20
// const UpgradePersianHeight = 25
// const UpgradeOrangeHeight = 27
// const UpgradeClausHeight = 30

const UpgradeBreezeHeight = 100000
const BreezeGasTampingDuration = 1000001

const UpgradeSmokeHeight = 1000002
const UpgradeIgnitionHeight = 1000003
const UpgradeRefuelHeight = 1000004
const UpgradeTapeHeight = 1000006

const UpgradeActorsV2Height = 1000005
const UpgradeLiftoffHeight = 1000007

const UpgradeKumquatHeight = 1000008
const UpgradeCalicoHeight = 1000009
const UpgradePersianHeight = 1000010
const UpgradeOrangeHeight = 1000011
const UpgradeClausHeight = 1000012

var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
	0: DrandMainnet,
}

func init() {
	InsecurePoStValidation = true
	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
	policy.SetPreCommitChallengeDelay(abi.ChainEpoch(15))

	BuildType |= Build2k
}

const BlockDelaySecs = uint64(4)

const PropagationDelaySecs = uint64(1)

// SlashablePowerDelay is the number of epochs after ElectionPeriodStart, after
// which the miner is slashed
//
// Epochs
const SlashablePowerDelay = 20

// Epochs
const InteractivePoRepConfidence = 6

const BootstrapPeerThreshold = 1

// ///// +build debug 2k

// package build

// import (
// 	"math"
// 	"os"

// 	"github.com/filecoin-project/go-state-types/abi"
// 	"github.com/filecoin-project/lotus/chain/actors/policy"
// )

// const UpgradeBreezeHeight = 10000
// const BreezeGasTampingDuration = 100000

// const UpgradeSmokeHeight = 100001
// const UpgradeIgnitionHeight = 100002
// const UpgradeRefuelHeight = 100003
// const UpgradeTapeHeight = 100004

// var UpgradeActorsV2Height = abi.ChainEpoch(100000)
// var UpgradeLiftoffHeight = abi.ChainEpoch(100045)

// const UpgradeKumquatHeight = 15
// const UpgradeCalicoHeight = 20
// const UpgradePersianHeight = 25

// var DrandSchedule = map[abi.ChainEpoch]DrandEnum{
// 	0: DrandMainnet,
// }

// func init() {
// 	InsecurePoStValidation = true

// 	policy.SetSupportedProofTypes(abi.RegisteredSealProof_StackedDrg2KiBV1)
// 	policy.SetConsensusMinerMinPower(abi.NewStoragePower(2048))
// 	policy.SetMinVerifiedDealSize(abi.NewStoragePower(256))
// 	policy.SetPreCommitChallengeDelay(abi.ChainEpoch(15))
// 	if os.Getenv("LOTUS_DISABLE_V2_ACTOR_MIGRATION") == "1" {
// 		UpgradeActorsV2Height = math.MaxInt64
// 		UpgradeLiftoffHeight = 100004
// 	}

// 	BuildType |= Build2k
// }

// const BlockDelaySecs = uint64(4)

// const PropagationDelaySecs = uint64(1)

// // SlashablePowerDelay is the number of epochs after ElectionPeriodStart, after
// // which the miner is slashed
// //
// // Epochs
// const SlashablePowerDelay = 20

// // Epochs
// const InteractivePoRepConfidence = 6

// const BootstrapPeerThreshold = 1
