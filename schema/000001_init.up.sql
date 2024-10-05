create table if not exists users(
    id bigserial primary key,
    name varchar(255) not null,
    age int not null
);
