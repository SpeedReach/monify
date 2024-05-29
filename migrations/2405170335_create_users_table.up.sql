CREATE TABLE user_identity(
                              user_id uuid primary key ,
                              created_at timestamp default CURRENT_TIMESTAMP,
                              activated boolean default false,
                              refresh_token varchar(100),
                              device_token varchar(100)
);

CREATE TABLE email_login(
                            email varchar(100) primary key,
                            user_id uuid REFERENCES user_identity(user_id)
)

