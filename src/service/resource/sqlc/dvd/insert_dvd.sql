with t as(
	select product_id from product 
	where barcode = '' and
	product_title=''
)
insert into dvd (product_id, category, director)
select
	t.product_id,'',''
from t
where t.product_id is not null
on conflict do nothing 