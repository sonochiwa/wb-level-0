package order

const (
	selectOrders = `SELECT order_uid FROM orders;`

	selectOrderByID = `
					SELECT order_uid, track_number, entry, delivery, payment, locale, internal_signature, 
						customer_id, delivery_service, shard_key, sm_id, date_created, oof_shard 
					FROM orders 
					WHERE order_uid = $1;
					`

	insertOrder = `
					INSERT INTO orders (order_uid, track_number)
					VALUES (DEFAULT, $1) 
					RETURNING order_uid;
					`

	deleteAllOrders = `DELETE FROM orders;`
)
