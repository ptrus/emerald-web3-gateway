package storage

import (
	"context"
	"fmt"

	"github.com/starfishlabs/oasis-evm-web3-gateway/model"
)

// ErrNoRoundsIndexed is the error returned by last indexed round when no rounds
// are indexed.
var ErrNoRoundsIndexed = fmt.Errorf("no rounds indexed")

type Storage interface {
	// Upsert upserts a record.
	Upsert(value interface{}) error

	// Delete deletes all records with round less than the given round.
	Delete(table interface{}, round uint64) error

	// GetBlockRound returns block round by block hash.
	GetBlockRound(hash string) (uint64, error)

	// GetBlockHash returns block hash by block round.
	GetBlockHash(round uint64) (string, error)

	// GetLatestBlockHash returns the block hash of the latest indexed round.
	GetLatestBlockHash() (string, error)

	// GetTransactionRef returns block hash, round and index of the transaction.
	GetTransactionRef(ethTxHash string) (*model.TransactionRef, error)

	// GetLastIndexedRound query the last indexed round.
	GetLastIndexedRound() (uint64, error)

	// GetLastRetainedRound query the minimum round not pruned.
	GetLastRetainedRound() (uint64, error)

	// GetTransaction queries ethereum transaction by hash.
	GetTransaction(hash string) (*model.Transaction, error)

	// GetLatestBlockNumber returns the latest block round.
	GetLatestBlockNumber() (uint64, error)

	// GetBlockByHash returns the block for the given hash.
	GetBlockByHash(hash string) (*model.Block, error)

	// GetBlockByNumber returns the block for the given round.
	GetBlockByNumber(round uint64) (*model.Block, error)

	// GetBlockTransactionCountByNumber returns the transaction count of the block.
	GetBlockTransactionCountByNumber(round uint64) (int, error)

	// GetBlockTransactionCountByHash returns the transaction count of the block.
	GetBlockTransactionCountByHash(hash string) (int, error)

	// GetBlockTransaction returns the transaction of the block for the given index.
	GetBlockTransaction(blockHash string, txIndex int) (*model.Transaction, error)

	// GetTransactionReceipt returns the receipt of the transaction.
	GetTransactionReceipt(txHash string) (*model.Receipt, error)

	// GetLogs returns all logs between start and end rounds.
	GetLogs(startRound, endRound uint64) ([]*model.Log, error)

	// RunInTransaction runs a function in a transaction. If function
	// returns an error transaction is rolled back, otherwise transaction
	// is committed.
	RunInTransaction(ctx context.Context, fn func(Storage) error) error
}
