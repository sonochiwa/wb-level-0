CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE orders
(
    order_uid          UUID PRIMARY KEY,
    track_number       VARCHAR(255) UNIQUE,
    entry              VARCHAR(255),
    delivery           JSONB,
    payment            JSONB,
    items              JSONB,
    locale             VARCHAR(12),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shard_key          VARCHAR(255),
    sm_id              NUMERIC,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR(255)
);
