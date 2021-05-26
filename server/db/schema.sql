create table if not exists users (
    id          int not null auto_increment primary key,
    email       varchar(225) not null unique,
    username    varchar(225) not null unique,
    passhash    binary(64) not null,
    first_name  varchar(128) not null,
    last_name   varchar(128) not null,
    -- store_id    int not null references stores(id)
);

-- Updates: moved store_id because one user can have multiple stores

create table if not exists stores (
    id          int not null auto_increment primary key,
    user_id     int not null references users(id),
    store_name       varchar(225) not null,
    store_location     varchar(225) not null
);

create table if not exists customers (
    id          int not null auto_increment primary key,
    user_id     int not null references users(id),
    store_id    int not null references stores(id),
    first_name  varchar(128) not null,
    last_name   varchar(128) not null,
    ethnicity   varchar(225) not null,
    gender      varchar(64) not null,
    birthday    date not null,
    postal_code int not null,
    last_visited    date not null,
    dis_channel varchar(225) not null,
    fav_item    int references items(id)
);

create table if not exists items (
    id          int not null auto_increment primary key,
    store_id    int not null references stores(id),
    item_name   varchar(225) not null.
    price       int not null
);

/* TODO: user sign in log schema 