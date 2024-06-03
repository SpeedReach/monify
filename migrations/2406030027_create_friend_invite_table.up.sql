CREATE TABLE friend_invite (
  invite_id uuid PRIMARY KEY,
  sender uuid NOT NULL,
  receiver uuid NOT NULL,
  created_at timestamp default CURRENT_TIMESTAMP,
  FOREIGN KEY (sender) REFERENCES "user_identity"(user_id)
  FOREIGN KEY (receiver) REFERENCES "user_identity"(user_id)
);