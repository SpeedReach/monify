CREATE TABLE group_invite_code (
  group_id uuid references "group"(group_id),
  invite_code varchar(10) PRIMARY KEY ,
  created_at timestamp default CURRENT_TIMESTAMP
);