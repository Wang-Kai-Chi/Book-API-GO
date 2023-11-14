insert into admin_user (id, name, email, phone, password) 
values 
(
	gen_random_uuid(), 
  :name,
  :email,
  :phone,
  :password
);
