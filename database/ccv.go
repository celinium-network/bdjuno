package database

import (
	"encoding/json"
	"fmt"

	"github.com/forbole/bdjuno/v4/types"
)

// SaveCcvProviderParams saves the ccv provider params for the given height
func (db *Db) SaveCcvProviderParams(params *types.CcvProviderParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return err
	}

	stmt := `
INSERT INTO ccv_provider_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params, 
        height = excluded.height
WHERE ccv_provider_params.height <= excluded.height`

	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing ccv provider params: %s", err)
	}

	return nil
}

// SaveCcvConsumerParams saves the ccv consumer params for the given height
func (db *Db) SaveCcvConsumerParams(params *types.CcvConsumerParams) error {
	paramsBz, err := json.Marshal(&params.Params)
	if err != nil {
		return err
	}

	stmt := `
INSERT INTO ccv_consumer_params (params, height) 
VALUES ($1, $2)
ON CONFLICT (one_row_id) DO UPDATE 
    SET params = excluded.params, 
        height = excluded.height
WHERE ccv_consumer_params.height <= excluded.height`

	_, err = db.SQL.Exec(stmt, string(paramsBz), params.Height)
	if err != nil {
		return fmt.Errorf("error while storing ccv consumer params: %s", err)
	}

	return nil
}

// SaveCcvConsumerChain saves the ccv consumer chain info for the given height
func (db *Db) SaveCcvConsumerChain(consumerChain *types.CcvConsumerChain) error {
	providerClientState, err := json.Marshal(&consumerChain.ProviderClientState)
	if err != nil {
		return err
	}
	providerConsensusState, err := json.Marshal(&consumerChain.ProviderConsensusState)
	if err != nil {
		return err
	}
	maturingPackets, err := json.Marshal(&consumerChain.MaturingPackets)
	if err != nil {
		return err
	}
	initialValSet, err := json.Marshal(&consumerChain.InitialValSet)
	if err != nil {
		return err
	}
	heightToValsetUpdateID, err := json.Marshal(&consumerChain.HeightToValsetUpdateID)
	if err != nil {
		return err
	}
	outstandingDowntimeSlashing, err := json.Marshal(&consumerChain.OutstandingDowntimeSlashing)
	if err != nil {
		return err
	}
	pendingConsumerPackets, err := json.Marshal(&consumerChain.PendingConsumerPackets)
	if err != nil {
		return err
	}
	lastTransmissionBlockHeight, err := json.Marshal(&consumerChain.LastTransmissionBlockHeight)
	if err != nil {
		return err
	}

	stmt := `
INSERT INTO ccv_consumer_chain (provider_client_id, provider_channel_id, new_chain, provider_client_state,
	provider_consensus_state, maturing_packets, initial_val_set, height_to_valset_update_id, 
	outstanding_downtime_slashing, pending_consumer_packets, last_transmission_block_height, height) 
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
ON CONFLICT (provider_client_id) DO UPDATE 
	SET provider_channel_id = excluded.provider_channel_id,
		new_chain = excluded.new_chain,
		provider_client_state = excluded.provider_client_state,
		provider_consensus_state = excluded.provider_consensus_state,
		maturing_packets = excluded.maturing_packets,
		initial_val_set = excluded.initial_val_set,
		height_to_valset_update_id = excluded.height_to_valset_update_id,
		outstanding_downtime_slashing = excluded.outstanding_downtime_slashing,
		pending_consumer_packets = excluded.pending_consumer_packets,
		last_transmission_block_height = excluded.last_transmission_block_height,
		height = excluded.height
WHERE ccv_consumer_chain.height <= excluded.height`

	_, err = db.SQL.Exec(stmt, consumerChain.ProviderClientID, consumerChain.ProviderChannelID,
		consumerChain.NewChain, string(providerClientState), string(providerConsensusState),
		string(maturingPackets), string(initialValSet), string(heightToValsetUpdateID),
		string(outstandingDowntimeSlashing), string(pendingConsumerPackets),
		string(lastTransmissionBlockHeight), consumerChain.Height)

	if err != nil {
		return fmt.Errorf("error while storing ccv consumer chain info: %s", err)
	}

	return nil
}