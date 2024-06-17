CREATE TABLE TmpImage(
    imgId uuid PRIMARY KEY,
    url varchar(100) NOT NULL ,
    expected_usage int8 NOT NULL ,
    uploader uuid references user_identity(user_id) NOT NULL,
    uploaded_at timestamp NOT NULL
);

CREATE TABLE ConfirmedImage(
    imgId uuid PRIMARY KEY,
    url varchar(100) NOT NULL ,
    usage int8 NOT NULL ,
    uploader uuid references user_identity(user_id) NOT NULL,
    uploaded_at timestamp NOT NULL,
    confirmed_at timestamp NOT NULL
)