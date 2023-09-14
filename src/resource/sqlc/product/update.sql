UPDATE product SET 
	barcode =:barcode,
	product_title =:product_title,
	publisher =:publisher,
	publication_date =:publication_date,
	quantity =:quantity,
	description =:description,
	price =:price 
WHERE product_id =:product_id
RETURNING product_id
