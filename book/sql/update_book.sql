with parent as(
	UPDATE product SET 
		barcode = :barcode,
		product_title = :product_title,
		publisher = :publisher,
		publication_date = :publication_date,
		quantity = :quantity,
		description = :description,
		price = :price 
	WHERE product_id = :product_id
	RETURNING product_id
)
update book set
	author=:author,
	translator=:translator,
	language=:language,
	category=:category
where product_id = (select product_id from parent)