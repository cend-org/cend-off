create table phone_number (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP ,
    updated_at DATETIME default CURRENT_TIMESTAMP  ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    mobile_phone_number varchar(13) ,
    is_urgency boolean default 0
);

create table user_phone_number(
    id int  primary key auto_increment not null ,
    created_at DATETIME default CURRENT_TIMESTAMP ,
    updated_at DATETIME default CURRENT_TIMESTAMP  ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    user_id int unique,
    phone_number_id int unique ,
    foreign key (user_id) references user(id),
    foreign key (phone_number_id) references phone_number(id)
)