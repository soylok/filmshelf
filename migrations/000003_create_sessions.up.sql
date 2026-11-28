create table filmshelf.sessions(
    token_hash text primary key,
    user_id bigint not null
        references filmshelf.users(id) on delete cascade,
    created_at timestamptz not null default now(),
    expires_at timestamptz not null,

    check(expires_at > created_at)
);

create index sessions_user_id_idx
    on filmshelf.sessions(user_id);
