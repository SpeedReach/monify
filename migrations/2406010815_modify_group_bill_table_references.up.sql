ALTER TABLE group_bill DROP COLUMN created_by;
ALTER TABLE group_bill ADD COLUMN created_by uuid REFERENCES user_identity (user_id);
ALTER TABLE group_prepaid_bill DROP COLUMN person;
ALTER TABLE group_prepaid_bill ADD COLUMN person uuid REFERENCES user_identity (user_id);
ALTER TABLE group_split_bill DROP COLUMN person;
ALTER TABLE group_split_bill ADD COLUMN person uuid REFERENCES user_identity (user_id)