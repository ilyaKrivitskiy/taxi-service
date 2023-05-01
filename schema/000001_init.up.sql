CREATE TABLE users
(
    id            serial primary key,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE drivers
(
    id            serial primary key,
    name          varchar(255) not null,
    email         varchar(255) not null unique,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null,
    taxi_type     varchar(255) not null
);

CREATE TABLE orders
(
    id         serial primary key,
    user_id    int          not null unique,
    driver_id  int          not null unique,
    place_from varchar(255) not null,
    place_to   varchar(255) not null,
    taxi_type  varchar(255) not null,
    date_start date         not null,
    date_end   date         not null,
    status     varchar(255) not null,
    info_order text
);

CREATE TABLE trips
(
    id        serial primary key,
    user_id   int references users (id) on update cascade on delete cascade,
    driver_id int references drivers (id) on update cascade on delete cascade
);
