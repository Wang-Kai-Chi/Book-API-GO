select t.product_id,t.author,t.translator,
			t.language,t.category,t.barcode,t.product_title,t.publisher,
			t.publication_date,t.quantity,t.description,
	t.price from (
		select 
			product_id,barcode,product_title,publisher,
			publication_date,quantity,description,
			cast(rtrim(price, ' 元特價') as integer) 
	as num_price,price,author,translator,
			language,category
		from (select * from book natural join 
		product) p where p.price != ''
	) t 
where t.num_price >= ? and 
t.num_price <= ? and
t.product_title like ? and 
t.publisher like ? and
t.author like ? and 
t.translator like ? and
t.language like ? and
t.category like ?
order by t.num_price
