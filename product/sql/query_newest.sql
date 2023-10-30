with t as (
	select max(product_id) as id from product
)
select * from product where product_id <= 
(select t.id from t) and
product_id > (select t.id from t) - ?