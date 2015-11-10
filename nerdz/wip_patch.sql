begin;

    -- https://news.ycombinator.com/item?id=9512912
    -- https://blog.lateral.io/2015/05/full-text-search-in-milliseconds-with-postgresql/
/*
    alter table posts add column tsv tsvector;
    alter table posts_comments add column tsv tsvector;
    alter table groups_posts add column tsv tsvector;
    alter table groups_posts_comments add column tsv tsvector;
*/
    drop view if exists messages cascade;
    create view messages as
    select "hpid","from","to","pid","message","time","news","lang","closed", 0 as type from groups_posts
    union all
    select "hpid","from","to","pid","message","time","news","lang","closed", 1 as type from posts;

    drop table if exists oauth2_clients cascade;
    create table oauth2_clients(
        id bigserial not null primary key,
        secret text not null unique,
        redirect_uri varchar(350) not null,
        user_id bigint not null references users(counter) on delete cascade
    );

    drop table if exists oauth2_authorize cascade;
    create table oauth2_authorize (
        id bigserial not null primary key, -- surrogated id
        code text not null unique, --logical id
        client_id bigint not null references oauth2_clients(id),
        Created_At timestamp with time zone not null,
        expires_in bigint not null,
        state text not null,
        scope text not null,
        redirect_uri varchar(350) not null,
        user_id bigint not null references users(counter) on delete cascade
    );

    drop table if exists oauth2_refresh cascade;
    create table oauth2_refresh(
        id bigserial not null primary key, -- surrogated id
        token text not null unique, -- logical id
        oauth2_access_id bigint not null references oauth2_access(id)
    );

    drop table if exists oauth2_access cascade;
    create table oauth2_access(
        id bigserial not null primary key, -- surrogated id
        client_id bigint not null references oauth2_clients(id),
        access_token text not null unique, -- logical id
        Created_At timestamp with time zone not null,
        expires_in bigint not null,
        redirect_uri varchar(350) not null,
        oauth2_authorize_id bigint not null references oauth2_authorize(id),
        oauth2_access_id bigint references oauth2_access(id),
        refresh_token bigint references oauth2_refresh(id),
        scope text not null,
        user_id bigint not null references users(counter)
    );

commit;
