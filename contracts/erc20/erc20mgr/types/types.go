package types

import (
	"crypto/sha256"
	"encoding/hex"
	"github.com/cosmos/cosmos-sdk/codec"
	authtypes "github.com/datachainlab/cross/x/core/auth/types"
	txtypes "github.com/datachainlab/cross/x/core/tx/types"
)

type ID string
type ClientID string
type MSPID string

// NewContractCallRequest creates a new instance of ContractCallRequest
func NewContractCallRequest(method string, args ...string) ContractCallRequest {
	return ContractCallRequest{
		Method: method,
		Args:   args,
	}
}

// ContractCallInfo converts the ContractCallRequest to a ContractCallInfo
func (r ContractCallRequest) ContractCallInfo(m codec.Codec) txtypes.ContractCallInfo {
	return m.MustMarshalJSON(&r)
}

// NOTE: This ID is an example. Do not use in production.
func NewID(mspID MSPID, clientID ClientID) (ID, error) {
	fabAcc := FabricAccount{mspID, clientID}
	fa, err := fabAcc.Marshal()
	if err != nil {
		return "", err
	}
	hash := sha256.Sum256(fa)
	id := hex.EncodeToString(hash[:])
	return ID(id), nil
}

func IDFromAccount(accountID authtypes.AccountID) ID {
	encoded := hex.EncodeToString(accountID)
	return ID(encoded)
}

func (i ID) Marshal() ([]byte, error) {
	return hex.DecodeString(string(i))
}
