CREATE TABLE orders
(
    order_uid          VARCHAR(36) PRIMARY KEY,
    track_number       VARCHAR(255) UNIQUE,
    entry              VARCHAR(255),
    delivery           JSONB,
    payment            JSONB,
    locale             VARCHAR(12),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shard_key          VARCHAR(255),
    sm_id              NUMERIC,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR(255)
);

CREATE TABLE order_items
(
    chrt_id      INT PRIMARY KEY,
    order_uid    VARCHAR(255) REFERENCES orders (order_uid),
    track_number VARCHAR(255),
    price        NUMERIC,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         NUMERIC,
    size         VARCHAR(10),
    total_price  NUMERIC,
    nm_id        NUMERIC,
    brand        VARCHAR(255),
    status       INT
);