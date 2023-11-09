with tid as (
	select product_id as id from product 
	where barcode =:barcode and
	product_title=:product_title
)
insert into book (
	product_id, author, translator, 
	language, category
)
select
	tid.id,:author,:translator,
	:language,:category
from tid
on conflict do nothing
returning product_id 