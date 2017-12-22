create table urls (
    id serial primary key,
    url text not null,
    path text not null,
    created_at timestamp default current_timestamp
);

create table analytics (
    id serial primary key,
    url_id int not null,
    created_at timestamp default current_timestamp
);
