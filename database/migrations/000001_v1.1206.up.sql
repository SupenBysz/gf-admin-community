-- ----------------------------
-- Table structure for fd_account
-- ----------------------------
DROP TABLE IF EXISTS "public"."fd_account";
CREATE TABLE "public"."fd_account" (
                                       "id" int8 NOT NULL,
                                       "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                       "union_license_id" int8 DEFAULT 0,
                                       "union_user_id" int8 NOT NULL,
                                       "currency_code" varchar(8) COLLATE "pg_catalog"."default" NOT NULL,
                                       "is_enabled" int2 DEFAULT 1,
                                       "limit_state" int4 DEFAULT 0,
                                       "precision_of_balance" int4 NOT NULL DEFAULT 100,
                                       "balance" int8 DEFAULT 0,
                                       "created_at" timestamp(6),
                                       "updated_at" timestamp(6),
                                       "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."fd_account"."id" IS 'ID';
COMMENT ON COLUMN "public"."fd_account"."name" IS '账户名称';
COMMENT ON COLUMN "public"."fd_account"."union_license_id" IS '关联资质ID，大于0时必须保值与 union_user_id 关联得上';
COMMENT ON COLUMN "public"."fd_account"."union_user_id" IS '关联用户ID';
COMMENT ON COLUMN "public"."fd_account"."currency_code" IS '货币代码';
COMMENT ON COLUMN "public"."fd_account"."is_enabled" IS '是否启用：1启用，0禁用';
COMMENT ON COLUMN "public"."fd_account"."limit_state" IS '限制状态：0不限制，1限制支出、2限制收入';
COMMENT ON COLUMN "public"."fd_account"."precision_of_balance" IS '货币单位精度：1:元，10:角，100:分，1000:厘，10000:毫，……';
COMMENT ON COLUMN "public"."fd_account"."balance" IS '当前余额，必须要与账单最后一笔交易余额对应得上';

-- ----------------------------
-- Records of fd_account
-- ----------------------------

-- ----------------------------
-- Indexes structure for table fd_account
-- ----------------------------
CREATE INDEX "fd_account_currency_code_idx" ON "public"."fd_account" USING btree (
    "currency_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_id_idx" ON "public"."fd_account" USING btree (
    "id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_name_idx" ON "public"."fd_account" USING btree (
    "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table fd_account
-- ----------------------------
ALTER TABLE "public"."fd_account" ADD CONSTRAINT "fd_account_pkey" PRIMARY KEY ("id");




-- ----------------------------
-- Table structure for fd_account_bill
-- ----------------------------
DROP TABLE IF EXISTS "public"."fd_account_bill";
CREATE TABLE "public"."fd_account_bill" (
                                            "id" int8 NOT NULL,
                                            "from_user_id" int8 NOT NULL,
                                            "to_user_id" int8 NOT NULL,
                                            "fd_account_id" int8 NOT NULL,
                                            "before_balance" int8 NOT NULL,
                                            "amount" int8 NOT NULL,
                                            "after_balance" int8 NOT NULL,
                                            "union_order_id" int8,
                                            "in_out_type" int4 NOT NULL,
                                            "trade_type" int4 NOT NULL,
                                            "trade_at" timestamp(6) NOT NULL,
                                            "remark" varchar(255) COLLATE "pg_catalog"."default",
                                            "trade_state" int4 NOT NULL,
                                            "created_at" timestamp(6) NOT NULL,
                                            "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."fd_account_bill"."id" IS 'ID';
COMMENT ON COLUMN "public"."fd_account_bill"."from_user_id" IS '交易发起方UserID，如果是系统则固定为-1';
COMMENT ON COLUMN "public"."fd_account_bill"."to_user_id" IS '交易对象UserID';
COMMENT ON COLUMN "public"."fd_account_bill"."fd_account_id" IS '财务账户ID';
COMMENT ON COLUMN "public"."fd_account_bill"."before_balance" IS '交易前账户余额';
COMMENT ON COLUMN "public"."fd_account_bill"."amount" IS '交易金额';
COMMENT ON COLUMN "public"."fd_account_bill"."after_balance" IS '交易后账户余额';
COMMENT ON COLUMN "public"."fd_account_bill"."union_order_id" IS '关联业务订单ID';
COMMENT ON COLUMN "public"."fd_account_bill"."in_out_type" IS '收支类型：1收入，2支出';
COMMENT ON COLUMN "public"."fd_account_bill"."trade_type" IS '交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费、128提现、256充值、512罚金、1024营收、2048生活缴费，8192其它';
COMMENT ON COLUMN "public"."fd_account_bill"."trade_at" IS '交易时间';
COMMENT ON COLUMN "public"."fd_account_bill"."remark" IS '备注信息';
COMMENT ON COLUMN "public"."fd_account_bill"."trade_state" IS '交易状态：1待支付、2支付中、4已支付、8支付失败、16交易完成、';

-- ----------------------------
-- Records of fd_account_bill
-- ----------------------------

-- ----------------------------
-- Indexes structure for table fd_account_bill
-- ----------------------------
CREATE INDEX "fd_account_bill_fd_account_id_idx" ON "public"."fd_account_bill" USING btree (
    "fd_account_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_bill_from_user_id_idx" ON "public"."fd_account_bill" USING btree (
    "from_user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "fd_account_bill_id_idx" ON "public"."fd_account_bill" USING btree (
    "id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_bill_in_out_type_idx" ON "public"."fd_account_bill" USING btree (
    "in_out_type" "pg_catalog"."int4_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_bill_to_user_id_idx" ON "public"."fd_account_bill" USING btree (
    "to_user_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_bill_trade_at_idx" ON "public"."fd_account_bill" USING btree (
    "trade_at" "pg_catalog"."timestamp_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_bill_trade_type_idx" ON "public"."fd_account_bill" USING btree (
    "trade_type" "pg_catalog"."int4_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_account_bill_union_order_id_idx" ON "public"."fd_account_bill" USING btree (
    "union_order_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table fd_account_bill
-- ----------------------------
ALTER TABLE "public"."fd_account_bill" ADD CONSTRAINT "fd_account_bill_pkey" PRIMARY KEY ("id");




-- ----------------------------
-- Table structure for fd_bank_card
-- ----------------------------
DROP TABLE IF EXISTS "public"."fd_bank_card";
CREATE TABLE "public"."fd_bank_card" (
                                         "id" int8 NOT NULL,
                                         "bank_name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
                                         "card_type" int2 NOT NULL DEFAULT 1,
                                         "card_number" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                         "expired_at" timestamp(6),
                                         "holder_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                         "bank_of_account" varchar(255) COLLATE "pg_catalog"."default",
                                         "state" int4 DEFAULT 1,
                                         "remark" varchar(255) COLLATE "pg_catalog"."default" DEFAULT ''::character varying,
                                         "created_at" timestamp(6),
                                         "updated_at" timestamp(6),
                                         "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."fd_bank_card"."id" IS 'ID';
COMMENT ON COLUMN "public"."fd_bank_card"."bank_name" IS '银行名称';
COMMENT ON COLUMN "public"."fd_bank_card"."card_type" IS '银行卡类型：1借记卡，2储蓄卡';
COMMENT ON COLUMN "public"."fd_bank_card"."card_number" IS '银行卡号';
COMMENT ON COLUMN "public"."fd_bank_card"."expired_at" IS '有效期';
COMMENT ON COLUMN "public"."fd_bank_card"."holder_name" IS '银行卡开户名';
COMMENT ON COLUMN "public"."fd_bank_card"."bank_of_account" IS '开户行';
COMMENT ON COLUMN "public"."fd_bank_card"."state" IS '状态：0禁用，1正常';
COMMENT ON COLUMN "public"."fd_bank_card"."remark" IS '备注信息';

-- ----------------------------
-- Records of fd_bank_card
-- ----------------------------

-- ----------------------------
-- Indexes structure for table fd_bank_card
-- ----------------------------
CREATE INDEX "fd_bank_card_bank_name_idx" ON "public"."fd_bank_card" USING btree (
    "bank_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_bank_card_card_number_idx" ON "public"."fd_bank_card" USING btree (
    "card_number" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_bank_card_holder_name_idx" ON "public"."fd_bank_card" USING btree (
    "holder_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_bank_card_id_idx" ON "public"."fd_bank_card" USING btree (
    "id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table fd_bank_card
-- ----------------------------
ALTER TABLE "public"."fd_bank_card" ADD CONSTRAINT "fd_bank_card_card_number_key" UNIQUE ("card_number");

-- ----------------------------
-- Primary Key structure for table fd_bank_card
-- ----------------------------
ALTER TABLE "public"."fd_bank_card" ADD CONSTRAINT "fd_bank_card_pkey" PRIMARY KEY ("id");



-- ----------------------------
-- Table structure for fd_currenty
-- ----------------------------
DROP TABLE IF EXISTS "public"."fd_currenty";
CREATE TABLE "public"."fd_currenty" (
                                        "code" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                        "en_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                        "cn_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                        "currency_code" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                        "currency_cn" varchar(32) COLLATE "pg_catalog"."default",
                                        "currency_en" varchar(32) COLLATE "pg_catalog"."default",
                                        "symbol" varchar(16) COLLATE "pg_catalog"."default",
                                        "symbol_native" varchar(16) COLLATE "pg_catalog"."default",
                                        "is_legal_tender" int2 NOT NULL
)
;
COMMENT ON COLUMN "public"."fd_currenty"."code" IS '国家编码';
COMMENT ON COLUMN "public"."fd_currenty"."en_name" IS '国家英文名称';
COMMENT ON COLUMN "public"."fd_currenty"."cn_name" IS '国家中文名称';
COMMENT ON COLUMN "public"."fd_currenty"."currency_code" IS '货币编码';
COMMENT ON COLUMN "public"."fd_currenty"."currency_cn" IS '货币中文名称';
COMMENT ON COLUMN "public"."fd_currenty"."currency_en" IS '货币英文名称';
COMMENT ON COLUMN "public"."fd_currenty"."symbol" IS '货币符号';
COMMENT ON COLUMN "public"."fd_currenty"."symbol_native" IS '货币原生符号';
COMMENT ON COLUMN "public"."fd_currenty"."is_legal_tender" IS '是否法定货币：1是，0否';

-- ----------------------------
-- Records of fd_currenty
-- ----------------------------
INSERT INTO "public"."fd_currenty" VALUES ('US', 'United States of America (USA)', '美国', 'USD', '美元', 'US Dollar', '$', '$', 1);
INSERT INTO "public"."fd_currenty" VALUES ('HK', 'Hong Kong', '香港', 'HKD', '港元', 'Hong Kong Dollar', 'HK$', '$', 1);
INSERT INTO "public"."fd_currenty" VALUES ('TW', 'Taiwan', '台湾', 'TWD', '新台币', 'New Taiwan Dollar', 'NT$', '$', 1);
INSERT INTO "public"."fd_currenty" VALUES ('JP', 'Japan', '日本', 'JPY', '日元', 'Japanese Yen', 'JP¥', '¥', 1);
INSERT INTO "public"."fd_currenty" VALUES ('CN', 'China', '中国', 'CNY', '人民币元', 'Chinese Yuan', '¥', '¥', 1);
INSERT INTO "public"."fd_currenty" VALUES ('EU', 'European Union', '欧盟', 'EUR', '欧元', 'Euro', '€', '€', 1);
INSERT INTO "public"."fd_currenty" VALUES ('BTC', 'Bitcoin', 'BTC', 'BTC', '比特币', 'BTC', '฿', '฿', 0);
INSERT INTO "public"."fd_currenty" VALUES ('USDT', 'TetherUSD', 'USDT', 'USDT', '泰达币', 'USDT', 'USDT$', '$', 0);
INSERT INTO "public"."fd_currenty" VALUES ('ETH', 'Ethereum', 'ETH', 'ETH', '以太币', 'ETH', 'ETH$', '$', 0);

-- ----------------------------
-- Indexes structure for table fd_currenty
-- ----------------------------
CREATE INDEX "fd_currenty_cn_name_idx" ON "public"."fd_currenty" USING btree (
    "cn_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_currenty_code_idx" ON "public"."fd_currenty" USING btree (
    "code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_currenty_currency_code_idx" ON "public"."fd_currenty" USING btree (
    "currency_code" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "fd_currenty_en_name_idx" ON "public"."fd_currenty" USING btree (
    "en_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table fd_currenty
-- ----------------------------
ALTER TABLE "public"."fd_currenty" ADD CONSTRAINT "fd_currenty_code_key" UNIQUE ("code");
ALTER TABLE "public"."fd_currenty" ADD CONSTRAINT "fd_currenty_currency_code_key" UNIQUE ("currency_code");

-- ----------------------------
-- Primary Key structure for table fd_currenty
-- ----------------------------
ALTER TABLE "public"."fd_currenty" ADD CONSTRAINT "fd_currenty_pkey" PRIMARY KEY ("currency_code");




-- ----------------------------
-- Table structure for fd_invoice
-- ----------------------------
DROP TABLE IF EXISTS "public"."fd_invoice";
CREATE TABLE "public"."fd_invoice" (
                                       "id" int8 NOT NULL,
                                       "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                       "tax_id" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                       "addr" varchar(255) COLLATE "pg_catalog"."default",
                                       "email" varchar(128) COLLATE "pg_catalog"."default",
                                       "user_id" int8 NOT NULL,
                                       "audit_user_id" int8 NOT NULL,
                                       "audit_replay_msg" varchar(255) COLLATE "pg_catalog"."default",
                                       "audit_at" timestamp(6) NOT NULL,
                                       "state" int4 NOT NULL DEFAULT 1,
                                       "created_at" timestamp(6),
                                       "updated_at" timestamp(6),
                                       "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."fd_invoice"."name" IS '发票抬头名称';
COMMENT ON COLUMN "public"."fd_invoice"."tax_id" IS '纳税识别号';
COMMENT ON COLUMN "public"."fd_invoice"."addr" IS '发票收件地址，限纸质';
COMMENT ON COLUMN "public"."fd_invoice"."email" IS '发票收件邮箱，限电子发票';
COMMENT ON COLUMN "public"."fd_invoice"."user_id" IS '申请人UserID';
COMMENT ON COLUMN "public"."fd_invoice"."audit_user_id" IS '审核人UserID';
COMMENT ON COLUMN "public"."fd_invoice"."audit_replay_msg" IS '审核回复，仅审核不通过时才有值';
COMMENT ON COLUMN "public"."fd_invoice"."audit_at" IS '审核时间';
COMMENT ON COLUMN "public"."fd_invoice"."state" IS '状态：1待审核、2已通过、3不通过';

-- ----------------------------
-- Records of fd_invoice
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table fd_invoice
-- ----------------------------
ALTER TABLE "public"."fd_invoice" ADD CONSTRAINT "fd_invoice_pkey" PRIMARY KEY ("id");



-- ----------------------------
-- Table structure for fd_invoice_detail
-- ----------------------------
DROP TABLE IF EXISTS "public"."fd_invoice_detail";
CREATE TABLE "public"."fd_invoice_detail" (
                                              "id" int8 NOT NULL,
                                              "tax_number" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                              "tax_name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                              "bill_ids" text COLLATE "pg_catalog"."default" NOT NULL,
                                              "amount" int8 NOT NULL,
                                              "rate" int4 NOT NULL,
                                              "rate_mount" int8 NOT NULL,
                                              "remark" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
                                              "type" int2 NOT NULL,
                                              "state" int4 NOT NULL,
                                              "audit_user_ids" int8,
                                              "make_type" int4,
                                              "make_user_id" int8,
                                              "make_at" timestamp(6),
                                              "courier_name" varchar(32) COLLATE "pg_catalog"."default",
                                              "courier_number" varchar(64) COLLATE "pg_catalog"."default",
                                              "fd_invoice_id" int8,
                                              "audit_user_id" int8,
                                              "audit_reply_msg" varchar(255) COLLATE "pg_catalog"."default",
                                              "audit_at" timestamp(6),
                                              "created_at" timestamp(6),
                                              "updated_at" timestamp(6),
                                              "deleted_at" timestamp(6)
)
;
COMMENT ON COLUMN "public"."fd_invoice_detail"."id" IS 'ID';
COMMENT ON COLUMN "public"."fd_invoice_detail"."tax_number" IS '纳税识别号';
COMMENT ON COLUMN "public"."fd_invoice_detail"."tax_name" IS '纳税人名称';
COMMENT ON COLUMN "public"."fd_invoice_detail"."bill_ids" IS '账单ID组';
COMMENT ON COLUMN "public"."fd_invoice_detail"."amount" IS '开票金额，单位精度：分';
COMMENT ON COLUMN "public"."fd_invoice_detail"."rate" IS '税率，如3% 则填入3';
COMMENT ON COLUMN "public"."fd_invoice_detail"."rate_mount" IS '税额，单位精度：分';
COMMENT ON COLUMN "public"."fd_invoice_detail"."remark" IS '发布内容描述';
COMMENT ON COLUMN "public"."fd_invoice_detail"."type" IS '发票类型：1电子发票，2纸质发票';
COMMENT ON COLUMN "public"."fd_invoice_detail"."state" IS '状态：1待审核、2待开票、4开票失败、8已开票、16已撤销';
COMMENT ON COLUMN "public"."fd_invoice_detail"."audit_user_ids" IS '审核者UserID，多个用逗号隔开';
COMMENT ON COLUMN "public"."fd_invoice_detail"."make_type" IS '出票类型：1普通发票、2增值税专用发票、3专业发票';
COMMENT ON COLUMN "public"."fd_invoice_detail"."make_user_id" IS '出票人UserID，如果是系统出票则默认-1';
COMMENT ON COLUMN "public"."fd_invoice_detail"."make_at" IS '出票时间';
COMMENT ON COLUMN "public"."fd_invoice_detail"."courier_name" IS '快递名称';
COMMENT ON COLUMN "public"."fd_invoice_detail"."courier_number" IS '快递编号';
COMMENT ON COLUMN "public"."fd_invoice_detail"."fd_invoice_id" IS '发票抬头ID';
COMMENT ON COLUMN "public"."fd_invoice_detail"."audit_user_id" IS '审核者UserID';
COMMENT ON COLUMN "public"."fd_invoice_detail"."audit_reply_msg" IS '审核回复，仅审核不通过时才有值';
COMMENT ON COLUMN "public"."fd_invoice_detail"."audit_at" IS '审核时间';

-- ----------------------------
-- Records of fd_invoice_detail
-- ----------------------------

-- ----------------------------
-- Primary Key structure for table fd_invoice_detail
-- ----------------------------
ALTER TABLE "public"."fd_invoice_detail" ADD CONSTRAINT "fd_invoice_detail_pkey" PRIMARY KEY ("id");
