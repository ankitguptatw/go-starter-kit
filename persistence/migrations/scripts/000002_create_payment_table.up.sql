create table if not exists payments (
    id serial primary key,

    amount numeric(10, 4) not null,

    beneficiary_name varchar(100) not null,
    beneficiary_account_number integer not null,
    beneficiary_code varchar(100) not null,

    payee_name varchar(100) not null,
    payee_account_number integer not null,
    payee_code varchar(100) not null,

    status varchar(100) not null,

    created_at timestamp without time zone default (now() at time zone 'utc'),
    updated_at timestamp without time zone default (now() at time zone 'utc'),
    deleted_at timestamp without time zone default (now() at time zone 'utc')
);