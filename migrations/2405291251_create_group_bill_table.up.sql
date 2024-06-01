

CREATE TABLE group_bill (
    bill_id uuid PRIMARY KEY,
    group_id uuid NOT NULL ,
    created_by uuid NOT NULL ,
    total_money double precision NOT NULL ,
    title varchar(30) NOT NULL ,
    description varchar(200) default '',
    FOREIGN KEY (group_id) REFERENCES "group"(group_id),
    FOREIGN KEY (created_by) REFERENCES group_member(group_member_id)
);

CREATE TABLE group_prepaid_bill(
    bill_id uuid NOT NULL ,
    person uuid NOT NULL ,
    amount double precision NOT NULL,
    PRIMARY KEY (bill_id, person),
    FOREIGN KEY (person) REFERENCES group_member(group_member_id),
    FOREIGN KEY (bill_id) REFERENCES group_bill(bill_id)
);

CREATE TABLE group_split_bill(
    bill_id uuid NOT NULL ,
    person uuid NOT NULL ,
    amount double precision NOT NULL,
    PRIMARY KEY (bill_id, person),
    FOREIGN KEY (person) REFERENCES group_member(group_member_id),
    FOREIGN KEY (bill_id) REFERENCES group_bill(bill_id)
)
