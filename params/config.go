// Copyright 2016 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import (
	"encoding/json"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
)

// Genesis hashes to enforce below configs on.
var (
	MainnetGenesisHash      = common.HexToHash("0x0c6e644fcbd396f7b235ecef44551c45afd9274e87cd77ec6e9778cf8bfb46fc")
	MainnetSnailGenesisHash = common.HexToHash("0xf82fd9c0c8a53474c9e40e4f1c0583a94609eaf88dae01a5496da459398485c6")

	TestnetGenesisHash      = common.HexToHash("0x4b82a68ebbf32f2e816754f2b50eda0ae2c0a71dd5f4e0ecd93ccbfb7dba00b8")
	TestnetSnailGenesisHash = common.HexToHash("0x4ab1748c057b744de202d6ebea64e8d3a0b2ec4c19abbc59e8639967b14b7c96")

	DevnetSnailGenesisHash = common.HexToHash("0xdf819f11beead767f91a6c05d74e5f902fc2988e9039a969a023bc75e467cdeb")
)

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetSnailGenesisHash: MainnetTrustedCheckpoint,
	TestnetSnailGenesisHash: TestnetTrustedCheckpoint,
	DevnetSnailGenesisHash:  DevnetTrustedCheckpoint,
}

// TrustedCheckpoints associates each known checkpoint with the genesis hash of
// the chain it belongs to.
var TrustedBloomCheckpoints = map[common.Hash]*TrustedCheckpoint{
	MainnetSnailGenesisHash: MainnetTrustedCheckpoint,
	TestnetSnailGenesisHash: TestnetTrustedCheckpoint,
	DevnetSnailGenesisHash:  DevnetTrustedBloomCheckpoint,
}

var (
	// MainnetChainConfig is the chain parameters to run a node on the main network.
	MainnetChainConfig = &ChainConfig{
		ChainID: big.NewInt(19330),
		Minerva: &(MinervaConfig{
			MinimumDifficulty:      big.NewInt(134217728),
			MinimumFruitDifficulty: big.NewInt(262144),
			DurationLimit:          big.NewInt(600),
		}),
		TIP3: &BlockConfig{FastNumber: big.NewInt(1500000)},
		TIP5: &BlockConfig{SnailNumber: big.NewInt(12800)},
	}

	// MainnetTrustedCheckpoint contains the light client trusted checkpoint for the main network.
	MainnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "mainnet",
		SectionIndex: 227,
		SectionHead:  common.HexToHash("0xa2e0b25d72c2fc6e35a7f853cdacb193b4b4f95c606accf7f8fa8415283582c7"),
		CHTRoot:      common.HexToHash("0xf69bdd4053b95b61a27b106a0e86103d791edd8574950dc96aa351ab9b9f1aa0"),
		BloomRoot:    common.HexToHash("0xec1b454d4c6322c78ccedf76ac922a8698c3cac4d98748a84af4995b7bd3d744"),
	}

	// TestnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	TestnetChainConfig = &ChainConfig{
		ChainID: big.NewInt(18928),
		Minerva: &(MinervaConfig{
			MinimumDifficulty:      big.NewInt(60000),
			MinimumFruitDifficulty: big.NewInt(200),
			DurationLimit:          big.NewInt(600),
		}),
		TIP3: &BlockConfig{FastNumber: big.NewInt(450000)},
		TIP5: &BlockConfig{SnailNumber: big.NewInt(4000)},
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	TestnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "testnet",
		SectionIndex: 161,
		SectionHead:  common.HexToHash("0x5378afa734e1feafb34bcca1534c4d96952b754579b96a4afb23d5301ecececc"),
		CHTRoot:      common.HexToHash("0x1cf2b071e7443a62914362486b613ff30f60cea0d9c268ed8c545f876a3ee60c"),
		BloomRoot:    common.HexToHash("0x5ac25c84bd18a9cbe878d4609a80220f57f85037a112644532412ba0d498a31b"),
	}

	// DevnetChainConfig contains the chain parameters to run a node on the Ropsten test network.
	DevnetChainConfig = &ChainConfig{
		ChainID: big.NewInt(100),
		Minerva: &(MinervaConfig{
			MinimumDifficulty:      big.NewInt(10000),
			MinimumFruitDifficulty: big.NewInt(100),
			DurationLimit:          big.NewInt(150),
		}),
		TIP3: &BlockConfig{FastNumber: big.NewInt(380000)},
		TIP5: &BlockConfig{SnailNumber: big.NewInt(5000)},
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	DevnetTrustedCheckpoint = &TrustedCheckpoint{
		Name:         "devnet",
		SectionIndex: 55,
		SectionHead:  common.HexToHash("0x9f88b11036f4302ea84bf1065af0f0e7c78a5b03098f6eecc22c69a43a2a3357"),
		CHTRoot:      common.HexToHash("0x204dc825e47f728ff911e303e46849216f21452577727f328f089de815983a66"),
		BloomRoot:    common.HexToHash("0x5ac25c84bd18a9cbe878d4609a80220f57f85037a112644532412ba0d498a31b"),
	}

	// TestnetTrustedCheckpoint contains the light client trusted checkpoint for the Ropsten test network.
	DevnetTrustedBloomCheckpoint = &TrustedCheckpoint{
		Name:         "devnet",
		SectionIndex: 2,
		SectionHead:  common.HexToHash("9cd5e2ec8a8505dd5b8dc7bc248daceafbb135a84e68e23ba0d9b12e1be2651a"),
		CHTRoot:      common.HexToHash("1a91e78eb6772ed586a1a31b5058876b664d96bddbf5d9f2942ea302da403d2d"),
		BloomRoot:    common.HexToHash("0x5ac25c84bd18a9cbe878d4609a80220f57f85037a112644532412ba0d498a31b"),
	}

	chainId = big.NewInt(9223372036854775790)
	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.
	AllMinervaProtocolChanges = &ChainConfig{ChainID: chainId, Minerva: new(MinervaConfig), TIP3: &BlockConfig{FastNumber: big.NewInt(0)}, TIP5: nil}

	// This configuration is intentionally not using keyed fields to force anyone
	// adding flags to the config to also have to set these fields.

	TestChainConfig = &ChainConfig{ChainID: chainId, Minerva: &MinervaConfig{MinimumDifficulty, MinimumFruitDifficulty, DurationLimit}, TIP3: &BlockConfig{FastNumber: big.NewInt(0)}, TIP5: nil}
)

// TrustedCheckpoint represents a set of post-processed trie roots (CHT and
// BloomTrie) associated with the appropriate section index and head hash. It is
// used to start light syncing from this checkpoint and avoid downloading the
// entire header chain while still being able to securely access old headers/logs.
type TrustedCheckpoint struct {
	Name         string      `json:"-"`
	SectionIndex uint64      `json:"sectionIndex"`
	SectionHead  common.Hash `json:"sectionHead"`
	CHTRoot      common.Hash `json:"chtRoot"`
	BloomRoot    common.Hash `json:"bloomRoot"`
}

// ChainConfig is the core config which determines the blockchain settings.
//
// ChainConfig is stored in the database on a per block basis. This means
// that any network, identified by its genesis block, can have its own
// set of configuration options.
type ChainConfig struct {
	ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

	// Various consensus engines
	Minerva *MinervaConfig `json:"minerva"`
	//Clique *CliqueConfig  `json:"clique,omitempty"`

	TIP3 *BlockConfig `json:"tip3"`

	TIP5 *BlockConfig `json:"tip5"`
}

type BlockConfig struct {
	FastNumber  *big.Int
	SnailNumber *big.Int
}

func (c *ChainConfig) UnmarshalJSON(input []byte) error {
	type ChainConfig struct {
		ChainID *big.Int `json:"chainId"` // chainId identifies the current chain and is used for replay protection

		Minerva *MinervaConfig `json:"minerva"`
	}
	var dec ChainConfig
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	c.ChainID = dec.ChainID
	if dec.Minerva == nil {
		c.Minerva = &(MinervaConfig{
			MinimumDifficulty:      MinimumDifficulty,
			MinimumFruitDifficulty: MinimumFruitDifficulty,
			DurationLimit:          DurationLimit,
		})
	} else {
		c.Minerva = dec.Minerva
	}

	return nil
}

// MinervaConfig is the consensus engine configs for proof-of-work based sealing.
type MinervaConfig struct {
	MinimumDifficulty      *big.Int `json:"minimumDifficulty"`
	MinimumFruitDifficulty *big.Int `json:"minimumFruitDifficulty"`
	DurationLimit          *big.Int `json:"durationLimit"`
}

func (c *MinervaConfig) UnmarshalJSON(input []byte) error {
	type MinervaConfig struct {
		MinimumDifficulty      *math.HexOrDecimal256 `json:"minimumDifficulty"`
		MinimumFruitDifficulty *math.HexOrDecimal256 `json:"minimumFruitDifficulty"`
		DurationLimit          *math.HexOrDecimal256 `json:"durationLimit"`
	}
	var dec MinervaConfig
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.MinimumDifficulty == nil {
		c.MinimumDifficulty = MinimumDifficulty
		//return errors.New("missing required field 'MinimumDifficulty' for Genesis")
	} else {
		c.MinimumDifficulty = (*big.Int)(dec.MinimumDifficulty)
	}
	if dec.MinimumFruitDifficulty == nil {
		c.MinimumFruitDifficulty = MinimumFruitDifficulty
		//return errors.New("missing required field 'MinimumFruitDifficulty' for Genesis")
	} else {
		c.MinimumFruitDifficulty = (*big.Int)(dec.MinimumFruitDifficulty)
	}
	if dec.DurationLimit == nil {
		c.DurationLimit = DurationLimit
		//return errors.New("missing required field 'DurationLimit' for Genesis")
	} else {
		c.DurationLimit = (*big.Int)(dec.DurationLimit)
	}
	return nil
}

func (c MinervaConfig) MarshalJSON() ([]byte, error) {
	type MinervaConfig struct {
		MinimumDifficulty      *math.HexOrDecimal256 `json:"minimumDifficulty,omitempty"`
		MinimumFruitDifficulty *math.HexOrDecimal256 `json:"minimumFruitDifficulty,omitempty"`
		DurationLimit          *math.HexOrDecimal256 `json:"durationLimit,omitempty"`
	}
	var enc MinervaConfig
	enc.MinimumDifficulty = (*math.HexOrDecimal256)(c.MinimumDifficulty)
	enc.MinimumFruitDifficulty = (*math.HexOrDecimal256)(c.MinimumFruitDifficulty)
	enc.DurationLimit = (*math.HexOrDecimal256)(c.DurationLimit)
	return json.Marshal(&enc)
}

// String implements the stringer interface, returning the consensus engine details.
func (c *MinervaConfig) String() string {
	return fmt.Sprintf("{MinimumDifficulty: %v MinimumFruitDifficulty: %v DurationLimit: %v}",
		c.MinimumDifficulty,
		c.MinimumFruitDifficulty,
		c.DurationLimit,
	)
}

// String implements the fmt.Stringer interface.
func (c *ChainConfig) String() string {
	var engine interface{}
	switch {
	case c.Minerva != nil:
		engine = c.Minerva
		// case c.Clique != nil:
		// 	engine = c.Clique
	default:
		engine = "unknown"
	}
	return fmt.Sprintf("{ChainID: %v Engine: %v}",
		c.ChainID,
		engine,
	)
}

// GasTable returns the gas table corresponding to the current phase (homestead or homestead reprice).
//
// The returned GasTable's fields shouldn't, under any circumstances, be changed.
func (c *ChainConfig) GasTable(num *big.Int) GasTable {
	return GasTableConstantinople

}

// CheckCompatible checks whether scheduled fork transitions have been imported
// with a mismatching chain configuration.
func (c *ChainConfig) CheckCompatible(newcfg *ChainConfig, height uint64) *ConfigCompatError {
	bhead := new(big.Int).SetUint64(height)

	// Iterate checkCompatible to find the lowest conflict.
	var lasterr *ConfigCompatError
	for {
		err := c.checkCompatible(newcfg, bhead)
		if err == nil || (lasterr != nil && err.RewindTo == lasterr.RewindTo) {
			break
		}
		lasterr = err
		bhead.SetUint64(err.RewindTo)
	}
	return lasterr
}

func (c *ChainConfig) checkCompatible(newcfg *ChainConfig, head *big.Int) *ConfigCompatError {
	return nil
}

// isForkIncompatible returns true if a fork scheduled at s1 cannot be rescheduled to
// block s2 because head is already past the fork.
func isForkIncompatible(s1, s2, head *big.Int) bool {
	return (isForked(s1, head) || isForked(s2, head)) && !configNumEqual(s1, s2)
}

// isForked returns whether a fork scheduled at block s is active at the given head block.
func isForked(s, head *big.Int) bool {
	if s == nil || head == nil {
		return false
	}
	return s.Cmp(head) <= 0
}

func configNumEqual(x, y *big.Int) bool {
	if x == nil {
		return y == nil
	}
	if y == nil {
		return x == nil
	}
	return x.Cmp(y) == 0
}

// ConfigCompatError is raised if the locally-stored blockchain is initialised with a
// ChainConfig that would alter the past.
type ConfigCompatError struct {
	What string
	// block numbers of the stored and new configurations
	StoredConfig, NewConfig *big.Int
	// the block number to which the local chain must be rewound to correct the error
	RewindTo uint64
}

func newCompatError(what string, storedblock, newblock *big.Int) *ConfigCompatError {
	var rew *big.Int
	switch {
	case storedblock == nil:
		rew = newblock
	case newblock == nil || storedblock.Cmp(newblock) < 0:
		rew = storedblock
	default:
		rew = newblock
	}
	err := &ConfigCompatError{what, storedblock, newblock, 0}
	if rew != nil && rew.Sign() > 0 {
		err.RewindTo = rew.Uint64() - 1
	}
	return err
}

func (err *ConfigCompatError) Error() string {
	return fmt.Sprintf("mismatching %s in database (have %d, want %d, rewindto %d)", err.What, err.StoredConfig, err.NewConfig, err.RewindTo)
}

// Rules wraps ChainConfig and is merely syntatic sugar or can be used for functions
// that do not have or require information about the block.
//
// Rules is a one time interface meaning that it shouldn't be used in between transition
// phases.
type Rules struct {
	ChainID *big.Int
	IsTIP3  bool
}

// Rules ensures c's ChainID is not nil.
func (c *ChainConfig) Rules(num *big.Int) Rules {
	chainID := c.ChainID
	if chainID == nil {
		chainID = new(big.Int)
	}
	return Rules{
		ChainID: new(big.Int).Set(chainID),
		IsTIP3:  c.IsTIP3(num),
	}
}

// IsTIP3 returns whether num is either equal to the IsTIP3 fork block or greater.
func (c *ChainConfig) IsTIP3(num *big.Int) bool {
	if c.TIP3 == nil {
		return false
	}
	return isForked(c.TIP3.FastNumber, num)
}

// IsTIP5 returns whether num is either equal to the IsTIP5 fork block or greater.
func (c *ChainConfig) IsTIP5(num *big.Int) bool {
	if c.TIP5 == nil {
		return false
	}
	return isForked(c.TIP5.SnailNumber, num)
}
