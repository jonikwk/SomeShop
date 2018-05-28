
drop trigger check_items_quantity on tables.order_product cascade;
drop function tables.check_items_quantity(); 

create or replace function tables.check_items_quantity() 
returns trigger as $$
declare
	quantity integer;
begin
    if TG_OP = 'INSERT'  then
	    SELECT tables.order_product.quantity FROM tables.order_product WHERE id_product = new.id_product and id_order=new.id_order and id_size=new.id_size into quantity;
	    if (quantity is not null) then 
            update tables.order_product 
            set quantity = (select tables.order_product.quantity from tables.order_product where id_product = new.id_product and id_order=new.id_order and id_size=new.id_size) + 1 
            where id_product = new.id_product and id_order=new.id_order and id_size=new.id_size;
        elsif  (quantity is null) then 
            return new;
        end if;	
    end if;     
    return null;
end;   
$$ 
language plpgsql;
create trigger check_items_quantity before insert on tables.order_product for each row execute procedure tables.check_items_quantity(); 



/*create or replace function tables.check_negative_quantity() 
returns trigger as $$
declare
	quantity integer;
begin
    if TG_OP = 'UPDATE'  then
	    SELECT tables.order_product.quantity FROM tables.order_product WHERE id_product = new.id_product and id_order=new.id_order and id_size=new.id_size into quantity;
	    if (quantity <= 0) then 
            update tables.order_product 
            set quantity = 1 where id_product = new.id_product and id_order=new.id_order and id_size=new.id_size;
        elsif  (quantity is null) then 
            return new;
        end if;	
    end if;     
    return null;
end;   
$$ 
language plpgsql;
create trigger check_negative_quantity before update on tables.order_product for each row execute procedure tables.check_negative_quantity(); 
*/