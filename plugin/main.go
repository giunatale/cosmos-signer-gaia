package main

import (
	pfmroutertypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"
	icatypes "github.com/cosmos/ibc-go/v8/modules/apps/27-interchain-accounts/types"
	ibctransfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	ibctypes "github.com/cosmos/ibc-go/v8/modules/core/types"
	ibctmtypes "github.com/cosmos/ibc-go/v8/modules/light-clients/07-tendermint"
	icsprovidertypes "github.com/cosmos/interchain-security/v5/x/ccv/provider/types"

	"github.com/cosmos/cosmos-sdk/codec"
	//evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	//upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	vestingtypes "github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
	authztypes "github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

var Cosmos_auth_v1beta1_RegisterLegacyAminoCodec = authtypes.RegisterLegacyAminoCodec
var Cosmos_auth_v1beta1_RegisterInterfaces = authtypes.RegisterInterfaces
var Cosmos_bank_v1beta1_RegisterLegacyAminoCodec = banktypes.RegisterLegacyAminoCodec
var Cosmos_bank_v1beta1_RegisterInterfaces = banktypes.RegisterInterfaces
var Cosmos_staking_v1beta1_RegisterLegacyAminoCodec = stakingtypes.RegisterLegacyAminoCodec
var Cosmos_staking_v1beta1_RegisterInterfaces = stakingtypes.RegisterInterfaces
var Cosmos_mint_v1beta1_RegisterLegacyAminoCodec = minttypes.RegisterLegacyAminoCodec
var Cosmos_mint_v1beta1_RegisterInterfaces = minttypes.RegisterInterfaces
var Cosmos_distr_v1beta1_RegisterLegacyAminoCodec = distrtypes.RegisterLegacyAminoCodec
var Cosmos_distr_v1beta1_RegisterInterfaces = distrtypes.RegisterInterfaces
var Cosmos_gov_v1beta1_RegisterLegacyAminoCodec = govtypes.RegisterLegacyAminoCodec
var Cosmos_gov_v1beta1_RegisterInterfaces = govtypes.RegisterInterfaces
var Cosmos_feegrant_v1beta1_RegisterLegacyAminoCodec = feegranttypes.RegisterLegacyAminoCodec
var Cosmos_feegrant_v1beta1_RegisterInterfaces = feegranttypes.RegisterInterfaces
var Cosmos_authz_v1beta1_RegisterLegacyAminoCodec = authztypes.RegisterLegacyAminoCodec
var Cosmos_authz_v1beta1_RegisterInterfaces = authztypes.RegisterInterfaces
var Cosmos_ibc_v1beta1_RegisterLegacyAminoCodec = func(*codec.LegacyAmino) {}
var Cosmos_ibc_v1beta1_RegisterInterfaces = ibctypes.RegisterInterfaces
var Cosmos_ibctm_v1_RegisterLegacyAminoCodec = func(*codec.LegacyAmino) {}
var Cosmos_ibctm_v1_RegisterInterfaces = ibctmtypes.RegisterInterfaces

/*
var Cosmos_upgrade_v1_RegisterLegacyAminoCodec = upgradetypes.RegisterLegacyAminoCodec
var Cosmos_upgrade_v1_RegisterInterfaces = upgradetypes.RegisterInterfaces
var Cosmos_evidence_v1_RegisterLegacyAminoCodec = evidencetypes.RegisterLegacyAminoCodec
var Cosmos_evidence_v1_RegisterInterfaces = evidencetypes.RegisterInterfaces
*/
var Cosmos_transfer_v1_RegisterLegacyAminoCodec = ibctransfertypes.RegisterLegacyAminoCodec
var Cosmos_transfer_v1_RegisterInterfaces = ibctransfertypes.RegisterInterfaces
var Cosmos_vesting_v1_RegisterLegacyAminoCodec = vestingtypes.RegisterLegacyAminoCodec
var Cosmos_vesting_v1_RegisterInterfaces = vestingtypes.RegisterInterfaces
var Cosmos_pfmrouter_v1_RegisterLegacyAminoCodec = pfmroutertypes.RegisterLegacyAminoCodec
var Cosmos_pfmrouter_v1_RegisterInterfaces = pfmroutertypes.RegisterInterfaces
var Cosmos_icsprovider_v1_RegisterLegacyAminoCodec = icsprovidertypes.RegisterLegacyAminoCodec
var Cosmos_icsprovider_v1_RegisterInterfaces = icsprovidertypes.RegisterInterfaces
var Cosmos_ica_v1_RegisterLegacyAminoCodec = func(*codec.LegacyAmino) {}
var Cosmos_ica_v1_RegisterInterfaces = icatypes.RegisterInterfaces

func main() {}
