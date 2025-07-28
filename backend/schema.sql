create table if not exists users (
    id serial primary key,
    username text not null unique,
    email text not null unique,
    password text not null
);

create table if not exists notes (
    id serial primary key,
    user_id integer not null references users(id),
    title text,
    content text,
    created_at timestamp default now()
);
