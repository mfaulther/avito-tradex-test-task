-- name: create-statistics-table
create table if not exists statistics
(
    stat_date date,
    views integer,
    clicks integer,
    cost double precision,
    cpc double precision,
    cpm double precision
);