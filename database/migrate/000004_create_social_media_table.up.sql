create table social_media(
sm_id SERIAL primary key,
sm_name varchar(50) not null,
sm_url text not null,
user_id int not null references users(u_id) on delete cascade on update cascade,
sm_created_date date,
sm_updated_date date
);