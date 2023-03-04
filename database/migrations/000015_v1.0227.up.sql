ALTER TABLE "public"."co_fd_invoice_detail"
    ADD COLUMN "belong_to" int2;

COMMENT ON COLUMN "public"."co_fd_invoice_detail"."belong_to" IS '发票拥有者类型：1个人  2主体';