CREATE TABLE "public"."sys_delivery_company" (
                                                 "id" int8 NOT NULL,
                                                 "name" varchar(255) NOT NULL,
                                                 "logo" varchar(255),
                                                 "site" varchar(255),
                                                 "express_no" varchar(32),
                                                 "express_no_electronic_sheet" varchar(32),
                                                 "print_style_json" json,
                                                 "exp_type_json" json,
                                                 "updated_at" timestamp,
                                                 "created_at" timestamp,
                                                 PRIMARY KEY ("id"),
                                                 UNIQUE ("name")
)
;

ALTER TABLE "public"."sys_delivery_company"
    OWNER TO "kuaimk";

CREATE INDEX ON "public"."sys_delivery_company" (
    "name"
    );

COMMENT ON COLUMN "public"."sys_delivery_company"."name" IS '物流公司';

COMMENT ON COLUMN "public"."sys_delivery_company"."logo" IS 'LOGO';

COMMENT ON COLUMN "public"."sys_delivery_company"."site" IS '网址';

COMMENT ON COLUMN "public"."sys_delivery_company"."express_no" IS '物流跟踪编号';

COMMENT ON COLUMN "public"."sys_delivery_company"."express_no_electronic_sheet" IS '电子面单编号';

COMMENT ON COLUMN "public"."sys_delivery_company"."print_style_json" IS '打印模板样式';

COMMENT ON COLUMN "public"."sys_delivery_company"."exp_type_json" IS '业务类型';