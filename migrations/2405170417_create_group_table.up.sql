CREATE TABLE "group" (
    group_id uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    description varchar(255) default '',
    avatar_url varchar(255) default ''
);

CREATE TABLE group_member(
    group_member_id uuid PRIMARY KEY,
    group_id uuid references "group"(group_id),
    user_id uuid references user_identity(user_id)
)