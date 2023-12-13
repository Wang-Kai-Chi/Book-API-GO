update admin_user 
set auth=$1
where id=$2
returning id