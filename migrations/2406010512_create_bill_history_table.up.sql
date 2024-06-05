
CREATE TABLE group_bill_history_type(
    name char(10),
    type int PRIMARY KEY
);

INSERT INTO group_bill_history_type (name, type) VALUES  ('create', 0);
INSERT INTO group_bill_history_type (name, type) VALUES  ('delete', 1);
INSERT INTO group_bill_history_type (name, type) VALUES  ('modify', 2);

CREATE TABLE group_bill_history(
    history_id uuid PRIMARY KEY,
    type int REFERENCES group_bill_history_type(type),
    bill_id uuid,
    group_id uuid REFERENCES "group"(group_id),
    title varchar(30) NOT NULL,
    operator uuid REFERENCES group_member(group_member_id),
    timestamp timestamp default CURRENT_TIMESTAMP
)

