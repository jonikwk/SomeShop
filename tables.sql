/*create schema tables;

create table tables.users(
    id bigint not null primary key,
    username text not null,
    phone text not null, 
    fullname text not null,
    address text not null
);

create table tables.catalog(
    id int not null primary key, 
    title text not null,
    parent int not null
);

create table tables.products(
    id serial not null primary key,
    title text not null,
    price int not null,
    color text not null,
    id_category int not null, 
    description text not null,
    photo text not null
);

create table tables.sizes(
    id int not null primary key,
    title text not null
);

create table tables.reviews(
    id int not null primary key,
    id_product int not null,
    id_user bigint not null,
    date text not null,
    description text not null
);

create table tables.products_sizes(
    id_product int not null,
    id_size int not null
);

create table tables.orders(
    id int not null primary key,
    number text not null,
    cost money not null,
    id_user bigint not null
);

create table tables.order_product(
    id_product int not null,
    id_order int not null,
    quantity int not null
);*/