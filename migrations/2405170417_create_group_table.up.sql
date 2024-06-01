CREATE TABLE "group" (
    group_id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255) default '',
    avatar_url varchar(255) default ''
);

CREATE TABLE group_member(
    group_member_id uuid PRIMARY KEY,
    group_id uuid NOT NULL ,
    user_id uuid NOT NULL,
    UNIQUE (group_id, user_id),
    FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    FOREIGN KEY (user_id) REFERENCES user_identity(user_id)
)