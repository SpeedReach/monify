

CREATE TABLE group_bill (
    bill_id uuid PRIMARY KEY,
    group_id uuid REFERENCES "group" (group_id),
    created_by uuid REFERENCES "group_member" (group_member_id),
    total_money double precision NOT NULL ,
    title varchar(30) NOT NULL ,
    description varchar(200) default ''
);

CREATE TABLE group_prepaid_bill(
    bill_id uuid REFERENCES group_bill(bill_id),
    person uuid REFERENCES group_member(group_member_id),
    amount double precision NOT NULL,
    PRIMARY KEY (bill_id, person)
);

CREATE TABLE group_split_bill(
    bill_id uuid REFERENCES group_bill(bill_id),
    person uuid REFERENCES group_member(group_member_id),
    amount double precision NOT NULL,
    PRIMARY KEY (bill_id, person)
)
