create table translator (
                            id int primary key auto_increment not null,
                            created_at DATETIME default CURRENT_TIMESTAMP not null ,
                            updated_at DATETIME default CURRENT_TIMESTAMP not null ,
                            deleted_at DATETIME default '0000-00-00 00:00:00',
                            msg varchar(5000) default '',
                            number int default 0,
                            language int  default 0
);

alter table translator
    add constraint translator_pk
        unique (number, language);

alter table translator add column menu_parent_number int default 0 after language;

create table website (
                            id int primary key auto_increment not null,
                            created_at DATETIME default CURRENT_TIMESTAMP not null ,
                            updated_at DATETIME default CURRENT_TIMESTAMP not null ,
                            deleted_at DATETIME default '0000-00-00 00:00:00',
                            name varchar(500) default '',
                            xid varchar(500) default ''
);