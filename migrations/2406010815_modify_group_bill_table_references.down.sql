ALTER TABLE group_bill DROP COLUMN created_by;
ALTER TABLE group_bill ADD COLUMN created_by uuid REFERENCES group_member(group_member_id);
ALTER TABLE group_prepaid_bill DROP COLUMN person;
ALTER TABLE group_prepaid_bill ADD COLUMN person uuid REFERENCES group_member(group_member_id);
ALTER TABLE group_split_bill DROP COLUMN person;
ALTER TABLE group_split_bill ADD COLUMN person uuid REFERENCES group_member(group_member_id)