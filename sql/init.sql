CREATE TABLE orders
(
    order_uid          VARCHAR(36) PRIMARY KEY,
    track_number       VARCHAR(255) UNIQUE,
    entry              VARCHAR(255),
    locale             VARCHAR(12),
    internal_signature VARCHAR(255),
    customer_id        VARCHAR(255),
    delivery_service   VARCHAR(255),
    shard_key          VARCHAR(255),
    sm_id              NUMERIC,
    date_created       TIMESTAMP,
    oof_shard          VARCHAR(255)
);

CREATE TABLE deliveries
(
    delivery_id SERIAL PRIMARY KEY,
    order_uid   VARCHAR(255) REFERENCES orders (order_uid),
    name        VARCHAR(255),
    phone       VARCHAR(15),
    zip         VARCHAR(10),
    city        VARCHAR(255),
    address     VARCHAR(255),
    region      VARCHAR(255),
    email       VARCHAR(255)
);

CREATE TABLE payments
(
    payment_id    SERIAL PRIMARY KEY,
    order_uid     VARCHAR(255) REFERENCES orders (order_uid),
    transaction   VARCHAR(255),
    request_id    VARCHAR(255),
    currency      VARCHAR(3),
    provider      VARCHAR(255),
    amount        NUMERIC,
    payment_dt    TIMESTAMP,
    bank          VARCHAR(255),
    delivery_cost NUMERIC,
    goods_total   NUMERIC,
    custom_fee    NUMERIC
);

CREATE TABLE order_items
(
    chrt_id      INT PRIMARY KEY,
    track_number VARCHAR(255),
    price        NUMERIC,
    rid          VARCHAR(255),
    name         VARCHAR(255),
    sale         NUMERIC,
    size         VARCHAR(10),
    total_price  NUMERIC,
    nm_id        NUMERIC,
    brand        VARCHAR(255),
    status       INT,
    FOREIGN KEY (track_number) REFERENCES orders (track_number)
);