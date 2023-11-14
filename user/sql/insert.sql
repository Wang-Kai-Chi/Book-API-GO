insert into admin_user (id, name, email, phone, password) 
select
	gen_random_uuid(), 
 	$1,
	$2,
	$3,
	$4
where 
not exists (
	select name
	from admin_user
	where 
	name=$1
)
on conflict do nothing;