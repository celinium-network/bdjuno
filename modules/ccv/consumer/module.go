package consumer

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/forbole/juno/v5/modules"

	"github.com/forbole/bdjuno/v4/database"
)

var (
	_ modules.GenesisModule            = &Module{}
	_ modules.PeriodicOperationsModule = &Module{}
)

// Module represent database/ccv/consumer module
type Module struct {
	cdc codec.Codec
	db  *database.Db
}

// NewModule returns a new Module instance
func NewModule(cdc codec.Codec, db *database.Db) *Module {
	return &Module{
		cdc: cdc,
		db:  db,
	}
}

// Name implements modules.Module
func (m *Module) Name() string {
	return "ccvconsumer"
}
