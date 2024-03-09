create table phone_number (
    id int primary key auto_increment not null,
    created_at DATETIME default CURRENT_TIMESTAMP not null ,
    updated_at DATETIME default CURRENT_TIMESTAMP not null ,
    deleted_at DATETIME default '0000-00-00 00:00:00',
    mobile_phone_number int not null unique ,
    is_urgency boolean not null
);

create table user_phone_number(
    user_id int,
    phone_number_id varchar(10) ,
    foreign key (user_id) references user(id),
    foreign key (phone_number_id) references phone_number(id)
)