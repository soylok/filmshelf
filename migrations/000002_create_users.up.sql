create table filmshelf.users(
    id bigint primary key generated always as identity,
    first_name varchar(100) not null,
    second_name varchar(100) not null,
    email text not null unique,
    email_verified boolean not null default false,
    password_hash text not null,
    role text not null default 'user',
    version bigint not null default 1,

    check(length(trim(first_name)) > 0),
    check(length(trim(second_name)) > 0),
    check(length(trim(email)) > 0),
    check(length(trim(password_hash)) > 0),
    check(role in ('user', 'admin'))
);
