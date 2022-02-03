package erc721

import (
	"crypto/ecdsa"
	"github.com/datachainlab/fabric-besu-cross-demo/cmds/erc721/erc721/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
)

type ERC721CMD interface {
	Mint(receiver common.Address, tokenID *big.Int) (*types.Transaction, error)
	Approve(spender common.Address, tokenID *big.Int) (*types.Transaction, error)
	GetApproved(tokenID *big.Int) (common.Address, error)
	BalanceOf(owner common.Address) (*big.Int, error)
	OwnerOf(tokenID *big.Int) (common.Address, error)
}

type ERC721CMDImpl struct {
	conn    *ethclient.Client
	chainID int64
	pvtKey  *ecdsa.PrivateKey
	token   *contract.ERC721Token
}

func NewERC721CMDImpl(conn *ethclient.Client, chainID int64, pvtKey *ecdsa.PrivateKey, token *contract.ERC721Token) *ERC721CMDImpl {
	return &ERC721CMDImpl{
		conn,
		chainID,
		pvtKey,
		token,
	}
}

func (e *ERC721CMDImpl) Mint(receiver common.Address, tokenID *big.Int) (*types.Transaction, error) {
	signer, err := CreateTransactionSigner(e.conn, e.chainID, e.pvtKey)
	if err != nil {
		return nil, err
	}

	tx, err := e.token.Mint(signer, receiver, tokenID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (e *ERC721CMDImpl) Approve(spender common.Address, tokenID *big.Int) (*types.Transaction, error) {
	signer, err := CreateTransactionSigner(e.conn, e.chainID, e.pvtKey)
	if err != nil {
		return nil, err
	}

	tx, err := e.token.Approve(signer, spender, tokenID)
	if err != nil {
		return nil, err
	}

	return tx, nil
}

func (e *ERC721CMDImpl) GetApproved(tokenID *big.Int) (common.Address, error) {
	spender, err := e.token.GetApproved(nil, tokenID)
	if err != nil {
		return common.Address{}, err
	}

	return spender, nil
}

func (e *ERC721CMDImpl) BalanceOf(owner common.Address) (*big.Int, error) {
	count, err := e.token.BalanceOf(nil, owner)
	if err != nil {
		return nil, err
	}

	return count, nil
}

func (e *ERC721CMDImpl) OwnerOf(tokenID *big.Int) (common.Address, error) {
	owner, err := e.token.OwnerOf(nil, tokenID)
	if err != nil {
		return common.Address{}, err
	}

	return owner, nil
}
