create table if not exists users (
    id          int not null auto_increment,
    email       varchar(225) not null unique,
    username    varchar(225) not null unique,
    passhash    binary(64) not null,
    first_name  varchar(128) not null,
    last_name   varchar(128) not null,
    primary key(id)
);

-- store_id    int not null references stores(id)
-- Updates: moved store_id because one user can have multiple stores

create table if not exists stores (
    id          int not null auto_increment,
    user_id int not null,
    store_name       varchar(225) not null,
    store_location     varchar(225) not null,
    primary key(id),
    foreign key(user_id) references users(id)
);

create table if not exists items (
    id          int not null auto_increment,
    store_id int not null,
    user_id     int not null,
    item_name   varchar(225) not null,
    price       int not null,
    primary key(id),
    foreign key(store_id) references stores(id),
    foreign key(user_id) references users(id)
);

create table if not exists customers (
    id          int not null auto_increment,
    user_id int not null,
    store_id int not null,
    first_name  varchar(128) not null,
    last_name   varchar(128) not null,
    ethnicity   varchar(225) not null,
    gender      varchar(64) not null,
    birthday    date not null,
    postal_code int not null,
    last_visited    date not null,
    dis_channel varchar(225) not null,
    fav_item    int not null,
    primary key(id),
    foreign key(user_id) references users(id),
    foreign key(store_id) references stores (id),
    foreign key(fav_item) references items (id)
);

create table if not exists UserSignIn (
    UserID varchar(20) not null,
    SignInTime datetime not null,
    ClientIP varchar(255) not null
);




-- TODO: user sign in log schema 