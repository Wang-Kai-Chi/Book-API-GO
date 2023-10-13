insert into product(
    barcode, product_title, publisher, 
    publication_date, price, quantity, description
) 
select 
    :barcode, :product_title, :publisher, 
    :publication_date, :price, :quantity, :description
on conflict do nothing