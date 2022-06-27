create table commentss(
c_id SERIAL primary key,
c_message text not null,
user_id int not null references users(u_id) on delete cascade on update cascade,
photo_id int references photo(p_id) on delete cascade on update cascade,
c_created_date date,
c_updated_date date
);