select max(t.numprice) from 
(
	select 
	cast(rtrim(p.price, ' 元特價') as integer) as numprice 
	from product p WHERE p.price != ''
) t