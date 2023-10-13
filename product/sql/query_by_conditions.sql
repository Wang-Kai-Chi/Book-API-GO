select t.product_id,
		t.barcode,
		t.product_title,
		t.publisher,
		t.publication_date,
		t.quantity,
		t.description, 
		t.price
	from (
		select 
			product_id,
			barcode,
			product_title,
			publisher,
			publication_date,
			quantity,
			description,
			cast(rtrim(price, ' 元特價') as integer) as num_price,
			price
		from product p where p.price != ''
	) t 
where t.num_price >= ? and 
t.num_price <= ? and
t.product_title like ? and 
t.publisher like ?
order by t.num_price
