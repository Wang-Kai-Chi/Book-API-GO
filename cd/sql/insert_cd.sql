with t as(
	select product_id from product 
	where barcode = :barcode and
	product_title=:product_title
)
insert into cd (product_id, performer, genre)
select
	t.product_id,:performer,:genre
from t
where t.product_id is not null
on conflict do nothing 