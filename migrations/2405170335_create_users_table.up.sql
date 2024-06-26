
CREATE TABLE user_identity(
                              user_id uuid primary key ,
                              created_at timestamp default CURRENT_TIMESTAMP,
                              activated boolean default false,
                              refresh_token varchar(100),
                              device_token varchar(100),
                              nick_id VARCHAR(150) UNIQUE
);

CREATE TABLE email_login(
                            email varchar(100) primary key,
                            user_id uuid,
                            FOREIGN KEY(user_id) REFERENCES user_identity(user_id)
)

