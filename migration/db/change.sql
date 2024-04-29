--
DROP DATABASE IF EXISTS duval;
create database duval;
use duval;

CREATE TABLE asset (
    id int auto_increment primary key,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    description varchar(5000) default ''
);


CREATE TABLE user (
    id int auto_increment primary key ,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    name varchar(500) default '',
    last_name varchar(500) default '',
    email varchar(100) default ''
);

CREATE UNIQUE INDEX email_ui on user (email);

CREATE TABLE password(
    id int auto_increment primary key ,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    user_id int default 0 ,
    hash varchar(5000) default ''
);

alter table password
    add constraint password_user_id_fk
        foreign key (user_id) references user (id);

CREATE TABLE authorization (
    id int auto_increment primary key ,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    user_id int default 0,
    access_level int default 0
);

ALTER TABLE authorization
    add constraint authorization_user_id_fk
        foreign key (user_id) references user(id);


CREATE TABLE school (
    id int auto_increment primary key ,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    name varchar(500) default ''
);

CREATE TABLE school_subject (
    id int auto_increment primary key ,
    created_at timestamp default CURRENT_TIMESTAMP,
    updated_at timestamp default CURRENT_TIMESTAMP,
    deleted_at timestamp default '0000-00-00 00:00:00',
    school_number int default 0,
    name varchar(500) default ''
);

ALTER TABLE school_subject
    add constraint school_fk
        foreign key (school_number) references school(id);