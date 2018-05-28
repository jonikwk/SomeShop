drop trigger calculate_cost on tables.order_product cascade;
drop function tables.calculate_cost(); 

create or replace function tables.calculate_cost() 
returns trigger as $$
declare
	quantity integer;
    price integer;
    cost1 integer;
    count integer;
    ofset integer;
    temp integer;
begin
cost1=0;
temp=1;
ofset=0;
    if TG_OP = 'INSERT' OR TG_OP = 'UPDATE'  then
        SELECT count(*) FROM tables.order_product WHERE id_order = new.id_order into count;
       /* SELECT tables.order_product.quantity FROM tables.order_product WHERE id_order = new.id_order LIMIT 1 OFFSET offset into quantity[count];*/
	    WHILE temp < count+1 LOOP

   
            SELECT tables.products.price,tables.order_product.quantity FROM tables.products INNER JOIN tables.order_product 
            ON tables.order_product.id_product=tables.products.id WHERE tables.order_product.id_order=new.id_order LIMIT 1 OFFSET ofset into price,quantity;
           
            cost1 = cost1+quantity*price;
            ofset=ofset+1;
            temp=temp+1;
        END LOOP;
        
    end if;     
    update tables.orders set cost = cost1 where tables.orders.id = new.id_order;
    return null;
end;   
$$ 
language plpgsql;
create trigger calculate_cost after insert or update on tables.order_product for each row execute procedure tables.calculate_cost(); 
