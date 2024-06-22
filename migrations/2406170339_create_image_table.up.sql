CREATE TABLE tmp_file(
                         file_id uuid PRIMARY KEY,
                         path varchar(100) NOT NULL ,
                         expected_usage int8 NOT NULL ,
                         uploader uuid references user_identity(user_id) NOT NULL,
                         uploaded_at timestamp NOT NULL
);

CREATE TABLE confirmed_file(
                               file_id uuid PRIMARY KEY,
                               path varchar(100) NOT NULL ,
                               usage int8 NOT NULL ,
                               uploader uuid references user_identity(user_id) NOT NULL,
                               uploaded_at timestamp NOT NULL,
                               confirmed_at timestamp NOT NULL
)

