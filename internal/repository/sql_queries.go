package repository

const (
	selectOrders = `SELECT order_uid FROM orders;`

	selectOrderByID = `
					SELECT order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, 
						customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard 
					FROM orders 
					WHERE order_uid = $1;
					`

	insertOrder = `
					INSERT INTO orders (order_uid, track_number, entry, delivery, payment, items, locale, internal_signature, 
					                    customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard)
					VALUES (DEFAULT, $1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13) 
					RETURNING order_uid;
					`

	deleteAllOrders = `DELETE FROM orders;`
)
