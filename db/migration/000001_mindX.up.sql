CREATE TABLE "user"
(
    "id"                bigserial PRIMARY KEY,
    "name"              varchar not null,
    "permanent_address" varchar not null,
    "current_address"   varchar not null,
    "current_status"    int     not null
);

CREATE TABLE "location"
(
    "id"       bigserial PRIMARY KEY,
    "location" varchar not null
);

CREATE TABLE "location_history"
(
    "id"           bigserial PRIMARY KEY,
    "user_id"      bigint      not null,
    "type"         int         not null,
    "location_id"  int         not null,
    "manual_input" varchar     not null,
    "date"         timestamptz not null
);

ALTER TABLE "location_history"
    ADD FOREIGN KEY ("user_id") REFERENCES "user" ("id");

ALTER TABLE "location_history"
    ADD FOREIGN KEY ("location_id") REFERENCES "location" ("id");


COMMENT
ON COLUMN "user"."current_status" IS '
    -1: unknow,
    0: F0,
    1: F1,
    2: F2,
  ';

COMMENT
ON COLUMN "location_history"."type" IS '
    1 - manual_input,
    2 - agent_location,
  ';

COMMENT
ON COLUMN "location_history"."location_id" IS '
    = 0: user manual input,
    > 0: user scan code
    from agent location
  ';
