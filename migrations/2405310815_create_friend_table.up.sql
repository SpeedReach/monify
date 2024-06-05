CREATE TABLE friend(
    relation_id uuid PRIMARY KEY,
    user1_id uuid REFERENCES user_identity(user_id),
    user2_id uuid REFERENCES user_identity(user_id),
    UNIQUE (user1_id, user2_id),
    CONSTRAINT user1_less_user2 CHECK (
        user1_id < friend.user2_id
    )
);

CREATE TABLE friend_bill(
    friend_bill_id uuid PRIMARY KEY,
    relation_id uuid NOT NULL REFERENCES friend(relation_id),
    amount double precision NOT NULL,
    title varchar(50) NOT NULL ,
    description varchar(100) NOT NULL default '',
    created_at timestamp NOT NULL default CURRENT_TIMESTAMP
)

