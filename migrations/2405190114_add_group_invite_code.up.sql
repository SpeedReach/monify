ALTER TABLE "group" ADD COLUMN "invite_code" CHAR(6) default '';
ALTER TABLE "group" ADD COLUMN "invite_code_expires" TIMESTAMP default 0;

