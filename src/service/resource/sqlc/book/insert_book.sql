with t as(
	select product_id from product 
	where barcode = :barcode and
	product_title=:product_title
)
insert into book (
	product_id, author, translator, 
	language, category
)
select
	t.product_id,:author,:translator,:language,:category
from t
where t.product_id is not null
on conflict do nothing