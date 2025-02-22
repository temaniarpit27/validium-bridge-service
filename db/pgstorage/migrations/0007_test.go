package migrations_test

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

type migrationTest0007 struct{}

func (m migrationTest0007) InsertData(db *sql.DB) error {
	block := "INSERT INTO sync.block (id, block_num, block_hash, parent_hash, network_id, received_at) VALUES(2, 2803824, decode('27474F16174BBE50C294FE13C190B92E42B2368A6D4AEB8A4A015F52816296C3','hex'), decode('C9B5033799ADF3739383A0489EFBE8A0D4D5E4478778A4F4304562FD51AE4C07','hex'), 1, '0001-01-01 01:00:00.000');"
	if _, err := db.Exec(block); err != nil {
		return err
	}
	insertDeposit := "INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim) VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 1, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 2, 0, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 1, true);"
	if _, err := db.Exec(insertDeposit); err != nil {
		return err
	}
	insertDeposit2 := "INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim) VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 1, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 2, 1, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 2, true);"
	if _, err := db.Exec(insertDeposit2); err != nil {
		return err
	}
	insertDeposit3 := "INSERT INTO sync.deposit(leaf_type, network_id, orig_net, orig_addr, amount, dest_net, dest_addr, block_id, deposit_cnt, tx_hash, metadata, id, ready_for_claim) VALUES(0, 0, 0, decode('0000000000000000000000000000000000000000','hex'), '10000000000000000000', 1, decode('C949254D682D8C9AD5682521675B8F43B102AEC4','hex'), 2, 2, decode('C2D6575EA98EB55E36B5AC6E11196800362594458A4B3143DB50E4995CB2422E','hex'), decode('','hex'), 3, true);"
	if _, err := db.Exec(insertDeposit3); err != nil {
		return err
	}
	insertRoot := "INSERT INTO mt.root (root, deposit_cnt, network, deposit_id) VALUES(decode('16C571C7A60CF3694BA81AFF143E8A8C9A393D351213DBFD4D539F39F1C4648E','hex'), 0, 0, 1);"
	if _, err := db.Exec(insertRoot); err != nil {
		return err
	}
	insertRoot2 := "INSERT INTO mt.root (root, deposit_cnt, network, deposit_id) VALUES(decode('16C571C7A60CF3694BA81AFF143E8A8C9A393D351213DBFD4D539F39F1C4648D','hex'), 1, 0, 2);"
	if _, err := db.Exec(insertRoot2); err != nil {
		return err
	}
	insertRoot3 := "INSERT INTO mt.root (root, deposit_cnt, network, deposit_id) VALUES(decode('16C571C7A60CF3694BA81AFF143E8A8C9A393D351213DBFD4D539F39F1C4648C','hex'), 2, 0, 3);"
	if _, err := db.Exec(insertRoot3); err != nil {
		return err
	}
	return nil
}

func (m migrationTest0007) RunAssertsAfterMigrationUp(t *testing.T, db *sql.DB) {
	queryDepositCount := "select deposit_cnt from mt.root where deposit_id = 0;"
	row := db.QueryRow(queryDepositCount)
	var depositCnt int
	assert.Error(t, row.Scan(&depositCnt))
}

func (m migrationTest0007) RunAssertsAfterMigrationDown(t *testing.T, db *sql.DB) {
	for i := 0; i < 3; i++ {
		queryDepositCount := "select deposit_cnt from mt.root where deposit_id = $1;"
		row := db.QueryRow(queryDepositCount, i+1)
		var depositCnt int
		assert.NoError(t, row.Scan(&depositCnt))
		assert.Equal(t, i, depositCnt)
	}
}

func TestMigration0007(t *testing.T) {
	runMigrationTest(t, 7, migrationTest0007{})
}
