create table customers(
    cust_id INTEGER NOT NULL AUTO_INCREMENT,
    cust_name varchar(100) NOT NULL,
    cust_address varchar(500) NOT NULL,
    cust_pincode INTEGER NOT NULL,
    cust_phone varchar(15) NOT NULL,
    cust_email varchar(100) NOT NULL UNIQUE,
    cust_hash_pwd char(60) NOT NULL,
    primary key(cust_id)
);

create table vendors(
    vendor_id INTEGER NOT NULL AUTO_INCREMENT,
    vendor_name varchar(100) NOT NULL,
    vendor_pincode INTEGER NOT NULL,
    vendor_gps_lat decimal(12,8) NOT NULL,
    vendor_gps_long decimal(12,8) NOT NULL,
    vendor_email varchar(100) NOT NULL UNIQUE,
    vendor_pwd char(60) NOT NULL,
    vendor_address varchar(100) NOT NULL,
    vendor_phone varchar(15) NOT NULL,
    primary key(vendor_id)
);

create table listings(
    list_id INTEGER NOT NULL ,
    vendor_id INTEGER NOT NULL,
    listing_price INTEGER NOT NULL,
    listing_desc varchar(500) NOT NULL,
    listing_name varchar(100) NOT NULL,
    primary key (list_id),
    foreign key (vendor_id) references vendors(vendor_id)
);

create table deliveries(
    delivery_id INTEGER NOT NULL AUTO_INCREMENT,
    cust_id INTEGER NOT NULL,
    vendor_id INTEGER NOT NULL,
    timeofdelivery DATETIME NOT NULL,
    drop_gps_lat decimal(12,8) NOT NULL,
    drop_gps_long decimal(12,8) NOT NULL,
    delivery_status varchar(50) NOT NULL,
    primary key (delivery_id),
    foreign key (cust_id) references customers(cust_id),
    foreign key (vendor_id) references vendors(vendor_id)
);

create table orders(
    order_id INTEGER NOT NULL AUTO_INCREMENT,
    delivery_id INTEGER NOT NULL,
    list_id INTEGER NOT NULL,
    order_quantity INTEGER NOT NULL,
    order_amount INTEGER NOT NULL,
    primary key (order_id),
    foreign key (delivery_id) references deliveries(delivery_id),
    foreign key (list_id) references listings(list_id)
);
