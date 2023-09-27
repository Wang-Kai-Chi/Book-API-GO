insert into book (
	product_id, author, translator, 
	language, category
)
values(
	(
		select product_id from product 
		where barcode =:barcode and
		product_title=:product_title
	)
	,:author,:translator,:language,:category
)