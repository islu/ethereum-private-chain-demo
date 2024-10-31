
/*
    block_tx
*/

-- Create block transaction
-- name: CreateBlockTx :one
INSERT INTO block_tx(
    block_number, from_address, to_address, tx_nonce, tx_hash, tx_value, tx_gas, tx_gas_price, tx_time, tx_data, create_time
) VALUES (
    $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11
)
RETURNING *;

-- Get block transaction by tx_hash
-- name: GetBlockTxByTxHash :one
SELECT * FROM block_tx
WHERE tx_hash = $1;

--Get max block number by from_address
-- name: GetMaxBlockNumberByFromAddress :one
SELECT MAX(block_number) AS max_block_number
FROM block_tx
WHERE from_address = $1
GROUP BY from_address;

-- List block transaction
-- name: ListBlockTx :many
SELECT * FROM block_tx
ORDER BY seqno DESC
LIMIT $1
;

-- List block transaction by from_address
-- name: ListBlockTxByFromAddress :many
SELECT * FROM block_tx
WHERE from_address = $2
ORDER BY seqno DESC
LIMIT $1
;
