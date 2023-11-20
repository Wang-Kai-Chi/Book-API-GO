insert into admin_user (id, name, email, phone, password) 
select
	gen_random_uuid(), 
 	$1,
	$2,
	$3,
	$4
where 
not exists (
	select name,email,phone
	from admin_user
	where 
	name=$1 and
	email=$2 and
	phone=$3
)
on conflict do nothing;