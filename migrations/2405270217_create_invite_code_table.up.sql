CREATE TABLE group_invite_code (
  group_id uuid NOT NULL,
  invite_code varchar(10) PRIMARY KEY ,
  created_at timestamp default CURRENT_TIMESTAMP,
  FOREIGN KEY (group_id) REFERENCES "group"(group_id)
);