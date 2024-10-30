CREATE TABLE block_tx (
    seqno serial NOT NULL PRIMARY KEY, -- 流水號
    block_number bigint NOT NULL,
    from_address varchar(160) NOT NULL,
    to_address varchar(160) NOT NULL,
    tx_nonce int NOT NULL,
    tx_hash varchar(255) NOT NULL,
    tx_value bigint NOT NULL,
    tx_gas bigint NOT NULL,
    tx_gas_price bigint NOT NULL,
    tx_time timestamptz NOT NULL,
    tx_data text NOT NULL,
    create_time timestamptz NOT NULL -- 建立時間
);

CREATE INDEX idx_block_tx_block_number ON block_tx(
    block_number
);

ALTER TABLE block_tx ADD CONSTRAINT unique_tx_hash UNIQUE (tx_hash);
