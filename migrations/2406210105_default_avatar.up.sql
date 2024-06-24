INSERT INTO user_identity(user_id, name) VALUES ('00000000-0000-0000-0000-000000000000', 'system');
INSERT INTO confirmed_file(
                           file_id,
                           path,
                           usage,
                           uploader,
                           uploaded_at,
                           confirmed_at
) VALUES (
             '00000000-0000-0000-0000-000000000000',
                            'default_avatar.png',
                            1,
             '00000000-0000-0000-0000-000000000000',
                            CURRENT_TIMESTAMP,
                            CURRENT_TIMESTAMP
);


ALTER TABLE user_identity ADD COLUMN avatar_id uuid DEFAULT '00000000-0000-0000-0000-000000000000' NOT NULL;