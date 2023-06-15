package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	ethermint "github.com/evmos/ethermint/types"
)

// Creates a new instance of the CSR object
func NewCSR(contracts []string, id uint64) CSR {
	return CSR{
		Contracts: contracts,
		Id:        id,
		Txs:       0,
		Revenue:   sdk.Int(sdk.ZeroUint()),
	}
}

// Validate performs stateless validation of a CSR object. This will check if
// there are duplicate smart contracts entered in the CSR, whether each smart contract
// is a valid eth address and check if the number of contracts is greater than 1.
func (csr CSR) Validate() error {
	seenSmartContracts := make(map[string]bool)
	for _, smartContract := range csr.Contracts {
		if err := ethermint.ValidateNonZeroAddress(smartContract); err != nil {
			return sdkerrors.Wrapf(ErrInvalidSmartContractAddress, "CSR::Validate one or more of the entered smart contract address are invalid: %s", smartContract)
		}

		if seenSmartContracts[smartContract] {
			return sdkerrors.Wrapf(ErrDuplicateSmartContracts, "CSR::Validate there are duplicate smart contracts in this CSR: %s", smartContract)
		}
		seenSmartContracts[smartContract] = true
	}

	// Ensure that there is at least one smart contract in the CSR Pool
	numSmartContracts := len(csr.Contracts)
	if numSmartContracts < 1 {
		return sdkerrors.Wrapf(ErrSmartContractSupply, "CSR::Validate # of smart contracts must be greater than 0 got: %d", numSmartContracts)
	}
	return nil
}
