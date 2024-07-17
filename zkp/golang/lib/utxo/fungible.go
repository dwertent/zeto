// Copyright © 2024 Kaleido, Inc.
//
// SPDX-License-Identifier: Apache-2.0
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package utxo

import (
	"math/big"

	"github.com/hyperledger-labs/zeto/lib/smt"
	"github.com/iden3/go-iden3-crypto/babyjub"
	"github.com/iden3/go-iden3-crypto/poseidon"
)

type Fungible struct {
	Amount *big.Int
	Owner  *babyjub.PublicKey
	Salt   *big.Int
}

func NewFungible(amount *big.Int, owner *babyjub.PublicKey, salt *big.Int) *Fungible {
	return &Fungible{
		Amount: amount,
		Owner:  owner,
		Salt:   salt,
	}
}

func (f *Fungible) CalculateIndex() (smt.NodeIndex, error) {
	hash, err := poseidon.Hash([]*big.Int{f.Amount, f.Salt, f.Owner.X, f.Owner.Y})
	if err != nil {
		return nil, err
	}
	return smt.NewNodeIndexFromBigInt(hash)
}