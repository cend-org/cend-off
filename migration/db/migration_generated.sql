drop database if exists cend;
create database cend;
use cend;

create table address (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	country	varchar(255) default '',
	city	varchar(255) default '',
	latitude	float default 0,
	longitude	float default 0,
	street	varchar(255) default '',
	full_address	varchar(255) default '',
	xid	varchar(255) default '',
	constraint address_pk
		unique (id)
);

create table asset (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	description	varchar(255) default '',
	constraint asset_pk
		unique (id)
);

create table authorization (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	level	int default 0,
	constraint authorization_pk
		unique (id)
);

alter table authorization
	add constraint authorization_user_id_fk
		foreign key (user_id) references user (id);

create table calendar_planning (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	authorization_id	int default 0,
	start_date_time	datetime default '0000-00-00 00:00:00',
	end_date_time	datetime default '0000-00-00 00:00:00',
	description	varchar(255) default '',
	constraint calendar_planning_pk
		unique (id)
);

alter table calendar_planning
	add constraint calendar_planning_authorization_id_fk
		foreign key (authorization_id) references authorization (id);

create table calendar_planning_actor (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	authorization_id	int default 0,
	calendar_planning_id	int default 0,
	constraint calendar_planning_actor_pk
		unique (id)
);

alter table calendar_planning_actor
	add constraint calendar_planning_actor_authorization_id_fk
		foreign key (authorization_id) references authorization (id);

alter table calendar_planning_actor
	add constraint calendar_planning_actor_calendar_planning_id_fk
		foreign key (calendar_planning_id) references calendar_planning (id);

create table code (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	verification_code	int default 0,
	constraint code_pk
		unique (id)
);

alter table code
	add constraint code_user_id_fk
		foreign key (user_id) references user (id);

create table contract (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	tutor_id	int default 0,
	parent_id	int default 0,
	student_id	int default 0,
	start_date	datetime default '0000-00-00 00:00:00',
	end_date	datetime default '0000-00-00 00:00:00',
	payment_type	int default 0,
	salary_value	float default 0,
	payment_method	int default 0,
	constraint contract_pk
		unique (id)
);

alter table contract
	add constraint contract_tutor_id_fk
		foreign key (tutor_id) references user (id);

alter table contract
	add constraint contract_parent_id_fk
		foreign key (parent_id) references user (id);

alter table contract
	add constraint contract_student_id_fk
		foreign key (student_id) references user (id);

create table contract_timesheet_detail (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	contract_id	int default 0,
	date	datetime default '0000-00-00 00:00:00',
	hours	float default 0,
	constraint contract_timesheet_detail_pk
		unique (id)
);

alter table contract_timesheet_detail
	add constraint contract_timesheet_detail_contract_id_fk
		foreign key (contract_id) references contract (id);

create table education (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	name	varchar(255) default '',
	constraint education_pk
		unique (id)
);

create table mark (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	author_id	int default 0,
	author_comment	varchar(255) default '',
	author_mark	int default 0,
	constraint mark_pk
		unique (id)
);

alter table mark
	add constraint mark_user_id_fk
		foreign key (user_id) references user (id);

alter table mark
	add constraint mark_author_id_fk
		foreign key (author_id) references user (id);

create table media (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	file_name	varchar(255) default '',
	extension	varchar(255) default '',
	xid	varchar(255) default '',
	content_type	int default 0,
	constraint media_pk
		unique (id)
);

create table media_thumb (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	extension	varchar(255) default '',
	media_xid	varchar(255) default '',
	xid	varchar(255) default '',
	constraint media_thumb_pk
		unique (id)
);

alter table media_thumb
	add constraint media_thumb_media_xid_fk
		foreign key (media_xid) references media (xid);

create table message (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	resource_type	int default 0,
	resource_number	int default 0,
	resource_value	int default 0,
	resource_label	varchar(255) default '',
	resource_language	int default 0,
	constraint message_pk
		unique (id)
);

create table mutation (
	id int auto_increment primary key,
);

create table password (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	hash	varchar(255) default '',
	constraint password_pk
		unique (id)
);

alter table password
	add constraint password_user_id_fk
		foreign key (user_id) references user (id);

create table phone_number (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	mobile_phone_number	varchar(255) default '',
	is_urgency	bool default '',
	constraint phone_number_pk
		unique (id)
);

create table post (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	publisher_id	int default 0,
	content	varchar(255) default '',
	expiration_date	datetime default '0000-00-00 00:00:00',
	constraint post_pk
		unique (id)
);

alter table post
	add constraint post_publisher_id_fk
		foreign key (publisher_id) references user (id);

create table post_tag (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	post_id	int default 0,
	tag_content	varchar(255) default '',
	constraint post_tag_pk
		unique (id)
);

alter table post_tag
	add constraint post_tag_post_id_fk
		foreign key (post_id) references post (id);

create table post_view (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	post_id	int default 0,
	viewer_id	int default 0,
	constraint post_view_pk
		unique (id)
);

alter table post_view
	add constraint post_view_post_id_fk
		foreign key (post_id) references post (id);

alter table post_view
	add constraint post_view_viewer_id_fk
		foreign key (viewer_id) references user (id);

create table qr_code_registry (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	xid	varchar(255) default '',
	is_used	bool default '',
	constraint qr_code_registry_pk
		unique (id)
);

alter table qr_code_registry
	add constraint qr_code_registry_user_id_fk
		foreign key (user_id) references user (id);

create table query (
	id int auto_increment primary key,
);

create table subject (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	education_level_id	int default 0,
	name	varchar(255) default '',
	constraint subject_pk
		unique (id)
);

alter table subject
	add constraint subject_education_level_id_fk
		foreign key (education_level_id) references education_level (id);

create table user (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	name	varchar(255) default '',
	family_name	varchar(255) default '',
	nick_name	varchar(255) default '',
	email	varchar(255) default '',
	matricule	varchar(255) default '',
	age	int default 0,
	birth_date	datetime default '0000-00-00 00:00:00',
	sex	int default 0,
	lang	int default 0,
	status	int default 0,
	profile_image_xid	varchar(255) default '',
	description	varchar(255) default '',
	cover_text	varchar(255) default '',
	profile	varchar(255) default '',
	experience_detail	varchar(255) default '',
	additional_description	varchar(255) default '',
	add_on_title	varchar(255) default '',
	constraint user_pk
		unique (id)
);

alter table user
	add constraint user_profile_image_xid_fk
		foreign key (profile_image_xid) references profile_image (xid);

create table user_address (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	address_id	int default 0,
	address_type	varchar(255) default '',
	constraint user_address_pk
		unique (id)
);

alter table user_address
	add constraint user_address_user_id_fk
		foreign key (user_id) references user (id);

alter table user_address
	add constraint user_address_address_id_fk
		foreign key (address_id) references address (id);

create table user_authorization_link (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	link_type	int default 0,
	constraint user_authorization_link_pk
		unique (id)
);

create table user_authorization_link_actor (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_authorization_link_id	int default 0,
	authorization_id	int default 0,
	constraint user_authorization_link_actor_pk
		unique (id)
);

alter table user_authorization_link_actor
	add constraint user_authorization_link_actor_user_authorization_link_id_fk
		foreign key (user_authorization_link_id) references user_authorization_link (id);

alter table user_authorization_link_actor
	add constraint user_authorization_link_actor_authorization_id_fk
		foreign key (authorization_id) references authorization (id);

create table user_education_level_subject (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	subject_id	int default 0,
	constraint user_education_level_subject_pk
		unique (id)
);

alter table user_education_level_subject
	add constraint user_education_level_subject_user_id_fk
		foreign key (user_id) references user (id);

alter table user_education_level_subject
	add constraint user_education_level_subject_subject_id_fk
		foreign key (subject_id) references subject (id);

create table user_media_detail (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	owner_id	int default 0,
	document_type	int default 0,
	document_xid	varchar(255) default '',
	constraint user_media_detail_pk
		unique (id)
);

alter table user_media_detail
	add constraint user_media_detail_owner_id_fk
		foreign key (owner_id) references user (id);

alter table user_media_detail
	add constraint user_media_detail_document_xid_fk
		foreign key (document_xid) references document (xid);

create table user_phone_number (
	id int auto_increment primary key,
	created_at	datetime default CURRENT_TIMESTAMP,
	updated_at	datetime default CURRENT_TIMESTAMP,
	deleted_at	datetime default '0000-00-00 00:00:00',
	user_id	int default 0,
	phone_number_id	int default 0,
	constraint user_phone_number_pk
		unique (id)
);

alter table user_phone_number
	add constraint user_phone_number_user_id_fk
		foreign key (user_id) references user (id);

alter table user_phone_number
	add constraint user_phone_number_phone_number_id_fk
		foreign key (phone_number_id) references phone_number (id);

