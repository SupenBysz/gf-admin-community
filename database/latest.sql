/*
 Navicat Premium Data Transfer

 Source Server         : 225-gf-admin
 Source Server Type    : PostgreSQL
 Source Server Version : 140005 (140005)
 Source Host           : 182.43.195.225:15432
 Source Catalog        : gf-admin
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140005 (140005)
 File Encoding         : 65001

 Date: 18/02/2023 10:01:30
*/


-- ----------------------------
-- Sequence structure for sys_area_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_area_id_seq";
CREATE SEQUENCE "public"."sys_area_id_seq"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."sys_area_id_seq" OWNER TO "kysion";

-- ----------------------------
-- Sequence structure for sys_area_id_seq1
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_area_id_seq1";
CREATE SEQUENCE "public"."sys_area_id_seq1"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."sys_area_id_seq1" OWNER TO "kysion";

-- ----------------------------
-- Sequence structure for sys_area_id_seq2
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_area_id_seq2";
CREATE SEQUENCE "public"."sys_area_id_seq2"
    INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;
ALTER SEQUENCE "public"."sys_area_id_seq2" OWNER TO "kysion";

-- ----------------------------
-- Table structure for co_company
-- ----------------------------
DROP TABLE IF EXISTS "public"."co_company";
CREATE TABLE "public"."co_company" (
                                       "id" int8 NOT NULL,
                                       "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                       "contact_name" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                       "contact_mobile" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                       "user_id" int8 NOT NULL,
                                       "parent_id" int8 DEFAULT 0,
                                       "state" int4,
                                       "remark" text COLLATE "pg_catalog"."default",
                                       "created_by" int8,
                                       "created_at" timestamp(6),
                                       "updated_by" int8,
                                       "updated_at" timestamp(6),
                                       "deleted_by" int8,
                                       "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."co_company" OWNER TO "kysion";
COMMENT ON COLUMN "public"."co_company"."id" IS 'ID';
COMMENT ON COLUMN "public"."co_company"."name" IS '名称';
COMMENT ON COLUMN "public"."co_company"."contact_name" IS '商务联系人';
COMMENT ON COLUMN "public"."co_company"."contact_mobile" IS '商务联系电话';
COMMENT ON COLUMN "public"."co_company"."user_id" IS '管理员ID';
COMMENT ON COLUMN "public"."co_company"."parent_id" IS '父级ID';
COMMENT ON COLUMN "public"."co_company"."state" IS '状态：0未启用，1正常';
COMMENT ON COLUMN "public"."co_company"."remark" IS '备注';
COMMENT ON COLUMN "public"."co_company"."created_by" IS '创建者';
COMMENT ON COLUMN "public"."co_company"."created_at" IS '创建时间';
COMMENT ON COLUMN "public"."co_company"."updated_by" IS '更新者';
COMMENT ON COLUMN "public"."co_company"."updated_at" IS '更新时间';
COMMENT ON COLUMN "public"."co_company"."deleted_by" IS '删除者';
COMMENT ON COLUMN "public"."co_company"."deleted_at" IS '删除时间';

-- ----------------------------
-- Records of co_company
-- ----------------------------
BEGIN;
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987075795648581, '羊羊村', '喜洋洋', '18625837911', 0, 0, NULL, '我是喜洋洋与灰太狼的羊村', 5302581852373061, '2023-01-12 10:52:59', NULL, '2023-01-12 10:52:59', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5994601065021509, '熊出没', '光头强', '18613552505', 0, 0, NULL, '熊熊出没', 5977706248405061, '2023-01-13 18:45:55', NULL, '2023-01-13 18:45:55', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5994616775639109, '熊出没2', '光头强', '18613552505', 0, 0, NULL, '熊熊出没', 5977706248405061, '2023-01-13 18:49:54', NULL, '2023-01-13 18:49:54', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5994608681549893, '熊出没1', '光头强', '18613552505', 0, 0, NULL, '熊熊出没', 5977706248405061, '2023-01-13 18:47:51', NULL, '2023-01-13 18:47:51', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5994620680011845, '熊出没3', '光头强', '18613552505', 0, 0, NULL, '熊熊出没', 5977706248405061, '2023-01-13 18:50:54', NULL, '2023-01-13 18:50:54', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5994627731292229, '熊出没4', '光头强', '18613552505', 0, 0, NULL, '熊熊出没', 5977706248405061, '2023-01-13 18:52:42', NULL, '2023-01-13 18:52:42', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5994632857124933, '熊出没5', '光头强', '18613552505', 0, 0, NULL, '熊熊出没', 5977706248405061, '2023-01-13 18:54:00', NULL, '2023-01-13 18:54:00', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987167197659205, '魔法屋', '蓝精灵', '18667667382', 5987167558369349, 0, NULL, '在山的那边海的那边有一群蓝精灵', 5302581852373061, '2023-01-12 11:15:43', 5977706248405061, '2023-01-14 10:45:11', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987149274939461, '魔法小镇', '女巫', '18625837911', 5987149735460933, 0, NULL, '我是魔法小镇', 5302581852373061, '2023-01-12 11:10:59', NULL, '2023-01-12 11:10:59', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5977719809835077, '魔法城堡', '游乐王子', '18661667015', 5977719818551365, 0, NULL, '我是魔法城堡', 5302581852373061, '2023-01-10 19:12:51', NULL, '2023-01-10 19:12:51', NULL, NULL);
INSERT INTO "public"."co_company" ("id", "name", "contact_name", "contact_mobile", "user_id", "parent_id", "state", "remark", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5977706236739653, '快乐星球', '小小怪', '18661667165', 5977706248405061, 0, NULL, '我是快乐星球', 5302581852373061, '2023-01-10 19:09:24', NULL, '2023-01-10 19:09:24', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for co_company_employee
-- ----------------------------
DROP TABLE IF EXISTS "public"."co_company_employee";
CREATE TABLE "public"."co_company_employee" (
                                                "id" int8 NOT NULL,
                                                "no" varchar(16) COLLATE "pg_catalog"."default",
                                                "avatar" text COLLATE "pg_catalog"."default",
                                                "name" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                                "mobile" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                                "union_main_id" int8 DEFAULT 0,
                                                "state" int2 NOT NULL DEFAULT 0,
                                                "last_active_ip" varchar(64) COLLATE "pg_catalog"."default",
                                                "hired_at" timestamp(6),
                                                "created_by" int8,
                                                "created_at" timestamp(6),
                                                "updated_by" int8,
                                                "updated_at" timestamp(6),
                                                "deleted_by" int8,
                                                "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."co_company_employee" OWNER TO "kysion";
COMMENT ON COLUMN "public"."co_company_employee"."id" IS 'ID，保持与USERID一致';
COMMENT ON COLUMN "public"."co_company_employee"."no" IS '工号';
COMMENT ON COLUMN "public"."co_company_employee"."avatar" IS '头像';
COMMENT ON COLUMN "public"."co_company_employee"."name" IS '姓名';
COMMENT ON COLUMN "public"."co_company_employee"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."co_company_employee"."union_main_id" IS '所属主体';
COMMENT ON COLUMN "public"."co_company_employee"."state" IS '状态： -1已离职，0待确认，1已入职';
COMMENT ON COLUMN "public"."co_company_employee"."last_active_ip" IS '最后活跃IP';
COMMENT ON COLUMN "public"."co_company_employee"."hired_at" IS '入职时间';

-- ----------------------------
-- Records of co_company_employee
-- ----------------------------
BEGIN;
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5977719818551365, '001', '', '游乐王子', '18661667015', 5977719809835077, 0, '', '2023-01-10 19:12:47.675531', 5302581852373061, '2023-01-10 19:12:51', 0, '2023-01-10 19:12:51', 0, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5978295989436485, '990', '', '开心超人', '18148481217', 5977706236739653, 0, '', NULL, 5977706248405061, '2023-01-10 21:39:20', 5977706248405061, '2023-01-10 21:48:51', 0, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5977932491456581, '554', '', '大大怪', '18627331021', 5977706236739653, 0, '', NULL, 5977706248405061, '2023-01-10 20:06:54', 0, '2023-01-10 20:06:54', 0, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5981433713131589, '912', '', '甜心超人', '18163386366', 5977706236739653, 0, '', NULL, 5977706248405061, '2023-01-11 10:57:18', 0, '2023-01-11 10:57:18', 0, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5981495520723013, '913', '', '粗心超人', '18163386366', 5977706236739653, 0, '', NULL, 5977706248405061, '2023-01-11 11:13:01', 0, '2023-01-11 11:13:01', 0, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987149735460933, '001', '', '女巫', '18625837911', 5987149274939461, 0, NULL, '2023-01-12 11:10:50.983602', 5302581852373061, '2023-01-12 11:10:59', NULL, '2023-01-12 11:10:59', NULL, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987229831856197, '915', '', '花心超人', '18155476184', 5977706236739653, 0, NULL, NULL, 5977706248405061, '2023-01-12 11:31:37', NULL, '2023-01-12 11:31:37', NULL, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987317698330693, '784', '', '暖心超人', '18355476184', 5977706236739653, 0, NULL, NULL, 5977706248405061, '2023-01-12 11:53:40', NULL, '2023-01-12 11:53:40', NULL, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5987167558369349, '', '', '蓝精灵', '18148314149', 5977706236739653, 0, NULL, '2023-01-12 11:15:24.627681', 5302581852373061, '2023-01-12 11:15:42', 5977706248405061, '2023-01-12 14:51:30', NULL, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (5977706248405061, '', '', '小小怪', '18678153083', 5977706236739653, 0, '', '2023-01-10 19:09:20.568787', 5302581852373061, '2023-01-10 19:09:24', 5977706248405061, '2023-01-12 21:18:57', 0, NULL);
INSERT INTO "public"."co_company_employee" ("id", "no", "avatar", "name", "mobile", "union_main_id", "state", "last_active_ip", "hired_at", "created_by", "created_at", "updated_by", "updated_at", "deleted_by", "deleted_at") VALUES (6000781163036741, '019', '', '灰心超人', '19890738726', 5977706236739653, 0, NULL, NULL, 5977706248405061, '2023-01-14 20:57:36', NULL, '2023-01-14 20:57:36', NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for co_company_team
-- ----------------------------
DROP TABLE IF EXISTS "public"."co_company_team";
CREATE TABLE "public"."co_company_team" (
                                            "id" int8 NOT NULL,
                                            "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                            "owner_employee_id" int8,
                                            "captain_employee_id" int8,
                                            "union_main_id" int8 NOT NULL,
                                            "parent_id" int8,
                                            "remark" text COLLATE "pg_catalog"."default",
                                            "created_at" timestamp(6),
                                            "updated_at" timestamp(6),
                                            "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."co_company_team" OWNER TO "kysion";
COMMENT ON COLUMN "public"."co_company_team"."name" IS '团队名称，公司维度下唯一';
COMMENT ON COLUMN "public"."co_company_team"."owner_employee_id" IS '团队所有者/业务总监/业务经理/团队队长';
COMMENT ON COLUMN "public"."co_company_team"."captain_employee_id" IS '团队队长编号/小组组长';
COMMENT ON COLUMN "public"."co_company_team"."union_main_id" IS '所属主体单位ID';
COMMENT ON COLUMN "public"."co_company_team"."parent_id" IS '父级ID';
COMMENT ON COLUMN "public"."co_company_team"."remark" IS '备注';

-- ----------------------------
-- Records of co_company_team
-- ----------------------------
BEGIN;
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5978661809225797, '快乐星球Team2', 5978295989436485, 5978295989436485, 5977706236739653, 0, '我是团队', '2023-01-10 23:12:25', '2023-01-10 23:12:25', NULL);
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5987253519122501, '快乐星球Team3', 5987229831856197, 5987229831856197, 5977706236739653, 0, '我是花心超人带领的团队', '2023-01-12 11:37:20', '2023-01-12 11:56:06', NULL);
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5978345486090309, '快乐星球Team1', 5981495520723013, 5981495520723013, 5977706236739653, 0, '', '2023-01-10 21:51:54', '2023-01-12 12:14:11', NULL);
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5987413533196357, '快乐星球Team3的Group1', 5987229831856197, 5987229831856197, 5977706236739653, 5987253519122501, '我是花心超人带领的小组', '2023-01-12 12:18:24', '2023-01-12 12:18:24', NULL);
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5989573813796933, '快乐星球Team5', 0, 0, 5977706236739653, 0, '我是Team5', '2023-01-12 21:27:25', '2023-01-12 21:27:25', NULL);
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5989544611676229, '快乐星球Team4', 5981495520723013, 0, 5977706236739653, 0, '', '2023-01-12 21:19:59', '2023-01-14 17:09:07', NULL);
INSERT INTO "public"."co_company_team" ("id", "name", "owner_employee_id", "captain_employee_id", "union_main_id", "parent_id", "remark", "created_at", "updated_at", "deleted_at") VALUES (5999856856006725, '快乐星球Team6', 0, 0, 5977706236739653, 0, 'ut ut ex reprehenderit sed', '2023-01-14 17:02:35', '2023-01-14 17:09:23', NULL);
COMMIT;

-- ----------------------------
-- Table structure for co_company_team_member
-- ----------------------------
DROP TABLE IF EXISTS "public"."co_company_team_member";
CREATE TABLE "public"."co_company_team_member" (
                                                   "id" int8 NOT NULL,
                                                   "team_id" int8 NOT NULL,
                                                   "employee_id" int8 NOT NULL,
                                                   "invite_user_id" int8,
                                                   "union_main_id" int8 NOT NULL,
                                                   "join_at" timestamp(6)
)
;
ALTER TABLE "public"."co_company_team_member" OWNER TO "kysion";
COMMENT ON COLUMN "public"."co_company_team_member"."id" IS 'ID';
COMMENT ON COLUMN "public"."co_company_team_member"."team_id" IS '团队ID';
COMMENT ON COLUMN "public"."co_company_team_member"."employee_id" IS '成员ID';
COMMENT ON COLUMN "public"."co_company_team_member"."invite_user_id" IS '邀约人ID';
COMMENT ON COLUMN "public"."co_company_team_member"."union_main_id" IS '关联主体ID';
COMMENT ON COLUMN "public"."co_company_team_member"."join_at" IS '加入时间';

-- ----------------------------
-- Records of co_company_team_member
-- ----------------------------
BEGIN;
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5978345486090310, 5978345486090309, 5977932491456581, NULL, 5977706236739653, '2023-01-10 21:51:54.715808');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5981476842176581, 5978345486090309, 5981433713131589, NULL, 5977706236739653, '2023-01-11 11:08:16.428508');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5981815571611717, 5978345486090309, 5981495520723013, NULL, 5977706236739653, '2023-01-11 12:34:24.421719');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5987253519122502, 5987253519122501, 5987229831856197, NULL, 5977706236739653, '2023-01-12 11:37:20.512541');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5987319132192837, 5987253519122501, 5987317698330693, NULL, 5977706236739653, '2023-01-12 11:54:02.013551');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5987413717876805, 5987413533196357, 5987229831856197, NULL, 5977706236739653, '2023-01-12 12:18:05.894468');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5987416869175365, 5987413533196357, 5987317698330693, NULL, 5977706236739653, '2023-01-12 12:18:53.170126');
INSERT INTO "public"."co_company_team_member" ("id", "team_id", "employee_id", "invite_user_id", "union_main_id", "join_at") VALUES (5989544611676230, 5989544611676229, 5981495520723013, NULL, 5977706236739653, '2023-01-12 21:19:59.810376');
COMMIT;

-- ----------------------------
-- Table structure for sys_area
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_area";
CREATE TABLE "public"."sys_area" (
                                     "id" int8 NOT NULL GENERATED ALWAYS AS IDENTITY (
                                         INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1
),
                                     "area_code" int4,
                                     "area_name" varchar(64) COLLATE "pg_catalog"."default",
                                     "level" int2,
                                     "city_code" varchar(16) COLLATE "pg_catalog"."default",
                                     "lat_long_center" varchar(64) COLLATE "pg_catalog"."default",
                                     "parent_id" int8
)
;
ALTER TABLE "public"."sys_area" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_area"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_area"."area_code" IS '地区编码';
COMMENT ON COLUMN "public"."sys_area"."area_name" IS '地区名称';
COMMENT ON COLUMN "public"."sys_area"."level" IS '1:省份province,2:市city,3:区县district,4:街道street';
COMMENT ON COLUMN "public"."sys_area"."city_code" IS '城市编码';
COMMENT ON COLUMN "public"."sys_area"."lat_long_center" IS '城市中心点（即经纬度）';
COMMENT ON COLUMN "public"."sys_area"."parent_id" IS '地区父节点';

-- ----------------------------
-- Records of sys_area
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2, 110100, '北京城区', 2, '010', '116.407394,39.904211', 1);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3, 110101, '东城区', 1, '010', '116.41649,39.928341', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (4, 110102, '西城区', 1, '010', '116.365873,39.912235', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (19, 120000, '天津市', 4, '022', '117.200983,39.084158', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (20, 120100, '天津城区', 2, '022', '117.200983,39.084158', 19);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2726, 540000, '西藏自治区', 4, '[]', '91.117525,29.647535', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2205, 460000, '海南省', 4, '[]', '110.349228,20.017377', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2580, 530000, '云南省', 4, '[]', '102.710002,25.045806', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2079, 450000, '广西壮族自治区', 4, '[]', '108.327546,22.815478', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3228, 710000, '台湾省', 4, '1886', '121.509062,25.044332', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1143, 350000, '福建省', 4, '[]', '119.295143,26.100779', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (920, 330000, '浙江省', 4, '[]', '120.152585,30.266597', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1350, 370000, '山东省', 4, '[]', '117.019915,36.671156', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (649, 230000, '黑龙江省', 4, '[]', '126.661665,45.742366', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1681, 420000, '湖北省', 4, '[]', '114.341745,30.546557', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3251, 820000, '澳门特别行政区', 4, '1853', '113.543028,22.186835', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (38, 130100, '石家庄市', 2, '0311', '114.514793,38.042228', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (61, 130200, '唐山市', 2, '0315', '118.180193,39.630867', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (76, 130300, '秦皇岛市', 2, '0335', '119.518197,39.888701', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (84, 130400, '邯郸市', 2, '0310', '114.538959,36.625594', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (16, 110117, '平谷区', 1, '010', '117.121351,40.140595', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (17, 110118, '密云区', 1, '010', '116.843047,40.376894', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (18, 110119, '延庆区', 1, '010', '115.974981,40.456591', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (21, 120101, '和平区', 1, '022', '117.214699,39.117196', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (22, 120102, '河东区', 1, '022', '117.251584,39.128294', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (23, 120103, '河西区', 1, '022', '117.223371,39.109563', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (24, 120104, '南开区', 1, '022', '117.150738,39.138205', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (25, 120105, '河北区', 1, '022', '117.196648,39.147869', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (26, 120106, '红桥区', 1, '022', '117.151533,39.167345', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (28, 120111, '西青区', 1, '022', '117.008826,39.141152', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (29, 120112, '津南区', 1, '022', '117.35726,38.937928', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (30, 120113, '北辰区', 1, '022', '117.135488,39.224791', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (31, 120114, '武清区', 1, '022', '117.044387,39.384119', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (103, 130500, '邢台市', 2, '0319', '114.504677,37.070834', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (32, 120115, '宝坻区', 1, '022', '117.309874,39.717564', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (33, 120116, '滨海新区', 1, '022', '117.698407,39.01727', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (34, 120117, '宁河区', 1, '022', '117.826724,39.330087', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (35, 120118, '静海区', 1, '022', '116.974232,38.94745', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (36, 120119, '蓟州区', 1, '022', '117.408296,40.045851', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (39, 130102, '长安区', 1, '0311', '114.539395,38.036347', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (40, 130104, '桥西区', 1, '0311', '114.461088,38.004193', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (41, 130105, '新华区', 1, '0311', '114.463377,38.05095', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (42, 130107, '井陉矿区', 1, '0311', '114.062062,38.065151', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (43, 130108, '裕华区', 1, '0311', '114.531202,38.00643', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (44, 130109, '藁城区', 1, '0311', '114.847023,38.021453', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (45, 130110, '鹿泉区', 1, '0311', '114.313654,38.085953', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (46, 130111, '栾城区', 1, '0311', '114.648318,37.900199', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (47, 130121, '井陉县', 1, '0311', '114.145242,38.032118', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (48, 130123, '正定县', 1, '0311', '114.570941,38.146444', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (49, 130125, '行唐县', 1, '0311', '114.552714,38.438377', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (50, 130126, '灵寿县', 1, '0311', '114.382614,38.308665', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (51, 130127, '高邑县', 1, '0311', '114.611121,37.615534', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (52, 130128, '深泽县', 1, '0311', '115.20092,38.184033', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (123, 130600, '保定市', 2, '0312', '115.464589,38.874434', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (53, 130129, '赞皇县', 1, '0311', '114.386111,37.665663', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (54, 130130, '无极县', 1, '0311', '114.97634,38.179192', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (55, 130131, '平山县', 1, '0311', '114.195918,38.247888', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (56, 130132, '元氏县', 1, '0311', '114.525409,37.766513', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (57, 130133, '赵县', 1, '0311', '114.776297,37.756578', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (58, 130181, '辛集市', 1, '0311', '115.217658,37.943121', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (59, 130183, '晋州市', 1, '0311', '115.044213,38.033671', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (60, 130184, '新乐市', 1, '0311', '114.683776,38.343319', 38);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (62, 130202, '路南区', 1, '0315', '118.154354,39.625058', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (63, 130203, '路北区', 1, '0315', '118.200692,39.624437', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (64, 130204, '古冶区', 1, '0315', '118.447635,39.733578', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (65, 130205, '开平区', 1, '0315', '118.261841,39.671001', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (66, 130207, '丰南区', 1, '0315', '118.085169,39.576031', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (67, 130208, '丰润区', 1, '0315', '118.162215,39.832582', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (68, 130209, '曹妃甸区', 1, '0315', '118.460379,39.27307', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (69, 130223, '滦县', 1, '0315', '118.703598,39.740593', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (70, 130224, '滦南县', 1, '0315', '118.682379,39.518996', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (71, 130225, '乐亭县', 1, '0315', '118.912571,39.425608', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (72, 130227, '迁西县', 1, '0315', '118.314715,40.1415', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (73, 130229, '玉田县', 1, '0315', '117.738658,39.900401', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (74, 130281, '遵化市', 1, '0315', '117.965892,40.189201', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (75, 130283, '迁安市', 1, '0315', '118.701144,39.999174', 61);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (77, 130302, '海港区', 1, '0335', '119.564962,39.94756', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (78, 130303, '山海关区', 1, '0335', '119.775799,39.978848', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (148, 130700, '张家口市', 2, '0313', '114.886252,40.768493', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (79, 130304, '北戴河区', 1, '0335', '119.484522,39.834596', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (80, 130306, '抚宁区', 1, '0335', '119.244847,39.876253', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (81, 130321, '青龙满族自治县', 1, '0335', '118.949684,40.407578', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (82, 130322, '昌黎县', 1, '0335', '119.199551,39.700911', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (83, 130324, '卢龙县', 1, '0335', '118.892986,39.891946', 76);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (85, 130402, '邯山区', 1, '0310', '114.531002,36.594313', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (86, 130403, '丛台区', 1, '0310', '114.492896,36.636409', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (87, 130404, '复兴区', 1, '0310', '114.462061,36.639033', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (88, 130406, '峰峰矿区', 1, '0310', '114.212802,36.419739', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (89, 130423, '临漳县', 1, '0310', '114.619536,36.335025', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (90, 130424, '成安县', 1, '0310', '114.670032,36.444317', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (91, 130425, '大名县', 1, '0310', '115.147814,36.285616', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (92, 130426, '涉县', 1, '0310', '113.6914,36.584994', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (93, 130427, '磁县', 1, '0310', '114.373946,36.374011', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (94, 130407, '肥乡区', 1, '0310', '114.800166,36.548131', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (95, 130408, '永年区', 1, '0310', '114.543832,36.743966', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (165, 130800, '承德市', 2, '0314', '117.962749,40.952942', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (96, 130430, '邱县', 1, '0310', '115.200589,36.811148', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (97, 130431, '鸡泽县', 1, '0310', '114.889376,36.91034', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (98, 130432, '广平县', 1, '0310', '114.948606,36.483484', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (99, 130433, '馆陶县', 1, '0310', '115.282467,36.547556', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (100, 130434, '魏县', 1, '0310', '114.93892,36.359868', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (101, 130435, '曲周县', 1, '0310', '114.957504,36.76607', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (102, 130481, '武安市', 1, '0310', '114.203697,36.696506', 84);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (104, 130502, '桥东区', 1, '0319', '114.507058,37.071287', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (105, 130503, '桥西区', 1, '0319', '114.468601,37.059827', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (106, 130521, '邢台县', 1, '0319', '114.561132,37.05073', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (107, 130522, '临城县', 1, '0319', '114.498761,37.444498', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (177, 130900, '沧州市', 2, '0317', '116.838834,38.304477', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (108, 130523, '内丘县', 1, '0319', '114.512128,37.286669', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (179, 130903, '运河区', 1, '0317', '116.843673,38.283749', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (109, 130524, '柏乡县', 1, '0319', '114.693425,37.482422', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (110, 130525, '隆尧县', 1, '0319', '114.770419,37.350172', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (111, 130526, '任县', 1, '0319', '114.671936,37.120982', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (112, 130527, '南和县', 1, '0319', '114.683863,37.005017', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (113, 130528, '宁晋县', 1, '0319', '114.93992,37.624564', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (114, 130529, '巨鹿县', 1, '0319', '115.037477,37.221112', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (115, 130530, '新河县', 1, '0319', '115.250907,37.520862', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (116, 130531, '广宗县', 1, '0319', '115.142626,37.074661', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (117, 130532, '平乡县', 1, '0319', '115.030075,37.063148', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (118, 130533, '威县', 1, '0319', '115.266703,36.975478', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (119, 130534, '清河县', 1, '0319', '115.667208,37.039991', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (120, 130535, '临西县', 1, '0319', '115.501048,36.870811', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (121, 130581, '南宫市', 1, '0319', '115.408747,37.359264', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (194, 131000, '廊坊市', 2, '0316', '116.683752,39.538047', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (122, 130582, '沙河市', 1, '0319', '114.503339,36.854929', 103);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (124, 130602, '竞秀区', 1, '0312', '115.45877,38.877449', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (125, 130606, '莲池区', 1, '0312', '115.497097,38.883582', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (126, 130607, '满城区', 1, '0312', '115.322334,38.949119', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (127, 130608, '清苑区', 1, '0312', '115.489959,38.765148', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (128, 130609, '徐水区', 1, '0312', '115.655774,39.018736', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (129, 130623, '涞水县', 1, '0312', '115.713904,39.394316', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (130, 130624, '阜平县', 1, '0312', '114.195104,38.849152', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (131, 130626, '定兴县', 1, '0312', '115.808296,39.263145', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (132, 130627, '唐县', 1, '0312', '114.982972,38.748203', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (205, 131100, '衡水市', 2, '0318', '115.670177,37.73892', 37);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (133, 130628, '高阳县', 1, '0312', '115.778965,38.700088', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (134, 130629, '容城县', 1, '0312', '115.861657,39.042784', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (135, 130630, '涞源县', 1, '0312', '114.694283,39.360247', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (136, 130631, '望都县', 1, '0312', '115.155128,38.695842', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (137, 130632, '安新县', 1, '0312', '115.935603,38.935369', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (138, 130633, '易县', 1, '0312', '115.497457,39.349393', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (139, 130634, '曲阳县', 1, '0312', '114.745008,38.622248', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (140, 130635, '蠡县', 1, '0312', '115.583854,38.488055', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (141, 130636, '顺平县', 1, '0312', '115.13547,38.837487', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (142, 130637, '博野县', 1, '0312', '115.46438,38.457364', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (143, 130638, '雄县', 1, '0312', '116.10865,38.99455', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (144, 130681, '涿州市', 1, '0312', '115.974422,39.485282', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (218, 140100, '太原市', 2, '0351', '112.548879,37.87059', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (145, 130682, '定州市', 1, '0312', '114.990392,38.516461', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (146, 130683, '安国市', 1, '0312', '115.326646,38.418439', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (147, 130684, '高碑店市', 1, '0312', '115.873886,39.326839', 123);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (149, 130702, '桥东区', 1, '0313', '114.894189,40.788434', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (150, 130703, '桥西区', 1, '0313', '114.869657,40.819581', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (151, 130705, '宣化区', 1, '0313', '115.099494,40.608793', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (152, 130706, '下花园区', 1, '0313', '115.287352,40.502652', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (153, 130708, '万全区', 1, '0313', '114.740557,40.766965', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (154, 130709, '崇礼区', 1, '0313', '115.282668,40.974675', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (155, 130722, '张北县', 1, '0313', '114.720077,41.158596', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (229, 140200, '大同市', 2, '0352', '113.300129,40.076763', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (156, 130723, '康保县', 1, '0313', '114.600404,41.852368', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (157, 130724, '沽源县', 1, '0313', '115.688692,41.669668', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (158, 130725, '尚义县', 1, '0313', '113.969618,41.076226', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (159, 130726, '蔚县', 1, '0313', '114.588903,39.840842', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (160, 130727, '阳原县', 1, '0313', '114.150348,40.104663', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (161, 130728, '怀安县', 1, '0313', '114.385791,40.674193', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (162, 130730, '怀来县', 1, '0313', '115.517861,40.415343', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (163, 130731, '涿鹿县', 1, '0313', '115.205345,40.379562', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (164, 130732, '赤城县', 1, '0313', '115.831498,40.912921', 148);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (166, 130802, '双桥区', 1, '0314', '117.943466,40.974643', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (167, 130803, '双滦区', 1, '0314', '117.799888,40.959236', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (241, 140300, '阳泉市', 2, '0353', '113.580519,37.856971', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (168, 130804, '鹰手营子矿区', 1, '0314', '117.659499,40.546361', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (169, 130821, '承德县', 1, '0314', '118.173824,40.768238', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (170, 130822, '兴隆县', 1, '0314', '117.500558,40.417358', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (171, 130881, '平泉市', 1, '0314', '118.701951,41.018405', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (172, 130824, '滦平县', 1, '0314', '117.332801,40.941482', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (247, 140400, '长治市', 2, '0355', '113.116404,36.195409', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (173, 130825, '隆化县', 1, '0314', '117.738937,41.313791', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (174, 130826, '丰宁满族自治县', 1, '0314', '116.646051,41.209069', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (175, 130827, '宽城满族自治县', 1, '0314', '118.485313,40.611391', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (176, 130828, '围场满族蒙古族自治县', 1, '0314', '117.760159,41.938529', 165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (178, 130902, '新华区', 1, '0317', '116.866284,38.314416', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (180, 130921, '沧县', 1, '0317', '117.007478,38.219856', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (181, 130922, '青县', 1, '0317', '116.804305,38.583021', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (182, 130923, '东光县', 1, '0317', '116.537067,37.888248', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (183, 130924, '海兴县', 1, '0317', '117.497651,38.143169', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (184, 130925, '盐山县', 1, '0317', '117.230602,38.058087', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (185, 130926, '肃宁县', 1, '0317', '115.829758,38.422801', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (186, 130927, '南皮县', 1, '0317', '116.708347,38.038421', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (187, 130928, '吴桥县', 1, '0317', '116.391508,37.627661', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (261, 140500, '晋城市', 2, '0356', '112.851486,35.490684', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (189, 130930, '孟村回族自治县', 1, '0317', '117.104298,38.053409', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (190, 130981, '泊头市', 1, '0317', '116.578367,38.083437', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (191, 130982, '任丘市', 1, '0317', '116.082917,38.683591', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (192, 130983, '黄骅市', 1, '0317', '117.329949,38.371402', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (268, 140600, '朔州市', 2, '0349', '112.432991,39.331855', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (272, 140622, '应县', 1, '0349', '113.191098,39.554247', 268);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (193, 130984, '河间市', 1, '0317', '116.099517,38.446624', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (195, 131002, '安次区', 1, '0316', '116.694544,39.502569', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (196, 131003, '广阳区', 1, '0316', '116.71069,39.522786', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (197, 131022, '固安县', 1, '0316', '116.298657,39.438214', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (198, 131023, '永清县', 1, '0316', '116.50568,39.330689', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (275, 140700, '晋中市', 2, '0354', '112.752652,37.687357', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (199, 131024, '香河县', 1, '0316', '117.006093,39.761424', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (200, 131025, '大城县', 1, '0316', '116.653793,38.705449', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (201, 131026, '文安县', 1, '0316', '116.457898,38.87292', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (202, 131028, '大厂回族自治县', 1, '0316', '116.989574,39.886547', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (203, 131081, '霸州市', 1, '0316', '116.391484,39.125744', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (204, 131082, '三河市', 1, '0316', '117.078294,39.982718', 194);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (206, 131102, '桃城区', 1, '0318', '115.67545,37.735465', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (207, 131103, '冀州区', 1, '0318', '115.579308,37.550856', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (208, 131121, '枣强县', 1, '0318', '115.724259,37.513417', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (209, 131122, '武邑县', 1, '0318', '115.887531,37.801665', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (210, 131123, '武强县', 1, '0318', '115.982461,38.041368', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (287, 140800, '运城市', 2, '0359', '111.00746,35.026516', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (211, 131124, '饶阳县', 1, '0318', '115.725833,38.235892', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (212, 131125, '安平县', 1, '0318', '115.519278,38.234501', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (213, 131126, '故城县', 1, '0318', '115.965874,37.347409', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (214, 131127, '景县', 1, '0318', '116.270648,37.69229', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (215, 131128, '阜城县', 1, '0318', '116.175262,37.862505', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (216, 131182, '深州市', 1, '0318', '115.559574,38.001535', 205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (219, 140105, '小店区', 1, '0351', '112.565659,37.736525', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (220, 140106, '迎泽区', 1, '0351', '112.5634,37.863451', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (221, 140107, '杏花岭区', 1, '0351', '112.570604,37.893955', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (222, 140108, '尖草坪区', 1, '0351', '112.486691,37.940387', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (223, 140109, '万柏林区', 1, '0351', '112.515937,37.85958', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (224, 140110, '晋源区', 1, '0351', '112.47794,37.715193', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (225, 140121, '清徐县', 1, '0351', '112.358667,37.607443', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (301, 140900, '忻州市', 2, '0350', '112.734174,38.416663', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (226, 140122, '阳曲县', 1, '0351', '112.672952,38.058488', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (227, 140123, '娄烦县', 1, '0351', '111.797083,38.067932', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (228, 140181, '古交市', 1, '0351', '112.175853,37.907129', 218);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (230, 140202, '城区', 1, '0352', '113.298026,40.075666', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (231, 140203, '矿区', 1, '0352', '113.177206,40.036858', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (232, 140211, '南郊区', 1, '0352', '113.149693,40.005404', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (233, 140212, '新荣区', 1, '0352', '113.140004,40.255866', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (234, 140221, '阳高县', 1, '0352', '113.748944,40.361059', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (235, 140222, '天镇县', 1, '0352', '114.090867,40.420237', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (236, 140223, '广灵县', 1, '0352', '114.282758,39.760281', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (237, 140224, '灵丘县', 1, '0352', '114.23435,39.442406', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (238, 140225, '浑源县', 1, '0352', '113.699475,39.693406', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (239, 140226, '左云县', 1, '0352', '112.703008,40.013442', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (240, 140227, '大同县', 1, '0352', '113.61244,40.040294', 229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (316, 141000, '临汾市', 2, '0357', '111.518975,36.088005', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (242, 140302, '城区', 1, '0353', '113.600669,37.847436', 241);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (243, 140303, '矿区', 1, '0353', '113.555279,37.868494', 241);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (244, 140311, '郊区', 1, '0353', '113.594163,37.944679', 241);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (245, 140321, '平定县', 1, '0353', '113.630107,37.804988', 241);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (246, 140322, '盂县', 1, '0353', '113.41233,38.085619', 241);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (248, 140402, '城区', 1, '0355', '113.123088,36.20353', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (249, 140411, '郊区', 1, '0355', '113.101211,36.218388', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (250, 140421, '长治县', 1, '0355', '113.051407,36.052858', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (251, 140423, '襄垣县', 1, '0355', '113.051491,36.535817', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (252, 140424, '屯留县', 1, '0355', '112.891998,36.315663', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (253, 140425, '平顺县', 1, '0355', '113.435961,36.200179', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (254, 140426, '黎城县', 1, '0355', '113.387155,36.502328', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (255, 140427, '壶关县', 1, '0355', '113.207049,36.115448', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (256, 140428, '长子县', 1, '0355', '112.8779,36.122334', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (257, 140429, '武乡县', 1, '0355', '112.864561,36.837625', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (258, 140430, '沁县', 1, '0355', '112.699226,36.756063', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (259, 140431, '沁源县', 1, '0355', '112.337446,36.5002', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (334, 141100, '吕梁市', 2, '0358', '111.144699,37.519126', 217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (260, 140481, '潞城市', 1, '0355', '113.228852,36.334104', 247);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (262, 140502, '城区', 1, '0356', '112.853555,35.501571', 261);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (263, 140521, '沁水县', 1, '0356', '112.186738,35.690141', 261);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (264, 140522, '阳城县', 1, '0356', '112.414738,35.486029', 261);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (265, 140524, '陵川县', 1, '0356', '113.280688,35.775685', 261);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (266, 140525, '泽州县', 1, '0356', '112.899137,35.617221', 261);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (267, 140581, '高平市', 1, '0356', '112.92392,35.797997', 261);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (269, 140602, '朔城区', 1, '0349', '112.432312,39.319519', 268);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (270, 140603, '平鲁区', 1, '0349', '112.28833,39.512155', 268);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (271, 140621, '山阴县', 1, '0349', '112.816413,39.527893', 268);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (273, 140623, '右玉县', 1, '0349', '112.466989,39.989063', 268);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (274, 140624, '怀仁县', 1, '0349', '113.131717,39.821627', 268);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (276, 140702, '榆次区', 1, '0354', '112.708224,37.697794', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (277, 140721, '榆社县', 1, '0354', '112.975209,37.070916', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (349, 150100, '呼和浩特市', 2, '0471', '111.749995,40.842356', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (278, 140722, '左权县', 1, '0354', '113.379403,37.082943', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (279, 140723, '和顺县', 1, '0354', '113.570415,37.32957', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (280, 140724, '昔阳县', 1, '0354', '113.706977,37.61253', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (281, 140725, '寿阳县', 1, '0354', '113.176373,37.895191', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (282, 140726, '太谷县', 1, '0354', '112.551305,37.421307', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (283, 140727, '祁县', 1, '0354', '112.335542,37.357869', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (284, 140728, '平遥县', 1, '0354', '112.176136,37.189421', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (285, 140729, '灵石县', 1, '0354', '111.77864,36.847927', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (359, 150200, '包头市', 2, '0472', '109.953504,40.621157', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (286, 140781, '介休市', 1, '0354', '111.916711,37.026944', 275);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (288, 140802, '盐湖区', 1, '0359', '110.998272,35.015101', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (289, 140821, '临猗县', 1, '0359', '110.774547,35.144277', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (290, 140822, '万荣县', 1, '0359', '110.838024,35.415253', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (291, 140823, '闻喜县', 1, '0359', '111.22472,35.356644', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (292, 140824, '稷山县', 1, '0359', '110.983333,35.604025', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (293, 140825, '新绛县', 1, '0359', '111.224734,35.616251', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (294, 140826, '绛县', 1, '0359', '111.568236,35.49119', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (295, 140827, '垣曲县', 1, '0359', '111.670108,35.297369', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (369, 150300, '乌海市', 2, '0473', '106.794216,39.655248', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (296, 140828, '夏县', 1, '0359', '111.220456,35.141363', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (297, 140829, '平陆县', 1, '0359', '111.194133,34.82926', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (298, 140830, '芮城县', 1, '0359', '110.694369,34.693579', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (373, 150400, '赤峰市', 2, '0476', '118.88694,42.257843', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (299, 140881, '永济市', 1, '0359', '110.447543,34.8671', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (300, 140882, '河津市', 1, '0359', '110.712063,35.596383', 287);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (302, 140902, '忻府区', 1, '0350', '112.746046,38.404242', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (303, 140921, '定襄县', 1, '0350', '112.957237,38.473506', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (304, 140922, '五台县', 1, '0350', '113.255309,38.728315', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (305, 140923, '代县', 1, '0350', '112.960282,39.066917', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (306, 140924, '繁峙县', 1, '0350', '113.265563,39.188811', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (307, 140925, '宁武县', 1, '0350', '112.304722,39.001524', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (308, 140926, '静乐县', 1, '0350', '111.939498,38.359306', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (309, 140927, '神池县', 1, '0350', '112.211296,39.090552', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (310, 140928, '五寨县', 1, '0350', '111.846904,38.910726', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (311, 140929, '岢岚县', 1, '0350', '111.57285,38.70418', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (386, 150500, '通辽市', 2, '0475', '122.243444,43.652889', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (312, 140930, '河曲县', 1, '0350', '111.138472,39.384482', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (313, 140931, '保德县', 1, '0350', '111.086564,39.022487', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (314, 140932, '偏关县', 1, '0350', '111.508831,39.436306', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (315, 140981, '原平市', 1, '0350', '112.711058,38.731402', 301);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (317, 141002, '尧都区', 1, '0357', '111.579554,36.07884', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (318, 141021, '曲沃县', 1, '0357', '111.47586,35.641086', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (319, 141022, '翼城县', 1, '0357', '111.718951,35.738576', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (320, 141023, '襄汾县', 1, '0357', '111.441725,35.876293', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (395, 150600, '鄂尔多斯市', 2, '0477', '109.781327,39.608266', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (321, 141024, '洪洞县', 1, '0357', '111.674965,36.253747', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (322, 141025, '古县', 1, '0357', '111.920465,36.266914', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (323, 141026, '安泽县', 1, '0357', '112.250144,36.147787', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (324, 141027, '浮山县', 1, '0357', '111.848883,35.968124', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (325, 141028, '吉县', 1, '0357', '110.681763,36.098188', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (326, 141029, '乡宁县', 1, '0357', '110.847021,35.970389', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (327, 141030, '大宁县', 1, '0357', '110.75291,36.465102', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (328, 141031, '隰县', 1, '0357', '110.940637,36.69333', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (329, 141032, '永和县', 1, '0357', '110.632006,36.759507', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (405, 150700, '呼伦贝尔市', 2, '0470', '119.765558,49.211576', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (330, 141033, '蒲县', 1, '0357', '111.096439,36.411826', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (331, 141034, '汾西县', 1, '0357', '111.56395,36.652854', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (332, 141081, '侯马市', 1, '0357', '111.372002,35.619105', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (333, 141082, '霍州市', 1, '0357', '111.755398,36.56893', 316);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (335, 141102, '离石区', 1, '0358', '111.150695,37.51786', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (336, 141121, '文水县', 1, '0358', '112.028866,37.438101', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (337, 141122, '交城县', 1, '0358', '112.156064,37.551963', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (338, 141123, '兴县', 1, '0358', '111.127667,38.462389', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (339, 141124, '临县', 1, '0358', '110.992093,37.950758', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (340, 141125, '柳林县', 1, '0358', '110.889007,37.429772', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (341, 141126, '石楼县', 1, '0358', '110.834634,36.99857', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (342, 141127, '岚县', 1, '0358', '111.671917,38.279299', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (343, 141128, '方山县', 1, '0358', '111.244098,37.894631', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (344, 141129, '中阳县', 1, '0358', '111.179657,37.357058', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (420, 150800, '巴彦淖尔市', 2, '0478', '107.387657,40.743213', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (345, 141130, '交口县', 1, '0358', '111.181151,36.982186', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (346, 141181, '孝义市', 1, '0358', '111.778818,37.146294', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (347, 141182, '汾阳市', 1, '0358', '111.770477,37.261756', 334);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (350, 150102, '新城区', 1, '0471', '111.665544,40.858289', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (351, 150103, '回民区', 1, '0471', '111.623692,40.808608', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (352, 150104, '玉泉区', 1, '0471', '111.673881,40.753655', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (353, 150105, '赛罕区', 1, '0471', '111.701355,40.792667', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (428, 150900, '乌兰察布市', 2, '0474', '113.132584,40.994785', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (354, 150121, '土默特左旗', 1, '0471', '111.163902,40.729572', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (355, 150122, '托克托县', 1, '0471', '111.194312,40.277431', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (356, 150123, '和林格尔县', 1, '0471', '111.821843,40.378787', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (357, 150124, '清水河县', 1, '0471', '111.647609,39.921095', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (358, 150125, '武川县', 1, '0471', '111.451303,41.096471', 349);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (360, 150202, '东河区', 1, '0472', '110.044106,40.576319', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (361, 150203, '昆都仑区', 1, '0472', '109.837707,40.642578', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (362, 150204, '青山区', 1, '0472', '109.901572,40.643246', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (363, 150205, '石拐区', 1, '0472', '110.060254,40.681748', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (364, 150206, '白云鄂博矿区', 1, '0472', '109.973803,41.769511', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (365, 150207, '九原区', 1, '0472', '109.967449,40.610561', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (440, 152200, '兴安盟', 2, '0482', '122.037657,46.082462', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (366, 150221, '土默特右旗', 1, '0472', '110.524262,40.569426', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (367, 150222, '固阳县', 1, '0472', '110.060514,41.034105', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (368, 150223, '达尔罕茂明安联合旗', 1, '0472', '110.432626,41.698992', 359);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (370, 150302, '海勃湾区', 1, '0473', '106.822778,39.691156', 369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (371, 150303, '海南区', 1, '0473', '106.891424,39.441364', 369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (372, 150304, '乌达区', 1, '0473', '106.726099,39.505925', 369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (374, 150402, '红山区', 1, '0476', '118.953854,42.296588', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (447, 152500, '锡林郭勒盟', 2, '0479', '116.048222,43.933454', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (446, 152224, '突泉县', 1, '0482', '121.593799,45.38193', 440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (375, 150403, '元宝山区', 1, '0476', '119.288611,42.038902', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (376, 150404, '松山区', 1, '0476', '118.916208,42.299798', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (377, 150421, '阿鲁科尔沁旗', 1, '0476', '120.0657,43.872298', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (378, 150422, '巴林左旗', 1, '0476', '119.362931,43.960889', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (379, 150423, '巴林右旗', 1, '0476', '118.66518,43.534414', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (380, 150424, '林西县', 1, '0476', '118.05545,43.61812', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (381, 150425, '克什克腾旗', 1, '0476', '117.545797,43.264988', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (382, 150426, '翁牛特旗', 1, '0476', '119.00658,42.936188', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (383, 150428, '喀喇沁旗', 1, '0476', '118.701937,41.927363', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (384, 150429, '宁城县', 1, '0476', '119.318876,41.601375', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (385, 150430, '敖汉旗', 1, '0476', '119.921603,42.290781', 373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (460, 152900, '阿拉善盟', 2, '0483', '105.728957,38.851921', 348);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (387, 150502, '科尔沁区', 1, '0475', '122.255671,43.623078', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (388, 150521, '科尔沁左翼中旗', 1, '0475', '123.312264,44.126625', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (389, 150522, '科尔沁左翼后旗', 1, '0475', '122.35677,42.935105', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (390, 150523, '开鲁县', 1, '0475', '121.319308,43.601244', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (465, 210100, '沈阳市', 2, '024', '123.465035,41.677284', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (391, 150524, '库伦旗', 1, '0475', '121.8107,42.735656', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (392, 150525, '奈曼旗', 1, '0475', '120.658282,42.867226', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (393, 150526, '扎鲁特旗', 1, '0475', '120.911676,44.556389', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (394, 150581, '霍林郭勒市', 1, '0475', '119.68187,45.533962', 386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (396, 150602, '东胜区', 1, '0477', '109.963333,39.822593', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (397, 150603, '康巴什区', 1, '0477', '109.790076,39.607472', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (398, 150621, '达拉特旗', 1, '0477', '110.033833,40.412438', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (399, 150622, '准格尔旗', 1, '0477', '111.240171,39.864361', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (400, 150623, '鄂托克前旗', 1, '0477', '107.477514,38.182362', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (401, 150624, '鄂托克旗', 1, '0477', '107.97616,39.08965', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (402, 150625, '杭锦旗', 1, '0477', '108.736208,39.833309', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (403, 150626, '乌审旗', 1, '0477', '108.817607,38.604136', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (404, 150627, '伊金霍洛旗', 1, '0477', '109.74774,39.564659', 395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (479, 210200, '大连市', 2, '0411', '121.614848,38.914086', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (406, 150702, '海拉尔区', 1, '0470', '119.736176,49.212188', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (407, 150703, '扎赉诺尔区', 1, '0470', '117.670248,49.510375', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (408, 150721, '阿荣旗', 1, '0470', '123.459049,48.126584', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (409, 150722, '莫力达瓦达斡尔族自治旗', 1, '0470', '124.519023,48.477728', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (410, 150723, '鄂伦春自治旗', 1, '0470', '123.726201,50.591842', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (411, 150724, '鄂温克族自治旗', 1, '0470', '119.755239,49.146592', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (412, 150725, '陈巴尔虎旗', 1, '0470', '119.424026,49.328916', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (413, 150726, '新巴尔虎左旗', 1, '0470', '118.269819,48.218241', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (414, 150727, '新巴尔虎右旗', 1, '0470', '116.82369,48.672101', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (415, 150781, '满洲里市', 1, '0470', '117.378529,49.597841', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (490, 210300, '鞍山市', 2, '0412', '122.994329,41.108647', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (416, 150782, '牙克石市', 1, '0470', '120.711775,49.285629', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (417, 150783, '扎兰屯市', 1, '0470', '122.737467,48.013733', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (418, 150784, '额尔古纳市', 1, '0470', '120.180506,50.243102', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (419, 150785, '根河市', 1, '0470', '121.520388,50.780344', 405);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (421, 150802, '临河区', 1, '0478', '107.363918,40.751187', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (422, 150821, '五原县', 1, '0478', '108.267561,41.088421', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (423, 150822, '磴口县', 1, '0478', '107.008248,40.330523', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (498, 210400, '抚顺市', 2, '0413', '123.957208,41.880872', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (424, 150823, '乌拉特前旗', 1, '0478', '108.652114,40.737018', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (425, 150824, '乌拉特中旗', 1, '0478', '108.513645,41.587732', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (426, 150825, '乌拉特后旗', 1, '0478', '107.074621,41.084282', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (427, 150826, '杭锦后旗', 1, '0478', '107.151245,40.88602', 420);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (429, 150902, '集宁区', 1, '0474', '113.116453,41.034134', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (430, 150921, '卓资县', 1, '0474', '112.577528,40.894691', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (431, 150922, '化德县', 1, '0474', '114.010437,41.90456', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (506, 210500, '本溪市', 2, '0414', '123.685142,41.486981', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (432, 150923, '商都县', 1, '0474', '113.577816,41.562113', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (433, 150924, '兴和县', 1, '0474', '113.834173,40.872301', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (434, 150925, '凉城县', 1, '0474', '112.503971,40.531555', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (435, 150926, '察哈尔右翼前旗', 1, '0474', '113.214733,40.785631', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (436, 150927, '察哈尔右翼中旗', 1, '0474', '112.635577,41.277462', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (437, 150928, '察哈尔右翼后旗', 1, '0474', '113.191035,41.436069', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (513, 210600, '丹东市', 2, '0415', '124.35445,40.000787', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (438, 150929, '四子王旗', 1, '0474', '111.706617,41.533462', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (439, 150981, '丰镇市', 1, '0474', '113.109892,40.436983', 428);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (441, 152201, '乌兰浩特市', 1, '0482', '122.093123,46.072731', 440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (442, 152202, '阿尔山市', 1, '0482', '119.943575,47.17744', 440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (443, 152221, '科尔沁右翼前旗', 1, '0482', '121.952621,46.079833', 440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (444, 152222, '科尔沁右翼中旗', 1, '0482', '121.47653,45.060837', 440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (520, 210700, '锦州市', 2, '0416', '121.126846,41.095685', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (445, 152223, '扎赉特旗', 1, '0482', '122.899656,46.723237', 440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (448, 152501, '二连浩特市', 1, '0479', '111.951002,43.6437', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (449, 152502, '锡林浩特市', 1, '0479', '116.086029,43.933403', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (450, 152522, '阿巴嘎旗', 1, '0479', '114.950248,44.022995', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (451, 152523, '苏尼特左旗', 1, '0479', '113.667248,43.85988', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (452, 152524, '苏尼特右旗', 1, '0479', '112.641783,42.742892', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (453, 152525, '东乌珠穆沁旗', 1, '0479', '116.974494,45.498221', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (528, 210800, '营口市', 2, '0417', '122.219458,40.625364', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (454, 152526, '西乌珠穆沁旗', 1, '0479', '117.608911,44.587882', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (455, 152527, '太仆寺旗', 1, '0479', '115.282986,41.877135', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (456, 152528, '镶黄旗', 1, '0479', '113.847287,42.232371', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (457, 152529, '正镶白旗', 1, '0479', '115.029848,42.28747', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (535, 210900, '阜新市', 2, '0418', '121.670273,42.021602', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (534, 210882, '大石桥市', 1, '0417', '122.509006,40.644482', 528);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (458, 152530, '正蓝旗', 1, '0479', '115.99247,42.241638', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (459, 152531, '多伦县', 1, '0479', '116.485555,42.203591', 447);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (461, 152921, '阿拉善左旗', 1, '0483', '105.666275,38.833389', 460);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (462, 152922, '阿拉善右旗', 1, '0483', '101.666917,39.216185', 460);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (463, 152923, '额济纳旗', 1, '0483', '101.055731,41.95455', 460);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (466, 210102, '和平区', 1, '024', '123.420368,41.789833', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (543, 211000, '辽阳市', 2, '0419', '123.236974,41.267794', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (467, 210103, '沈河区', 1, '024', '123.458691,41.796177', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (468, 210104, '大东区', 1, '024', '123.469948,41.805137', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (469, 210105, '皇姑区', 1, '024', '123.442378,41.824516', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (470, 210106, '铁西区', 1, '024', '123.333968,41.820807', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (471, 210111, '苏家屯区', 1, '024', '123.344062,41.664757', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (472, 210112, '浑南区', 1, '024', '123.449714,41.714914', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (473, 210113, '沈北新区', 1, '024', '123.583196,41.912487', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (551, 211100, '盘锦市', 2, '0427', '122.170584,40.719847', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (474, 210114, '于洪区', 1, '024', '123.308119,41.793721', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (475, 210115, '辽中区', 1, '024', '122.765409,41.516826', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (476, 210123, '康平县', 1, '024', '123.343699,42.72793', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (477, 210124, '法库县', 1, '024', '123.440294,42.50108', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (556, 211200, '铁岭市', 2, '0410', '123.726035,42.223828', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (478, 210181, '新民市', 1, '024', '122.836723,41.985186', 465);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (480, 210202, '中山区', 1, '0411', '121.644926,38.918574', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (481, 210203, '西岗区', 1, '0411', '121.612324,38.914687', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (482, 210204, '沙河口区', 1, '0411', '121.594297,38.904788', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (483, 210211, '甘井子区', 1, '0411', '121.525466,38.953343', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (484, 210212, '旅顺口区', 1, '0411', '121.261953,38.851705', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (485, 210213, '金州区', 1, '0411', '121.782655,39.050001', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (564, 211300, '朝阳市', 2, '0421', '120.450879,41.573762', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (486, 210214, '普兰店区', 1, '0411', '121.938269,39.392095', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (487, 210224, '长海县', 1, '0411', '122.588494,39.272728', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (488, 210281, '瓦房店市', 1, '0411', '121.979543,39.626897', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (489, 210283, '庄河市', 1, '0411', '122.967424,39.680843', 479);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (491, 210302, '铁东区', 1, '0412', '122.991052,41.089933', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (492, 210303, '铁西区', 1, '0412', '122.969629,41.119884', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (493, 210304, '立山区', 1, '0412', '123.029091,41.150401', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (572, 211400, '葫芦岛市', 2, '0429', '120.836939,40.71104', 464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (494, 210311, '千山区', 1, '0412', '122.944751,41.068901', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (495, 210321, '台安县', 1, '0412', '122.436196,41.412767', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (496, 210323, '岫岩满族自治县', 1, '0412', '123.280935,40.29088', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (497, 210381, '海城市', 1, '0412', '122.685217,40.882377', 490);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (499, 210402, '新抚区', 1, '0413', '123.912872,41.862026', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (500, 210403, '东洲区', 1, '0413', '124.038685,41.853191', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (501, 210404, '望花区', 1, '0413', '123.784225,41.853641', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (580, 220100, '长春市', 2, '0431', '125.323513,43.817251', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (502, 210411, '顺城区', 1, '0413', '123.945075,41.883235', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (503, 210421, '抚顺县', 1, '0413', '124.097978,41.922644', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (504, 210422, '新宾满族自治县', 1, '0413', '125.039978,41.734256', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (505, 210423, '清原满族自治县', 1, '0413', '124.924083,42.100538', 498);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (507, 210502, '平山区', 1, '0414', '123.769088,41.299587', 506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (508, 210503, '溪湖区', 1, '0414', '123.767646,41.329219', 506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (509, 210504, '明山区', 1, '0414', '123.817214,41.308719', 506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (510, 210505, '南芬区', 1, '0414', '123.744802,41.100445', 506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (511, 210521, '本溪满族自治县', 1, '0414', '124.120635,41.302009', 506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (512, 210522, '桓仁满族自治县', 1, '0414', '125.361007,41.267127', 506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (591, 220200, '吉林市', 2, '0432', '126.549572,43.837883', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (514, 210602, '元宝区', 1, '0415', '124.395661,40.136434', 513);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (515, 210603, '振兴区', 1, '0415', '124.383237,40.129944', 513);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (516, 210604, '振安区', 1, '0415', '124.470034,40.201553', 513);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (517, 210624, '宽甸满族自治县', 1, '0415', '124.783659,40.731316', 513);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (518, 210681, '东港市', 1, '0415', '124.152705,39.863008', 513);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (519, 210682, '凤城市', 1, '0415', '124.066919,40.452297', 513);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (521, 210702, '古塔区', 1, '0416', '121.128279,41.117245', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (522, 210703, '凌河区', 1, '0416', '121.150877,41.114989', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (523, 210711, '太和区', 1, '0416', '121.103892,41.109147', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (601, 220300, '四平市', 2, '0434', '124.350398,43.166419', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (524, 210726, '黑山县', 1, '0416', '122.126292,41.653593', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (526, 210781, '凌海市', 1, '0416', '121.35549,41.160567', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (527, 210782, '北镇市', 1, '0416', '121.777395,41.58844', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (529, 210802, '站前区', 1, '0417', '122.259033,40.672563', 528);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (530, 210803, '西市区', 1, '0417', '122.206419,40.666213', 528);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (531, 210804, '鲅鱼圈区', 1, '0417', '122.121521,40.226661', 528);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (608, 220400, '辽源市', 2, '0437', '125.14366,42.887766', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (532, 210811, '老边区', 1, '0417', '122.380087,40.680191', 528);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (533, 210881, '盖州市', 1, '0417', '122.349012,40.40074', 528);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (536, 210902, '海州区', 1, '0418', '121.657638,42.011162', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (537, 210903, '新邱区', 1, '0418', '121.792535,42.087632', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (613, 220500, '通化市', 2, '0435', '125.939697,41.728401', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (538, 210904, '太平区', 1, '0418', '121.678604,42.010669', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (539, 210905, '清河门区', 1, '0418', '121.416105,41.7831', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (540, 210911, '细河区', 1, '0418', '121.68054,42.025494', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (541, 210921, '阜新蒙古族自治县', 1, '0418', '121.757901,42.065175', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (542, 210922, '彰武县', 1, '0418', '122.538793,42.386543', 535);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (544, 211002, '白塔区', 1, '0419', '123.174325,41.270347', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (545, 211003, '文圣区', 1, '0419', '123.231408,41.283754', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (621, 220600, '白山市', 2, '0439', '126.41473,41.943972', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (622, 220602, '浑江区', 1, '0439', '126.416093,41.945409', 621);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (546, 211004, '宏伟区', 1, '0419', '123.196672,41.217649', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (547, 211005, '弓长岭区', 1, '0419', '123.419803,41.151847', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (548, 211011, '太子河区', 1, '0419', '123.18144,41.295023', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (549, 211021, '辽阳县', 1, '0419', '123.105694,41.205329', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (628, 220700, '松原市', 2, '0438', '124.825042,45.141548', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (550, 211081, '灯塔市', 1, '0419', '123.339312,41.426372', 543);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (552, 211102, '双台子区', 1, '0427', '122.039787,41.19965', 551);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (553, 211103, '兴隆台区', 1, '0427', '122.070769,41.119898', 551);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (554, 211104, '大洼区', 1, '0427', '122.082574,41.002279', 551);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (555, 211122, '盘山县', 1, '0427', '121.996411,41.242639', 551);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (634, 220800, '白城市', 2, '0436', '122.838714,45.619884', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (557, 211202, '银州区', 1, '0410', '123.842305,42.286129', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (558, 211204, '清河区', 1, '0410', '124.159191,42.546565', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (559, 211221, '铁岭县', 1, '0410', '123.728933,42.223395', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (560, 211223, '西丰县', 1, '0410', '124.727392,42.73803', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (561, 211224, '昌图县', 1, '0410', '124.111099,42.785791', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (640, 222400, '延边朝鲜族自治州', 2, '1433', '129.471868,42.909408', 579);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (562, 211281, '调兵山市', 1, '0410', '123.567117,42.467521', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (563, 211282, '开原市', 1, '0410', '124.038268,42.546307', 556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (565, 211302, '双塔区', 1, '0421', '120.453744,41.565627', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (566, 211303, '龙城区', 1, '0421', '120.413376,41.576749', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (567, 211321, '朝阳县', 1, '0421', '120.389754,41.497825', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (568, 211322, '建平县', 1, '0421', '119.64328,41.403128', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (569, 211324, '喀喇沁左翼蒙古族自治县', 1, '0421', '119.741223,41.12815', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (570, 211381, '北票市', 1, '0421', '120.77073,41.800683', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (571, 211382, '凌源市', 1, '0421', '119.401574,41.245445', 564);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (650, 230100, '哈尔滨市', 2, '0451', '126.534967,45.803775', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (573, 211402, '连山区', 1, '0429', '120.869231,40.774461', 572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (574, 211403, '龙港区', 1, '0429', '120.893786,40.735519', 572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (575, 211404, '南票区', 1, '0429', '120.749727,41.107107', 572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (576, 211421, '绥中县', 1, '0429', '120.344311,40.32558', 572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (577, 211422, '建昌县', 1, '0429', '119.837124,40.824367', 572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (578, 211481, '兴城市', 1, '0429', '120.756479,40.609731', 572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (581, 220102, '南关区', 1, '0431', '125.350173,43.863989', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (582, 220103, '宽城区', 1, '0431', '125.326581,43.943612', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (583, 220104, '朝阳区', 1, '0431', '125.288254,43.833762', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (584, 220105, '二道区', 1, '0431', '125.374327,43.865577', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (585, 220106, '绿园区', 1, '0431', '125.256135,43.880975', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (586, 220112, '双阳区', 1, '0431', '125.664662,43.525311', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (587, 220113, '九台区', 1, '0431', '125.839573,44.151742', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (588, 220122, '农安县', 1, '0431', '125.184887,44.432763', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (589, 220182, '榆树市', 1, '0431', '126.533187,44.840318', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (590, 220183, '德惠市', 1, '0431', '125.728755,44.522056', 580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (592, 220202, '昌邑区', 1, '0432', '126.574709,43.881818', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (593, 220203, '龙潭区', 1, '0432', '126.562197,43.910802', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (669, 230200, '齐齐哈尔市', 2, '0452', '123.918186,47.354348', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (594, 220204, '船营区', 1, '0432', '126.540966,43.833445', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (595, 220211, '丰满区', 1, '0432', '126.562274,43.821601', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (596, 220221, '永吉县', 1, '0432', '126.497741,43.672582', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (597, 220281, '蛟河市', 1, '0432', '127.344229,43.724007', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (598, 220282, '桦甸市', 1, '0432', '126.746309,42.972096', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (599, 220283, '舒兰市', 1, '0432', '126.965607,44.406105', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (600, 220284, '磐石市', 1, '0432', '126.060427,42.946285', 591);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (602, 220302, '铁西区', 1, '0434', '124.345722,43.146155', 601);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (603, 220303, '铁东区', 1, '0434', '124.409591,43.162105', 601);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (604, 220322, '梨树县', 1, '0434', '124.33539,43.30706', 601);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (605, 220323, '伊通满族自治县', 1, '0434', '125.305393,43.345754', 601);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (606, 220381, '公主岭市', 1, '0434', '124.822929,43.504676', 601);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (607, 220382, '双辽市', 1, '0434', '123.502723,43.518302', 601);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (609, 220402, '龙山区', 1, '0437', '125.136627,42.90158', 608);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (610, 220403, '西安区', 1, '0437', '125.149281,42.927324', 608);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (611, 220421, '东丰县', 1, '0437', '125.531021,42.677371', 608);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (686, 230300, '鸡西市', 2, '0467', '130.969333,45.295075', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (612, 220422, '东辽县', 1, '0437', '124.991424,42.92625', 608);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (614, 220502, '东昌区', 1, '0435', '125.927101,41.702859', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (615, 220503, '二道江区', 1, '0435', '126.042678,41.774044', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (616, 220521, '通化县', 1, '0435', '125.759259,41.679808', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (617, 220523, '辉南县', 1, '0435', '126.046783,42.684921', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (618, 220524, '柳河县', 1, '0435', '125.744735,42.284605', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (619, 220581, '梅河口市', 1, '0435', '125.710859,42.539253', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (620, 220582, '集安市', 1, '0435', '126.19403,41.125307', 613);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (623, 220605, '江源区', 1, '0439', '126.591178,42.056747', 621);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (696, 230400, '鹤岗市', 2, '0468', '130.297943,47.350189', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (624, 220621, '抚松县', 1, '0439', '127.449763,42.221207', 621);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (625, 220622, '靖宇县', 1, '0439', '126.813583,42.388896', 621);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (626, 220623, '长白朝鲜族自治县', 1, '0439', '128.200789,41.420018', 621);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (627, 220681, '临江市', 1, '0439', '126.918087,41.811979', 621);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (629, 220702, '宁江区', 1, '0438', '124.86562,45.209915', 628);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (630, 220721, '前郭尔罗斯蒙古族自治县', 1, '0438', '124.823417,45.118061', 628);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (631, 220722, '长岭县', 1, '0438', '123.967483,44.275895', 628);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (632, 220723, '乾安县', 1, '0438', '124.041139,45.003773', 628);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (705, 230500, '双鸭山市', 2, '0469', '131.141195,46.676418', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (633, 220781, '扶余市', 1, '0438', '126.049803,44.9892', 628);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (635, 220802, '洮北区', 1, '0436', '122.851029,45.621716', 634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (636, 220821, '镇赉县', 1, '0436', '123.199607,45.84835', 634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (710, 230521, '集贤县', 1, '0469', '131.141311,46.728412', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (637, 220822, '通榆县', 1, '0436', '123.088238,44.81291', 634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (638, 220881, '洮南市', 1, '0436', '122.798579,45.356807', 634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (714, 230600, '大庆市', 2, '0459', '125.103784,46.589309', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (639, 220882, '大安市', 1, '0436', '124.292626,45.506996', 634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (641, 222401, '延吉市', 1, '1433', '129.508804,42.89125', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (642, 222402, '图们市', 1, '1433', '129.84371,42.968044', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (643, 222403, '敦化市', 1, '1433', '128.232131,43.372642', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (644, 222404, '珲春市', 1, '1433', '130.366036,42.862821', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (645, 222405, '龙井市', 1, '1433', '129.427066,42.76631', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (646, 222406, '和龙市', 1, '1433', '129.010106,42.546675', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (647, 222424, '汪清县', 1, '1433', '129.771607,43.312522', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (648, 222426, '安图县', 1, '1433', '128.899772,43.11195', 640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (724, 230700, '伊春市', 2, '0458', '128.841125,47.727535', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (651, 230102, '道里区', 1, '0451', '126.616973,45.75577', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (652, 230103, '南岗区', 1, '0451', '126.668784,45.760174', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (653, 230104, '道外区', 1, '0451', '126.64939,45.792057', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (654, 230108, '平房区', 1, '0451', '126.637611,45.597911', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (655, 230109, '松北区', 1, '0451', '126.516914,45.794504', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (656, 230110, '香坊区', 1, '0451', '126.662593,45.707716', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (657, 230111, '呼兰区', 1, '0451', '126.587905,45.889457', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (658, 230112, '阿城区', 1, '0451', '126.958098,45.548669', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (659, 230113, '双城区', 1, '0451', '126.312624,45.383218', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (660, 230123, '依兰县', 1, '0451', '129.567877,46.325419', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (661, 230124, '方正县', 1, '0451', '128.829536,45.851694', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (662, 230125, '宾县', 1, '0451', '127.466634,45.745917', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (663, 230126, '巴彦县', 1, '0451', '127.403781,46.086549', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (664, 230127, '木兰县', 1, '0451', '128.043466,45.950582', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (665, 230128, '通河县', 1, '0451', '128.746124,45.990205', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (666, 230129, '延寿县', 1, '0451', '128.331643,45.451897', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (667, 230183, '尚志市', 1, '0451', '128.009894,45.209586', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (742, 230800, '佳木斯市', 2, '0454', '130.318878,46.799777', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (668, 230184, '五常市', 1, '0451', '127.167618,44.931991', 650);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (670, 230202, '龙沙区', 1, '0452', '123.957531,47.317308', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (671, 230203, '建华区', 1, '0452', '123.955464,47.354364', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (672, 230204, '铁锋区', 1, '0452', '123.978293,47.340517', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (673, 230205, '昂昂溪区', 1, '0452', '123.8224,47.15516', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (674, 230206, '富拉尔基区', 1, '0452', '123.629189,47.208843', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (675, 230207, '碾子山区', 1, '0452', '122.887775,47.516872', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (676, 230208, '梅里斯达斡尔族区', 1, '0452', '123.75291,47.309537', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (677, 230221, '龙江县', 1, '0452', '123.205323,47.338665', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (678, 230223, '依安县', 1, '0452', '125.306278,47.893548', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (753, 230900, '七台河市', 2, '0464', '131.003082,45.771396', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (679, 230224, '泰来县', 1, '0452', '123.416631,46.393694', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (680, 230225, '甘南县', 1, '0452', '123.507429,47.922405', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (681, 230227, '富裕县', 1, '0452', '124.473793,47.774347', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (682, 230229, '克山县', 1, '0452', '125.875705,48.037031', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (758, 231000, '牡丹江市', 2, '0453', '129.633168,44.551653', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (683, 230230, '克东县', 1, '0452', '126.24872,48.04206', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (684, 230231, '拜泉县', 1, '0452', '126.100213,47.595851', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (685, 230281, '讷河市', 1, '0452', '124.88287,48.466592', 669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (687, 230302, '鸡冠区', 1, '0467', '130.981185,45.304355', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (688, 230303, '恒山区', 1, '0467', '130.904963,45.210668', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (689, 230304, '滴道区', 1, '0467', '130.843613,45.348763', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (690, 230305, '梨树区', 1, '0467', '130.69699,45.092046', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (691, 230306, '城子河区', 1, '0467', '131.011304,45.33697', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (692, 230307, '麻山区', 1, '0467', '130.478187,45.212088', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (693, 230321, '鸡东县', 1, '0467', '131.124079,45.260412', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (769, 231100, '黑河市', 2, '0456', '127.528293,50.245129', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (694, 230381, '虎林市', 1, '0467', '132.93721,45.762685', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (695, 230382, '密山市', 1, '0467', '131.846635,45.529774', 686);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (697, 230402, '向阳区', 1, '0468', '130.294235,47.342468', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (698, 230403, '工农区', 1, '0468', '130.274684,47.31878', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (699, 230404, '南山区', 1, '0468', '130.286788,47.315174', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (700, 230405, '兴安区', 1, '0468', '130.239245,47.252849', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (776, 231200, '绥化市', 2, '0455', '126.968887,46.653845', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (701, 230406, '东山区', 1, '0468', '130.317002,47.338537', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (702, 230407, '兴山区', 1, '0468', '130.303481,47.357702', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (703, 230421, '萝北县', 1, '0468', '130.85155,47.576444', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (704, 230422, '绥滨县', 1, '0468', '131.852759,47.289115', 696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (706, 230502, '尖山区', 1, '0469', '131.158415,46.64635', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (707, 230503, '岭东区', 1, '0469', '131.164723,46.592721', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (708, 230505, '四方台区', 1, '0469', '131.337592,46.597264', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (709, 230506, '宝山区', 1, '0469', '131.401589,46.577167', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (711, 230522, '友谊县', 1, '0469', '131.808063,46.767299', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (712, 230523, '宝清县', 1, '0469', '132.196853,46.327457', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (787, 232700, '大兴安岭地区', 2, '0457', '124.711526,52.335262', 649);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (713, 230524, '饶河县', 1, '0469', '134.013872,46.798163', 705);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (715, 230602, '萨尔图区', 1, '0459', '125.135591,46.629092', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (716, 230603, '龙凤区', 1, '0459', '125.135326,46.562247', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (717, 230604, '让胡路区', 1, '0459', '124.870596,46.652357', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (718, 230605, '红岗区', 1, '0459', '124.891039,46.398418', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (793, 310100, '上海城区', 2, '021', '121.473662,31.230372', 792);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (719, 230606, '大同区', 1, '0459', '124.812364,46.039827', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (720, 230621, '肇州县', 1, '0459', '125.268643,45.699066', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (721, 230622, '肇源县', 1, '0459', '125.078223,45.51932', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (722, 230623, '林甸县', 1, '0459', '124.863603,47.171717', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (798, 310107, '普陀区', 1, '021', '121.395514,31.249603', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (723, 230624, '杜尔伯特蒙古族自治县', 1, '0459', '124.442572,46.862817', 714);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (725, 230702, '伊春区', 1, '0458', '128.907257,47.728237', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (726, 230703, '南岔区', 1, '0458', '129.283467,47.138034', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (727, 230704, '友好区', 1, '0458', '128.836291,47.841032', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (728, 230705, '西林区', 1, '0458', '129.312851,47.480735', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (729, 230706, '翠峦区', 1, '0458', '128.669754,47.726394', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (730, 230707, '新青区', 1, '0458', '129.533599,48.290455', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (731, 230708, '美溪区', 1, '0458', '129.129314,47.63509', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (732, 230709, '金山屯区', 1, '0458', '129.429117,47.413074', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (733, 230710, '五营区', 1, '0458', '129.245343,48.10791', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (734, 230711, '乌马河区', 1, '0458', '128.799477,47.727687', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (811, 320100, '南京市', 2, '025', '118.796682,32.05957', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (735, 230712, '汤旺河区', 1, '0458', '129.571108,48.454651', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (736, 230713, '带岭区', 1, '0458', '129.020888,47.028379', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (737, 230714, '乌伊岭区', 1, '0458', '129.43792,48.590322', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (738, 230715, '红星区', 1, '0458', '129.390983,48.239431', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (739, 230716, '上甘岭区', 1, '0458', '129.02426,47.974707', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (740, 230722, '嘉荫县', 1, '0458', '130.403134,48.888972', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (741, 230781, '铁力市', 1, '0458', '128.032424,46.986633', 724);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (743, 230803, '向阳区', 1, '0454', '130.365346,46.80779', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (744, 230804, '前进区', 1, '0454', '130.375062,46.814102', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (745, 230805, '东风区', 1, '0454', '130.403664,46.822571', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (746, 230811, '郊区', 1, '0454', '130.327194,46.810085', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (823, 320200, '无锡市', 2, '0510', '120.31191,31.491169', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (747, 230822, '桦南县', 1, '0454', '130.553343,46.239184', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (748, 230826, '桦川县', 1, '0454', '130.71908,47.023001', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (749, 230828, '汤原县', 1, '0454', '129.905072,46.730706', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (750, 230881, '同江市', 1, '0454', '132.510919,47.642707', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (751, 230882, '富锦市', 1, '0454', '132.037686,47.250107', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (752, 230883, '抚远市', 1, '0454', '134.307884,48.364687', 742);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (754, 230902, '新兴区', 1, '0464', '130.932143,45.81593', 753);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (831, 320300, '徐州市', 2, '0516', '117.284124,34.205768', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (755, 230903, '桃山区', 1, '0464', '131.020202,45.765705', 753);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (756, 230904, '茄子河区', 1, '0464', '131.068075,45.785215', 753);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (757, 230921, '勃利县', 1, '0464', '130.59217,45.755063', 753);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (759, 231002, '东安区', 1, '0453', '129.626641,44.58136', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (760, 231003, '阳明区', 1, '0453', '129.635615,44.596104', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (761, 231004, '爱民区', 1, '0453', '129.591537,44.596042', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (762, 231005, '西安区', 1, '0453', '129.616058,44.577625', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (763, 231025, '林口县', 1, '0453', '130.284033,45.278046', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (764, 231081, '绥芬河市', 1, '0453', '131.152545,44.412308', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (765, 231083, '海林市', 1, '0453', '129.380481,44.594213', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (842, 320400, '常州市', 2, '0519', '119.974061,31.811226', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (766, 231084, '宁安市', 1, '0453', '129.482851,44.34072', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (767, 231085, '穆棱市', 1, '0453', '130.524436,44.918813', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (768, 231086, '东宁市', 1, '0453', '131.122915,44.087585', 758);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (770, 231102, '爱辉区', 1, '0456', '127.50045,50.252105', 769);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (771, 231121, '嫩江县', 1, '0456', '125.221192,49.185766', 769);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (772, 231123, '逊克县', 1, '0456', '128.478749,49.564252', 769);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (849, 320500, '苏州市', 2, '0512', '120.585728,31.2974', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (773, 231124, '孙吴县', 1, '0456', '127.336303,49.425647', 769);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (774, 231181, '北安市', 1, '0456', '126.490864,48.241365', 769);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (775, 231182, '五大连池市', 1, '0456', '126.205516,48.517257', 769);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (777, 231202, '北林区', 1, '0455', '126.985504,46.6375', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (778, 231221, '望奎县', 1, '0455', '126.486075,46.832719', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (779, 231222, '兰西县', 1, '0455', '126.288117,46.25245', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (780, 231223, '青冈县', 1, '0455', '126.099195,46.70391', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (781, 231224, '庆安县', 1, '0455', '127.507824,46.880102', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (782, 231225, '明水县', 1, '0455', '125.906301,47.173426', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (859, 320600, '南通市', 2, '0513', '120.894676,31.981143', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (783, 231226, '绥棱县', 1, '0455', '127.114832,47.236015', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (784, 231281, '安达市', 1, '0455', '125.346156,46.419633', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (785, 231282, '肇东市', 1, '0455', '125.961814,46.051126', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (786, 231283, '海伦市', 1, '0455', '126.930106,47.45117', 776);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (788, 232701, '加格达奇区', 1, '0457', '124.139595,50.408735', 787);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (789, 232721, '呼玛县', 1, '0457', '126.652396,51.726091', 787);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (790, 232722, '塔河县', 1, '0457', '124.709996,52.334456', 787);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (791, 232723, '漠河县', 1, '0457', '122.538591,52.972272', 787);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (868, 320700, '连云港市', 2, '0518', '119.221611,34.596653', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (794, 310101, '黄浦区', 1, '021', '121.484428,31.231739', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (795, 310104, '徐汇区', 1, '021', '121.436128,31.188464', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (796, 310105, '长宁区', 1, '021', '121.424622,31.220372', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (797, 310106, '静安区', 1, '021', '121.447453,31.227906', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (799, 310109, '虹口区', 1, '021', '121.505133,31.2646', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (800, 310110, '杨浦区', 1, '021', '121.525727,31.259822', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (875, 320800, '淮安市', 2, '0517', '119.113185,33.551052', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (801, 310112, '闵行区', 1, '021', '121.380831,31.1129', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (802, 310113, '宝山区', 1, '021', '121.489612,31.405457', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (803, 310114, '嘉定区', 1, '021', '121.265374,31.375869', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (804, 310115, '浦东新区', 1, '021', '121.544379,31.221517', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (805, 310116, '金山区', 1, '021', '121.342455,30.741798', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (806, 310117, '松江区', 1, '021', '121.227747,31.032243', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (807, 310118, '青浦区', 1, '021', '121.124178,31.150681', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (883, 320900, '盐城市', 2, '0515', '120.163107,33.347708', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (808, 310120, '奉贤区', 1, '021', '121.474055,30.917766', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (809, 310151, '崇明区', 1, '021', '121.397421,31.623728', 793);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (812, 320102, '玄武区', 1, '025', '118.797757,32.048498', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (887, 320921, '响水县', 1, '0515', '119.578364,34.199479', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (813, 320104, '秦淮区', 1, '025', '118.79476,32.039113', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (814, 320105, '建邺区', 1, '025', '118.731793,32.003731', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (815, 320106, '鼓楼区', 1, '025', '118.770182,32.066601', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (816, 320111, '浦口区', 1, '025', '118.628003,32.058903', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (893, 321000, '扬州市', 2, '0514', '119.412939,32.394209', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (817, 320113, '栖霞区', 1, '025', '118.909153,32.096388', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (818, 320114, '雨花台区', 1, '025', '118.779051,31.99126', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (819, 320115, '江宁区', 1, '025', '118.840015,31.952612', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (820, 320116, '六合区', 1, '025', '118.822132,32.323584', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (821, 320117, '溧水区', 1, '025', '119.028288,31.651099', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (822, 320118, '高淳区', 1, '025', '118.89222,31.327586', 811);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (900, 321100, '镇江市', 2, '0511', '119.425836,32.187849', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (824, 320205, '锡山区', 1, '0510', '120.357858,31.589715', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (825, 320206, '惠山区', 1, '0510', '120.298433,31.680335', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (826, 320211, '滨湖区', 1, '0510', '120.283811,31.527276', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (827, 320213, '梁溪区', 1, '0510', '120.303108,31.566155', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (828, 320214, '新吴区', 1, '0510', '120.352782,31.550966', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (829, 320281, '江阴市', 1, '0510', '120.286129,31.921345', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (907, 321200, '泰州市', 2, '0523', '119.922933,32.455536', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (830, 320282, '宜兴市', 1, '0510', '119.823308,31.340637', 823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (832, 320302, '鼓楼区', 1, '0516', '117.185576,34.288646', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (833, 320303, '云龙区', 1, '0516', '117.251076,34.253164', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (834, 320305, '贾汪区', 1, '0516', '117.464958,34.436936', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (835, 320311, '泉山区', 1, '0516', '117.194469,34.225522', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (836, 320312, '铜山区', 1, '0516', '117.169461,34.180779', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (914, 321300, '宿迁市', 2, '0527', '118.275198,33.963232', 810);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (837, 320321, '丰县', 1, '0516', '116.59539,34.693906', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (838, 320322, '沛县', 1, '0516', '116.936353,34.760761', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (839, 320324, '睢宁县', 1, '0516', '117.941563,33.912597', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (840, 320381, '新沂市', 1, '0516', '118.354537,34.36958', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (841, 320382, '邳州市', 1, '0516', '118.012531,34.338888', 831);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (843, 320402, '天宁区', 1, '0519', '119.999219,31.792787', 842);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (921, 330100, '杭州市', 2, '0571', '120.209789,30.24692', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (844, 320404, '钟楼区', 1, '0519', '119.902369,31.802089', 842);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (845, 320411, '新北区', 1, '0519', '119.971697,31.830427', 842);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (846, 320412, '武进区', 1, '0519', '119.942437,31.701187', 842);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (847, 320413, '金坛区', 1, '0519', '119.597811,31.723219', 842);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (848, 320481, '溧阳市', 1, '0519', '119.48421,31.416911', 842);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (850, 320505, '虎丘区', 1, '0512', '120.434238,31.329601', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (851, 320506, '吴中区', 1, '0512', '120.632308,31.263183', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (852, 320507, '相城区', 1, '0512', '120.642626,31.369089', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (853, 320508, '姑苏区', 1, '0512', '120.617369,31.33565', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (854, 320509, '吴江区', 1, '0512', '120.645157,31.138677', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (855, 320581, '常熟市', 1, '0512', '120.752481,31.654375', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (856, 320582, '张家港市', 1, '0512', '120.555982,31.875571', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (857, 320583, '昆山市', 1, '0512', '120.980736,31.385597', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (935, 330200, '宁波市', 2, '0574', '121.622485,29.859971', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (858, 320585, '太仓市', 1, '0512', '121.13055,31.457735', 849);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (860, 320602, '崇川区', 1, '0513', '120.857434,32.009875', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (861, 320611, '港闸区', 1, '0513', '120.818526,32.032441', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (862, 320612, '通州区', 1, '0513', '121.073828,32.06568', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (863, 320621, '海安县', 1, '0513', '120.467343,32.533572', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (864, 320623, '如东县', 1, '0513', '121.185201,32.331765', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (865, 320681, '启东市', 1, '0513', '121.655432,31.793278', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (866, 320682, '如皋市', 1, '0513', '120.573803,32.371562', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (867, 320684, '海门市', 1, '0513', '121.18181,31.869483', 859);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (869, 320703, '连云区', 1, '0518', '119.338788,34.760249', 868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (946, 330300, '温州市', 2, '0577', '120.699361,27.993828', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (870, 320706, '海州区', 1, '0518', '119.163509,34.572274', 868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (871, 320707, '赣榆区', 1, '0518', '119.17333,34.841348', 868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (872, 320722, '东海县', 1, '0518', '118.752842,34.542308', 868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (873, 320723, '灌云县', 1, '0518', '119.239381,34.284381', 868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (874, 320724, '灌南县', 1, '0518', '119.315651,34.087134', 868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (876, 320812, '清江浦区', 1, '0517', '119.026718,33.552627', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (877, 320803, '淮安区', 1, '0517', '119.141099,33.502868', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (878, 320804, '淮阴区', 1, '0517', '119.034725,33.631892', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (879, 320813, '洪泽区', 1, '0517', '118.873241,33.294214', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (880, 320826, '涟水县', 1, '0517', '119.260227,33.781331', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (881, 320830, '盱眙县', 1, '0517', '118.54436,33.011971', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (958, 330400, '嘉兴市', 2, '0573', '120.75547,30.746191', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (882, 320831, '金湖县', 1, '0517', '119.020584,33.025433', 875);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (884, 320902, '亭湖区', 1, '0515', '120.197358,33.390536', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (885, 320903, '盐都区', 1, '0515', '120.153712,33.338283', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (886, 320904, '大丰区', 1, '0515', '120.50085,33.200333', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (888, 320922, '滨海县', 1, '0515', '119.82083,33.990334', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (889, 320923, '阜宁县', 1, '0515', '119.802527,33.759325', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (890, 320924, '射阳县', 1, '0515', '120.229986,33.758361', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (966, 330500, '湖州市', 2, '0572', '120.086809,30.89441', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (891, 320925, '建湖县', 1, '0515', '119.7886,33.439067', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (892, 320981, '东台市', 1, '0515', '120.320328,32.868426', 883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (894, 321002, '广陵区', 1, '0514', '119.431849,32.39472', 893);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (895, 321003, '邗江区', 1, '0514', '119.397994,32.377655', 893);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (896, 321012, '江都区', 1, '0514', '119.569989,32.434672', 893);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (972, 330600, '绍兴市', 2, '0575', '120.580364,30.030192', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (897, 321023, '宝应县', 1, '0514', '119.360729,33.240391', 893);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (898, 321081, '仪征市', 1, '0514', '119.184766,32.272258', 893);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (899, 321084, '高邮市', 1, '0514', '119.459161,32.781659', 893);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (976, 330624, '新昌县', 1, '0575', '120.903866,29.499831', 972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (901, 321102, '京口区', 1, '0511', '119.47016,32.19828', 900);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (979, 330700, '金华市', 2, '0579', '119.647229,29.079208', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (902, 321111, '润州区', 1, '0511', '119.411959,32.195264', 900);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (903, 321112, '丹徒区', 1, '0511', '119.433853,32.131962', 900);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (904, 321181, '丹阳市', 1, '0511', '119.606439,32.010214', 900);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (905, 321182, '扬中市', 1, '0511', '119.797634,32.23483', 900);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (906, 321183, '句容市', 1, '0511', '119.168695,31.944998', 900);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (908, 321202, '海陵区', 1, '0523', '119.919424,32.491016', 907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (909, 321203, '高港区', 1, '0523', '119.881717,32.318821', 907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (910, 321204, '姜堰区', 1, '0523', '120.127934,32.509155', 907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (911, 321281, '兴化市', 1, '0523', '119.852541,32.910459', 907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (989, 330800, '衢州市', 2, '0570', '118.859457,28.970079', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (912, 321282, '靖江市', 1, '0523', '120.277138,31.982751', 907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (913, 321283, '泰兴市', 1, '0523', '120.051743,32.171853', 907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (915, 321302, '宿城区', 1, '0527', '118.242533,33.963029', 914);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (916, 321311, '宿豫区', 1, '0527', '118.330781,33.946822', 914);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (917, 321322, '沭阳县', 1, '0527', '118.804784,34.111022', 914);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (918, 321323, '泗阳县', 1, '0527', '118.703424,33.722478', 914);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (996, 330900, '舟山市', 2, '0580', '122.207106,29.985553', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (919, 321324, '泗洪县', 1, '0527', '118.223591,33.476051', 914);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (922, 330102, '上城区', 1, '0571', '120.169312,30.242404', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (923, 330103, '下城区', 1, '0571', '120.180891,30.281677', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (924, 330104, '江干区', 1, '0571', '120.205001,30.257012', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1001, 331000, '台州市', 2, '0576', '121.42076,28.65638', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (925, 330105, '拱墅区', 1, '0571', '120.141406,30.319037', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (926, 330106, '西湖区', 1, '0571', '120.130194,30.259463', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (927, 330108, '滨江区', 1, '0571', '120.211623,30.208847', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (928, 330109, '萧山区', 1, '0571', '120.264253,30.183806', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (929, 330110, '余杭区', 1, '0571', '120.299401,30.419045', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (930, 330111, '富阳区', 1, '0571', '119.960076,30.048692', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (931, 330122, '桐庐县', 1, '0571', '119.691467,29.79299', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (932, 330127, '淳安县', 1, '0571', '119.042037,29.608886', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (933, 330182, '建德市', 1, '0571', '119.281231,29.474759', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1011, 331100, '丽水市', 2, '0578', '119.922796,28.46763', 920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (934, 330185, '临安市', 1, '0571', '119.724734,30.233873', 921);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (936, 330203, '海曙区', 1, '0574', '121.550752,29.874903', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (937, 330205, '江北区', 1, '0574', '121.555081,29.886781', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (938, 330206, '北仑区', 1, '0574', '121.844172,29.899778', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (939, 330211, '镇海区', 1, '0574', '121.596496,29.965203', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (940, 330212, '鄞州区', 1, '0574', '121.546603,29.816511', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (941, 330225, '象山县', 1, '0574', '121.869339,29.476705', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (942, 330226, '宁海县', 1, '0574', '121.429477,29.287939', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (943, 330281, '余姚市', 1, '0574', '121.154629,30.037106', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (944, 330282, '慈溪市', 1, '0574', '121.266561,30.170261', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1022, 340100, '合肥市', 2, '0551', '117.227219,31.820591', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (945, 330213, '奉化区', 1, '0574', '121.406997,29.655144', 935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (947, 330302, '鹿城区', 1, '0577', '120.655271,28.015737', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (948, 330303, '龙湾区', 1, '0577', '120.811213,27.932747', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (949, 330304, '瓯海区', 1, '0577', '120.61491,27.966844', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (950, 330305, '洞头区', 1, '0577', '121.157249,27.836154', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (951, 330324, '永嘉县', 1, '0577', '120.692025,28.153607', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (952, 330326, '平阳县', 1, '0577', '120.565793,27.661918', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (953, 330327, '苍南县', 1, '0577', '120.427619,27.519773', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (954, 330328, '文成县', 1, '0577', '120.091498,27.786996', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1032, 340200, '芜湖市', 2, '0553', '118.432941,31.352859', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (955, 330329, '泰顺县', 1, '0577', '119.717649,27.556884', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (956, 330381, '瑞安市', 1, '0577', '120.655148,27.778657', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (957, 330382, '乐清市', 1, '0577', '120.983906,28.113725', 946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (959, 330402, '南湖区', 1, '0573', '120.783024,30.747842', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (960, 330411, '秀洲区', 1, '0573', '120.710082,30.765188', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (961, 330421, '嘉善县', 1, '0573', '120.926028,30.830864', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (962, 330424, '海盐县', 1, '0573', '120.946263,30.526435', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (963, 330481, '海宁市', 1, '0573', '120.680239,30.511539', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1041, 340300, '蚌埠市', 2, '0552', '117.388512,32.91663', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (964, 330482, '平湖市', 1, '0573', '121.015142,30.677233', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (965, 330483, '桐乡市', 1, '0573', '120.565098,30.630173', 958);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (967, 330502, '吴兴区', 1, '0572', '120.185838,30.857151', 966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (968, 330503, '南浔区', 1, '0572', '120.418513,30.849689', 966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (969, 330521, '德清县', 1, '0572', '119.9774,30.54251', 966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (970, 330522, '长兴县', 1, '0572', '119.910952,31.026665', 966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (971, 330523, '安吉县', 1, '0572', '119.680353,30.638674', 966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1049, 340400, '淮南市', 2, '0554', '117.018399,32.587117', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (973, 330602, '越城区', 1, '0575', '120.582633,29.988244', 972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (974, 330603, '柯桥区', 1, '0575', '120.495085,30.081929', 972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (975, 330604, '上虞区', 1, '0575', '120.868122,30.03312', 972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (977, 330681, '诸暨市', 1, '0575', '120.246863,29.708692', 972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (978, 330683, '嵊州市', 1, '0575', '120.831025,29.56141', 972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (980, 330702, '婺城区', 1, '0579', '119.571728,29.0872', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (981, 330703, '金东区', 1, '0579', '119.69278,29.099723', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1057, 340500, '马鞍山市', 2, '0555', '118.507011,31.67044', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (982, 330723, '武义县', 1, '0579', '119.816562,28.89267', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (983, 330726, '浦江县', 1, '0579', '119.892222,29.452476', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (984, 330727, '磐安县', 1, '0579', '120.450005,29.054548', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (985, 330781, '兰溪市', 1, '0579', '119.460472,29.2084', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (986, 330782, '义乌市', 1, '0579', '120.075106,29.306775', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (987, 330783, '东阳市', 1, '0579', '120.241566,29.289648', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1064, 340600, '淮北市', 2, '0561', '116.798265,33.955844', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1065, 340602, '杜集区', 1, '0561', '116.828133,33.991451', 1064);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (988, 330784, '永康市', 1, '0579', '120.047651,28.888555', 979);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (990, 330802, '柯城区', 1, '0570', '118.871516,28.96862', 989);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1069, 340700, '铜陵市', 2, '0562', '117.81154,30.945515', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (991, 330803, '衢江区', 1, '0570', '118.95946,28.97978', 989);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (992, 330822, '常山县', 1, '0570', '118.511235,28.901462', 989);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (993, 330824, '开化县', 1, '0570', '118.415495,29.137336', 989);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (994, 330825, '龙游县', 1, '0570', '119.172189,29.028439', 989);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1074, 340800, '安庆市', 2, '0556', '117.115101,30.531919', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (995, 330881, '江山市', 1, '0570', '118.626991,28.737331', 989);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (997, 330902, '定海区', 1, '0580', '122.106773,30.019858', 996);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (998, 330903, '普陀区', 1, '0580', '122.323867,29.97176', 996);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (999, 330921, '岱山县', 1, '0580', '122.226237,30.264139', 996);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1000, 330922, '嵊泗县', 1, '0580', '122.451382,30.725686', 996);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1002, 331002, '椒江区', 1, '0576', '121.442978,28.672981', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1003, 331003, '黄岩区', 1, '0576', '121.261972,28.650083', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1004, 331004, '路桥区', 1, '0576', '121.365123,28.582654', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1005, 331021, '玉环市', 1, '0576', '121.231805,28.135929', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1006, 331022, '三门县', 1, '0576', '121.395711,29.104789', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1085, 341000, '黄山市', 2, '0559', '118.338272,29.715185', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1007, 331023, '天台县', 1, '0576', '121.006595,29.144064', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1008, 331024, '仙居县', 1, '0576', '120.728801,28.846966', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1009, 331081, '温岭市', 1, '0576', '121.385604,28.372506', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1010, 331082, '临海市', 1, '0576', '121.144556,28.858881', 1001);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1012, 331102, '莲都区', 1, '0578', '119.912626,28.445928', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1013, 331121, '青田县', 1, '0578', '120.289478,28.139837', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1014, 331122, '缙云县', 1, '0578', '120.091572,28.659283', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1093, 341100, '滁州市', 2, '0550', '118.327944,32.255636', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1015, 331123, '遂昌县', 1, '0578', '119.276103,28.592148', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1016, 331124, '松阳县', 1, '0578', '119.481511,28.448803', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1017, 331125, '云和县', 1, '0578', '119.573397,28.11579', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1018, 331126, '庆元县', 1, '0578', '119.06259,27.61922', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1019, 331127, '景宁畲族自治县', 1, '0578', '119.635739,27.9733', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1020, 331181, '龙泉市', 1, '0578', '119.141473,28.074649', 1011);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1023, 340102, '瑶海区', 1, '0551', '117.309546,31.857917', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1024, 340103, '庐阳区', 1, '0551', '117.264786,31.878589', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1102, 341200, '阜阳市', 2, '1558', '115.814504,32.890479', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1025, 340104, '蜀山区', 1, '0551', '117.260521,31.85124', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1026, 340111, '包河区', 1, '0551', '117.309519,31.793859', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1027, 340121, '长丰县', 1, '0551', '117.167564,32.478018', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1028, 340122, '肥东县', 1, '0551', '117.469382,31.88794', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1029, 340123, '肥西县', 1, '0551', '117.157981,31.706809', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1030, 340124, '庐江县', 1, '0551', '117.2882,31.256524', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1031, 340181, '巢湖市', 1, '0551', '117.890354,31.624522', 1022);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1033, 340202, '镜湖区', 1, '0553', '118.385009,31.340728', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1111, 341300, '宿州市', 2, '0557', '116.964195,33.647309', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1034, 340203, '弋江区', 1, '0553', '118.372655,31.311756', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1035, 340207, '鸠江区', 1, '0553', '118.391734,31.369373', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1036, 340208, '三山区', 1, '0553', '118.268101,31.219568', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1037, 340221, '芜湖县', 1, '0553', '118.576124,31.134809', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1038, 340222, '繁昌县', 1, '0553', '118.198703,31.101782', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1117, 341500, '六安市', 2, '0564', '116.520139,31.735456', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1039, 340223, '南陵县', 1, '0553', '118.334359,30.914922', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1040, 340225, '无为县', 1, '0553', '117.902366,31.303167', 1032);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1042, 340302, '龙子湖区', 1, '0552', '117.379778,32.950611', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1043, 340303, '蚌山区', 1, '0552', '117.373595,32.917048', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1044, 340304, '禹会区', 1, '0552', '117.342155,32.929799', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1045, 340311, '淮上区', 1, '0552', '117.35933,32.965435', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1046, 340321, '怀远县', 1, '0552', '117.205237,32.970031', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1125, 341600, '亳州市', 2, '0558', '115.77867,33.844592', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1047, 340322, '五河县', 1, '0552', '117.879486,33.127823', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1048, 340323, '固镇县', 1, '0552', '117.316913,33.31688', 1041);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1050, 340402, '大通区', 1, '0554', '117.053314,32.631519', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1051, 340403, '田家庵区', 1, '0554', '117.017349,32.647277', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1130, 341700, '池州市', 2, '0566', '117.491592,30.664779', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1052, 340404, '谢家集区', 1, '0554', '116.859188,32.600037', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1053, 340405, '八公山区', 1, '0554', '116.83349,32.631379', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1054, 340406, '潘集区', 1, '0554', '116.834715,32.77208', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1055, 340421, '凤台县', 1, '0554', '116.71105,32.709444', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1135, 341800, '宣城市', 2, '0563', '118.75868,30.940195', 1021);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1056, 340422, '寿县', 1, '0554', '116.798232,32.545109', 1049);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1058, 340503, '花山区', 1, '0555', '118.492565,31.71971', 1057);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1059, 340504, '雨山区', 1, '0555', '118.498578,31.682132', 1057);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1060, 340506, '博望区', 1, '0555', '118.844538,31.558471', 1057);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1061, 340521, '当涂县', 1, '0555', '118.497972,31.571213', 1057);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1062, 340522, '含山县', 1, '0555', '118.101421,31.735598', 1057);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1063, 340523, '和县', 1, '0555', '118.353667,31.742293', 1057);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1066, 340603, '相山区', 1, '0561', '116.794344,33.959892', 1064);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1144, 350100, '福州市', 2, '0591', '119.296389,26.074268', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1067, 340604, '烈山区', 1, '0561', '116.813042,33.895139', 1064);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1068, 340621, '濉溪县', 1, '0561', '116.766298,33.915477', 1064);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1070, 340705, '铜官区', 1, '0562', '117.85616,30.936272', 1069);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1071, 340706, '义安区', 1, '0562', '117.791544,30.952823', 1069);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1072, 340711, '郊区', 1, '0562', '117.768026,30.821069', 1069);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1073, 340722, '枞阳县', 1, '0562', '117.250594,30.706018', 1069);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1075, 340802, '迎江区', 1, '0556', '117.09115,30.511548', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1076, 340803, '大观区', 1, '0556', '117.013469,30.553697', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1077, 340811, '宜秀区', 1, '0556', '116.987542,30.613332', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1179, 350427, '沙县', 1, '0598', '117.792396,26.397199', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1078, 340822, '怀宁县', 1, '0556', '116.829475,30.733824', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1079, 340824, '潜山县', 1, '0556', '116.581371,30.631136', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1080, 340825, '太湖县', 1, '0556', '116.308795,30.45422', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1158, 350200, '厦门市', 2, '0592', '118.089204,24.479664', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1081, 340826, '宿松县', 1, '0556', '116.129105,30.153746', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1082, 340827, '望江县', 1, '0556', '116.706498,30.128002', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1083, 340828, '岳西县', 1, '0556', '116.359692,30.849762', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1084, 340881, '桐城市', 1, '0556', '116.936748,31.035848', 1074);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1086, 341002, '屯溪区', 1, '0559', '118.315329,29.696108', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1087, 341003, '黄山区', 1, '0559', '118.141567,30.272942', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1165, 350300, '莆田市', 2, '0594', '119.007777,25.454084', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1088, 341004, '徽州区', 1, '0559', '118.336743,29.827271', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1089, 341021, '歙县', 1, '0559', '118.415345,29.861379', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1090, 341022, '休宁县', 1, '0559', '118.193618,29.784124', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1091, 341023, '黟县', 1, '0559', '117.938373,29.924805', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1092, 341024, '祁门县', 1, '0559', '117.717396,29.854055', 1085);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1171, 350400, '三明市', 2, '0598', '117.638678,26.263406', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1094, 341102, '琅琊区', 1, '0550', '118.305961,32.294631', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1095, 341103, '南谯区', 1, '0550', '118.41697,32.200197', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1096, 341122, '来安县', 1, '0550', '118.435718,32.452199', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1097, 341124, '全椒县', 1, '0550', '118.274149,32.08593', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1098, 341125, '定远县', 1, '0550', '117.698562,32.530981', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1099, 341126, '凤阳县', 1, '0550', '117.531622,32.874735', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1100, 341181, '天长市', 1, '0550', '119.004816,32.667571', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1101, 341182, '明光市', 1, '0550', '118.018193,32.78196', 1093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1103, 341202, '颍州区', 1, '1558', '115.806942,32.883468', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1104, 341203, '颍东区', 1, '1558', '115.856762,32.912477', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1105, 341204, '颍泉区', 1, '1558', '115.80835,32.925211', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1106, 341221, '临泉县', 1, '1558', '115.263115,33.039715', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1184, 350500, '泉州市', 2, '0595', '118.675676,24.874132', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1107, 341222, '太和县', 1, '1558', '115.621941,33.160327', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1108, 341225, '阜南县', 1, '1558', '115.595643,32.658297', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1109, 341226, '颍上县', 1, '1558', '116.256772,32.653211', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1110, 341282, '界首市', 1, '1558', '115.374821,33.258244', 1102);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1112, 341302, '埇桥区', 1, '0557', '116.977203,33.64059', 1111);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1113, 341321, '砀山县', 1, '0557', '116.367095,34.442561', 1111);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1114, 341322, '萧县', 1, '0557', '116.947349,34.188732', 1111);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1115, 341323, '灵璧县', 1, '0557', '117.549395,33.554604', 1111);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1116, 341324, '泗县', 1, '0557', '117.910629,33.482982', 1111);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1118, 341502, '金安区', 1, '0564', '116.539173,31.750119', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1119, 341503, '裕安区', 1, '0564', '116.479829,31.738183', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1120, 341504, '叶集区', 1, '0564', '115.925271,31.863693', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1197, 350600, '漳州市', 2, '0596', '117.647093,24.513025', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1121, 341522, '霍邱县', 1, '0564', '116.277911,32.353038', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1122, 341523, '舒城县', 1, '0564', '116.948736,31.462234', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1123, 341524, '金寨县', 1, '0564', '115.934366,31.72717', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1124, 341525, '霍山县', 1, '0564', '116.351892,31.410561', 1117);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1126, 341602, '谯城区', 1, '0558', '115.779025,33.876235', 1125);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1127, 341621, '涡阳县', 1, '0558', '116.215665,33.492921', 1125);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1128, 341622, '蒙城县', 1, '0558', '116.564247,33.26583', 1125);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1129, 341623, '利辛县', 1, '0558', '116.208564,33.144515', 1125);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1131, 341702, '贵池区', 1, '0566', '117.567264,30.687219', 1130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1132, 341721, '东至县', 1, '0566', '117.027618,30.111163', 1130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1133, 341722, '石台县', 1, '0566', '117.486306,30.210313', 1130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1209, 350700, '南平市', 2, '0599', '118.17771,26.641774', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1134, 341723, '青阳县', 1, '0566', '117.84743,30.63923', 1130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1136, 341802, '宣州区', 1, '0563', '118.785561,30.944076', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1137, 341821, '郎溪县', 1, '0563', '119.179656,31.126412', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1138, 341822, '广德县', 1, '0563', '119.420935,30.877555', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1139, 341823, '泾县', 1, '0563', '118.419859,30.688634', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1140, 341824, '绩溪县', 1, '0563', '118.578519,30.067533', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1141, 341825, '旌德县', 1, '0563', '118.549861,30.298142', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1142, 341881, '宁国市', 1, '0563', '118.983171,30.633882', 1135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1145, 350102, '鼓楼区', 1, '0591', '119.303917,26.081983', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1146, 350103, '台江区', 1, '0591', '119.314041,26.052843', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1220, 350800, '龙岩市', 2, '0597', '117.017295,25.075119', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1147, 350104, '仓山区', 1, '0591', '119.273545,26.046743', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1148, 350105, '马尾区', 1, '0591', '119.455588,25.9895', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1149, 350111, '晋安区', 1, '0591', '119.328521,26.082107', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1150, 350121, '闽侯县', 1, '0591', '119.131724,26.150047', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1151, 350122, '连江县', 1, '0591', '119.539704,26.197364', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1152, 350123, '罗源县', 1, '0591', '119.549776,26.489558', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1153, 350124, '闽清县', 1, '0591', '118.863361,26.221197', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1228, 350900, '宁德市', 2, '0593', '119.547932,26.665617', 1143);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1154, 350125, '永泰县', 1, '0591', '118.932592,25.866694', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1155, 350128, '平潭县', 1, '0591', '119.790168,25.49872', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1156, 350181, '福清市', 1, '0591', '119.384201,25.72071', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1157, 350182, '长乐市', 1, '0591', '119.523266,25.962888', 1144);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1159, 350203, '思明区', 1, '0592', '118.082649,24.445484', 1158);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1160, 350205, '海沧区', 1, '0592', '118.032984,24.484685', 1158);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1161, 350206, '湖里区', 1, '0592', '118.146768,24.512904', 1158);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1162, 350211, '集美区', 1, '0592', '118.097337,24.575969', 1158);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1163, 350212, '同安区', 1, '0592', '118.152041,24.723234', 1158);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1164, 350213, '翔安区', 1, '0592', '118.248034,24.618543', 1158);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1239, 360100, '南昌市', 2, '0791', '115.858198,28.682892', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1166, 350302, '城厢区', 1, '0594', '118.993884,25.419319', 1165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1167, 350303, '涵江区', 1, '0594', '119.116289,25.45872', 1165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1168, 350304, '荔城区', 1, '0594', '119.015061,25.431941', 1165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1169, 350305, '秀屿区', 1, '0594', '119.105494,25.31836', 1165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1170, 350322, '仙游县', 1, '0594', '118.691637,25.362093', 1165);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1172, 350402, '梅列区', 1, '0598', '117.645855,26.271711', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1173, 350403, '三元区', 1, '0598', '117.608044,26.234019', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1174, 350421, '明溪县', 1, '0598', '117.202226,26.355856', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1249, 360200, '景德镇市', 2, '0798', '117.178222,29.268945', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1175, 350423, '清流县', 1, '0598', '116.816909,26.177796', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1176, 350424, '宁化县', 1, '0598', '116.654365,26.261754', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1177, 350425, '大田县', 1, '0598', '117.847115,25.692699', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1178, 350426, '尤溪县', 1, '0598', '118.190467,26.170171', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1254, 360300, '萍乡市', 2, '0799', '113.887083,27.658373', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1180, 350428, '将乐县', 1, '0598', '117.471372,26.728952', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1181, 350429, '泰宁县', 1, '0598', '117.17574,26.900259', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1182, 350430, '建宁县', 1, '0598', '116.848443,26.833588', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1183, 350481, '永安市', 1, '0598', '117.365052,25.941937', 1171);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1185, 350502, '鲤城区', 1, '0595', '118.587097,24.907424', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1260, 360400, '九江市', 2, '0792', '115.952914,29.662117', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1186, 350503, '丰泽区', 1, '0595', '118.613172,24.891173', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1187, 350504, '洛江区', 1, '0595', '118.671193,24.939796', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1188, 350505, '泉港区', 1, '0595', '118.916309,25.119815', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1189, 350521, '惠安县', 1, '0595', '118.796607,25.030801', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1190, 350524, '安溪县', 1, '0595', '118.186288,25.055954', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1191, 350525, '永春县', 1, '0595', '118.294048,25.321565', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1192, 350526, '德化县', 1, '0595', '118.241094,25.491493', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1193, 350527, '金门县', 1, '0595', '118.323221,24.436417', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1194, 350581, '石狮市', 1, '0595', '118.648066,24.732204', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1195, 350582, '晋江市', 1, '0595', '118.551682,24.781636', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1196, 350583, '南安市', 1, '0595', '118.386279,24.960385', 1184);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1198, 350602, '芗城区', 1, '0596', '117.653968,24.510787', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1199, 350603, '龙文区', 1, '0596', '117.709754,24.503113', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1274, 360500, '新余市', 2, '0790', '114.917346,27.817808', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1200, 350622, '云霄县', 1, '0596', '117.339573,23.957936', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1201, 350623, '漳浦县', 1, '0596', '117.613808,24.117102', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1277, 360600, '鹰潭市', 2, '0701', '117.042173,28.272537', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1202, 350624, '诏安县', 1, '0596', '117.175184,23.711579', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1203, 350625, '长泰县', 1, '0596', '117.759153,24.625449', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1204, 350626, '东山县', 1, '0596', '117.430061,23.701262', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1281, 360700, '赣州市', 2, '0797', '114.933546,25.830694', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1205, 350627, '南靖县', 1, '0596', '117.35732,24.514654', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1206, 350628, '平和县', 1, '0596', '117.315017,24.363508', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1207, 350629, '华安县', 1, '0596', '117.534103,25.004425', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1208, 350681, '龙海市', 1, '0596', '117.818197,24.446706', 1197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1210, 350702, '延平区', 1, '0599', '118.182036,26.637438', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1211, 350703, '建阳区', 1, '0599', '118.120464,27.331876', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1212, 350721, '顺昌县', 1, '0599', '117.810357,26.793288', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1213, 350722, '浦城县', 1, '0599', '118.541256,27.917263', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1214, 350723, '光泽县', 1, '0599', '117.334106,27.540987', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1215, 350724, '松溪县', 1, '0599', '118.785468,27.526232', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1216, 350725, '政和县', 1, '0599', '118.857642,27.366104', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1217, 350781, '邵武市', 1, '0599', '117.492533,27.340326', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1218, 350782, '武夷山市', 1, '0599', '118.035309,27.756647', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1219, 350783, '建瓯市', 1, '0599', '118.304966,27.022774', 1209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1221, 350802, '新罗区', 1, '0597', '117.037155,25.098312', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1222, 350803, '永定区', 1, '0597', '116.732091,24.723961', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1223, 350821, '长汀县', 1, '0597', '116.357581,25.833531', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1224, 350823, '上杭县', 1, '0597', '116.420098,25.049518', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1300, 360800, '吉安市', 2, '0796', '114.966567,27.090763', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1225, 350824, '武平县', 1, '0597', '116.100414,25.095386', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1226, 350825, '连城县', 1, '0597', '116.754472,25.710538', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1227, 350881, '漳平市', 1, '0597', '117.419998,25.290184', 1220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1229, 350902, '蕉城区', 1, '0593', '119.526299,26.66061', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1230, 350921, '霞浦县', 1, '0593', '120.005146,26.885703', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1231, 350922, '古田县', 1, '0593', '118.746284,26.577837', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1232, 350923, '屏南县', 1, '0593', '118.985895,26.908276', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1233, 350924, '寿宁县', 1, '0593', '119.514986,27.454479', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1234, 350925, '周宁县', 1, '0593', '119.339025,27.104591', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1235, 350926, '柘荣县', 1, '0593', '119.900609,27.233933', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1237, 350982, '福鼎市', 1, '0593', '120.216977,27.324479', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1240, 360102, '东湖区', 1, '0791', '115.903526,28.698731', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1241, 360103, '西湖区', 1, '0791', '115.877233,28.657595', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1314, 360900, '宜春市', 2, '0795', '114.416785,27.815743', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1242, 360104, '青云谱区', 1, '0791', '115.925749,28.621169', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1243, 360105, '湾里区', 1, '0791', '115.730847,28.714796', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1244, 360111, '青山湖区', 1, '0791', '115.962144,28.682984', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1245, 360112, '新建区', 1, '0791', '115.815277,28.692864', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1246, 360121, '南昌县', 1, '0791', '115.933742,28.558296', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1247, 360123, '安义县', 1, '0791', '115.548658,28.846', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1248, 360124, '进贤县', 1, '0791', '116.241288,28.377343', 1239);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1250, 360202, '昌江区', 1, '0798', '117.18363,29.273565', 1249);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1251, 360203, '珠山区', 1, '0798', '117.202919,29.299938', 1249);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1252, 360222, '浮梁县', 1, '0798', '117.215066,29.352253', 1249);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1325, 361000, '抚州市', 2, '0794', '116.358181,27.949217', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1253, 360281, '乐平市', 1, '0798', '117.151796,28.97844', 1249);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1255, 360302, '安源区', 1, '0799', '113.870704,27.61511', 1254);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1256, 360313, '湘东区', 1, '0799', '113.733047,27.640075', 1254);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1257, 360321, '莲花县', 1, '0799', '113.961488,27.127664', 1254);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1258, 360322, '上栗县', 1, '0799', '113.795311,27.880301', 1254);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1331, 361025, '乐安县', 1, '0794', '115.83048,27.428765', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1259, 360323, '芦溪县', 1, '0799', '114.029827,27.630806', 1254);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1261, 360402, '濂溪区', 1, '0792', '115.992842,29.668064', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1262, 360403, '浔阳区', 1, '0792', '115.990301,29.727593', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1263, 360421, '九江县', 1, '0792', '115.911323,29.608431', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1337, 361100, '上饶市', 2, '0793', '117.943433,28.454863', 1238);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1264, 360423, '武宁县', 1, '0792', '115.092757,29.246591', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1265, 360424, '修水县', 1, '0792', '114.546836,29.025726', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1266, 360425, '永修县', 1, '0792', '115.831956,29.011871', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1267, 360426, '德安县', 1, '0792', '115.767447,29.298696', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1268, 360483, '庐山市', 1, '0792', '116.04506,29.448128', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1269, 360428, '都昌县', 1, '0792', '116.203979,29.273239', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1270, 360429, '湖口县', 1, '0792', '116.251947,29.731101', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1271, 360430, '彭泽县', 1, '0792', '116.56438,29.876991', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1272, 360481, '瑞昌市', 1, '0792', '115.681335,29.675834', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1273, 360482, '共青城市', 1, '0792', '115.808844,29.248316', 1260);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1275, 360502, '渝水区', 1, '0790', '114.944549,27.800148', 1274);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1276, 360521, '分宜县', 1, '0790', '114.692049,27.814757', 1274);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1278, 360602, '月湖区', 1, '0701', '117.102475,28.267018', 1277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1351, 370100, '济南市', 2, '0531', '117.120098,36.6512', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1279, 360622, '余江县', 1, '0701', '116.85926,28.198652', 1277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1280, 360681, '贵溪市', 1, '0701', '117.245497,28.292519', 1277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1282, 360702, '章贡区', 1, '0797', '114.921171,25.817816', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1283, 360703, '南康区', 1, '0797', '114.765412,25.66145', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1284, 360704, '赣县区', 1, '0797', '115.011561,25.86069', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1285, 360722, '信丰县', 1, '0797', '114.922922,25.386379', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1286, 360723, '大余县', 1, '0797', '114.362112,25.401313', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1287, 360724, '上犹县', 1, '0797', '114.551138,25.785172', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1288, 360725, '崇义县', 1, '0797', '114.308267,25.681784', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1289, 360726, '安远县', 1, '0797', '115.393922,25.136927', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1362, 370200, '青岛市', 2, '0532', '120.382621,36.067131', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1290, 360727, '龙南县', 1, '0797', '114.789873,24.911069', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1291, 360728, '定南县', 1, '0797', '115.027845,24.78441', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1292, 360729, '全南县', 1, '0797', '114.530125,24.742403', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1293, 360730, '宁都县', 1, '0797', '116.009472,26.470116', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1294, 360731, '于都县', 1, '0797', '115.415508,25.952068', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1295, 360732, '兴国县', 1, '0797', '115.363189,26.337937', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1296, 360733, '会昌县', 1, '0797', '115.786056,25.600272', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1297, 360734, '寻乌县', 1, '0797', '115.637933,24.969167', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1298, 360735, '石城县', 1, '0797', '116.346995,26.314775', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1299, 360781, '瑞金市', 1, '0797', '116.027134,25.885555', 1281);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1373, 370300, '淄博市', 2, '0533', '118.055019,36.813546', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1301, 360802, '吉州区', 1, '0796', '114.994763,27.143801', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1302, 360803, '青原区', 1, '0796', '115.014811,27.081977', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1303, 360821, '吉安县', 1, '0796', '114.907875,27.039787', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1304, 360822, '吉水县', 1, '0796', '115.135507,27.229632', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1305, 360823, '峡江县', 1, '0796', '115.316566,27.582901', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1306, 360824, '新干县', 1, '0796', '115.387052,27.740191', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1307, 360825, '永丰县', 1, '0796', '115.421344,27.316939', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1308, 360826, '泰和县', 1, '0796', '114.92299,26.801628', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1382, 370400, '枣庄市', 2, '0632', '117.323725,34.810488', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1309, 360827, '遂川县', 1, '0796', '114.520537,26.313737', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1310, 360828, '万安县', 1, '0796', '114.759364,26.456553', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1311, 360829, '安福县', 1, '0796', '114.619893,27.392873', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1312, 360830, '永新县', 1, '0796', '114.243072,26.944962', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1313, 360881, '井冈山市', 1, '0796', '114.289228,26.748081', 1300);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1315, 360902, '袁州区', 1, '0795', '114.427858,27.797091', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1389, 370500, '东营市', 2, '0546', '118.674614,37.433963', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1316, 360921, '奉新县', 1, '0795', '115.400491,28.688423', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1317, 360922, '万载县', 1, '0795', '114.444854,28.105689', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1318, 360923, '上高县', 1, '0795', '114.947683,28.238061', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1319, 360924, '宜丰县', 1, '0795', '114.802852,28.394565', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1320, 360925, '靖安县', 1, '0795', '115.362628,28.861478', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1395, 370600, '烟台市', 2, '0535', '121.447852,37.464539', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1321, 360926, '铜鼓县', 1, '0795', '114.371172,28.520769', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1322, 360981, '丰城市', 1, '0795', '115.771093,28.159141', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1323, 360982, '樟树市', 1, '0795', '115.546152,28.055853', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1324, 360983, '高安市', 1, '0795', '115.360619,28.441152', 1314);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1326, 361002, '临川区', 1, '0794', '116.312166,27.934572', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1327, 361021, '南城县', 1, '0794', '116.63704,27.569678', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1328, 361022, '黎川县', 1, '0794', '116.907681,27.282333', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1329, 361023, '南丰县', 1, '0794', '116.525725,27.218444', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1330, 361024, '崇仁县', 1, '0794', '116.07626,27.754466', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1332, 361026, '宜黄县', 1, '0794', '116.236201,27.554886', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1333, 361027, '金溪县', 1, '0794', '116.755058,27.918959', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1334, 361028, '资溪县', 1, '0794', '117.060263,27.706101', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1408, 370700, '潍坊市', 2, '0536', '119.161748,36.706962', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1335, 361003, '东乡区', 1, '0794', '116.603559,28.247696', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1336, 361030, '广昌县', 1, '0794', '116.335686,26.843684', 1325);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1338, 361102, '信州区', 1, '0793', '117.966268,28.431006', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1339, 361103, '广丰区', 1, '0793', '118.19124,28.436285', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1340, 361121, '上饶县', 1, '0793', '117.907849,28.448982', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1341, 361123, '玉山县', 1, '0793', '118.244769,28.682309', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1342, 361124, '铅山县', 1, '0793', '117.709659,28.315664', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1343, 361125, '横峰县', 1, '0793', '117.596452,28.407117', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1344, 361126, '弋阳县', 1, '0793', '117.449588,28.378044', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1345, 361127, '余干县', 1, '0793', '116.695646,28.702302', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1421, 370800, '济宁市', 2, '0537', '116.587282,35.414982', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1420, 370786, '昌邑市', 1, '0536', '119.403069,36.843319', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1346, 361128, '鄱阳县', 1, '0793', '116.70359,29.004847', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1347, 361129, '万年县', 1, '0793', '117.058445,28.694582', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1348, 361130, '婺源县', 1, '0793', '117.861797,29.248085', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1349, 361181, '德兴市', 1, '0793', '117.578713,28.946464', 1337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1352, 370102, '历下区', 1, '0531', '117.076441,36.666465', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1353, 370103, '市中区', 1, '0531', '116.997845,36.651335', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1354, 370104, '槐荫区', 1, '0531', '116.901224,36.651441', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1355, 370105, '天桥区', 1, '0531', '116.987153,36.678589', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1356, 370112, '历城区', 1, '0531', '117.06523,36.680259', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1357, 370113, '长清区', 1, '0531', '116.751843,36.55371', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1433, 370900, '泰安市', 2, '0538', '117.087614,36.200252', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1358, 370124, '平阴县', 1, '0531', '116.456006,36.289251', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1359, 370125, '济阳县', 1, '0531', '117.173524,36.978537', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1360, 370126, '商河县', 1, '0531', '117.157232,37.309041', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1361, 370114, '章丘区', 1, '0531', '117.526228,36.681258', 1351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1363, 370202, '市南区', 1, '0532', '120.412392,36.075651', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1364, 370203, '市北区', 1, '0532', '120.374701,36.0876', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1440, 371000, '威海市', 2, '0631', '122.120282,37.513412', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1365, 370211, '黄岛区', 1, '0532', '120.198055,35.960933', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1366, 370212, '崂山区', 1, '0532', '120.468956,36.107538', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1367, 370213, '李沧区', 1, '0532', '120.432922,36.145519', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1368, 370214, '城阳区', 1, '0532', '120.396256,36.307559', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1445, 371100, '日照市', 2, '0633', '119.526925,35.416734', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1369, 370281, '胶州市', 1, '0532', '120.033382,36.26468', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1370, 370282, '即墨市', 1, '0532', '120.447158,36.389408', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1371, 370283, '平度市', 1, '0532', '119.98842,36.776357', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1372, 370285, '莱西市', 1, '0532', '120.51769,36.889084', 1362);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1450, 371200, '莱芜市', 2, '0634', '117.676723,36.213813', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1374, 370302, '淄川区', 1, '0533', '117.966723,36.643452', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1375, 370303, '张店区', 1, '0533', '118.017938,36.806669', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1453, 371300, '临沂市', 2, '0539', '118.356414,35.104673', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1376, 370304, '博山区', 1, '0533', '117.861851,36.494701', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1377, 370305, '临淄区', 1, '0533', '118.309118,36.826981', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1378, 370306, '周村区', 1, '0533', '117.869886,36.803072', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1379, 370321, '桓台县', 1, '0533', '118.097922,36.959804', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1380, 370322, '高青县', 1, '0533', '117.826924,37.170979', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1381, 370323, '沂源县', 1, '0533', '118.170855,36.185038', 1373);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1383, 370402, '市中区', 1, '0632', '117.556139,34.863554', 1382);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1384, 370403, '薛城区', 1, '0632', '117.263164,34.795062', 1382);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1385, 370404, '峄城区', 1, '0632', '117.590816,34.773263', 1382);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1386, 370405, '台儿庄区', 1, '0632', '117.734414,34.56244', 1382);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1387, 370406, '山亭区', 1, '0632', '117.461517,35.099528', 1382);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1388, 370481, '滕州市', 1, '0632', '117.165824,35.114155', 1382);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1466, 371400, '德州市', 2, '0534', '116.359381,37.436657', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1390, 370502, '东营区', 1, '0546', '118.582184,37.448964', 1389);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1391, 370503, '河口区', 1, '0546', '118.525543,37.886162', 1389);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1392, 370505, '垦利区', 1, '0546', '118.575228,37.573054', 1389);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1393, 370522, '利津县', 1, '0546', '118.255287,37.490328', 1389);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1394, 370523, '广饶县', 1, '0546', '118.407107,37.053555', 1389);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1396, 370602, '芝罘区', 1, '0535', '121.400445,37.541475', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1397, 370611, '福山区', 1, '0535', '121.267741,37.498246', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1398, 370612, '牟平区', 1, '0535', '121.600455,37.387061', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1399, 370613, '莱山区', 1, '0535', '121.445301,37.511291', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1400, 370634, '长岛县', 1, '0535', '120.73658,37.921368', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1401, 370681, '龙口市', 1, '0535', '120.477813,37.646107', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1478, 371500, '聊城市', 2, '0635', '115.985389,36.456684', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1402, 370682, '莱阳市', 1, '0535', '120.711672,36.978941', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1403, 370683, '莱州市', 1, '0535', '119.942274,37.177129', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1404, 370684, '蓬莱市', 1, '0535', '120.758848,37.810661', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1405, 370685, '招远市', 1, '0535', '120.434071,37.355469', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1406, 370686, '栖霞市', 1, '0535', '120.849675,37.335123', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1407, 370687, '海阳市', 1, '0535', '121.173793,36.688', 1395);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1409, 370702, '潍城区', 1, '0536', '119.024835,36.7281', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1410, 370703, '寒亭区', 1, '0536', '119.211157,36.755623', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1487, 371600, '滨州市', 2, '0543', '117.970699,37.38198', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1411, 370704, '坊子区', 1, '0536', '119.166485,36.654448', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1412, 370705, '奎文区', 1, '0536', '119.132482,36.70759', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1413, 370724, '临朐县', 1, '0536', '118.542982,36.512506', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1414, 370725, '昌乐县', 1, '0536', '118.829992,36.706964', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1415, 370781, '青州市', 1, '0536', '118.479654,36.684789', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1416, 370782, '诸城市', 1, '0536', '119.410103,35.995654', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1417, 370783, '寿光市', 1, '0536', '118.790739,36.85576', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1495, 371700, '菏泽市', 2, '0530', '115.480656,35.23375', 1350);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1418, 370784, '安丘市', 1, '0536', '119.218978,36.478493', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1419, 370785, '高密市', 1, '0536', '119.755597,36.382594', 1408);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1422, 370811, '任城区', 1, '0537', '116.606103,35.444028', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1423, 370812, '兖州区', 1, '0537', '116.783833,35.553144', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1424, 370826, '微山县', 1, '0537', '117.128827,34.806554', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1425, 370827, '鱼台县', 1, '0537', '116.650608,35.012749', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1426, 370828, '金乡县', 1, '0537', '116.311532,35.066619', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1427, 370829, '嘉祥县', 1, '0537', '116.342449,35.408824', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1428, 370830, '汶上县', 1, '0537', '116.49708,35.712298', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1429, 370831, '泗水县', 1, '0537', '117.251195,35.664323', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1506, 410100, '郑州市', 2, '0371', '113.625328,34.746611', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1430, 370832, '梁山县', 1, '0537', '116.096044,35.802306', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1431, 370881, '曲阜市', 1, '0537', '116.986526,35.581108', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1432, 370883, '邹城市', 1, '0537', '117.007453,35.40268', 1421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1434, 370902, '泰山区', 1, '0538', '117.135354,36.192083', 1433);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1435, 370911, '岱岳区', 1, '0538', '117.041581,36.187989', 1433);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1436, 370921, '宁阳县', 1, '0538', '116.805796,35.758786', 1433);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1437, 370923, '东平县', 1, '0538', '116.470304,35.937102', 1433);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1438, 370982, '新泰市', 1, '0538', '117.767952,35.909032', 1433);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1439, 370983, '肥城市', 1, '0538', '116.768358,36.182571', 1433);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1441, 371002, '环翠区', 1, '0631', '122.123443,37.50199', 1440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1442, 371003, '文登区', 1, '0631', '122.05767,37.193735', 1440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1519, 410200, '开封市', 2, '0378', '114.307677,34.797966', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1443, 371082, '荣成市', 1, '0631', '122.486657,37.16516', 1440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1444, 371083, '乳山市', 1, '0631', '121.539764,36.919816', 1440);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1446, 371102, '东港区', 1, '0633', '119.462267,35.42548', 1445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1447, 371103, '岚山区', 1, '0633', '119.318928,35.121884', 1445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1448, 371121, '五莲县', 1, '0633', '119.213619,35.760228', 1445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1449, 371122, '莒县', 1, '0633', '118.837063,35.579868', 1445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1451, 371202, '莱城区', 1, '0634', '117.659884,36.203179', 1450);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1452, 371203, '钢城区', 1, '0634', '117.811354,36.058572', 1450);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1454, 371302, '兰山区', 1, '0539', '118.347842,35.051804', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1529, 410300, '洛阳市', 2, '0379', '112.453926,34.620202', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1455, 371311, '罗庄区', 1, '0539', '118.284786,34.996741', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1456, 371312, '河东区', 1, '0539', '118.402893,35.089916', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1457, 371321, '沂南县', 1, '0539', '118.465221,35.550217', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1458, 371322, '郯城县', 1, '0539', '118.367215,34.613586', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1459, 371323, '沂水县', 1, '0539', '118.627917,35.79045', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1460, 371324, '兰陵县', 1, '0539', '118.07065,34.857149', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1461, 371325, '费县', 1, '0539', '117.977325,35.26596', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1462, 371326, '平邑县', 1, '0539', '117.640352,35.505943', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1463, 371327, '莒南县', 1, '0539', '118.835163,35.174846', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1464, 371328, '蒙阴县', 1, '0539', '117.953621,35.719396', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1465, 371329, '临沭县', 1, '0539', '118.650781,34.919851', 1453);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1467, 371402, '德城区', 1, '0534', '116.29947,37.450804', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1468, 371403, '陵城区', 1, '0534', '116.576092,37.335794', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1469, 371422, '宁津县', 1, '0534', '116.800306,37.652189', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1470, 371423, '庆云县', 1, '0534', '117.385256,37.775349', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1545, 410400, '平顶山市', 2, '0375', '113.192661,33.766169', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1471, 371424, '临邑县', 1, '0534', '116.866799,37.189797', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1472, 371425, '齐河县', 1, '0534', '116.762893,36.784158', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1473, 371426, '平原县', 1, '0534', '116.434032,37.165323', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1474, 371427, '夏津县', 1, '0534', '116.001726,36.948371', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1475, 371428, '武城县', 1, '0534', '116.069302,37.213311', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1476, 371481, '乐陵市', 1, '0534', '117.231934,37.729907', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1477, 371482, '禹城市', 1, '0534', '116.638327,36.933812', 1466);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1479, 371502, '东昌府区', 1, '0635', '115.988349,36.434669', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1480, 371521, '阳谷县', 1, '0635', '115.79182,36.114392', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1481, 371522, '莘县', 1, '0635', '115.671191,36.233598', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1556, 410500, '安阳市', 2, '0372', '114.392392,36.097577', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1482, 371523, '茌平县', 1, '0635', '116.25527,36.580688', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1483, 371524, '东阿县', 1, '0635', '116.247579,36.334917', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1484, 371525, '冠县', 1, '0635', '115.442739,36.484009', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1485, 371526, '高唐县', 1, '0635', '116.23016,36.846762', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1486, 371581, '临清市', 1, '0635', '115.704881,36.838277', 1478);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1488, 371602, '滨城区', 1, '0543', '118.019326,37.430724', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1489, 371603, '沾化区', 1, '0543', '118.098902,37.69926', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1490, 371621, '惠民县', 1, '0543', '117.509921,37.489877', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1491, 371622, '阳信县', 1, '0543', '117.603339,37.632433', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1566, 410600, '鹤壁市', 2, '0392', '114.297309,35.748325', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1492, 371623, '无棣县', 1, '0543', '117.625696,37.77026', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1493, 371625, '博兴县', 1, '0543', '118.110709,37.15457', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1494, 371626, '邹平县', 1, '0543', '117.743109,36.862989', 1487);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1496, 371702, '牡丹区', 1, '0530', '115.417826,35.252512', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1497, 371703, '定陶区', 1, '0530', '115.57302,35.070995', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1572, 410700, '新乡市', 2, '0373', '113.926763,35.303704', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1499, 371722, '单县', 1, '0530', '116.107428,34.778808', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1500, 371723, '成武县', 1, '0530', '115.889764,34.952459', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1501, 371724, '巨野县', 1, '0530', '116.062394,35.388925', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1502, 371725, '郓城县', 1, '0530', '115.9389,35.575135', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1503, 371726, '鄄城县', 1, '0530', '115.510192,35.563408', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1504, 371728, '东明县', 1, '0530', '115.107404,35.276162', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1507, 410102, '中原区', 1, '0371', '113.613337,34.748256', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1508, 410103, '二七区', 1, '0371', '113.640211,34.724114', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1509, 410104, '管城回族区', 1, '0371', '113.6775,34.75429', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1510, 410105, '金水区', 1, '0371', '113.660617,34.800004', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1511, 410106, '上街区', 1, '0371', '113.30893,34.802752', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1512, 410108, '惠济区', 1, '0371', '113.6169,34.867457', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1585, 410800, '焦作市', 2, '0391', '113.241823,35.215893', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1513, 410122, '中牟县', 1, '0371', '113.976253,34.718936', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1514, 410181, '巩义市', 1, '0371', '113.022406,34.7481', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1515, 410182, '荥阳市', 1, '0371', '113.38324,34.786948', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1516, 410183, '新密市', 1, '0371', '113.391087,34.539376', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1517, 410184, '新郑市', 1, '0371', '113.740662,34.395949', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1518, 410185, '登封市', 1, '0371', '113.050581,34.454443', 1506);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1520, 410202, '龙亭区', 1, '0378', '114.356076,34.815565', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1521, 410203, '顺河回族区', 1, '0378', '114.364875,34.800458', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1522, 410204, '鼓楼区', 1, '0378', '114.348306,34.78856', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1523, 410205, '禹王台区', 1, '0378', '114.34817,34.777104', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1596, 410900, '濮阳市', 2, '0393', '115.029216,35.761829', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1524, 410212, '祥符区', 1, '0378', '114.441285,34.756916', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1598, 410922, '清丰县', 1, '0393', '115.104389,35.88518', 1596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1525, 410221, '杞县', 1, '0378', '114.783139,34.549174', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1526, 410222, '通许县', 1, '0378', '114.467467,34.480433', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1527, 410223, '尉氏县', 1, '0378', '114.193081,34.411494', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1603, 411000, '许昌市', 2, '0374', '113.852454,34.035771', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1528, 410225, '兰考县', 1, '0378', '114.821348,34.822211', 1519);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1530, 410302, '老城区', 1, '0379', '112.469766,34.6842', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1531, 410303, '西工区', 1, '0379', '112.427914,34.660378', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1532, 410304, '瀍河回族区', 1, '0379', '112.500131,34.679773', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1533, 410305, '涧西区', 1, '0379', '112.395756,34.658033', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1534, 410306, '吉利区', 1, '0379', '112.589112,34.900467', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1610, 411100, '漯河市', 2, '0395', '114.016536,33.580873', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1535, 410311, '洛龙区', 1, '0379', '112.463833,34.619711', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1536, 410322, '孟津县', 1, '0379', '112.445354,34.825638', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1537, 410323, '新安县', 1, '0379', '112.13244,34.728284', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1538, 410324, '栾川县', 1, '0379', '111.615768,33.785698', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1539, 410325, '嵩县', 1, '0379', '112.085634,34.134516', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1616, 411200, '三门峡市', 2, '0398', '111.200367,34.772792', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1540, 410326, '汝阳县', 1, '0379', '112.473139,34.153939', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1541, 410327, '宜阳县', 1, '0379', '112.179238,34.514644', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1542, 410328, '洛宁县', 1, '0379', '111.653111,34.389197', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1543, 410329, '伊川县', 1, '0379', '112.425676,34.421323', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1544, 410381, '偃师市', 1, '0379', '112.789534,34.72722', 1529);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1546, 410402, '新华区', 1, '0375', '113.293977,33.737251', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1623, 411300, '南阳市', 2, '0377', '112.528308,32.990664', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1547, 410403, '卫东区', 1, '0375', '113.335192,33.734706', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1548, 410404, '石龙区', 1, '0375', '112.898818,33.898713', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1549, 410411, '湛河区', 1, '0375', '113.320873,33.725681', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1550, 410421, '宝丰县', 1, '0375', '113.054801,33.868434', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1551, 410422, '叶县', 1, '0375', '113.357239,33.626731', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1552, 410423, '鲁山县', 1, '0375', '112.908202,33.738293', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1553, 410425, '郏县', 1, '0375', '113.212609,33.971787', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1554, 410481, '舞钢市', 1, '0375', '113.516343,33.314033', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1555, 410482, '汝州市', 1, '0375', '112.844517,34.167029', 1545);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1557, 410502, '文峰区', 1, '0372', '114.357082,36.090468', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1558, 410503, '北关区', 1, '0372', '114.355742,36.10766', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1559, 410505, '殷都区', 1, '0372', '114.303553,36.10989', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1560, 410506, '龙安区', 1, '0372', '114.301331,36.076225', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1637, 411400, '商丘市', 2, '0370', '115.656339,34.414961', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1561, 410522, '安阳县', 1, '0372', '114.130207,36.130584', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1562, 410523, '汤阴县', 1, '0372', '114.357763,35.924514', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1563, 410526, '滑县', 1, '0372', '114.519311,35.575417', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1564, 410527, '内黄县', 1, '0372', '114.901452,35.971704', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1565, 410581, '林州市', 1, '0372', '113.820129,36.083046', 1556);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1567, 410602, '鹤山区', 1, '0392', '114.163258,35.954611', 1566);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1568, 410603, '山城区', 1, '0392', '114.184318,35.898033', 1566);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1569, 410611, '淇滨区', 1, '0392', '114.298789,35.741592', 1566);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1570, 410621, '浚县', 1, '0392', '114.55091,35.67636', 1566);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1647, 411500, '信阳市', 2, '0376', '114.091193,32.147679', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1571, 410622, '淇县', 1, '0392', '114.208828,35.622507', 1566);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1573, 410702, '红旗区', 1, '0373', '113.875245,35.30385', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1574, 410703, '卫滨区', 1, '0373', '113.865663,35.301992', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1575, 410704, '凤泉区', 1, '0373', '113.915184,35.383978', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1576, 410711, '牧野区', 1, '0373', '113.908772,35.315039', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1577, 410721, '新乡县', 1, '0373', '113.805205,35.190836', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1578, 410724, '获嘉县', 1, '0373', '113.657433,35.259808', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1579, 410725, '原阳县', 1, '0373', '113.940046,35.065587', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1580, 410726, '延津县', 1, '0373', '114.20509,35.141889', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1581, 410727, '封丘县', 1, '0373', '114.418882,35.041198', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1658, 411600, '周口市', 2, '0394', '114.69695,33.626149', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1582, 410728, '长垣县', 1, '0373', '114.668936,35.201548', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1583, 410781, '卫辉市', 1, '0373', '114.064907,35.398494', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1584, 410782, '辉县市', 1, '0373', '113.805468,35.462312', 1572);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1586, 410802, '解放区', 1, '0391', '113.230816,35.240282', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1587, 410803, '中站区', 1, '0391', '113.182946,35.236819', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1588, 410804, '马村区', 1, '0391', '113.322332,35.256108', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1589, 410811, '山阳区', 1, '0391', '113.254881,35.214507', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1590, 410821, '修武县', 1, '0391', '113.447755,35.223514', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1591, 410822, '博爱县', 1, '0391', '113.064379,35.171045', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1592, 410823, '武陟县', 1, '0391', '113.401679,35.099378', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1669, 411700, '驻马店市', 2, '0396', '114.022247,33.012885', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1593, 410825, '温县', 1, '0391', '113.08053,34.940189', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1594, 410882, '沁阳市', 1, '0391', '112.950716,35.087539', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1595, 410883, '孟州市', 1, '0391', '112.791401,34.907315', 1585);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1597, 410902, '华龙区', 1, '0393', '115.074151,35.777346', 1596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1599, 410923, '南乐县', 1, '0393', '115.204675,36.069476', 1596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1600, 410926, '范县', 1, '0393', '115.504201,35.851906', 1596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1601, 410927, '台前县', 1, '0393', '115.871906,35.96939', 1596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1602, 410928, '濮阳县', 1, '0393', '115.029078,35.712193', 1596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1604, 411002, '魏都区', 1, '0374', '113.822647,34.025341', 1603);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1605, 411003, '建安区', 1, '0374', '113.822983,34.12466', 1603);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1680, 419001, '济源市', 2, '1391', '112.602256,35.067199', 1505);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1606, 411024, '鄢陵县', 1, '0374', '114.177399,34.102332', 1603);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1682, 420100, '武汉市', 2, '027', '114.305469,30.593175', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1607, 411025, '襄城县', 1, '0374', '113.505874,33.851459', 1603);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1608, 411081, '禹州市', 1, '0374', '113.488478,34.140701', 1603);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1609, 411082, '长葛市', 1, '0374', '113.813714,34.19592', 1603);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1611, 411102, '源汇区', 1, '0395', '114.017948,33.565441', 1610);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1612, 411103, '郾城区', 1, '0395', '114.006943,33.587409', 1610);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1613, 411104, '召陵区', 1, '0395', '114.093902,33.586565', 1610);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1614, 411121, '舞阳县', 1, '0395', '113.609286,33.437876', 1610);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1615, 411122, '临颍县', 1, '0395', '113.931261,33.828042', 1610);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1617, 411202, '湖滨区', 1, '0398', '111.188397,34.770886', 1616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1618, 411203, '陕州区', 1, '0398', '111.103563,34.720547', 1616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1619, 411221, '渑池县', 1, '0398', '111.761797,34.767951', 1616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1620, 411224, '卢氏县', 1, '0398', '111.047858,34.054324', 1616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1621, 411281, '义马市', 1, '0398', '111.87448,34.7474', 1616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1696, 420200, '黄石市', 2, '0714', '115.038962,30.201038', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1622, 411282, '灵宝市', 1, '0398', '110.89422,34.516828', 1616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1624, 411302, '宛城区', 1, '0377', '112.539558,33.003784', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1625, 411303, '卧龙区', 1, '0377', '112.528789,32.989877', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1626, 411321, '南召县', 1, '0377', '112.429133,33.489877', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1627, 411322, '方城县', 1, '0377', '113.012494,33.254391', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1628, 411323, '西峡县', 1, '0377', '111.47353,33.307294', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1703, 420300, '十堰市', 2, '0719', '110.799291,32.629462', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1629, 411324, '镇平县', 1, '0377', '112.234697,33.03411', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1630, 411325, '内乡县', 1, '0377', '111.849392,33.044864', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1631, 411326, '淅川县', 1, '0377', '111.490964,33.13782', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1632, 411327, '社旗县', 1, '0377', '112.948245,33.056109', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1633, 411328, '唐河县', 1, '0377', '112.807636,32.681335', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1634, 411329, '新野县', 1, '0377', '112.360026,32.520805', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1635, 411330, '桐柏县', 1, '0377', '113.428287,32.380073', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1636, 411381, '邓州市', 1, '0377', '112.087493,32.68758', 1623);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1712, 420500, '宜昌市', 2, '0717', '111.286445,30.691865', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1638, 411402, '梁园区', 1, '0370', '115.613965,34.443893', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1639, 411403, '睢阳区', 1, '0370', '115.653301,34.388389', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1640, 411421, '民权县', 1, '0370', '115.173971,34.648191', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1641, 411422, '睢县', 1, '0370', '115.071879,34.445655', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1642, 411423, '宁陵县', 1, '0370', '115.313743,34.460399', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1643, 411424, '柘城县', 1, '0370', '115.305708,34.091082', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1644, 411425, '虞城县', 1, '0370', '115.828319,34.400835', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1645, 411426, '夏邑县', 1, '0370', '116.131447,34.237553', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1646, 411481, '永城市', 1, '0370', '116.4495,33.929291', 1637);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1648, 411502, '浉河区', 1, '0376', '114.058713,32.116803', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1649, 411503, '平桥区', 1, '0376', '114.125656,32.101031', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1650, 411521, '罗山县', 1, '0376', '114.512872,32.203883', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1651, 411522, '光山县', 1, '0376', '114.919152,32.010002', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1726, 420600, '襄阳市', 2, '0710', '112.122426,32.009016', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1652, 411523, '新县', 1, '0376', '114.879239,31.643918', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1653, 411524, '商城县', 1, '0376', '115.406862,31.798377', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1654, 411525, '固始县', 1, '0376', '115.654481,32.168137', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1655, 411526, '潢川县', 1, '0376', '115.051908,32.131522', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1656, 411527, '淮滨县', 1, '0376', '115.419537,32.473258', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1657, 411528, '息县', 1, '0376', '114.740456,32.342792', 1647);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1659, 411602, '川汇区', 1, '0394', '114.650628,33.647598', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1660, 411621, '扶沟县', 1, '0394', '114.394821,34.059968', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1661, 411622, '西华县', 1, '0394', '114.529756,33.767407', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1736, 420700, '鄂州市', 2, '0711', '114.894935,30.391141', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1662, 411623, '商水县', 1, '0394', '114.611651,33.542138', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1663, 411624, '沈丘县', 1, '0394', '115.098583,33.409369', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1664, 411625, '郸城县', 1, '0394', '115.177188,33.644743', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1740, 420800, '荆门市', 2, '0724', '112.199427,31.035395', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1665, 411626, '淮阳县', 1, '0394', '114.886153,33.731561', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1666, 411627, '太康县', 1, '0394', '114.837888,34.064463', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1667, 411628, '鹿邑县', 1, '0394', '115.484454,33.86', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1668, 411681, '项城市', 1, '0394', '114.875333,33.465838', 1658);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1670, 411702, '驿城区', 1, '0396', '113.993914,32.973054', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1746, 420900, '孝感市', 2, '0712', '113.957037,30.917766', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1671, 411721, '西平县', 1, '0396', '114.021538,33.387684', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1672, 411722, '上蔡县', 1, '0396', '114.264381,33.262439', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1673, 411723, '平舆县', 1, '0396', '114.619159,32.96271', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1674, 411724, '正阳县', 1, '0396', '114.392773,32.605697', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1675, 411725, '确山县', 1, '0396', '114.026429,32.802064', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1676, 411726, '泌阳县', 1, '0396', '113.327144,32.723975', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1677, 411727, '汝南县', 1, '0396', '114.362379,33.006729', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1754, 421000, '荆州市', 2, '0716', '112.239746,30.335184', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1678, 411728, '遂平县', 1, '0396', '114.013182,33.145649', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1679, 411729, '新蔡县', 1, '0396', '114.96547,32.744896', 1669);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1683, 420102, '江岸区', 1, '027', '114.30911,30.600052', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1684, 420103, '江汉区', 1, '027', '114.270867,30.601475', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1685, 420104, '硚口区', 1, '027', '114.21492,30.582202', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1686, 420105, '汉阳区', 1, '027', '114.21861,30.553983', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1687, 420106, '武昌区', 1, '027', '114.31665,30.554408', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1688, 420107, '青山区', 1, '027', '114.384968,30.640191', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1763, 421100, '黄冈市', 2, '0713', '114.872199,30.453667', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1689, 420111, '洪山区', 1, '027', '114.343796,30.500247', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1690, 420112, '东西湖区', 1, '027', '114.137116,30.619917', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1691, 420113, '汉南区', 1, '027', '114.084597,30.308829', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1692, 420114, '蔡甸区', 1, '027', '114.087285,30.536454', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1693, 420115, '江夏区', 1, '027', '114.319097,30.376308', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1694, 420116, '黄陂区', 1, '027', '114.375725,30.882174', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1695, 420117, '新洲区', 1, '027', '114.801096,30.841425', 1682);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1697, 420202, '黄石港区', 1, '0714', '115.065849,30.222938', 1696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1698, 420203, '西塞山区', 1, '0714', '115.109955,30.204924', 1696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1699, 420204, '下陆区', 1, '0714', '114.961327,30.173912', 1696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1774, 421200, '咸宁市', 2, '0715', '114.322616,29.841362', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1700, 420205, '铁山区', 1, '0714', '114.891605,30.203118', 1696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1701, 420222, '阳新县', 1, '0714', '115.215227,29.830257', 1696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1702, 420281, '大冶市', 1, '0714', '114.980424,30.096147', 1696);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1704, 420302, '茅箭区', 1, '0719', '110.813719,32.591904', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1705, 420303, '张湾区', 1, '0719', '110.769132,32.652297', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1706, 420304, '郧阳区', 1, '0719', '110.81205,32.834775', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1781, 421300, '随州市', 2, '0722', '113.382515,31.690191', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1707, 420322, '郧西县', 1, '0719', '110.425983,32.993182', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1708, 420323, '竹山县', 1, '0719', '110.228747,32.224808', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1709, 420324, '竹溪县', 1, '0719', '109.715304,32.318255', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1785, 422800, '恩施土家族苗族自治州', 2, '0718', '109.488172,30.272156', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1710, 420325, '房县', 1, '0719', '110.733181,32.050378', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1711, 420381, '丹江口市', 1, '0719', '111.513127,32.540157', 1703);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1713, 420502, '西陵区', 1, '0717', '111.285646,30.710781', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1714, 420503, '伍家岗区', 1, '0717', '111.361037,30.644334', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1715, 420504, '点军区', 1, '0717', '111.268119,30.693247', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1716, 420505, '猇亭区', 1, '0717', '111.43462,30.530903', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1717, 420506, '夷陵区', 1, '0717', '111.32638,30.770006', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1718, 420525, '远安县', 1, '0717', '111.640508,31.060869', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1794, 429005, '潜江市', 2, '2728', '112.899762,30.402167', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1795, 429021, '神农架林区', 2, '1719', '110.675743,31.744915', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1796, 429006, '天门市', 2, '1728', '113.166078,30.663337', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1797, 429004, '仙桃市', 2, '0728', '113.423583,30.361438', 1681);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1799, 430100, '长沙市', 2, '0731', '112.938884,28.22808', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1719, 420526, '兴山县', 1, '0717', '110.746804,31.348196', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1720, 420527, '秭归县', 1, '0717', '110.977711,30.825897', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1721, 420528, '长阳土家族自治县', 1, '0717', '111.207242,30.472763', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1722, 420529, '五峰土家族自治县', 1, '0717', '111.07374,30.156741', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1723, 420581, '宜都市', 1, '0717', '111.450096,30.378299', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1724, 420582, '当阳市', 1, '0717', '111.788312,30.821266', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1725, 420583, '枝江市', 1, '0717', '111.76053,30.42594', 1712);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1727, 420602, '襄城区', 1, '0710', '112.134052,32.010366', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1728, 420606, '樊城区', 1, '0710', '112.135684,32.044832', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1809, 430200, '株洲市', 2, '0733', '113.133853,27.827986', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1729, 420607, '襄州区', 1, '0710', '112.211982,32.087127', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1730, 420624, '南漳县', 1, '0710', '111.838905,31.774636', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1731, 420625, '谷城县', 1, '0710', '111.652982,32.263849', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1732, 420626, '保康县', 1, '0710', '111.261308,31.87831', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1733, 420682, '老河口市', 1, '0710', '111.683861,32.359068', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1734, 420683, '枣阳市', 1, '0710', '112.771959,32.128818', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1735, 420684, '宜城市', 1, '0710', '112.257788,31.719806', 1726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1737, 420702, '梁子湖区', 1, '0711', '114.684731,30.100141', 1736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1738, 420703, '华容区', 1, '0711', '114.729878,30.534309', 1736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1819, 430300, '湘潭市', 2, '0732', '112.944026,27.829795', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1739, 420704, '鄂城区', 1, '0711', '114.891586,30.400651', 1736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1741, 420802, '东宝区', 1, '0724', '112.201493,31.051852', 1740);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1742, 420804, '掇刀区', 1, '0724', '112.207962,30.973451', 1740);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1743, 420821, '京山县', 1, '0724', '113.119566,31.018457', 1740);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1744, 420822, '沙洋县', 1, '0724', '112.588581,30.709221', 1740);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1825, 430400, '衡阳市', 2, '0734', '112.572018,26.893368', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1745, 420881, '钟祥市', 1, '0724', '112.58812,31.167819', 1740);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1747, 420902, '孝南区', 1, '0712', '113.910705,30.916812', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1748, 420921, '孝昌县', 1, '0712', '113.998009,31.258159', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1749, 420922, '大悟县', 1, '0712', '114.127022,31.561164', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1750, 420923, '云梦县', 1, '0712', '113.753554,31.020983', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1751, 420981, '应城市', 1, '0712', '113.572707,30.92837', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1752, 420982, '安陆市', 1, '0712', '113.688941,31.25561', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1753, 420984, '汉川市', 1, '0712', '113.839149,30.661243', 1746);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1755, 421002, '沙市区', 1, '0716', '112.25193,30.326009', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1756, 421003, '荆州区', 1, '0716', '112.190185,30.352853', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1757, 421022, '公安县', 1, '0716', '112.229648,30.058336', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1758, 421023, '监利县', 1, '0716', '112.904788,29.840179', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1838, 430500, '邵阳市', 2, '0739', '111.467674,27.23895', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1759, 421024, '江陵县', 1, '0716', '112.424664,30.041822', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1760, 421081, '石首市', 1, '0716', '112.425454,29.720938', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1761, 421083, '洪湖市', 1, '0716', '113.475801,29.826916', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1762, 421087, '松滋市', 1, '0716', '111.756781,30.174529', 1754);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1764, 421102, '黄州区', 1, '0713', '114.880104,30.434354', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1765, 421121, '团风县', 1, '0713', '114.872191,30.643569', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1766, 421122, '红安县', 1, '0713', '114.618236,31.288153', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1767, 421123, '罗田县', 1, '0713', '115.399222,30.78429', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1768, 421124, '英山县', 1, '0713', '115.681359,30.735157', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1769, 421125, '浠水县', 1, '0713', '115.265355,30.452115', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1770, 421126, '蕲春县', 1, '0713', '115.437007,30.225964', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1771, 421127, '黄梅县', 1, '0713', '115.944219,30.070453', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1851, 430600, '岳阳市', 2, '0730', '113.12873,29.356803', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1772, 421181, '麻城市', 1, '0713', '115.008163,31.172739', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1773, 421182, '武穴市', 1, '0713', '115.561217,29.844107', 1763);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1775, 421202, '咸安区', 1, '0715', '114.298711,29.852891', 1774);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1776, 421221, '嘉鱼县', 1, '0715', '113.939271,29.970676', 1774);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1777, 421222, '通城县', 1, '0715', '113.816966,29.245269', 1774);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1778, 421223, '崇阳县', 1, '0715', '114.039523,29.556688', 1774);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1779, 421224, '通山县', 1, '0715', '114.482622,29.606372', 1774);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1780, 421281, '赤壁市', 1, '0715', '113.90038,29.725184', 1774);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1782, 421303, '曾都区', 1, '0722', '113.37112,31.71628', 1781);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1861, 430700, '常德市', 2, '0736', '111.698784,29.031654', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1783, 421321, '随县', 1, '0722', '113.290634,31.883739', 1781);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1798, 430000, '湖南省', 4, '[]', '112.9836,28.112743', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1863, 430703, '鼎城区', 1, '0736', '111.680783,29.018593', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1784, 421381, '广水市', 1, '0722', '113.825889,31.616853', 1781);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1786, 422801, '恩施市', 1, '0718', '109.479664,30.29468', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1787, 422802, '利川市', 1, '0718', '108.936452,30.29098', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1788, 422822, '建始县', 1, '0718', '109.722109,30.602129', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1789, 422823, '巴东县', 1, '0718', '110.340756,31.042324', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1790, 422825, '宣恩县', 1, '0718', '109.489926,29.98692', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1871, 430800, '张家界市', 2, '0744', '110.479148,29.117013', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1791, 422826, '咸丰县', 1, '0718', '109.139726,29.665202', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1792, 422827, '来凤县', 1, '0718', '109.407828,29.493484', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1793, 422828, '鹤峰县', 1, '0718', '110.033662,29.890171', 1785);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1800, 430102, '芙蓉区', 1, '0731', '113.032539,28.185389', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1876, 430900, '益阳市', 2, '0737', '112.355129,28.554349', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1801, 430103, '天心区', 1, '0731', '112.989897,28.114526', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1802, 430104, '岳麓区', 1, '0731', '112.93132,28.234538', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1803, 430105, '开福区', 1, '0731', '112.985884,28.256298', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1804, 430111, '雨花区', 1, '0731', '113.03826,28.135722', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1805, 430112, '望城区', 1, '0731', '112.831176,28.353434', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1806, 430121, '长沙县', 1, '0731', '113.081097,28.246918', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1883, 431000, '郴州市', 2, '0735', '113.014984,25.770532', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1807, 430124, '宁乡市', 1, '0731', '112.551885,28.277483', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1808, 430181, '浏阳市', 1, '0731', '113.643076,28.162833', 1799);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1810, 430202, '荷塘区', 1, '0733', '113.173487,27.855928', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1811, 430203, '芦淞区', 1, '0733', '113.152724,27.78507', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1812, 430204, '石峰区', 1, '0733', '113.117731,27.875445', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1813, 430211, '天元区', 1, '0733', '113.082216,27.826866', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1814, 430221, '株洲县', 1, '0733', '113.144109,27.699232', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1815, 430223, '攸县', 1, '0733', '113.396385,27.014583', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1816, 430224, '茶陵县', 1, '0733', '113.539094,26.777521', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1817, 430225, '炎陵县', 1, '0733', '113.772655,26.489902', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1818, 430281, '醴陵市', 1, '0733', '113.496999,27.646096', 1809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1895, 431100, '永州市', 2, '0746', '111.613418,26.419641', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1820, 430302, '雨湖区', 1, '0732', '112.907162,27.856325', 1819);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1821, 430304, '岳塘区', 1, '0732', '112.969479,27.872028', 1819);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1822, 430321, '湘潭县', 1, '0732', '112.950831,27.778958', 1819);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1823, 430381, '湘乡市', 1, '0732', '112.550205,27.718549', 1819);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1824, 430382, '韶山市', 1, '0732', '112.52667,27.915008', 1819);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1826, 430405, '珠晖区', 1, '0734', '112.620209,26.894765', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1827, 430406, '雁峰区', 1, '0734', '112.6154,26.840602', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1828, 430407, '石鼓区', 1, '0734', '112.597992,26.943755', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1829, 430408, '蒸湘区', 1, '0734', '112.567107,26.911854', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1830, 430412, '南岳区', 1, '0734', '112.738604,27.232443', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1831, 430421, '衡阳县', 1, '0734', '112.370546,26.969577', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1907, 431200, '怀化市', 2, '0745', '110.001923,27.569517', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1832, 430422, '衡南县', 1, '0734', '112.677877,26.738247', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1833, 430423, '衡山县', 1, '0734', '112.868268,27.23029', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1834, 430424, '衡东县', 1, '0734', '112.953168,27.08117', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1835, 430426, '祁东县', 1, '0734', '112.090356,26.799896', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1836, 430481, '耒阳市', 1, '0734', '112.859759,26.422277', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1837, 430482, '常宁市', 1, '0734', '112.399878,26.421956', 1825);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1839, 430502, '双清区', 1, '0739', '111.496341,27.232708', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1840, 430503, '大祥区', 1, '0739', '111.439091,27.221452', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1841, 430511, '北塔区', 1, '0739', '111.452196,27.246489', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1842, 430521, '邵东县', 1, '0739', '111.74427,27.258987', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1843, 430522, '新邵县', 1, '0739', '111.458656,27.320917', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1844, 430523, '邵阳县', 1, '0739', '111.273805,26.990637', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1920, 431300, '娄底市', 2, '0738', '111.994482,27.70027', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1845, 430524, '隆回县', 1, '0739', '111.032437,27.113978', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1846, 430525, '洞口县', 1, '0739', '110.575846,27.06032', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1847, 430527, '绥宁县', 1, '0739', '110.155655,26.581954', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1848, 430528, '新宁县', 1, '0739', '110.856988,26.433367', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1849, 430529, '城步苗族自治县', 1, '0739', '110.322239,26.390598', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1926, 433100, '湘西土家族苗族自治州', 2, '0743', '109.738906,28.31195', 1798);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1850, 430581, '武冈市', 1, '0739', '110.631884,26.726599', 1838);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1852, 430602, '岳阳楼区', 1, '0730', '113.129684,29.371814', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1853, 430603, '云溪区', 1, '0730', '113.272312,29.472745', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1854, 430611, '君山区', 1, '0730', '113.006435,29.461106', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1855, 430621, '岳阳县', 1, '0730', '113.116418,29.144066', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1856, 430623, '华容县', 1, '0730', '112.540463,29.531057', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1857, 430624, '湘阴县', 1, '0730', '112.909426,28.689104', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1858, 430626, '平江县', 1, '0730', '113.581234,28.701868', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1859, 430681, '汨罗市', 1, '0730', '113.067251,28.806881', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1936, 440100, '广州市', 2, '020', '113.264385,23.12911', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1860, 430682, '临湘市', 1, '0730', '113.450423,29.476849', 1851);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1862, 430702, '武陵区', 1, '0736', '111.683153,29.055163', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1864, 430721, '安乡县', 1, '0736', '112.171131,29.411309', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1865, 430722, '汉寿县', 1, '0736', '111.970514,28.906106', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1866, 430723, '澧县', 1, '0736', '111.758702,29.633236', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1867, 430724, '临澧县', 1, '0736', '111.647517,29.440793', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1868, 430725, '桃源县', 1, '0736', '111.488925,28.902503', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1869, 430726, '石门县', 1, '0736', '111.380014,29.584292', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1870, 430781, '津市市', 1, '0736', '111.877499,29.60548', 1861);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1872, 430802, '永定区', 1, '0744', '110.537138,29.119855', 1871);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1873, 430811, '武陵源区', 1, '0744', '110.550433,29.34573', 1871);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1948, 440200, '韶关市', 2, '0751', '113.59762,24.810879', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1874, 430821, '慈利县', 1, '0744', '111.139775,29.429999', 1871);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1875, 430822, '桑植县', 1, '0744', '110.204652,29.414111', 1871);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1877, 430902, '资阳区', 1, '0737', '112.324272,28.59111', 1876);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1878, 430903, '赫山区', 1, '0737', '112.374145,28.579494', 1876);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1879, 430921, '南县', 1, '0737', '112.396337,29.362275', 1876);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1880, 430922, '桃江县', 1, '0737', '112.155822,28.518084', 1876);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1881, 430923, '安化县', 1, '0737', '111.212846,28.374107', 1876);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1882, 430981, '沅江市', 1, '0737', '112.355954,28.847045', 1876);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1884, 431002, '北湖区', 1, '0735', '113.011035,25.784054', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1885, 431003, '苏仙区', 1, '0735', '113.112105,25.797013', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1959, 440300, '深圳市', 2, '0755', '114.057939,22.543527', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1886, 431021, '桂阳县', 1, '0735', '112.734173,25.754172', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1887, 431022, '宜章县', 1, '0735', '112.948712,25.399938', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1888, 431023, '永兴县', 1, '0735', '113.116527,26.12715', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1889, 431024, '嘉禾县', 1, '0735', '112.36902,25.587519', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1890, 431025, '临武县', 1, '0735', '112.563456,25.27556', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1891, 431026, '汝城县', 1, '0735', '113.684727,25.532816', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1892, 431027, '桂东县', 1, '0735', '113.944614,26.077616', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1893, 431028, '安仁县', 1, '0735', '113.26932,26.709061', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1968, 440400, '珠海市', 2, '0756', '113.576677,22.270978', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1894, 431081, '资兴市', 1, '0735', '113.236146,25.976243', 1883);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1896, 431102, '零陵区', 1, '0746', '111.631109,26.221936', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1897, 431103, '冷水滩区', 1, '0746', '111.592343,26.46128', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1972, 440500, '汕头市', 2, '0754', '116.681972,23.354091', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1898, 431121, '祁阳县', 1, '0746', '111.840657,26.58012', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1899, 431122, '东安县', 1, '0746', '111.316464,26.392183', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1900, 431123, '双牌县', 1, '0746', '111.659967,25.961909', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1901, 431124, '道县', 1, '0746', '111.600795,25.526437', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1902, 431125, '江永县', 1, '0746', '111.343911,25.273539', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1903, 431126, '宁远县', 1, '0746', '111.945844,25.570888', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1904, 431127, '蓝山县', 1, '0746', '112.196567,25.369725', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1980, 440600, '佛山市', 2, '0757', '113.121435,23.021478', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1905, 431128, '新田县', 1, '0746', '112.203287,25.904305', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1906, 431129, '江华瑶族自治县', 1, '0746', '111.579535,25.185809', 1895);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1908, 431202, '鹤城区', 1, '0745', '110.040315,27.578926', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1909, 431221, '中方县', 1, '0745', '109.944711,27.440138', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1910, 431222, '沅陵县', 1, '0745', '110.393844,28.452686', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1986, 440700, '江门市', 2, '0750', '113.081542,22.57899', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1911, 431223, '辰溪县', 1, '0745', '110.183917,28.006336', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1912, 431224, '溆浦县', 1, '0745', '110.594879,27.908267', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1913, 431225, '会同县', 1, '0745', '109.735661,26.887238', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1914, 431226, '麻阳苗族自治县', 1, '0745', '109.81701,27.857569', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1915, 431227, '新晃侗族自治县', 1, '0745', '109.174932,27.352673', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1916, 431228, '芷江侗族自治县', 1, '0745', '109.684629,27.443499', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1917, 431229, '靖州苗族侗族自治县', 1, '0745', '109.696273,26.575107', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1994, 440800, '湛江市', 2, '0759', '110.356639,21.270145', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1918, 431230, '通道侗族自治县', 1, '0745', '109.784412,26.158054', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1919, 431281, '洪江市', 1, '0745', '109.836669,27.208609', 1907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1921, 431302, '娄星区', 1, '0738', '112.001914,27.729863', 1920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1922, 431321, '双峰县', 1, '0738', '112.175163,27.457172', 1920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1923, 431322, '新化县', 1, '0738', '111.327412,27.726514', 1920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1924, 431381, '冷水江市', 1, '0738', '111.434984,27.686251', 1920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1925, 431382, '涟源市', 1, '0738', '111.664329,27.692577', 1920);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1927, 433101, '吉首市', 1, '0743', '109.698015,28.262376', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1928, 433122, '泸溪县', 1, '0743', '110.21961,28.216641', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2004, 440900, '茂名市', 2, '0668', '110.925439,21.662991', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1929, 433123, '凤凰县', 1, '0743', '109.581083,27.958081', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1930, 433124, '花垣县', 1, '0743', '109.482078,28.572029', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1931, 433125, '保靖县', 1, '0743', '109.660559,28.699878', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1932, 433126, '古丈县', 1, '0743', '109.950728,28.616935', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1933, 433127, '永顺县', 1, '0743', '109.856933,28.979955', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2010, 441200, '肇庆市', 2, '0758', '112.465091,23.047191', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1934, 433130, '龙山县', 1, '0743', '109.443938,29.457663', 1926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1937, 440103, '荔湾区', 1, '020', '113.244258,23.125863', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1938, 440104, '越秀区', 1, '020', '113.266835,23.128537', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1939, 440105, '海珠区', 1, '020', '113.317443,23.083788', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1940, 440106, '天河区', 1, '020', '113.361575,23.124807', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1941, 440111, '白云区', 1, '020', '113.273238,23.157367', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1942, 440112, '黄埔区', 1, '020', '113.480541,23.181706', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1943, 440113, '番禺区', 1, '020', '113.384152,22.937556', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2019, 441300, '惠州市', 2, '0752', '114.415612,23.112381', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1944, 440114, '花都区', 1, '020', '113.220463,23.403744', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1945, 440115, '南沙区', 1, '020', '113.525165,22.801624', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1946, 440117, '从化区', 1, '020', '113.586679,23.548748', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1947, 440118, '增城区', 1, '020', '113.810627,23.261465', 1936);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1949, 440203, '武江区', 1, '0751', '113.587756,24.792926', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2025, 441400, '梅州市', 2, '0753', '116.122523,24.288578', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1950, 440204, '浈江区', 1, '0751', '113.611098,24.804381', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1951, 440205, '曲江区', 1, '0751', '113.604535,24.682501', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1952, 440222, '始兴县', 1, '0751', '114.061789,24.952976', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1953, 440224, '仁化县', 1, '0751', '113.749027,25.085621', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1954, 440229, '翁源县', 1, '0751', '114.130342,24.350346', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1955, 440232, '乳源瑶族自治县', 1, '0751', '113.275883,24.776078', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1956, 440233, '新丰县', 1, '0751', '114.206867,24.05976', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1957, 440281, '乐昌市', 1, '0751', '113.347545,25.130602', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2034, 441500, '汕尾市', 2, '0660', '115.375431,22.78705', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1958, 440282, '南雄市', 1, '0751', '114.311982,25.117753', 1948);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1960, 440303, '罗湖区', 1, '0755', '114.131459,22.548389', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1961, 440304, '福田区', 1, '0755', '114.055072,22.521521', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2039, 441600, '河源市', 2, '0762', '114.700961,23.743686', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2040, 441602, '源城区', 1, '0762', '114.702517,23.733969', 2039);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1962, 440305, '南山区', 1, '0755', '113.930413,22.533287', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1963, 440306, '宝安区', 1, '0755', '113.883802,22.554996', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1964, 440307, '龙岗区', 1, '0755', '114.246899,22.720974', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1965, 440308, '盐田区', 1, '0755', '114.236739,22.557001', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2046, 441700, '阳江市', 2, '0662', '111.982589,21.857887', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1966, 440309, '龙华区', 1, '0755', '114.045422,22.696667', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1967, 440310, '坪山区', 1, '0755', '114.350584,22.708881', 1959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1969, 440402, '香洲区', 1, '0756', '113.543784,22.265811', 1968);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1970, 440403, '斗门区', 1, '0756', '113.296467,22.2092', 1968);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2051, 441800, '清远市', 2, '0763', '113.056042,23.681774', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1971, 440404, '金湾区', 1, '0756', '113.362656,22.147471', 1968);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1973, 440507, '龙湖区', 1, '0754', '116.716446,23.372254', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1974, 440511, '金平区', 1, '0754', '116.70345,23.365556', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1975, 440512, '濠江区', 1, '0754', '116.726973,23.286079', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1976, 440513, '潮阳区', 1, '0754', '116.601509,23.265356', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1977, 440514, '潮南区', 1, '0754', '116.439178,23.23865', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1978, 440515, '澄海区', 1, '0754', '116.755992,23.466709', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1979, 440523, '南澳县', 1, '0754', '117.023374,23.421724', 1972);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2060, 441900, '东莞市', 2, '0769', '113.751799,23.020673', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2061, 442000, '中山市', 2, '0760', '113.39277,22.517585', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2062, 445100, '潮州市', 2, '0768', '116.622444,23.657262', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1981, 440604, '禅城区', 1, '0757', '113.122421,23.009551', 1980);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1982, 440605, '南海区', 1, '0757', '113.143441,23.028956', 1980);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1983, 440606, '顺德区', 1, '0757', '113.293359,22.80524', 1980);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2066, 445200, '揭阳市', 2, '0663', '116.372708,23.549701', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1984, 440607, '三水区', 1, '0757', '112.896685,23.155931', 1980);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1985, 440608, '高明区', 1, '0757', '112.892585,22.900139', 1980);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1987, 440703, '蓬江区', 1, '0750', '113.078521,22.595149', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1988, 440704, '江海区', 1, '0750', '113.111612,22.560473', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1989, 440705, '新会区', 1, '0750', '113.034187,22.4583', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2072, 445300, '云浮市', 2, '0766', '112.044491,22.915094', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1990, 440781, '台山市', 1, '0750', '112.794065,22.251924', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1991, 440783, '开平市', 1, '0750', '112.698545,22.376395', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1992, 440784, '鹤山市', 1, '0750', '112.964252,22.76545', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1993, 440785, '恩平市', 1, '0750', '112.305145,22.183206', 1986);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1995, 440802, '赤坎区', 1, '0759', '110.365899,21.266119', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2078, 442100, '东沙群岛', 2, '[]', '116.887613,20.617825', 1935);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1996, 440803, '霞山区', 1, '0759', '110.397656,21.192457', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2080, 450100, '南宁市', 2, '0771', '108.366543,22.817002', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1997, 440804, '坡头区', 1, '0759', '110.455332,21.244721', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1998, 440811, '麻章区', 1, '0759', '110.334387,21.263442', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1999, 440823, '遂溪县', 1, '0759', '110.250123,21.377246', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2000, 440825, '徐闻县', 1, '0759', '110.176749,20.325489', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2001, 440881, '廉江市', 1, '0759', '110.286208,21.6097', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2002, 440882, '雷州市', 1, '0759', '110.096586,20.914178', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2003, 440883, '吴川市', 1, '0759', '110.778411,21.441808', 1994);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2005, 440902, '茂南区', 1, '0668', '110.918026,21.641337', 2004);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2006, 440904, '电白区', 1, '0668', '111.013556,21.514163', 2004);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2007, 440981, '高州市', 1, '0668', '110.853299,21.918203', 2004);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2008, 440982, '化州市', 1, '0668', '110.639565,21.66463', 2004);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2009, 440983, '信宜市', 1, '0668', '110.947043,22.354385', 2004);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2093, 450200, '柳州市', 2, '0772', '109.428608,24.326291', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2011, 441202, '端州区', 1, '0758', '112.484848,23.052101', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2012, 441203, '鼎湖区', 1, '0758', '112.567588,23.158447', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2013, 441204, '高要区', 1, '0758', '112.457981,23.025305', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2014, 441223, '广宁县', 1, '0758', '112.44069,23.634675', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2015, 441224, '怀集县', 1, '0758', '112.167742,23.92035', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2016, 441225, '封开县', 1, '0758', '111.512343,23.424033', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2017, 441226, '德庆县', 1, '0758', '111.785937,23.143722', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2018, 441284, '四会市', 1, '0758', '112.734103,23.327001', 2010);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2020, 441302, '惠城区', 1, '0752', '114.382474,23.084137', 2019);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2021, 441303, '惠阳区', 1, '0752', '114.456176,22.789788', 2019);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2104, 450300, '桂林市', 2, '0773', '110.179953,25.234479', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2022, 441322, '博罗县', 1, '0752', '114.289528,23.172771', 2019);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2023, 441323, '惠东县', 1, '0752', '114.719988,22.985014', 2019);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2024, 441324, '龙门县', 1, '0752', '114.254863,23.727737', 2019);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2026, 441402, '梅江区', 1, '0753', '116.116695,24.31049', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2027, 441403, '梅县区', 1, '0753', '116.081656,24.265926', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2028, 441422, '大埔县', 1, '0753', '116.695195,24.347782', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2029, 441423, '丰顺县', 1, '0753', '116.181691,23.739343', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2030, 441424, '五华县', 1, '0753', '115.775788,23.932409', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2031, 441426, '平远县', 1, '0753', '115.891638,24.567261', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2032, 441427, '蕉岭县', 1, '0753', '116.171355,24.658699', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2033, 441481, '兴宁市', 1, '0753', '115.731167,24.136708', 2025);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2035, 441502, '城区', 1, '0660', '115.365058,22.779207', 2034);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2036, 441521, '海丰县', 1, '0660', '115.323436,22.966585', 2034);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2037, 441523, '陆河县', 1, '0660', '115.660143,23.301616', 2034);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2038, 441581, '陆丰市', 1, '0660', '115.652151,22.919228', 2034);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2041, 441621, '紫金县', 1, '0762', '115.184107,23.635745', 2039);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2042, 441622, '龙川县', 1, '0762', '115.259871,24.100066', 2039);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2122, 450400, '梧州市', 2, '0774', '111.279115,23.476962', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2043, 441623, '连平县', 1, '0762', '114.488556,24.369583', 2039);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2044, 441624, '和平县', 1, '0762', '114.938684,24.44218', 2039);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2045, 441625, '东源县', 1, '0762', '114.746344,23.788189', 2039);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2047, 441702, '江城区', 1, '0662', '111.955058,21.861786', 2046);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2048, 441704, '阳东区', 1, '0662', '112.006363,21.868337', 2046);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2049, 441721, '阳西县', 1, '0662', '111.61766,21.752771', 2046);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2128, 450423, '蒙山县', 1, '0774', '110.525003,24.19357', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2130, 450500, '北海市', 2, '0779', '109.120161,21.481291', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2050, 441781, '阳春市', 1, '0662', '111.791587,22.17041', 2046);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2052, 441802, '清城区', 1, '0763', '113.062692,23.697899', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2053, 441803, '清新区', 1, '0763', '113.017747,23.734677', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2054, 441821, '佛冈县', 1, '0763', '113.531607,23.879192', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2135, 450600, '防城港市', 2, '0770', '108.353846,21.68686', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2055, 441823, '阳山县', 1, '0763', '112.641363,24.465359', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2056, 441825, '连山壮族瑶族自治县', 1, '0763', '112.093617,24.570491', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2057, 441826, '连南瑶族自治县', 1, '0763', '112.287012,24.726017', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2058, 441881, '英德市', 1, '0763', '113.401701,24.206986', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2140, 450700, '钦州市', 2, '0777', '108.654146,21.979933', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2059, 441882, '连州市', 1, '0763', '112.377361,24.780966', 2051);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2063, 445102, '湘桥区', 1, '0768', '116.628627,23.674387', 2062);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2064, 445103, '潮安区', 1, '0768', '116.678203,23.462613', 2062);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2065, 445122, '饶平县', 1, '0768', '117.0039,23.663824', 2062);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2145, 450800, '贵港市', 2, '1755', '109.598926,23.11153', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2067, 445202, '榕城区', 1, '0663', '116.367012,23.525382', 2066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2068, 445203, '揭东区', 1, '0663', '116.412015,23.566126', 2066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2069, 445222, '揭西县', 1, '0663', '115.841837,23.431294', 2066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2070, 445224, '惠来县', 1, '0663', '116.29515,23.033266', 2066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2071, 445281, '普宁市', 1, '0663', '116.165777,23.297493', 2066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2151, 450900, '玉林市', 2, '0775', '110.18122,22.654032', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2073, 445302, '云城区', 1, '0766', '112.043945,22.92815', 2072);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2074, 445303, '云安区', 1, '0766', '112.003208,23.071019', 2072);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2075, 445321, '新兴县', 1, '0766', '112.225334,22.69569', 2072);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2076, 445322, '郁南县', 1, '0766', '111.535285,23.23456', 2072);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2077, 445381, '罗定市', 1, '0766', '111.569892,22.768285', 2072);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2081, 450102, '兴宁区', 1, '0771', '108.368871,22.854021', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2082, 450103, '青秀区', 1, '0771', '108.494024,22.785879', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2159, 451000, '百色市', 2, '0776', '106.618202,23.90233', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2083, 450105, '江南区', 1, '0771', '108.273133,22.78136', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2084, 450107, '西乡塘区', 1, '0771', '108.313494,22.833928', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2085, 450108, '良庆区', 1, '0771', '108.39301,22.752997', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2086, 450109, '邕宁区', 1, '0771', '108.487368,22.75839', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2087, 450110, '武鸣区', 1, '0771', '108.27467,23.158595', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2088, 450123, '隆安县', 1, '0771', '107.696153,23.166028', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2089, 450124, '马山县', 1, '0771', '108.177019,23.708321', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2090, 450125, '上林县', 1, '0771', '108.602846,23.431908', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2091, 450126, '宾阳县', 1, '0771', '108.810326,23.217786', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2092, 450127, '横县', 1, '0771', '109.261384,22.679931', 2080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2094, 450202, '城中区', 1, '0772', '109.4273,24.366', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2095, 450203, '鱼峰区', 1, '0772', '109.452442,24.318516', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2172, 451100, '贺州市', 2, '1774', '111.566871,24.403528', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2096, 450204, '柳南区', 1, '0772', '109.385518,24.336229', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2097, 450205, '柳北区', 1, '0772', '109.402049,24.362691', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2098, 450206, '柳江区', 1, '0772', '109.32638,24.254891', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2099, 450222, '柳城县', 1, '0772', '109.24473,24.651518', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2100, 450223, '鹿寨县', 1, '0772', '109.750638,24.472897', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2178, 451200, '河池市', 2, '0778', '108.085261,24.692931', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2101, 450224, '融安县', 1, '0772', '109.397538,25.224549', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2102, 450225, '融水苗族自治县', 1, '0772', '109.256334,25.065934', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2103, 450226, '三江侗族自治县', 1, '0772', '109.607675,25.783198', 2093);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2105, 450302, '秀峰区', 1, '0773', '110.264183,25.273625', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2106, 450303, '叠彩区', 1, '0773', '110.301723,25.314', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2107, 450304, '象山区', 1, '0773', '110.281082,25.261686', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2108, 450305, '七星区', 1, '0773', '110.317826,25.252701', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2109, 450311, '雁山区', 1, '0773', '110.28669,25.101934', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2110, 450312, '临桂区', 1, '0773', '110.212463,25.238628', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2111, 450321, '阳朔县', 1, '0773', '110.496593,24.77848', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2112, 450323, '灵川县', 1, '0773', '110.319897,25.394781', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2190, 451300, '来宾市', 2, '1772', '109.221465,23.750306', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2113, 450324, '全州县', 1, '0773', '111.072946,25.928387', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2114, 450325, '兴安县', 1, '0773', '110.67167,25.611704', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2115, 450326, '永福县', 1, '0773', '109.983076,24.979855', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2116, 450327, '灌阳县', 1, '0773', '111.160851,25.489383', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2117, 450328, '龙胜各族自治县', 1, '0773', '110.011238,25.797931', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2118, 450329, '资源县', 1, '0773', '110.6527,26.042443', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2197, 451400, '崇左市', 2, '1771', '107.365094,22.377253', 2079);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2119, 450330, '平乐县', 1, '0773', '110.643305,24.633362', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2120, 450331, '荔浦县', 1, '0773', '110.395104,24.488342', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2121, 450332, '恭城瑶族自治县', 1, '0773', '110.828409,24.831682', 2104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2123, 450403, '万秀区', 1, '0774', '111.320518,23.472991', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2124, 450405, '长洲区', 1, '0774', '111.274673,23.485944', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2125, 450406, '龙圩区', 1, '0774', '111.246606,23.404772', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2126, 450421, '苍梧县', 1, '0774', '111.544007,23.845097', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2127, 450422, '藤县', 1, '0774', '110.914849,23.374983', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2206, 469025, '白沙黎族自治县', 2, '0802', '109.451484,19.224823', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2207, 469029, '保亭黎族苗族自治县', 2, '0801', '109.70259,18.63913', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2208, 469026, '昌江黎族自治县', 2, '0803', '109.055739,19.298184', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2209, 469023, '澄迈县', 2, '0804', '110.006754,19.738521', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2210, 460100, '海口市', 2, '0898', '110.198286,20.044412', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2129, 450481, '岑溪市', 1, '0774', '110.994913,22.91835', 2122);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2131, 450502, '海城区', 1, '0779', '109.117209,21.475004', 2130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2132, 450503, '银海区', 1, '0779', '109.139862,21.449308', 2130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2133, 450512, '铁山港区', 1, '0779', '109.42158,21.529127', 2130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2134, 450521, '合浦县', 1, '0779', '109.207335,21.660935', 2130);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2215, 460200, '三亚市', 2, '0899', '109.511772,18.253135', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2216, 460202, '海棠区', 1, '0899', '109.752569,18.400106', 2215);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2136, 450602, '港口区', 1, '0770', '108.380143,21.643383', 2135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2220, 460300, '三沙市', 2, '2898', '112.338695,16.831839', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2137, 450603, '防城区', 1, '0770', '108.353499,21.769211', 2135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2138, 450621, '上思县', 1, '0770', '107.983627,22.153671', 2135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2139, 450681, '东兴市', 1, '0770', '107.971828,21.547821', 2135);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2224, 460400, '儋州市', 2, '0805', '109.580811,19.521134', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2225, 469021, '定安县', 2, '0806', '110.359339,19.681404', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2226, 469007, '东方市', 2, '0807', '108.651815,19.095351', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2227, 469027, '乐东黎族自治县', 2, '2802', '109.173054,18.750259', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2228, 469024, '临高县', 2, '1896', '109.690508,19.912025', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2229, 469028, '陵水黎族自治县', 2, '0809', '110.037503,18.506048', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2230, 469002, '琼海市', 2, '1894', '110.474497,19.259134', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2231, 469030, '琼中黎族苗族自治县', 2, '1899', '109.838389,19.033369', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2232, 469022, '屯昌县', 2, '1892', '110.103415,19.351765', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2233, 469006, '万宁市', 2, '1898', '110.391073,18.795143', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2234, 469005, '文昌市', 2, '1893', '110.797717,19.543422', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2235, 469001, '五指山市', 2, '1897', '109.516925,18.775146', 2205);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2237, 500100, '重庆城区', 2, '023', '106.551643,29.562849', 2236);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2141, 450702, '钦南区', 1, '0777', '108.657209,21.938859', 2140);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2142, 450703, '钦北区', 1, '0777', '108.44911,22.132761', 2140);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2143, 450721, '灵山县', 1, '0777', '109.291006,22.416536', 2140);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2144, 450722, '浦北县', 1, '0777', '109.556953,22.271651', 2140);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2146, 450802, '港北区', 1, '1755', '109.57224,23.11153', 2145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2147, 450803, '港南区', 1, '1755', '109.599556,23.075573', 2145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2148, 450804, '覃塘区', 1, '1755', '109.452662,23.127149', 2145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2149, 450821, '平南县', 1, '1755', '110.392311,23.539264', 2145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2150, 450881, '桂平市', 1, '1755', '110.079379,23.394325', 2145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2152, 450902, '玉州区', 1, '0775', '110.151153,22.628087', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2153, 450903, '福绵区', 1, '0775', '110.059439,22.585556', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2154, 450921, '容县', 1, '0775', '110.558074,22.857839', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2155, 450922, '陆川县', 1, '0775', '110.264052,22.321048', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2156, 450923, '博白县', 1, '0775', '109.975985,22.273048', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2157, 450924, '兴业县', 1, '0775', '109.875304,22.736421', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2158, 450981, '北流市', 1, '0775', '110.354214,22.70831', 2151);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2160, 451002, '右江区', 1, '0776', '106.618225,23.90097', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2161, 451021, '田阳县', 1, '0776', '106.915496,23.735692', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2162, 451022, '田东县', 1, '0776', '107.12608,23.597194', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2163, 451023, '平果县', 1, '0776', '107.589809,23.329376', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2164, 451024, '德保县', 1, '0776', '106.615373,23.32345', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2165, 451026, '那坡县', 1, '0776', '105.83253,23.387441', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2166, 451027, '凌云县', 1, '0776', '106.56131,24.347557', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2167, 451028, '乐业县', 1, '0776', '106.556519,24.776827', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2262, 500200, '重庆郊县', 2, '023', '108.165537,29.293902', 2236);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2168, 451029, '田林县', 1, '0776', '106.228538,24.294487', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2169, 451030, '西林县', 1, '0776', '105.093825,24.489823', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2170, 451031, '隆林各族自治县', 1, '0776', '105.34404,24.770896', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2171, 451081, '靖西市', 1, '0776', '106.417805,23.134117', 2159);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2173, 451102, '八步区', 1, '1774', '111.552095,24.411805', 2172);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2174, 451103, '平桂区', 1, '1774', '111.479923,24.453845', 2172);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2175, 451121, '昭平县', 1, '1774', '110.811325,24.169385', 2172);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2176, 451122, '钟山县', 1, '1774', '111.303009,24.525957', 2172);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2177, 451123, '富川瑶族自治县', 1, '1774', '111.27745,24.814443', 2172);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2179, 451202, '金城江区', 1, '0778', '108.037276,24.689703', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2180, 451221, '南丹县', 1, '0778', '107.541244,24.975631', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2181, 451222, '天峨县', 1, '0778', '107.173802,24.999108', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2182, 451223, '凤山县', 1, '0778', '107.04219,24.546876', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2183, 451224, '东兰县', 1, '0778', '107.374293,24.510842', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2184, 451225, '罗城仫佬族自治县', 1, '0778', '108.904706,24.777411', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2278, 510100, '成都市', 2, '028', '104.066794,30.572893', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2185, 451226, '环江毛南族自治县', 1, '0778', '108.258028,24.825664', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2186, 451227, '巴马瑶族自治县', 1, '0778', '107.258588,24.142298', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2187, 451228, '都安瑶族自治县', 1, '0778', '108.105311,23.932675', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2188, 451229, '大化瑶族自治县', 1, '0778', '107.998149,23.736457', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2189, 451203, '宜州区', 1, '0778', '108.636414,24.485214', 2178);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2191, 451302, '兴宾区', 1, '1772', '109.183333,23.72892', 2190);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2192, 451321, '忻城县', 1, '1772', '108.665666,24.066234', 2190);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2193, 451322, '象州县', 1, '1772', '109.705065,23.973793', 2190);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2194, 451323, '武宣县', 1, '1772', '109.663206,23.59411', 2190);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2195, 451324, '金秀瑶族自治县', 1, '1772', '110.189462,24.130374', 2190);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2196, 451381, '合山市', 1, '1772', '108.886082,23.806535', 2190);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2198, 451402, '江州区', 1, '1771', '107.353437,22.405325', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2199, 451421, '扶绥县', 1, '1771', '107.904186,22.635012', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2200, 451422, '宁明县', 1, '1771', '107.076456,22.140192', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2201, 451423, '龙州县', 1, '1771', '106.854482,22.342778', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2202, 451424, '大新县', 1, '1771', '107.200654,22.829287', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2203, 451425, '天等县', 1, '1771', '107.143432,23.081394', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2204, 451481, '凭祥市', 1, '1771', '106.766293,22.094484', 2197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2211, 460105, '秀英区', 1, '0898', '110.293603,20.007494', 2210);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2212, 460106, '龙华区', 1, '0898', '110.328492,20.031006', 2210);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2299, 510300, '自贡市', 2, '0813', '104.778442,29.33903', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2213, 460107, '琼山区', 1, '0898', '110.353972,20.003169', 2210);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2214, 460108, '美兰区', 1, '0898', '110.366358,20.029083', 2210);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2236, 500000, '重庆市', 4, '023', '106.551643,29.562849', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2217, 460203, '吉阳区', 1, '0899', '109.578336,18.281406', 2215);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2218, 460204, '天涯区', 1, '0899', '109.452378,18.298156', 2215);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2219, 460205, '崖州区', 1, '0899', '109.171841,18.357291', 2215);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2306, 510400, '攀枝花市', 2, '0812', '101.718637,26.582347', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2221, 460321, '西沙群岛', 1, '1895', '111.792944,16.204546', 2220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2222, 460322, '南沙群岛', 1, '1891', '116.749997,11.471888', 2220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2223, 460323, '中沙群岛的岛礁及其海域', 1, '2801', '117.740071,15.112855', 2220);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2238, 500101, '万州区', 1, '023', '108.408661,30.807667', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2239, 500102, '涪陵区', 1, '023', '107.38977,29.703022', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2312, 510500, '泸州市', 2, '0830', '105.442285,28.871805', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2240, 500103, '渝中区', 1, '023', '106.568896,29.552736', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2241, 500104, '大渡口区', 1, '023', '106.482346,29.484527', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2242, 500105, '江北区', 1, '023', '106.574271,29.606703', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2243, 500106, '沙坪坝区', 1, '023', '106.456878,29.541144', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2244, 500107, '九龙坡区', 1, '023', '106.510676,29.502272', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2245, 500108, '南岸区', 1, '023', '106.644447,29.50126', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2246, 500109, '北碚区', 1, '023', '106.395612,29.805107', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2320, 510600, '德阳市', 2, '0838', '104.397894,31.126855', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2247, 500110, '綦江区', 1, '023', '106.651361,29.028066', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2248, 500111, '大足区', 1, '023', '105.721733,29.707032', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2249, 500112, '渝北区', 1, '023', '106.631187,29.718142', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2250, 500113, '巴南区', 1, '023', '106.540256,29.402408', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2251, 500114, '黔江区', 1, '023', '108.770677,29.533609', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2252, 500115, '长寿区', 1, '023', '107.080734,29.857912', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2327, 510700, '绵阳市', 2, '0816', '104.679004,31.467459', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2253, 500116, '江津区', 1, '023', '106.259281,29.290069', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2254, 500117, '合川区', 1, '023', '106.27613,29.972084', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2255, 500118, '永川区', 1, '023', '105.927001,29.356311', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2256, 500119, '南川区', 1, '023', '107.099266,29.15789', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2257, 500120, '璧山区', 1, '023', '106.227305,29.592024', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2258, 500151, '铜梁区', 1, '023', '106.056404,29.844811', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2259, 500152, '潼南区', 1, '023', '105.840431,30.190992', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2260, 500153, '荣昌区', 1, '023', '105.594623,29.405002', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2261, 500154, '开州区', 1, '023', '108.393135,31.160711', 2237);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2337, 510800, '广元市', 2, '0839', '105.843357,32.435435', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2263, 500155, '梁平区', 1, '023', '107.769568,30.654233', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2264, 500229, '城口县', 1, '023', '108.664214,31.947633', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2265, 500230, '丰都县', 1, '023', '107.730894,29.8635', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2266, 500231, '垫江县', 1, '023', '107.33339,30.327716', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2267, 500156, '武隆区', 1, '023', '107.760025,29.325601', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2268, 500233, '忠县', 1, '023', '108.039002,30.299559', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2269, 500235, '云阳县', 1, '023', '108.697324,30.930612', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2345, 510900, '遂宁市', 2, '0825', '105.592803,30.53292', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2270, 500236, '奉节县', 1, '023', '109.400403,31.018363', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2271, 500237, '巫山县', 1, '023', '109.879153,31.074834', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2272, 500238, '巫溪县', 1, '023', '109.570062,31.398604', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2273, 500240, '石柱土家族自治县', 1, '023', '108.114069,29.999285', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2274, 500241, '秀山土家族苗族自治县', 1, '023', '109.007094,28.447997', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2351, 511000, '内江市', 2, '1832', '105.058432,29.580228', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2275, 500242, '酉阳土家族苗族自治县', 1, '023', '108.767747,28.841244', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2276, 500243, '彭水苗族土家族自治县', 1, '023', '108.165537,29.293902', 2262);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2279, 510104, '锦江区', 1, '028', '104.117022,30.598158', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2280, 510105, '青羊区', 1, '028', '104.061442,30.673914', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2281, 510106, '金牛区', 1, '028', '104.052236,30.691359', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2357, 511100, '乐山市', 2, '0833', '103.765678,29.552115', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2282, 510107, '武侯区', 1, '028', '104.043235,30.641907', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2283, 510108, '成华区', 1, '028', '104.101515,30.659966', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2284, 510112, '龙泉驿区', 1, '028', '104.274632,30.556506', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2285, 510113, '青白江区', 1, '028', '104.250945,30.878629', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2286, 510114, '新都区', 1, '028', '104.158705,30.823498', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2287, 510115, '温江区', 1, '028', '103.856646,30.682203', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2288, 510116, '双流区', 1, '028', '103.923566,30.574449', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2289, 510121, '金堂县', 1, '028', '104.411976,30.861979', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2290, 510117, '郫都区', 1, '028', '103.901091,30.795854', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2291, 510129, '大邑县', 1, '028', '103.511865,30.572268', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2292, 510131, '蒲江县', 1, '028', '103.506498,30.196788', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2369, 511300, '南充市', 2, '0817', '106.110698,30.837793', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2293, 510132, '新津县', 1, '028', '103.811286,30.410346', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2294, 510185, '简阳市', 1, '028', '104.54722,30.411264', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2295, 510181, '都江堰市', 1, '028', '103.647153,30.988767', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2296, 510182, '彭州市', 1, '028', '103.957983,30.990212', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2297, 510183, '邛崃市', 1, '028', '103.464207,30.410324', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2298, 510184, '崇州市', 1, '028', '103.673001,30.630122', 2278);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2300, 510302, '自流井区', 1, '0813', '104.777191,29.337429', 2299);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2301, 510303, '贡井区', 1, '0813', '104.715288,29.345313', 2299);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2302, 510304, '大安区', 1, '0813', '104.773994,29.363702', 2299);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2379, 511400, '眉山市', 2, '1833', '103.848403,30.076994', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2303, 510311, '沿滩区', 1, '0813', '104.874079,29.272586', 2299);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2304, 510321, '荣县', 1, '0813', '104.417493,29.445479', 2299);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2305, 510322, '富顺县', 1, '0813', '104.975048,29.181429', 2299);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2307, 510402, '东区', 1, '0812', '101.704109,26.546491', 2306);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2308, 510403, '西区', 1, '0812', '101.630619,26.597781', 2306);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2309, 510411, '仁和区', 1, '0812', '101.738528,26.497765', 2306);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2386, 511500, '宜宾市', 2, '0831', '104.642845,28.752134', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2310, 510421, '米易县', 1, '0812', '102.112895,26.897694', 2306);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2311, 510422, '盐边县', 1, '0812', '101.855071,26.683213', 2306);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2313, 510502, '江阳区', 1, '0830', '105.434982,28.87881', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2314, 510503, '纳溪区', 1, '0830', '105.371505,28.773134', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2315, 510504, '龙马潭区', 1, '0830', '105.437751,28.913257', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2316, 510521, '泸县', 1, '0830', '105.381893,29.151534', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2317, 510522, '合江县', 1, '0830', '105.830986,28.811164', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2318, 510524, '叙永县', 1, '0830', '105.444765,28.155801', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2319, 510525, '古蔺县', 1, '0830', '105.812601,28.038801', 2312);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2321, 510603, '旌阳区', 1, '0838', '104.416966,31.142633', 2320);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2397, 511600, '广安市', 2, '0826', '106.633088,30.456224', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2322, 510623, '中江县', 1, '0838', '104.678751,31.03307', 2320);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2323, 510626, '罗江县', 1, '0838', '104.510249,31.317045', 2320);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2324, 510681, '广汉市', 1, '0838', '104.282429,30.977119', 2320);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2325, 510682, '什邡市', 1, '0838', '104.167501,31.12678', 2320);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2326, 510683, '绵竹市', 1, '0838', '104.22075,31.338077', 2320);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2328, 510703, '涪城区', 1, '0816', '104.756944,31.455101', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2404, 511700, '达州市', 2, '0818', '107.467758,31.209121', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2329, 510704, '游仙区', 1, '0816', '104.766392,31.473779', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2330, 510705, '安州区', 1, '0816', '104.567187,31.534886', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2331, 510722, '三台县', 1, '0816', '105.094586,31.095979', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2332, 510723, '盐亭县', 1, '0816', '105.389453,31.208362', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2333, 510725, '梓潼县', 1, '0816', '105.170845,31.642718', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2334, 510726, '北川羌族自治县', 1, '0816', '104.46797,31.617203', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2335, 510727, '平武县', 1, '0816', '104.555583,32.409675', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2412, 511800, '雅安市', 2, '0835', '103.042375,30.010602', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2336, 510781, '江油市', 1, '0816', '104.745915,31.778026', 2327);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2338, 510802, '利州区', 1, '0839', '105.845307,32.433756', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2339, 510811, '昭化区', 1, '0839', '105.962819,32.323256', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2340, 510812, '朝天区', 1, '0839', '105.882642,32.651336', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2341, 510821, '旺苍县', 1, '0839', '106.289983,32.229058', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2342, 510822, '青川县', 1, '0839', '105.238842,32.575484', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2343, 510823, '剑阁县', 1, '0839', '105.524766,32.287722', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2344, 510824, '苍溪县', 1, '0839', '105.934756,31.731709', 2337);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2421, 511900, '巴中市', 2, '0827', '106.747477,31.867903', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2346, 510903, '船山区', 1, '0825', '105.568297,30.525475', 2345);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2347, 510904, '安居区', 1, '0825', '105.456342,30.355379', 2345);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2348, 510921, '蓬溪县', 1, '0825', '105.70757,30.757575', 2345);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2349, 510922, '射洪县', 1, '0825', '105.388412,30.871131', 2345);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2350, 510923, '大英县', 1, '0825', '105.236923,30.594409', 2345);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2427, 512000, '资阳市', 2, '0832', '104.627636,30.128901', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2352, 511002, '市中区', 1, '1832', '105.067597,29.587053', 2351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2353, 511011, '东兴区', 1, '1832', '105.075489,29.592756', 2351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2354, 511024, '威远县', 1, '1832', '104.668879,29.52744', 2351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2431, 513200, '阿坝藏族羌族自治州', 2, '0837', '102.224653,31.899413', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2355, 511025, '资中县', 1, '1832', '104.851944,29.764059', 2351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2356, 511028, '隆昌市', 1, '1832', '105.287612,29.339476', 2351);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2358, 511102, '市中区', 1, '0833', '103.761329,29.555374', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2359, 511111, '沙湾区', 1, '0833', '103.549991,29.413091', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2360, 511112, '五通桥区', 1, '0833', '103.818014,29.406945', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2361, 511113, '金口河区', 1, '0833', '103.07862,29.244345', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2362, 511123, '犍为县', 1, '0833', '103.949326,29.20817', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2363, 511124, '井研县', 1, '0833', '104.069726,29.651287', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2364, 511126, '夹江县', 1, '0833', '103.571656,29.73763', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2365, 511129, '沐川县', 1, '0833', '103.902334,28.956647', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2366, 511132, '峨边彝族自治县', 1, '0833', '103.262048,29.230425', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2367, 511133, '马边彝族自治县', 1, '0833', '103.546347,28.83552', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2368, 511181, '峨眉山市', 1, '0833', '103.484503,29.601198', 2357);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2445, 513300, '甘孜藏族自治州', 2, '0836', '101.96231,30.04952', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2370, 511302, '顺庆区', 1, '0817', '106.09245,30.796803', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2371, 511303, '高坪区', 1, '0817', '106.118808,30.781623', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2372, 511304, '嘉陵区', 1, '0817', '106.071876,30.758823', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2373, 511321, '南部县', 1, '0817', '106.036584,31.347467', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2374, 511322, '营山县', 1, '0817', '106.565519,31.076579', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2375, 511323, '蓬安县', 1, '0817', '106.412136,31.029091', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2376, 511324, '仪陇县', 1, '0817', '106.303042,31.271561', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2377, 511325, '西充县', 1, '0817', '105.90087,30.995683', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2378, 511381, '阆中市', 1, '0817', '106.005046,31.558356', 2369);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2380, 511402, '东坡区', 1, '1833', '103.831863,30.042308', 2379);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2381, 511403, '彭山区', 1, '1833', '103.872949,30.193056', 2379);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2382, 511421, '仁寿县', 1, '1833', '104.133995,29.995635', 2379);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2383, 511423, '洪雅县', 1, '1833', '103.372863,29.90489', 2379);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2384, 511424, '丹棱县', 1, '1833', '103.512783,30.01521', 2379);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2385, 511425, '青神县', 1, '1833', '103.846688,29.831357', 2379);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2387, 511502, '翠屏区', 1, '0831', '104.620009,28.765689', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2388, 511503, '南溪区', 1, '0831', '104.969152,28.846382', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2389, 511521, '宜宾县', 1, '0831', '104.533212,28.690045', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2464, 513400, '凉山彝族自治州', 2, '0834', '102.267712,27.88157', 2277);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2390, 511523, '江安县', 1, '0831', '105.066879,28.723855', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2391, 511524, '长宁县', 1, '0831', '104.921174,28.582169', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2392, 511525, '高县', 1, '0831', '104.517748,28.436166', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2393, 511526, '珙县', 1, '0831', '104.709202,28.43863', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2394, 511527, '筠连县', 1, '0831', '104.512025,28.167831', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2395, 511528, '兴文县', 1, '0831', '105.236325,28.303614', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2396, 511529, '屏山县', 1, '0831', '104.345974,28.828482', 2386);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2398, 511602, '广安区', 1, '0826', '106.641662,30.473913', 2397);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2399, 511603, '前锋区', 1, '0826', '106.886143,30.495804', 2397);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2400, 511621, '岳池县', 1, '0826', '106.440114,30.537863', 2397);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2401, 511622, '武胜县', 1, '0826', '106.295764,30.348772', 2397);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2402, 511623, '邻水县', 1, '0826', '106.93038,30.334768', 2397);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2403, 511681, '华蓥市', 1, '0826', '106.7831,30.390188', 2397);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2405, 511702, '通川区', 1, '0818', '107.504928,31.214715', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2480, 513436, '美姑县', 1, '0834', '103.132179,28.32864', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2406, 511703, '达川区', 1, '0818', '107.511749,31.196157', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2483, 520100, '贵阳市', 2, '0851', '106.630153,26.647661', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2407, 511722, '宣汉县', 1, '0818', '107.72719,31.353835', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2408, 511723, '开江县', 1, '0818', '107.868736,31.082986', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2409, 511724, '大竹县', 1, '0818', '107.204795,30.73641', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2410, 511725, '渠县', 1, '0818', '106.97303,30.836618', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2411, 511781, '万源市', 1, '0818', '108.034657,32.081631', 2404);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2413, 511802, '雨城区', 1, '0835', '103.033026,30.005461', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2414, 511803, '名山区', 1, '0835', '103.109184,30.069954', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2415, 511822, '荥经县', 1, '0835', '102.846737,29.792931', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2416, 511823, '汉源县', 1, '0835', '102.645467,29.347192', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2417, 511824, '石棉县', 1, '0835', '102.359462,29.227874', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2494, 520200, '六盘水市', 2, '0858', '104.830458,26.592707', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2418, 511825, '天全县', 1, '0835', '102.758317,30.066712', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2419, 511826, '芦山县', 1, '0835', '102.932385,30.142307', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2420, 511827, '宝兴县', 1, '0835', '102.815403,30.37641', 2412);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2422, 511902, '巴州区', 1, '0827', '106.768878,31.851478', 2421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2499, 520300, '遵义市', 2, '0852', '106.927389,27.725654', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2423, 511903, '恩阳区', 1, '0827', '106.654386,31.787186', 2421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2424, 511921, '通江县', 1, '0827', '107.245033,31.911705', 2421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2425, 511922, '南江县', 1, '0827', '106.828697,32.346589', 2421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2426, 511923, '平昌县', 1, '0827', '107.104008,31.560874', 2421);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2428, 512002, '雁江区', 1, '0832', '104.677091,30.108216', 2427);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2429, 512021, '安岳县', 1, '0832', '105.35534,30.103107', 2427);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2430, 512022, '乐至县', 1, '0832', '105.02019,30.276121', 2427);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2432, 513201, '马尔康市', 1, '0837', '102.20652,31.905693', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2433, 513221, '汶川县', 1, '0837', '103.590179,31.476854', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2434, 513222, '理县', 1, '0837', '103.164661,31.435174', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2435, 513223, '茂县', 1, '0837', '103.853363,31.681547', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2436, 513224, '松潘县', 1, '0837', '103.604698,32.655325', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2437, 513225, '九寨沟县', 1, '0837', '104.243841,33.252056', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2438, 513226, '金川县', 1, '0837', '102.063829,31.476277', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2514, 520400, '安顺市', 2, '0853', '105.947594,26.253088', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2439, 513227, '小金县', 1, '0837', '102.362984,30.995823', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2440, 513228, '黑水县', 1, '0837', '102.990108,32.061895', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2441, 513230, '壤塘县', 1, '0837', '100.978526,32.265796', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2442, 513231, '阿坝县', 1, '0837', '101.706655,32.902459', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2443, 513232, '若尔盖县', 1, '0837', '102.967826,33.578159', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2444, 513233, '红原县', 1, '0837', '102.544405,32.790891', 2431);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2521, 520500, '毕节市', 2, '0857', '105.291702,27.283908', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2446, 513301, '康定市', 1, '0836', '101.957146,29.998435', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2447, 513322, '泸定县', 1, '0836', '102.234617,29.91416', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2448, 513323, '丹巴县', 1, '0836', '101.890358,30.878577', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2449, 513324, '九龙县', 1, '0836', '101.507294,29.000347', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2450, 513325, '雅江县', 1, '0836', '101.014425,30.031533', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2451, 513326, '道孚县', 1, '0836', '101.125237,30.979545', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2452, 513327, '炉霍县', 1, '0836', '100.676372,31.39179', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2453, 513328, '甘孜县', 1, '0836', '99.99267,31.622933', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2530, 520600, '铜仁市', 2, '0856', '109.189598,27.731514', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2454, 513329, '新龙县', 1, '0836', '100.311368,30.939169', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2455, 513330, '德格县', 1, '0836', '98.580914,31.806118', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2456, 513331, '白玉县', 1, '0836', '98.824182,31.209913', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2457, 513332, '石渠县', 1, '0836', '98.102914,32.97896', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2458, 513333, '色达县', 1, '0836', '100.332743,32.268129', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2459, 513334, '理塘县', 1, '0836', '100.269817,29.996049', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2460, 513335, '巴塘县', 1, '0836', '99.110712,30.004677', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2461, 513336, '乡城县', 1, '0836', '99.798435,28.931172', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2462, 513337, '稻城县', 1, '0836', '100.298403,29.037007', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2463, 513338, '得荣县', 1, '0836', '99.286335,28.713036', 2445);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2541, 522300, '黔西南布依族苗族自治州', 2, '0859', '104.906397,25.087856', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2465, 513401, '西昌市', 1, '0834', '102.264449,27.894504', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2466, 513422, '木里藏族自治县', 1, '0834', '101.280205,27.928835', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2467, 513423, '盐源县', 1, '0834', '101.509188,27.422645', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2468, 513424, '德昌县', 1, '0834', '102.17567,27.402839', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2469, 513425, '会理县', 1, '0834', '102.244683,26.655026', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2470, 513426, '会东县', 1, '0834', '102.57796,26.634669', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2471, 513427, '宁南县', 1, '0834', '102.751745,27.061189', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2472, 513428, '普格县', 1, '0834', '102.540901,27.376413', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2550, 522600, '黔东南苗族侗族自治州', 2, '0855', '107.982874,26.583457', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2473, 513429, '布拖县', 1, '0834', '102.812061,27.706061', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2474, 513430, '金阳县', 1, '0834', '103.248772,27.69686', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2475, 513431, '昭觉县', 1, '0834', '102.840264,28.015333', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2476, 513432, '喜德县', 1, '0834', '102.412518,28.306726', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2477, 513433, '冕宁县', 1, '0834', '102.17701,28.549656', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2478, 513434, '越西县', 1, '0834', '102.50768,28.639801', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2479, 513435, '甘洛县', 1, '0834', '102.771504,28.959157', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2481, 513437, '雷波县', 1, '0834', '103.571696,28.262682', 2464);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2484, 520102, '南明区', 1, '0851', '106.714374,26.567944', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2485, 520103, '云岩区', 1, '0851', '106.724494,26.604688', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2486, 520111, '花溪区', 1, '0851', '106.67026,26.409817', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2487, 520112, '乌当区', 1, '0851', '106.750625,26.630845', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2488, 520113, '白云区', 1, '0851', '106.623007,26.678561', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2489, 520115, '观山湖区', 1, '0851', '106.622453,26.60145', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2490, 520121, '开阳县', 1, '0851', '106.965089,27.057764', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2491, 520122, '息烽县', 1, '0851', '106.740407,27.090479', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2492, 520123, '修文县', 1, '0851', '106.592108,26.838926', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2589, 530124, '富民县', 1, '0871', '102.4976,25.221935', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2567, 522700, '黔南布依族苗族自治州', 2, '0854', '107.522171,26.253275', 2482);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2493, 520181, '清镇市', 1, '0851', '106.470714,26.556079', 2483);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2495, 520201, '钟山区', 1, '0858', '104.843555,26.574979', 2494);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2496, 520203, '六枝特区', 1, '0858', '105.476608,26.213108', 2494);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2497, 520221, '水城县', 1, '0858', '104.95783,26.547904', 2494);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2498, 520222, '盘州市', 1, '0858', '104.471375,25.709852', 2494);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2500, 520302, '红花岗区', 1, '0852', '106.8937,27.644754', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2501, 520303, '汇川区', 1, '0852', '106.93427,27.750125', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2502, 520304, '播州区', 1, '0852', '106.829574,27.536298', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2503, 520322, '桐梓县', 1, '0852', '106.825198,28.133311', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2504, 520323, '绥阳县', 1, '0852', '107.191222,27.946222', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2505, 520324, '正安县', 1, '0852', '107.453945,28.553285', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2506, 520325, '道真仡佬族苗族自治县', 1, '0852', '107.613133,28.862425', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2507, 520326, '务川仡佬族苗族自治县', 1, '0852', '107.898956,28.563086', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2581, 530100, '昆明市', 2, '0871', '102.832891,24.880095', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2508, 520327, '凤冈县', 1, '0852', '107.716355,27.954695', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2509, 520328, '湄潭县', 1, '0852', '107.465407,27.749055', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2510, 520329, '余庆县', 1, '0852', '107.905197,27.215491', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2511, 520330, '习水县', 1, '0852', '106.197137,28.33127', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2512, 520381, '赤水市', 1, '0852', '105.697472,28.590337', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2513, 520382, '仁怀市', 1, '0852', '106.40109,27.792514', 2499);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2515, 520402, '西秀区', 1, '0853', '105.965116,26.245315', 2514);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2516, 520403, '平坝区', 1, '0853', '106.256412,26.405715', 2514);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2517, 520422, '普定县', 1, '0853', '105.743277,26.301565', 2514);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2518, 520423, '镇宁布依族苗族自治县', 1, '0853', '105.770283,26.058086', 2514);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2519, 520424, '关岭布依族苗族自治县', 1, '0853', '105.61933,25.94361', 2514);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2520, 520425, '紫云苗族布依族自治县', 1, '0853', '106.084441,25.751047', 2514);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2522, 520502, '七星关区', 1, '0857', '105.30474,27.298458', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2523, 520521, '大方县', 1, '0857', '105.613037,27.141735', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2596, 530300, '曲靖市', 2, '0874', '103.796167,25.489999', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2524, 520522, '黔西县', 1, '0857', '106.033544,27.007713', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2525, 520523, '金沙县', 1, '0857', '106.220227,27.459214', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2526, 520524, '织金县', 1, '0857', '105.770542,26.663449', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2527, 520525, '纳雍县', 1, '0857', '105.382714,26.777645', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2528, 520526, '威宁彝族回族苗族自治县', 1, '0857', '104.253071,26.873806', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2529, 520527, '赫章县', 1, '0857', '104.727418,27.123078', 2521);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2531, 520602, '碧江区', 1, '0856', '109.263998,27.815927', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2532, 520603, '万山区', 1, '0856', '109.213644,27.517896', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2533, 520621, '江口县', 1, '0856', '108.839557,27.69965', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2606, 530400, '玉溪市', 2, '0877', '102.527197,24.347324', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2534, 520622, '玉屏侗族自治县', 1, '0856', '108.906411,27.235813', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2535, 520623, '石阡县', 1, '0856', '108.223612,27.513829', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2536, 520624, '思南县', 1, '0856', '108.253882,27.93755', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2537, 520625, '印江土家族苗族自治县', 1, '0856', '108.409751,27.994246', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2538, 520626, '德江县', 1, '0856', '108.119807,28.263963', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2539, 520627, '沿河土家族自治县', 1, '0856', '108.50387,28.563927', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2540, 520628, '松桃苗族自治县', 1, '0856', '109.202886,28.154071', 2530);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2542, 522301, '兴义市', 1, '0859', '104.895467,25.09204', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2543, 522322, '兴仁县', 1, '0859', '105.186237,25.435183', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2616, 530500, '保山市', 2, '0875', '99.161761,25.112046', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2544, 522323, '普安县', 1, '0859', '104.953062,25.784135', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2545, 522324, '晴隆县', 1, '0859', '105.218991,25.834783', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2546, 522325, '贞丰县', 1, '0859', '105.649864,25.38576', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2547, 522326, '望谟县', 1, '0859', '106.099617,25.178421', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2548, 522327, '册亨县', 1, '0859', '105.811592,24.983663', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2622, 530600, '昭通市', 2, '0870', '103.717465,27.338257', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2549, 522328, '安龙县', 1, '0859', '105.442701,25.099014', 2541);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2551, 522601, '凯里市', 1, '0855', '107.97754,26.582963', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2552, 522622, '黄平县', 1, '0855', '107.916411,26.905396', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2553, 522623, '施秉县', 1, '0855', '108.124379,27.03292', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2554, 522624, '三穗县', 1, '0855', '108.675267,26.952967', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2555, 522625, '镇远县', 1, '0855', '108.429534,27.049497', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2556, 522626, '岑巩县', 1, '0855', '108.81606,27.173887', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2557, 522627, '天柱县', 1, '0855', '109.207751,26.909639', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2558, 522628, '锦屏县', 1, '0855', '109.200534,26.676233', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2559, 522629, '剑河县', 1, '0855', '108.441501,26.728274', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2560, 522630, '台江县', 1, '0855', '108.321245,26.667525', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2634, 530700, '丽江市', 2, '0888', '100.22775,26.855047', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2561, 522631, '黎平县', 1, '0855', '109.136932,26.230706', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2562, 522632, '榕江县', 1, '0855', '108.52188,25.931893', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2563, 522633, '从江县', 1, '0855', '108.905329,25.753009', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2564, 522634, '雷山县', 1, '0855', '108.07754,26.378442', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2565, 522635, '麻江县', 1, '0855', '107.589359,26.491105', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2640, 530800, '普洱市', 2, '0879', '100.966156,22.825155', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2566, 522636, '丹寨县', 1, '0855', '107.788727,26.19832', 2550);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2568, 522701, '都匀市', 1, '0854', '107.518847,26.259427', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2569, 522702, '福泉市', 1, '0854', '107.520386,26.686335', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2570, 522722, '荔波县', 1, '0854', '107.898882,25.423895', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2571, 522723, '贵定县', 1, '0854', '107.232793,26.557089', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2572, 522725, '瓮安县', 1, '0854', '107.470942,27.078441', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2573, 522726, '独山县', 1, '0854', '107.545048,25.822132', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2574, 522727, '平塘县', 1, '0854', '107.322323,25.822349', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2575, 522728, '罗甸县', 1, '0854', '106.751589,25.426173', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2576, 522729, '长顺县', 1, '0854', '106.441805,26.025626', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2651, 530900, '临沧市', 2, '0883', '100.08879,23.883955', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2652, 530902, '临翔区', 1, '0883', '100.082523,23.895137', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2577, 522730, '龙里县', 1, '0854', '106.979524,26.453154', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2578, 522731, '惠水县', 1, '0854', '106.656442,26.13278', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2579, 522732, '三都水族自治县', 1, '0854', '107.869749,25.983202', 2567);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2582, 530102, '五华区', 1, '0871', '102.707262,25.043635', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2583, 530103, '盘龙区', 1, '0871', '102.751941,25.116465', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2584, 530111, '官渡区', 1, '0871', '102.749026,24.950231', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2660, 532300, '楚雄彝族自治州', 2, '0878', '101.527992,25.045513', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2585, 530112, '西山区', 1, '0871', '102.664382,25.038604', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2586, 530113, '东川区', 1, '0871', '103.187824,26.082873', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2587, 530114, '呈贡区', 1, '0871', '102.821675,24.885587', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2588, 530115, '晋宁区', 1, '0871', '102.595412,24.66974', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2590, 530125, '宜良县', 1, '0871', '103.141603,24.919839', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2591, 530126, '石林彝族自治县', 1, '0871', '103.290536,24.771761', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2592, 530127, '嵩明县', 1, '0871', '103.036908,25.338643', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2593, 530128, '禄劝彝族苗族自治县', 1, '0871', '102.471518,25.551332', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2594, 530129, '寻甸回族彝族自治县', 1, '0871', '103.256615,25.558201', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2595, 530181, '安宁市', 1, '0871', '102.478494,24.919493', 2581);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2671, 532500, '红河哈尼族彝族自治州', 2, '0873', '103.374893,23.363245', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2597, 530302, '麒麟区', 1, '0874', '103.80474,25.495326', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2598, 530303, '沾益区', 1, '0874', '103.822324,25.600507', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2599, 530321, '马龙县', 1, '0874', '103.578478,25.42805', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2600, 530322, '陆良县', 1, '0874', '103.666663,25.030051', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2601, 530323, '师宗县', 1, '0874', '103.985321,24.822233', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2602, 530324, '罗平县', 1, '0874', '104.308675,24.884626', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2603, 530325, '富源县', 1, '0874', '104.255014,25.674238', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2604, 530326, '会泽县', 1, '0874', '103.297386,26.417345', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2605, 530381, '宣威市', 1, '0874', '104.10455,26.219735', 2596);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2607, 530402, '红塔区', 1, '0877', '102.540122,24.341215', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2608, 530403, '江川区', 1, '0877', '102.75344,24.287485', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2609, 530422, '澄江县', 1, '0877', '102.904629,24.675689', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2610, 530423, '通海县', 1, '0877', '102.725452,24.111048', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2685, 532600, '文山壮族苗族自治州', 2, '0876', '104.216248,23.400733', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2611, 530424, '华宁县', 1, '0877', '102.928835,24.19276', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2612, 530425, '易门县', 1, '0877', '102.162531,24.671651', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2613, 530426, '峨山彝族自治县', 1, '0877', '102.405819,24.168957', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2614, 530427, '新平彝族傣族自治县', 1, '0877', '101.990157,24.07005', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2615, 530428, '元江哈尼族彝族傣族自治县', 1, '0877', '101.998103,23.596503', 2606);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2617, 530502, '隆阳区', 1, '0875', '99.165607,25.121154', 2616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2618, 530521, '施甸县', 1, '0875', '99.189221,24.723064', 2616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2619, 530523, '龙陵县', 1, '0875', '98.689261,24.586794', 2616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2694, 532800, '西双版纳傣族自治州', 2, '0691', '100.796984,22.009113', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2620, 530524, '昌宁县', 1, '0875', '99.605142,24.827839', 2616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2621, 530581, '腾冲市', 1, '0875', '98.490966,25.020439', 2616);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2623, 530602, '昭阳区', 1, '0870', '103.706539,27.320075', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2698, 532900, '大理白族自治州', 2, '0872', '100.267638,25.606486', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2624, 530621, '鲁甸县', 1, '0870', '103.558042,27.186659', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2625, 530622, '巧家县', 1, '0870', '102.930164,26.90846', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2626, 530623, '盐津县', 1, '0870', '104.234441,28.10871', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2627, 530624, '大关县', 1, '0870', '103.891146,27.747978', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2628, 530625, '永善县', 1, '0870', '103.638067,28.229112', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2629, 530626, '绥江县', 1, '0870', '103.968978,28.592099', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2630, 530627, '镇雄县', 1, '0870', '104.87376,27.441622', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2631, 530628, '彝良县', 1, '0870', '104.048289,27.625418', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2632, 530629, '威信县', 1, '0870', '105.049027,27.8469', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2633, 530630, '水富县', 1, '0870', '104.41603,28.62988', 2622);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2635, 530702, '古城区', 1, '0888', '100.225784,26.876927', 2634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2636, 530721, '玉龙纳西族自治县', 1, '0888', '100.236954,26.821459', 2634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2711, 533100, '德宏傣族景颇族自治州', 2, '0692', '98.584895,24.433353', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2637, 530722, '永胜县', 1, '0888', '100.750826,26.684225', 2634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2638, 530723, '华坪县', 1, '0888', '101.266195,26.629211', 2634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2639, 530724, '宁蒗彝族自治县', 1, '0888', '100.852001,27.28207', 2634);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2641, 530802, '思茅区', 1, '0879', '100.977256,22.787115', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2642, 530821, '宁洱哈尼族彝族自治县', 1, '0879', '101.045837,23.048401', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2717, 533300, '怒江傈僳族自治州', 2, '0886', '98.8566,25.817555', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2643, 530822, '墨江哈尼族自治县', 1, '0879', '101.692461,23.431894', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2644, 530823, '景东彝族自治县', 1, '0879', '100.833877,24.446731', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2645, 530824, '景谷傣族彝族自治县', 1, '0879', '100.702871,23.497028', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2646, 530825, '镇沅彝族哈尼族拉祜族自治县', 1, '0879', '101.108595,24.004441', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2722, 533400, '迪庆藏族自治州', 2, '0887', '99.702583,27.818807', 2580);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2647, 530826, '江城哈尼族彝族自治县', 1, '0879', '101.86212,22.585867', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2648, 530827, '孟连傣族拉祜族佤族自治县', 1, '0879', '99.584157,22.329099', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2649, 530828, '澜沧拉祜族自治县', 1, '0879', '99.931975,22.555904', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2650, 530829, '西盟佤族自治县', 1, '0879', '99.590123,22.644508', 2640);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2727, 540100, '拉萨市', 2, '0891', '91.172148,29.652341', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2653, 530921, '凤庆县', 1, '0883', '99.928459,24.580424', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2654, 530922, '云县', 1, '0883', '100.129354,24.44422', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2655, 530923, '永德县', 1, '0883', '99.259339,24.018357', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2656, 530924, '镇康县', 1, '0883', '98.825284,23.762584', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2657, 530925, '双江拉祜族佤族布朗族傣族自治县', 1, '0883', '99.827697,23.473499', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2658, 530926, '耿马傣族佤族自治县', 1, '0883', '99.397126,23.538092', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2659, 530927, '沧源佤族自治县', 1, '0883', '99.246196,23.146712', 2651);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2736, 540200, '日喀则市', 2, '0892', '88.880583,29.266869', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2738, 540221, '南木林县', 1, '0892', '89.099242,29.68233', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2661, 532301, '楚雄市', 1, '0878', '101.545906,25.032889', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2662, 532322, '双柏县', 1, '0878', '101.641937,24.688875', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2663, 532323, '牟定县', 1, '0878', '101.546566,25.313121', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2664, 532324, '南华县', 1, '0878', '101.273577,25.192293', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2665, 532325, '姚安县', 1, '0878', '101.241728,25.504173', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2666, 532326, '大姚县', 1, '0878', '101.336617,25.729513', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2667, 532327, '永仁县', 1, '0878', '101.666132,26.049464', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2668, 532328, '元谋县', 1, '0878', '101.87452,25.704338', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2669, 532329, '武定县', 1, '0878', '102.404337,25.530389', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2670, 532331, '禄丰县', 1, '0878', '102.079027,25.150111', 2660);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2672, 532501, '个旧市', 1, '0873', '103.160034,23.359121', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2673, 532502, '开远市', 1, '0873', '103.266624,23.714523', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2674, 532503, '蒙自市', 1, '0873', '103.364905,23.396201', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2675, 532504, '弥勒市', 1, '0873', '103.414874,24.411912', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2676, 532523, '屏边苗族自治县', 1, '0873', '103.687612,22.983559', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2755, 540300, '昌都市', 2, '0895', '97.17202,31.140969', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2677, 532524, '建水县', 1, '0873', '102.826557,23.6347', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2678, 532525, '石屏县', 1, '0873', '102.494983,23.705936', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2679, 532527, '泸西县', 1, '0873', '103.766196,24.532025', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2680, 532528, '元阳县', 1, '0873', '102.835223,23.219932', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2681, 532529, '红河县', 1, '0873', '102.4206,23.369161', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2682, 532530, '金平苗族瑶族傣族自治县', 1, '0873', '103.226448,22.779543', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2683, 532531, '绿春县', 1, '0873', '102.392463,22.993717', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2684, 532532, '河口瑶族自治县', 1, '0873', '103.93952,22.529645', 2671);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2686, 532601, '文山市', 1, '0876', '104.232665,23.386527', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2687, 532622, '砚山县', 1, '0876', '104.337211,23.605768', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2688, 532623, '西畴县', 1, '0876', '104.672597,23.437782', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2767, 540400, '林芝市', 2, '0894', '94.36149,29.649128', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2689, 532624, '麻栗坡县', 1, '0876', '104.702799,23.125714', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2690, 532625, '马关县', 1, '0876', '104.394157,23.012915', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2691, 532626, '丘北县', 1, '0876', '104.166587,24.051746', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2692, 532627, '广南县', 1, '0876', '105.055107,24.046386', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2693, 532628, '富宁县', 1, '0876', '105.630999,23.625283', 2685);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2695, 532801, '景洪市', 1, '0691', '100.799545,22.011928', 2694);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2696, 532822, '勐海县', 1, '0691', '100.452547,21.957353', 2694);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2775, 540500, '山南市', 2, '0893', '91.773134,29.237137', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2697, 532823, '勐腊县', 1, '0691', '101.564635,21.459233', 2694);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2699, 532901, '大理市', 1, '0872', '100.30127,25.678068', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2700, 532922, '漾濞彝族自治县', 1, '0872', '99.958015,25.670148', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2701, 532923, '祥云县', 1, '0872', '100.550945,25.48385', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2702, 532924, '宾川县', 1, '0872', '100.590473,25.829828', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2703, 532925, '弥渡县', 1, '0872', '100.49099,25.343804', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2704, 532926, '南涧彝族自治县', 1, '0872', '100.509035,25.04351', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2705, 532927, '巍山彝族回族自治县', 1, '0872', '100.307174,25.227212', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2706, 532928, '永平县', 1, '0872', '99.541236,25.464681', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2707, 532929, '云龙县', 1, '0872', '99.37112,25.885595', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2708, 532930, '洱源县', 1, '0872', '99.951053,26.11116', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2709, 532931, '剑川县', 1, '0872', '99.905559,26.537033', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2788, 542400, '那曲地区', 2, '0896', '92.052064,31.476479', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2710, 532932, '鹤庆县', 1, '0872', '100.176498,26.560231', 2698);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2712, 533102, '瑞丽市', 1, '0692', '97.85559,24.017958', 2711);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2713, 533103, '芒市', 1, '0692', '98.588086,24.43369', 2711);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2714, 533122, '梁河县', 1, '0692', '98.296657,24.804232', 2711);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2715, 533123, '盈江县', 1, '0692', '97.931936,24.705164', 2711);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2716, 533124, '陇川县', 1, '0692', '97.792104,24.182965', 2711);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2718, 533301, '泸水市', 1, '0886', '98.857977,25.822879', 2717);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2719, 533323, '福贡县', 1, '0886', '98.869132,26.901831', 2717);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2720, 533324, '贡山独龙族怒族自治县', 1, '0886', '98.665964,27.740999', 2717);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2721, 533325, '兰坪白族普米族自治县', 1, '0886', '99.416677,26.453571', 2717);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2723, 533401, '香格里拉市', 1, '0887', '99.700904,27.829578', 2722);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2800, 542500, '阿里地区', 2, '0897', '80.105804,32.501111', 2726);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2724, 533422, '德钦县', 1, '0887', '98.911559,28.486163', 2722);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2725, 533423, '维西傈僳族自治县', 1, '0887', '99.287173,27.177161', 2722);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2728, 540102, '城关区', 1, '0891', '91.140552,29.654838', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2729, 540103, '堆龙德庆区', 1, '0891', '91.003339,29.646063', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2730, 540121, '林周县', 1, '0891', '91.265287,29.893545', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2731, 540122, '当雄县', 1, '0891', '91.101162,30.473118', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2732, 540123, '尼木县', 1, '0891', '90.164524,29.431831', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2733, 540124, '曲水县', 1, '0891', '90.743853,29.353058', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2809, 610100, '西安市', 2, '029', '108.93977,34.341574', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2734, 540126, '达孜县', 1, '0891', '91.349867,29.66941', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2735, 540127, '墨竹工卡县', 1, '0891', '91.730732,29.834111', 2727);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2737, 540202, '桑珠孜区', 1, '0892', '88.898483,29.24779', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2739, 540222, '江孜县', 1, '0892', '89.605627,28.911626', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2740, 540223, '定日县', 1, '0892', '87.12612,28.658743', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2741, 540224, '萨迦县', 1, '0892', '88.021674,28.899664', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2742, 540225, '拉孜县', 1, '0892', '87.63704,29.081659', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2743, 540226, '昂仁县', 1, '0892', '87.236051,29.294802', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2744, 540227, '谢通门县', 1, '0892', '88.261664,29.432476', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2745, 540228, '白朗县', 1, '0892', '89.261977,29.107688', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2746, 540229, '仁布县', 1, '0892', '89.841983,29.230933', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2747, 540230, '康马县', 1, '0892', '89.681663,28.555627', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2748, 540231, '定结县', 1, '0892', '87.765872,28.364159', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2823, 610200, '铜川市', 2, '0919', '108.945019,34.897887', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2749, 540232, '仲巴县', 1, '0892', '84.03153,29.770279', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2750, 540233, '亚东县', 1, '0892', '88.907093,27.484806', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2828, 610300, '宝鸡市', 2, '0917', '107.237743,34.363184', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2751, 540234, '吉隆县', 1, '0892', '85.297534,28.852393', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2752, 540235, '聂拉木县', 1, '0892', '85.982237,28.155186', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2753, 540236, '萨嘎县', 1, '0892', '85.232941,29.328818', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2754, 540237, '岗巴县', 1, '0892', '88.520031,28.274601', 2736);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2756, 540302, '卡若区', 1, '0895', '97.196021,31.112087', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2757, 540321, '江达县', 1, '0895', '98.21843,31.499202', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2758, 540322, '贡觉县', 1, '0895', '98.27097,30.860099', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2759, 540323, '类乌齐县', 1, '0895', '96.600246,31.211601', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2760, 540324, '丁青县', 1, '0895', '95.619868,31.409024', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2761, 540325, '察雅县', 1, '0895', '97.568752,30.653943', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2762, 540326, '八宿县', 1, '0895', '96.917836,30.053209', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2763, 540327, '左贡县', 1, '0895', '97.841022,29.671069', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2841, 610400, '咸阳市', 2, '0910', '108.709136,34.32987', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2764, 540328, '芒康县', 1, '0895', '98.593113,29.679907', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2765, 540329, '洛隆县', 1, '0895', '95.825197,30.741845', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2766, 540330, '边坝县', 1, '0895', '94.7078,30.933652', 2755);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2768, 540402, '巴宜区', 1, '0894', '94.361094,29.636576', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2769, 540421, '工布江达县', 1, '0894', '93.246077,29.88528', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2770, 540422, '米林县', 1, '0894', '94.213679,29.213811', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2771, 540423, '墨脱县', 1, '0894', '95.333197,29.325298', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2772, 540424, '波密县', 1, '0894', '95.767913,29.859028', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2773, 540425, '察隅县', 1, '0894', '97.466919,28.66128', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2774, 540426, '朗县', 1, '0894', '93.074702,29.046337', 2767);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2776, 540502, '乃东区', 1, '0893', '91.761538,29.224904', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2777, 540521, '扎囊县', 1, '0893', '91.33725,29.245113', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2778, 540522, '贡嘎县', 1, '0893', '90.98414,29.289455', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2779, 540523, '桑日县', 1, '0893', '92.015818,29.259189', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2856, 610500, '渭南市', 2, '0913', '109.471094,34.52044', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2780, 540524, '琼结县', 1, '0893', '91.683881,29.024625', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2781, 540525, '曲松县', 1, '0893', '92.203738,29.062826', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2782, 540526, '措美县', 1, '0893', '91.433509,28.438202', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2783, 540527, '洛扎县', 1, '0893', '90.859992,28.385713', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2784, 540528, '加查县', 1, '0893', '92.593993,29.14029', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2785, 540529, '隆子县', 1, '0893', '92.463308,28.408548', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2786, 540530, '错那县', 1, '0893', '91.960132,27.991707', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2787, 540531, '浪卡子县', 1, '0893', '90.397977,28.968031', 2775);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2789, 542421, '那曲县', 1, '0896', '92.0535,31.469643', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2790, 542422, '嘉黎县', 1, '0896', '93.232528,30.640814', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2791, 542423, '比如县', 1, '0896', '93.679639,31.480249', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2868, 610600, '延安市', 2, '0911', '109.494112,36.651381', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2792, 542424, '聂荣县', 1, '0896', '92.303377,32.10775', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2793, 542425, '安多县', 1, '0896', '91.68233,32.265176', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2794, 542426, '申扎县', 1, '0896', '88.709852,30.930505', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2795, 542427, '索县', 1, '0896', '93.785516,31.886671', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2796, 542428, '班戈县', 1, '0896', '90.009957,31.392411', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2797, 542429, '巴青县', 1, '0896', '94.053438,31.91847', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2798, 542430, '尼玛县', 1, '0896', '87.236772,31.784701', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2799, 542431, '双湖县', 1, '0896', '88.837641,33.188514', 2788);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2801, 542521, '普兰县', 1, '0897', '81.176237,30.294402', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2802, 542522, '札达县', 1, '0897', '79.802706,31.479216', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2803, 542523, '噶尔县', 1, '0897', '80.096419,32.491488', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2804, 542524, '日土县', 1, '0897', '79.732427,33.381359', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2805, 542525, '革吉县', 1, '0897', '81.145433,32.387233', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2882, 610700, '汉中市', 2, '0916', '107.02305,33.067225', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2806, 542526, '改则县', 1, '0897', '84.06259,32.302713', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2807, 542527, '措勤县', 1, '0897', '85.151455,31.017312', 2800);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2810, 610102, '新城区', 1, '029', '108.960716,34.266447', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2811, 610103, '碑林区', 1, '029', '108.94059,34.256783', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2812, 610104, '莲湖区', 1, '029', '108.943895,34.265239', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2813, 610111, '灞桥区', 1, '029', '109.064646,34.272793', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2814, 610112, '未央区', 1, '029', '108.946825,34.29292', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2815, 610113, '雁塔区', 1, '029', '108.944644,34.214113', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2816, 610114, '阎良区', 1, '029', '109.226124,34.662232', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2817, 610115, '临潼区', 1, '029', '109.214237,34.367069', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2818, 610116, '长安区', 1, '029', '108.907173,34.158926', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2894, 610800, '榆林市', 2, '0912', '109.734474,38.285369', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2819, 610117, '高陵区', 1, '029', '109.088297,34.534829', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2820, 610122, '蓝田县', 1, '029', '109.32345,34.151298', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2821, 610124, '周至县', 1, '029', '108.222162,34.163669', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2822, 610118, '鄠邑区', 1, '029', '108.604894,34.109244', 2809);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2824, 610202, '王益区', 1, '0919', '109.075578,35.068964', 2823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2825, 610203, '印台区', 1, '0919', '109.099974,35.114492', 2823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2826, 610204, '耀州区', 1, '0919', '108.980102,34.909793', 2823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2827, 610222, '宜君县', 1, '0919', '109.116932,35.398577', 2823);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2829, 610302, '渭滨区', 1, '0917', '107.155344,34.355068', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2830, 610303, '金台区', 1, '0917', '107.146806,34.376069', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2831, 610304, '陈仓区', 1, '0917', '107.369987,34.35147', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2832, 610322, '凤翔县', 1, '0917', '107.400737,34.521217', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2907, 610900, '安康市', 2, '0915', '109.029113,32.68481', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2833, 610323, '岐山县', 1, '0917', '107.621053,34.443459', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2834, 610324, '扶风县', 1, '0917', '107.900219,34.37541', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2835, 610326, '眉县', 1, '0917', '107.749766,34.274246', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2836, 610327, '陇县', 1, '0917', '106.864397,34.89305', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2837, 610328, '千阳县', 1, '0917', '107.132441,34.642381', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2838, 610329, '麟游县', 1, '0917', '107.793524,34.677902', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2839, 610330, '凤县', 1, '0917', '106.515803,33.91091', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2840, 610331, '太白县', 1, '0917', '107.319116,34.058401', 2828);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2842, 610402, '秦都区', 1, '0910', '108.706272,34.329567', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2843, 610403, '杨陵区', 1, '0910', '108.084731,34.272117', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2918, 611000, '商洛市', 2, '0914', '109.91857,33.872726', 2808);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2917, 610929, '白河县', 1, '0915', '110.112629,32.809026', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2844, 610404, '渭城区', 1, '0910', '108.737204,34.36195', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2845, 610422, '三原县', 1, '0910', '108.940509,34.617381', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2846, 610423, '泾阳县', 1, '0910', '108.842622,34.527114', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2847, 610424, '乾县', 1, '0910', '108.239473,34.527551', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2848, 610425, '礼泉县', 1, '0910', '108.425018,34.481764', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2849, 610426, '永寿县', 1, '0910', '108.142311,34.691979', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2850, 610427, '彬县', 1, '0910', '108.077658,35.043911', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2927, 620100, '兰州市', 2, '0931', '103.834303,36.061089', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2851, 610428, '长武县', 1, '0910', '107.798757,35.205886', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2852, 610429, '旬邑县', 1, '0910', '108.333986,35.111978', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2853, 610430, '淳化县', 1, '0910', '108.580681,34.79925', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2854, 610431, '武功县', 1, '0910', '108.200398,34.260203', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2855, 610481, '兴平市', 1, '0910', '108.490475,34.29922', 2841);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2857, 610502, '临渭区', 1, '0913', '109.510175,34.499314', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2858, 610503, '华州区', 1, '0913', '109.775247,34.495915', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2859, 610522, '潼关县', 1, '0913', '110.246349,34.544296', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2936, 620200, '嘉峪关市', 2, '1937', '98.289419,39.772554', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2937, 620300, '金昌市', 2, '0935', '102.188117,38.520717', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2860, 610523, '大荔县', 1, '0913', '109.941734,34.797259', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2861, 610524, '合阳县', 1, '0913', '110.149453,35.237988', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2940, 620400, '白银市', 2, '0943', '104.138771,36.545261', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2862, 610525, '澄城县', 1, '0913', '109.93235,35.190245', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2863, 610526, '蒲城县', 1, '0913', '109.586403,34.955562', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2864, 610527, '白水县', 1, '0913', '109.590671,35.177451', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2865, 610528, '富平县', 1, '0913', '109.18032,34.751077', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2866, 610581, '韩城市', 1, '0913', '110.442846,35.476788', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2946, 620500, '天水市', 2, '0938', '105.724979,34.580885', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2867, 610582, '华阴市', 1, '0913', '110.092078,34.566079', 2856);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2869, 610602, '宝塔区', 1, '0911', '109.48976,36.585472', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2870, 610621, '延长县', 1, '0911', '110.012334,36.579313', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2871, 610622, '延川县', 1, '0911', '110.193514,36.878117', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2872, 610623, '子长县', 1, '0911', '109.675264,37.142535', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2873, 610603, '安塞区', 1, '0911', '109.328842,36.863853', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2874, 610625, '志丹县', 1, '0911', '108.768432,36.822194', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2954, 620600, '武威市', 2, '1935', '102.638201,37.928267', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2875, 610626, '吴起县', 1, '0911', '108.175933,36.927215', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2876, 610627, '甘泉县', 1, '0911', '109.351019,36.276526', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2877, 610628, '富县', 1, '0911', '109.379776,35.987953', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2878, 610629, '洛川县', 1, '0911', '109.432369,35.761974', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2959, 620700, '张掖市', 2, '0936', '100.449913,38.925548', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2879, 610630, '宜川县', 1, '0911', '110.168963,36.050178', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2880, 610631, '黄龙县', 1, '0911', '109.840314,35.584743', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2881, 610632, '黄陵县', 1, '0911', '109.262961,35.579427', 2868);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2883, 610702, '汉台区', 1, '0916', '107.031856,33.067771', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2884, 610721, '南郑县', 1, '0916', '106.93623,32.999333', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2885, 610722, '城固县', 1, '0916', '107.33393,33.157131', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2966, 620800, '平凉市', 2, '0933', '106.665061,35.542606', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2886, 610723, '洋县', 1, '0916', '107.545836,33.222738', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2887, 610724, '西乡县', 1, '0916', '107.766613,32.983101', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2888, 610725, '勉县', 1, '0916', '106.673221,33.153553', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2889, 610726, '宁强县', 1, '0916', '106.257171,32.829694', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2890, 610727, '略阳县', 1, '0916', '106.156718,33.327281', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2891, 610728, '镇巴县', 1, '0916', '107.895035,32.536704', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2892, 610729, '留坝县', 1, '0916', '106.920808,33.617571', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2974, 620900, '酒泉市', 2, '0937', '98.493927,39.732795', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2893, 610730, '佛坪县', 1, '0916', '107.990538,33.524359', 2882);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2895, 610802, '榆阳区', 1, '0912', '109.721069,38.277046', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2896, 610803, '横山区', 1, '0912', '109.294346,37.962208', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2897, 610881, '神木市', 1, '0912', '110.498939,38.842578', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2898, 610822, '府谷县', 1, '0912', '111.067276,39.028116', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2899, 610824, '靖边县', 1, '0912', '108.793988,37.599438', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2900, 610825, '定边县', 1, '0912', '107.601267,37.594612', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2982, 621000, '庆阳市', 2, '0934', '107.643571,35.70898', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2901, 610826, '绥德县', 1, '0912', '110.263362,37.50294', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2902, 610827, '米脂县', 1, '0912', '110.183754,37.755416', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2903, 610828, '佳县', 1, '0912', '110.491345,38.01951', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2904, 610829, '吴堡县', 1, '0912', '110.739673,37.452067', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2905, 610830, '清涧县', 1, '0912', '110.121209,37.088878', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2906, 610831, '子洲县', 1, '0912', '110.03525,37.610683', 2894);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2908, 610902, '汉滨区', 1, '0915', '109.026836,32.695172', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2909, 610921, '汉阴县', 1, '0915', '108.508745,32.893026', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2991, 621100, '定西市', 2, '0932', '104.592225,35.606978', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2910, 610922, '石泉县', 1, '0915', '108.247886,33.038408', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2911, 610923, '宁陕县', 1, '0915', '108.314283,33.310527', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2912, 610924, '紫阳县', 1, '0915', '108.534228,32.520246', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2913, 610925, '岚皋县', 1, '0915', '108.902049,32.307001', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2914, 610926, '平利县', 1, '0915', '109.361864,32.388854', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2915, 610927, '镇坪县', 1, '0915', '109.526873,31.883672', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2916, 610928, '旬阳县', 1, '0915', '109.361024,32.832012', 2907);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2999, 621200, '陇南市', 2, '2935', '104.960851,33.37068', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2919, 611002, '商州区', 1, '0914', '109.941839,33.862599', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2920, 611021, '洛南县', 1, '0914', '110.148508,34.090837', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2921, 611022, '丹凤县', 1, '0914', '110.32733,33.695783', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2922, 611023, '商南县', 1, '0914', '110.881807,33.530995', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2923, 611024, '山阳县', 1, '0914', '109.882289,33.532172', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2924, 611025, '镇安县', 1, '0914', '109.152892,33.423357', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3005, 621225, '西和县', 1, '2935', '105.298756,34.014215', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2925, 611026, '柞水县', 1, '0914', '109.114206,33.68611', 2918);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3009, 622900, '临夏回族自治州', 2, '0930', '103.210655,35.601352', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2928, 620102, '城关区', 1, '0931', '103.825307,36.057464', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2929, 620103, '七里河区', 1, '0931', '103.785949,36.066146', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2930, 620104, '西固区', 1, '0931', '103.627951,36.088552', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2931, 620105, '安宁区', 1, '0931', '103.719054,36.104579', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2932, 620111, '红古区', 1, '0931', '102.859323,36.345669', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2933, 620121, '永登县', 1, '0931', '103.26038,36.736513', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2934, 620122, '皋兰县', 1, '0931', '103.947377,36.332663', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2935, 620123, '榆中县', 1, '0931', '104.112527,35.843056', 2927);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3018, 623000, '甘南藏族自治州', 2, '0941', '102.910995,34.983409', 2926);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2938, 620302, '金川区', 1, '0935', '102.194015,38.521087', 2937);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2939, 620321, '永昌县', 1, '0935', '101.984458,38.243434', 2937);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2941, 620402, '白银区', 1, '0943', '104.148556,36.535398', 2940);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2942, 620403, '平川区', 1, '0943', '104.825208,36.728304', 2940);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2943, 620421, '靖远县', 1, '0943', '104.676774,36.571365', 2940);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2944, 620422, '会宁县', 1, '0943', '105.053358,35.692823', 2940);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2945, 620423, '景泰县', 1, '0943', '104.063091,37.183804', 2940);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2947, 620502, '秦州区', 1, '0938', '105.724215,34.580888', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2948, 620503, '麦积区', 1, '0938', '105.889556,34.570384', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3028, 630100, '西宁市', 2, '0971', '101.778223,36.617134', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2949, 620521, '清水县', 1, '0938', '106.137293,34.749864', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2950, 620522, '秦安县', 1, '0938', '105.674982,34.858916', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2951, 620523, '甘谷县', 1, '0938', '105.340747,34.745486', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2952, 620524, '武山县', 1, '0938', '104.890587,34.72139', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2953, 620525, '张家川回族自治县', 1, '0938', '106.204517,34.988037', 2946);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2955, 620602, '凉州区', 1, '1935', '102.642184,37.928224', 2954);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2956, 620621, '民勤县', 1, '1935', '103.093791,38.62435', 2954);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3036, 630200, '海东市', 2, '0972', '102.104287,36.502039', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2957, 620622, '古浪县', 1, '1935', '102.897533,37.47012', 2954);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2958, 620623, '天祝藏族自治县', 1, '1935', '103.141757,36.97174', 2954);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2960, 620702, '甘州区', 1, '0936', '100.415096,38.944662', 2959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2961, 620721, '肃南裕固族自治县', 1, '0936', '99.615601,38.836931', 2959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2962, 620722, '民乐县', 1, '0936', '100.812629,38.430347', 2959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2963, 620723, '临泽县', 1, '0936', '100.164283,39.152462', 2959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3043, 632200, '海北藏族自治州', 2, '0970', '100.900997,36.954413', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2964, 620724, '高台县', 1, '0936', '99.819519,39.378311', 2959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2965, 620725, '山丹县', 1, '0936', '101.088529,38.784505', 2959);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2967, 620802, '崆峒区', 1, '0933', '106.674767,35.542491', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2968, 620821, '泾川县', 1, '0933', '107.36785,35.332666', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3048, 632300, '黄南藏族自治州', 2, '0973', '102.015248,35.519548', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2969, 620822, '灵台县', 1, '0933', '107.595874,35.070027', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2970, 620823, '崇信县', 1, '0933', '107.025763,35.305596', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2971, 620824, '华亭县', 1, '0933', '106.653158,35.218292', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2972, 620825, '庄浪县', 1, '0933', '106.036686,35.202385', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3053, 632500, '海南藏族自治州', 2, '0974', '100.622692,36.296529', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2973, 620826, '静宁县', 1, '0933', '105.732556,35.521976', 2966);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2975, 620902, '肃州区', 1, '0937', '98.507843,39.744953', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2976, 620921, '金塔县', 1, '0937', '98.901252,39.983955', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2977, 620922, '瓜州县', 1, '0937', '95.782318,40.520538', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2978, 620923, '肃北蒙古族自治县', 1, '0937', '94.876579,39.51245', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3059, 632600, '果洛藏族自治州', 2, '0975', '100.244808,34.471431', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2979, 620924, '阿克塞哈萨克族自治县', 1, '0937', '94.340204,39.633943', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2980, 620981, '玉门市', 1, '0937', '97.045661,40.292106', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2981, 620982, '敦煌市', 1, '0937', '94.661941,40.142089', 2974);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2983, 621002, '西峰区', 1, '0934', '107.651077,35.730652', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2984, 621021, '庆城县', 1, '0934', '107.881802,36.016299', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2985, 621022, '环县', 1, '0934', '107.308501,36.568434', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3066, 632700, '玉树藏族自治州', 2, '0976', '97.091934,33.011674', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2986, 621023, '华池县', 1, '0934', '107.990062,36.461306', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2987, 621024, '合水县', 1, '0934', '108.019554,35.819194', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2988, 621025, '正宁县', 1, '0934', '108.359865,35.49178', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2989, 621026, '宁县', 1, '0934', '107.928371,35.502176', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2990, 621027, '镇原县', 1, '0934', '107.200832,35.677462', 2982);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2992, 621102, '安定区', 1, '0932', '104.610668,35.580629', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3073, 632800, '海西蒙古族藏族自治州', 2, '0977', '97.369751,37.377139', 3027);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2993, 621121, '通渭县', 1, '0932', '105.24206,35.210831', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2994, 621122, '陇西县', 1, '0932', '104.634983,35.00394', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2995, 621123, '渭源县', 1, '0932', '104.215467,35.136755', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2996, 621124, '临洮县', 1, '0932', '103.859565,35.394988', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2997, 621125, '漳县', 1, '0932', '104.471572,34.848444', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2998, 621126, '岷县', 1, '0932', '104.03688,34.438075', 2991);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3081, 640100, '银川市', 2, '0951', '106.230909,38.487193', 3080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3000, 621202, '武都区', 1, '2935', '104.926337,33.392211', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3001, 621221, '成县', 1, '2935', '105.742424,33.75061', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3002, 621222, '文县', 1, '2935', '104.683433,32.943815', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3003, 621223, '宕昌县', 1, '2935', '104.393385,34.047261', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3004, 621224, '康县', 1, '2935', '105.609169,33.329136', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3006, 621226, '礼县', 1, '2935', '105.17864,34.189345', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3088, 640200, '石嘴山市', 2, '0952', '106.383303,38.983236', 3080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3007, 621227, '徽县', 1, '2935', '106.08778,33.768826', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3008, 621228, '两当县', 1, '2935', '106.304966,33.908917', 2999);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3080, 640000, '宁夏回族自治区', 4, '[]', '106.259126,38.472641', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3010, 622901, '临夏市', 1, '0930', '103.243021,35.604376', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3092, 640300, '吴忠市', 2, '0953', '106.198913,37.997428', 3080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3091, 640221, '平罗县', 1, '0952', '106.523474,38.913544', 3088);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3131, 650421, '鄯善县', 1, '0995', '90.21333,42.868744', 3129);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3011, 622921, '临夏县', 1, '0930', '103.039826,35.478722', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3012, 622922, '康乐县', 1, '0930', '103.708354,35.370505', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3013, 622923, '永靖县', 1, '0930', '103.285853,35.958306', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3098, 640400, '固原市', 2, '0954', '106.24261,36.015855', 3080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3014, 622924, '广河县', 1, '0930', '103.575834,35.488051', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3015, 622925, '和政县', 1, '0930', '103.350997,35.424603', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3016, 622926, '东乡族自治县', 1, '0930', '103.389346,35.663752', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3017, 622927, '积石山保安族东乡族撒拉族自治县', 1, '0930', '102.875843,35.71766', 3009);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3019, 623001, '合作市', 1, '0941', '102.910484,35.000286', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3104, 640500, '中卫市', 2, '1953', '105.196902,37.499972', 3080);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3020, 623021, '临潭县', 1, '0941', '103.353919,34.692747', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3021, 623022, '卓尼县', 1, '0941', '103.507109,34.589588', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3022, 623023, '舟曲县', 1, '0941', '104.251482,33.793631', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3023, 623024, '迭部县', 1, '0941', '103.221869,34.055938', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3109, 659002, '阿拉尔市', 2, '1997', '81.280527,40.547653', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3110, 659005, '北屯市', 2, '1906', '87.837075,47.332643', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3111, 659008, '可克达拉市', 2, '1999', '81.044542,43.944798', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3112, 659009, '昆玉市', 2, '1903', '79.291083,37.209642', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3113, 659001, '石河子市', 2, '0993', '86.080602,44.306097', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3114, 659007, '双河市', 2, '1909', '82.353656,44.840524', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3115, 650100, '乌鲁木齐市', 2, '0991', '87.616848,43.825592', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3024, 623025, '玛曲县', 1, '0941', '102.072698,33.997712', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3025, 623026, '碌曲县', 1, '0941', '102.487327,34.590944', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3026, 623027, '夏河县', 1, '0941', '102.521807,35.202503', 3018);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3029, 630102, '城东区', 1, '0971', '101.803717,36.599744', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3030, 630103, '城中区', 1, '0971', '101.705298,36.545652', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3031, 630104, '城西区', 1, '0971', '101.765843,36.628304', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3032, 630105, '城北区', 1, '0971', '101.766228,36.650038', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3033, 630121, '大通回族土族自治县', 1, '0971', '101.685643,36.926954', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3124, 650200, '克拉玛依市', 2, '0990', '84.889207,45.579888', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3034, 630122, '湟中县', 1, '0971', '101.571667,36.500879', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3035, 630123, '湟源县', 1, '0971', '101.256464,36.682426', 3028);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3037, 630202, '乐都区', 1, '0972', '102.401724,36.482058', 3036);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3038, 630203, '平安区', 1, '0972', '102.108834,36.500563', 3036);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3129, 650400, '吐鲁番市', 2, '0995', '89.189752,42.951303', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3039, 630222, '民和回族土族自治县', 1, '0972', '102.830892,36.320321', 3036);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3040, 630223, '互助土族自治县', 1, '0972', '101.959271,36.844248', 3036);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3041, 630224, '化隆回族自治县', 1, '0972', '102.264143,36.094908', 3036);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3133, 650500, '哈密市', 2, '0902', '93.515224,42.819541', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3042, 630225, '循化撒拉族自治县', 1, '0972', '102.489135,35.851152', 3036);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3044, 632221, '门源回族自治县', 1, '0970', '101.611539,37.388746', 3043);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3045, 632222, '祁连县', 1, '0970', '100.253211,38.177112', 3043);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3137, 652300, '昌吉回族自治州', 2, '0994', '87.308224,44.011182', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3046, 632223, '海晏县', 1, '0970', '100.99426,36.896359', 3043);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3047, 632224, '刚察县', 1, '0970', '100.145833,37.32547', 3043);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3049, 632321, '同仁县', 1, '0973', '102.018323,35.516063', 3048);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3050, 632322, '尖扎县', 1, '0973', '102.04014,35.943156', 3048);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3051, 632323, '泽库县', 1, '0973', '101.466689,35.035313', 3048);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3052, 632324, '河南蒙古族自治县', 1, '0973', '101.617503,34.734568', 3048);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3054, 632521, '共和县', 1, '0974', '100.620031,36.284107', 3053);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3145, 652700, '博尔塔拉蒙古自治州', 2, '0909', '82.066363,44.906039', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3055, 632522, '同德县', 1, '0974', '100.578051,35.25479', 3053);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3056, 632523, '贵德县', 1, '0974', '101.433391,36.040166', 3053);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3057, 632524, '兴海县', 1, '0974', '99.987965,35.588612', 3053);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3058, 632525, '贵南县', 1, '0974', '100.747503,35.586714', 3053);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3150, 652800, '巴音郭楞蒙古自治州', 2, '0996', '86.145297,41.764115', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3060, 632621, '玛沁县', 1, '0975', '100.238888,34.477433', 3059);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3061, 632622, '班玛县', 1, '0975', '100.737138,32.932723', 3059);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3062, 632623, '甘德县', 1, '0975', '99.900923,33.969216', 3059);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3063, 632624, '达日县', 1, '0975', '99.651392,33.74892', 3059);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3064, 632625, '久治县', 1, '0975', '101.482831,33.429471', 3059);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3065, 632626, '玛多县', 1, '0975', '98.209206,34.915946', 3059);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3067, 632701, '玉树市', 1, '0976', '97.008784,32.993106', 3066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3068, 632722, '杂多县', 1, '0976', '95.300723,32.893185', 3066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3069, 632723, '称多县', 1, '0976', '97.110831,33.369218', 3066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3160, 652900, '阿克苏地区', 2, '0997', '80.260605,41.168779', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3070, 632724, '治多县', 1, '0976', '95.61896,33.844956', 3066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3071, 632725, '囊谦县', 1, '0976', '96.48943,32.203432', 3066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3072, 632726, '曲麻莱县', 1, '0976', '95.797367,34.126428', 3066);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3074, 632801, '格尔木市', 1, '0977', '94.928453,36.406367', 3073);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3075, 632802, '德令哈市', 1, '0977', '97.360984,37.369436', 3073);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3076, 632821, '乌兰县', 1, '0977', '98.480195,36.929749', 3073);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3077, 632822, '都兰县', 1, '0977', '98.095844,36.302496', 3073);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3078, 632823, '天峻县', 1, '0977', '99.022984,37.300851', 3073);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3079, 632825, '海西蒙古族藏族自治州直辖', 1, '0977', '95.356546,37.853328', 3073);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3170, 653000, '克孜勒苏柯尔克孜自治州', 2, '0908', '76.167819,39.714526', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3082, 640104, '兴庆区', 1, '0951', '106.28865,38.473609', 3081);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3083, 640105, '西夏区', 1, '0951', '106.161106,38.502605', 3081);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3084, 640106, '金凤区', 1, '0951', '106.239679,38.47436', 3081);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3085, 640121, '永宁县', 1, '0951', '106.253145,38.277372', 3081);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3175, 653100, '喀什地区', 2, '0998', '75.989741,39.47046', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3188, 653200, '和田地区', 2, '0903', '79.922211,37.114157', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3178, 653122, '疏勒县', 1, '0998', '76.048139,39.401384', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3179, 653123, '英吉沙县', 1, '0998', '76.175729,38.930381', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3180, 653124, '泽普县', 1, '0998', '77.259675,38.18529', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3181, 653125, '莎车县', 1, '0998', '77.245761,38.41422', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3182, 653126, '叶城县', 1, '0998', '77.413836,37.882989', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3183, 653127, '麦盖提县', 1, '0998', '77.610125,38.898001', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3197, 654000, '伊犁哈萨克自治州', 2, '0999', '81.324136,43.916823', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3184, 653128, '岳普湖县', 1, '0998', '76.8212,39.2198', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3185, 653129, '伽师县', 1, '0998', '76.723719,39.488181', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3186, 653130, '巴楚县', 1, '0998', '78.549296,39.785155', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3187, 653131, '塔什库尔干塔吉克自治县', 1, '0998', '75.229889,37.772094', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3189, 653201, '和田市', 1, '0903', '79.913534,37.112148', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3086, 640122, '贺兰县', 1, '0951', '106.349861,38.554599', 3081);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3087, 640181, '灵武市', 1, '0951', '106.340053,38.102655', 3081);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3089, 640202, '大武口区', 1, '0952', '106.367958,39.01918', 3088);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3090, 640205, '惠农区', 1, '0952', '106.781176,39.239302', 3088);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3093, 640302, '利通区', 1, '0953', '106.212613,37.98349', 3092);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3094, 640303, '红寺堡区', 1, '0953', '106.062113,37.425702', 3092);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3209, 654200, '塔城地区', 2, '0901', '82.980316,46.745364', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3095, 640323, '盐池县', 1, '0953', '107.407358,37.783205', 3092);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3096, 640324, '同心县', 1, '0953', '105.895309,36.95449', 3092);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3097, 640381, '青铜峡市', 1, '0953', '106.078817,38.021302', 3092);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3099, 640402, '原州区', 1, '0954', '106.287781,36.003739', 3098);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3100, 640422, '西吉县', 1, '0954', '105.729085,35.963912', 3098);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3101, 640423, '隆德县', 1, '0954', '106.111595,35.625914', 3098);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3102, 640424, '泾源县', 1, '0954', '106.330646,35.498159', 3098);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3217, 654300, '阿勒泰地区', 2, '0906', '88.141253,47.844924', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3103, 640425, '彭阳县', 1, '0954', '106.631809,35.858815', 3098);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3105, 640502, '沙坡头区', 1, '1953', '105.173721,37.516883', 3104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3106, 640521, '中宁县', 1, '1953', '105.685218,37.491546', 3104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3107, 640522, '海原县', 1, '1953', '105.643487,36.565033', 3104);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3116, 650102, '天山区', 1, '0991', '87.631676,43.794399', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3117, 650103, '沙依巴克区', 1, '0991', '87.598195,43.800939', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3118, 650104, '新市区', 1, '0991', '87.569431,43.855378', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3225, 659006, '铁门关市', 2, '1996', '85.501217,41.82725', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3226, 659003, '图木舒克市', 2, '1998', '79.073963,39.868965', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3227, 659004, '五家渠市', 2, '1994', '87.54324,44.166756', 3108);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3119, 650105, '水磨沟区', 1, '0991', '87.642481,43.832459', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3120, 650106, '头屯河区', 1, '0991', '87.428141,43.877664', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3121, 650107, '达坂城区', 1, '0991', '88.311099,43.363668', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3122, 650109, '米东区', 1, '0991', '87.655935,43.974784', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3123, 650121, '乌鲁木齐县', 1, '0991', '87.409417,43.47136', 3115);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3125, 650202, '独山子区', 1, '0990', '84.886974,44.328095', 3124);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3126, 650203, '克拉玛依区', 1, '0990', '84.867844,45.602525', 3124);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3127, 650204, '白碱滩区', 1, '0990', '85.131696,45.687854', 3124);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3128, 650205, '乌尔禾区', 1, '0990', '85.693742,46.089148', 3124);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3130, 650402, '高昌区', 1, '0995', '89.185877,42.942327', 3129);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3132, 650422, '托克逊县', 1, '0995', '88.653827,42.792526', 3129);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3134, 650502, '伊州区', 1, '0902', '93.514797,42.827254', 3133);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3135, 650521, '巴里坤哈萨克自治县', 1, '0902', '93.010383,43.599929', 3133);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3136, 650522, '伊吾县', 1, '0902', '94.697074,43.254978', 3133);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3138, 652301, '昌吉市', 1, '0994', '87.267532,44.014435', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3139, 652302, '阜康市', 1, '0994', '87.952991,44.164402', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3140, 652323, '呼图壁县', 1, '0994', '86.871584,44.179361', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3141, 652324, '玛纳斯县', 1, '0994', '86.20368,44.284722', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3142, 652325, '奇台县', 1, '0994', '89.593967,44.022066', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3143, 652327, '吉木萨尔县', 1, '0994', '89.180437,44.000497', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3144, 652328, '木垒哈萨克自治县', 1, '0994', '90.286028,43.834689', 3137);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3146, 652701, '博乐市', 1, '0909', '82.051004,44.853869', 3145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3147, 652702, '阿拉山口市', 1, '0909', '82.559396,45.172227', 3145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3148, 652722, '精河县', 1, '0909', '82.890656,44.599393', 3145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3149, 652723, '温泉县', 1, '0909', '81.024816,44.968856', 3145);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3151, 652801, '库尔勒市', 1, '0996', '86.174633,41.725891', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3152, 652822, '轮台县', 1, '0996', '84.252156,41.777702', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3153, 652823, '尉犁县', 1, '0996', '86.261321,41.343933', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3154, 652824, '若羌县', 1, '0996', '88.167152,39.023241', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3155, 652825, '且末县', 1, '0996', '85.529702,38.145485', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3156, 652826, '焉耆回族自治县', 1, '0996', '86.574067,42.059759', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3157, 652827, '和静县', 1, '0996', '86.384065,42.323625', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3158, 652828, '和硕县', 1, '0996', '86.876799,42.284331', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3159, 652829, '博湖县', 1, '0996', '86.631997,41.980152', 3150);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3161, 652901, '阿克苏市', 1, '0997', '80.263387,41.167548', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3162, 652922, '温宿县', 1, '0997', '80.238959,41.276688', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3027, 630000, '青海省', 4, '[]', '101.780268,36.620939', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10001, 110910, '华中地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10002, 110920, '华南地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10003, 110930, '华北地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10004, 110940, '华东地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10005, 110950, '华西地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10006, 110960, '西南地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10007, 110970, '西北地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10008, 110980, '港澳台地区', 8, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3163, 652923, '库车县', 1, '0997', '82.987312,41.714696', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3164, 652924, '沙雅县', 1, '0997', '82.781818,41.221666', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3165, 652925, '新和县', 1, '0997', '82.618736,41.551206', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3166, 652926, '拜城县', 1, '0997', '81.85148,41.795912', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3167, 652927, '乌什县', 1, '0997', '79.224616,41.222319', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3168, 652928, '阿瓦提县', 1, '0997', '80.375053,40.643647', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3169, 652929, '柯坪县', 1, '0997', '79.054497,40.501936', 3160);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3171, 653001, '阿图什市', 1, '0908', '76.1684,39.71616', 3170);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3172, 653022, '阿克陶县', 1, '0908', '75.947396,39.147785', 3170);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3173, 653023, '阿合奇县', 1, '0908', '78.446253,40.936936', 3170);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3174, 653024, '乌恰县', 1, '0908', '75.259227,39.71931', 3170);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3176, 653101, '喀什市', 1, '0998', '75.99379,39.467685', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3177, 653121, '疏附县', 1, '0998', '75.862813,39.375043', 3175);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (37, 130000, '河北省', 4, '[]', '114.530235,38.037433', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2277, 510000, '四川省', 4, '[]', '104.075809,30.651239', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1505, 410000, '河南省', 4, '[]', '113.753394,34.765869', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (464, 210000, '辽宁省', 4, '[]', '123.431382,41.836175', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1021, 340000, '安徽省', 4, '[]', '117.329949,31.733806', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (792, 310000, '上海市', 4, '021', '121.473662,31.230372', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3229, 810000, '香港特别行政区', 4, '1852', '114.171203,22.277468', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (348, 150000, '内蒙古自治区', 4, '[]', '111.76629,40.81739', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1935, 440000, '广东省', 4, '[]', '113.26641,23.132324', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3108, 650000, '新疆维吾尔自治区', 4, '[]', '87.627704,43.793026', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2482, 520000, '贵州省', 4, '[]', '106.70546,26.600055', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (579, 220000, '吉林省', 4, '[]', '125.32568,43.897016', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (217, 140000, '山西省', 4, '[]', '112.562678,37.873499', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (810, 320000, '江苏省', 4, '[]', '118.762765,32.060875', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2926, 620000, '甘肃省', 4, '[]', '103.826447,36.05956', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1238, 360000, '江西省', 4, '[]', '115.81635,28.63666', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (2808, 610000, '陕西省', 4, '[]', '108.954347,34.265502', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3190, 653221, '和田县', 1, '0903', '79.81907,37.120031', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3191, 653222, '墨玉县', 1, '0903', '79.728683,37.277143', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3192, 653223, '皮山县', 1, '0903', '78.283669,37.62145', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3193, 653224, '洛浦县', 1, '0903', '80.188986,37.073667', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3194, 653225, '策勒县', 1, '0903', '80.806159,36.998335', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3195, 653226, '于田县', 1, '0903', '81.677418,36.85708', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3196, 653227, '民丰县', 1, '0903', '82.695861,37.06408', 3188);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3198, 654002, '伊宁市', 1, '0999', '81.27795,43.908558', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3199, 654003, '奎屯市', 1, '0999', '84.903267,44.426529', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3200, 654004, '霍尔果斯市', 1, '0999', '80.411271,44.213941', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3201, 654021, '伊宁县', 1, '0999', '81.52745,43.977119', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3202, 654022, '察布查尔锡伯自治县', 1, '0999', '81.151337,43.840726', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3203, 654023, '霍城县', 1, '0999', '80.87898,44.055984', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3204, 654024, '巩留县', 1, '0999', '82.231718,43.482628', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3205, 654025, '新源县', 1, '0999', '83.232848,43.433896', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3206, 654026, '昭苏县', 1, '0999', '81.130974,43.157293', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3207, 654027, '特克斯县', 1, '0999', '81.836206,43.217183', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3208, 654028, '尼勒克县', 1, '0999', '82.511809,43.800247', 3197);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3210, 654201, '塔城市', 1, '0901', '82.986978,46.751428', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3211, 654202, '乌苏市', 1, '0901', '84.713396,44.41881', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3212, 654221, '额敏县', 1, '0901', '83.628303,46.524673', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3213, 654223, '沙湾县', 1, '0901', '85.619416,44.326388', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3214, 654224, '托里县', 1, '0901', '83.60695,45.947638', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3215, 654225, '裕民县', 1, '0901', '82.982667,46.201104', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3216, 654226, '和布克赛尔蒙古自治县', 1, '0901', '85.728328,46.793235', 3209);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3218, 654301, '阿勒泰市', 1, '0906', '88.131842,47.827308', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3219, 654321, '布尔津县', 1, '0906', '86.874923,47.702163', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (5, 110105, '朝阳区', 1, '010', '116.443205,39.921506', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (6, 110106, '丰台区', 1, '010', '116.287039,39.858421', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (7, 110107, '石景山区', 1, '010', '116.222933,39.906611', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (8, 110108, '海淀区', 1, '010', '116.298262,39.95993', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (9, 110109, '门头沟区', 1, '010', '116.101719,39.940338', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10, 110111, '房山区', 1, '010', '116.143486,39.748823', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (11, 110112, '通州区', 1, '010', '116.656434,39.909946', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (12, 110113, '顺义区', 1, '010', '116.654642,40.130211', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (13, 110114, '昌平区', 1, '010', '116.231254,40.220804', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (14, 110115, '大兴区', 1, '010', '116.341483,39.726917', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (15, 110116, '怀柔区', 1, '010', '116.631931,40.316053', 2);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (27, 120110, '东丽区', 1, '022', '117.31362,39.086802', 20);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1, 110000, '北京市', 4, '010', '116.407394,39.904211', 0);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (188, 130929, '献县', 1, '0317', '116.122725,38.190185', 177);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (525, 210727, '义县', 1, '0416', '121.23908,41.533086', 520);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1236, 350981, '福安市', 1, '0593', '119.64785,27.08834', 1228);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (1498, 371721, '曹县', 1, '0530', '115.542328,34.825508', 1495);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3220, 654322, '富蕴县', 1, '0906', '89.525504,46.994115', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3221, 654323, '福海县', 1, '0906', '87.486703,47.111918', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3222, 654324, '哈巴河县', 1, '0906', '86.418621,48.060846', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3223, 654325, '青河县', 1, '0906', '90.37555,46.679113', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3224, 654326, '吉木乃县', 1, '0906', '85.874096,47.443101', 3217);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3230, 810001, '中西区', 1, '1852', '114.154373,22.281981', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3231, 810002, '湾仔区', 1, '1852', '114.182915,22.276389', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3232, 810003, '东区', 1, '1852', '114.226003,22.279693', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3233, 810004, '南区', 1, '1852', '114.160012,22.245897', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3234, 810005, '油尖旺区', 1, '1852', '114.173332,22.311704', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3235, 810006, '深水埗区', 1, '1852', '114.163242,22.333854', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3236, 810007, '九龙城区', 1, '1852', '114.192847,22.31251', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3237, 810008, '黄大仙区', 1, '1852', '114.203886,22.336321', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3238, 810009, '观塘区', 1, '1852', '114.214054,22.320838', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3239, 810010, '荃湾区', 1, '1852', '114.121079,22.368306', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3240, 810011, '屯门区', 1, '1852', '113.976574,22.393844', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3241, 810012, '元朗区', 1, '1852', '114.032438,22.441428', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3243, 810013, '北区', 1, '1852', '114.147364,22.496104', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3245, 810014, '大埔区', 1, '1852', '114.171743,22.445653', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3246, 810015, '西贡区', 1, '1852', '114.264645,22.314213', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3247, 810016, '沙田区', 1, '1852', '114.195365,22.379532', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3248, 810017, '葵青区', 1, '1852', '114.139319,22.363877', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3249, 810018, '离岛区', 1, '1852', '113.94612,22.286408', 3229);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3252, 820001, '花地玛堂区', 1, '1853', '113.552896,22.20787', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3253, 820002, '花王堂区', 1, '1853', '113.548961,22.199207', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3254, 820003, '望德堂区', 1, '1853', '113.550183,22.193721', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3255, 820004, '大堂区', 1, '1853', '113.553647,22.188539', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3256, 820005, '风顺堂区', 1, '1853', '113.541928,22.187368', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3257, 820006, '嘉模堂区', 1, '1853', '113.558705,22.15376', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3258, 820007, '路凼填海区', 1, '1853', '113.569599,22.13663', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (3259, 820008, '圣方济各堂区', 1, '1853', '113.559954,22.123486', 3251);
INSERT INTO "public"."sys_area" ("id", "area_code", "area_name", "level", "city_code", "lat_long_center", "parent_id") OVERRIDING SYSTEM VALUE VALUES (10000, 110900, '全国', 16, '000', '116.407394,39.904211', 0);
COMMIT;

-- ----------------------------
-- Table structure for sys_audit
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_audit";
CREATE TABLE "public"."sys_audit" (
                                      "id" int8 NOT NULL,
                                      "state" int4 NOT NULL DEFAULT 0,
                                      "replay" text COLLATE "pg_catalog"."default",
                                      "union_main_id" int8 NOT NULL,
                                      "category" int4 NOT NULL,
                                      "audit_data" text COLLATE "pg_catalog"."default" NOT NULL,
                                      "expire_at" timestamp(6),
                                      "audit_replay_at" timestamp(6),
                                      "history_Items" text COLLATE "pg_catalog"."default",
                                      "created_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_audit" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_audit"."state" IS '审核状态：-1不通过，0待审核，1通过';
COMMENT ON COLUMN "public"."sys_audit"."replay" IS '不通过时回复的审核不通过原因';
COMMENT ON COLUMN "public"."sys_audit"."union_main_id" IS '关联主体ID';
COMMENT ON COLUMN "public"."sys_audit"."category" IS '业务类别';
COMMENT ON COLUMN "public"."sys_audit"."audit_data" IS '待审核的业务数据包';
COMMENT ON COLUMN "public"."sys_audit"."expire_at" IS '服务时限';
COMMENT ON COLUMN "public"."sys_audit"."audit_replay_at" IS '审核回复时间';
COMMENT ON COLUMN "public"."sys_audit"."history_Items" IS '历史申请记录';

-- ----------------------------
-- Records of sys_audit
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_casbin
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_casbin";
CREATE TABLE "public"."sys_casbin" (
                                       "ptype" varchar(10) COLLATE "pg_catalog"."default",
                                       "v0" varchar(256) COLLATE "pg_catalog"."default",
                                       "v1" varchar(256) COLLATE "pg_catalog"."default",
                                       "v2" varchar(256) COLLATE "pg_catalog"."default",
                                       "v3" varchar(256) COLLATE "pg_catalog"."default",
                                       "v4" varchar(256) COLLATE "pg_catalog"."default",
                                       "v5" varchar(256) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_casbin" OWNER TO "kysion";

-- ----------------------------
-- Records of sys_casbin
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5699282126110789', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5699282403065925', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5699303772323909', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977484017139781', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977479579435077', '5977484017139781', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977591512694853', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977591062921285', '5977591512694853', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977606853427269', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977604411621445', '5977606853427269', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977677758595141', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977677295255621', '5977677758595141', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977706705715269', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977706248405061', '5977706705715269', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977720232017989', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5977719818551365', '5977720232017989', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5987168481509445', 'Super.Admin', 'kysion.com', '', '', '');
INSERT INTO "public"."sys_casbin" ("ptype", "v0", "v1", "v2", "v3", "v4", "v5") VALUES ('g', '5987167558369349', '5987168481509445', 'kysion.com', '', '', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_config";
CREATE TABLE "public"."sys_config" (
                                       "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                       "value" json,
                                       "created_at" timestamp(6),
                                       "updated_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_config" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_config"."name" IS '配置名称';
COMMENT ON COLUMN "public"."sys_config"."value" IS '配置信息';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_config" ("name", "value", "created_at", "updated_at") VALUES ('baidu_sdk_conf', '[{"identifier":"certificate_orc","description":"百度OCR识别","appID":"28256014","aesKey":"b2be05d92c4a0450","apiKey":"n2VFMDG8UtqW24k22XEZiDhK","secretKey":"KTcFYIjxy6WdpoghrsyRNDCRebz5lstn"},{"identifier":"certificate_nlp","description":"SDK测试数据","appID":"875634231","aesKey":"b2be05d92c4a0450","apiKey":"t4cuMDG8UtqW24k22XEZiDhK","secretKey":"LyypYIjxu6WdpoghrsyRNDCRebz5lstn"},{"identifier":"test","description":"","appID":"234","aesKey":"234","apiKey":"234","secretKey":"234"},{"identifier":"112123","description":"","appID":"123","aesKey":"123","apiKey":"32","secretKey":"34234"},{"identifier":"34534","description":"345","appID":"345","aesKey":"345","apiKey":"345345","secretKey":"345"},{"identifier":"345","description":"345","appID":"345","aesKey":"345","apiKey":"345","secretKey":"345"},{"identifier":"222","description":"2","appID":"23","aesKey":"23","apiKey":"23","secretKey":"3233"}]', '2022-11-05 05:57:35', '2022-11-26 09:28:41');
INSERT INTO "public"."sys_config" ("name", "value", "created_at", "updated_at") VALUES ('aliyun_sdk_conf', '[{"identifier":"certificate_ssm","description":"阿里云SDK短信服务测试数据","appID":"876542083","aesKey":"qcfR43Jd92c4a0450","apiKey":"pli9TDG8UtqW24k22XEZiDhK","secretKey":"OOpeYIjxu6WdpoghrsyRNDCRebz5lstn"},{"identifier":"certificate_ocr","description":"阿里云SDK测试数据2","appID":"876542009","aesKey":"qP0243Jd92c4a0450","apiKey":"pOY9TDG8UtqW24k22XEZiDhK","secretKey":"OiiTrYIjxu6WdpoghrsyRNDCRebz5lstn"}]', '2022-11-25 16:59:35', '2022-11-25 17:06:31');
INSERT INTO "public"."sys_config" ("name", "value", "created_at", "updated_at") VALUES ('huawei_sdk_conf', '[{"identifier":"certificate_ocr","description":"华为云SDK测试","appID":"876542083","aesKey":"qcfR43Jd92c4a0450","apiKey":"pli9TDG8UtqW24k22XEZiDhK","secretKey":"OOpeYIjxu6WdpoghrsyRNDCRebz5lstn"}]', '2022-11-26 20:02:55', '2022-11-26 22:08:49');
INSERT INTO "public"."sys_config" ("name", "value", "created_at", "updated_at") VALUES ('tencent_sdk_conf', '[{"identifier":"certificate_nlp","description":"腾讯云SDK测试数据","appID":"876542083","aesKey":"qcfR43Jd92c4a0450","apiKey":"pli9TDG8UtqW24k22XEZiDhK","secretKey":"OOpeYIjxu6WdpoghrsyRNDCRebz5lstn","active":"certificate_nlp","version":"2019-12-09","region":"上海"}]', '2022-11-26 23:02:57', '2022-11-26 23:02:57');
INSERT INTO "public"."sys_config" ("name", "value", "created_at", "updated_at") VALUES ('ctyun_sdk_conf', '[{"identifier":"certificate_sms","description":"天翼云SDK测试数据","appID":"876542003","appKey":"OOpeYIjxu6WdpoghrsyRNDCRebz5lstn","appSecret":"pli9TDG8UtqW24k22XEZiDhK","accessKey":"aqcfR43Jd92c4a0450","security_key":"voluptateqcfR43Jd92c4a0450"}]', '2022-11-28 09:19:13', '2022-11-28 09:19:13');
COMMIT;

-- ----------------------------
-- Table structure for sys_file
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_file";
CREATE TABLE "public"."sys_file" (
                                     "id" int8 NOT NULL,
                                     "name" varchar(45) COLLATE "pg_catalog"."default",
                                     "src" varchar(500) COLLATE "pg_catalog"."default" NOT NULL,
                                     "url" varchar(500) COLLATE "pg_catalog"."default",
                                     "ext" varchar(16) COLLATE "pg_catalog"."default",
                                     "size" int8,
                                     "category" varchar(255) COLLATE "pg_catalog"."default",
                                     "user_id" int8,
                                     "union_main_id" int8,
                                     "created_at" timestamp(6),
                                     "updated_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_file" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_file"."id" IS '自增ID';
COMMENT ON COLUMN "public"."sys_file"."name" IS '文件名称';
COMMENT ON COLUMN "public"."sys_file"."src" IS '存储路径';
COMMENT ON COLUMN "public"."sys_file"."url" IS 'URL地址';
COMMENT ON COLUMN "public"."sys_file"."ext" IS '扩展名';
COMMENT ON COLUMN "public"."sys_file"."size" IS '文件大小';
COMMENT ON COLUMN "public"."sys_file"."category" IS '文件分类';
COMMENT ON COLUMN "public"."sys_file"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_file"."union_main_id" IS '关联主体ID';

-- ----------------------------
-- Records of sys_file
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_file" ("id", "name", "src", "url", "ext", "size", "category", "user_id", "union_main_id", "created_at", "updated_at") VALUES (5978271823298629, 'cpojrcqg5qzs90807a.webp', '/tmp/upload/20230110/20230110/5978271823298629/cpojrcqg5qzs90807a.webp', '/api/common/sys_file/getFileById?id=5978271823298629', '.webp', 36480, NULL, 5977706248405061, NULL, '2023-01-10 21:36:33', '2023-01-10 21:36:33');
INSERT INTO "public"."sys_file" ("id", "name", "src", "url", "ext", "size", "category", "user_id", "union_main_id", "created_at", "updated_at") VALUES (5988115924058181, 'cpq105tmivag4jdgdh.webp', '/tmp/upload/20230112/20230112/5988115924058181/cpq105tmivag4jdgdh.webp', '/api/common/file/getFileById?id=5988115924058181', '.webp', 36480, NULL, 5977706248405061, 5977706236739653, '2023-01-12 15:16:39', '2023-01-12 15:26:41');
COMMIT;

-- ----------------------------
-- Table structure for sys_license
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_license";
CREATE TABLE "public"."sys_license" (
                                        "id" int8 NOT NULL,
                                        "idcard_front_path" varchar(256) COLLATE "pg_catalog"."default" NOT NULL,
                                        "idcard_back_path" varchar(256) COLLATE "pg_catalog"."default" NOT NULL,
                                        "idcard_no" varchar(18) COLLATE "pg_catalog"."default" NOT NULL,
                                        "idcard_expired_date" date,
                                        "idcard_address" varchar(128) COLLATE "pg_catalog"."default",
                                        "person_contact_name" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                        "person_contact_mobile" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                        "business_license_name" varchar(128) COLLATE "pg_catalog"."default",
                                        "business_license_address" varchar(128) COLLATE "pg_catalog"."default",
                                        "business_license_path" varchar(256) COLLATE "pg_catalog"."default",
                                        "business_license_scope" varchar(512) COLLATE "pg_catalog"."default",
                                        "business_license_reg_capital" varchar(32) COLLATE "pg_catalog"."default",
                                        "business_license_term_time" varchar(64) COLLATE "pg_catalog"."default",
                                        "business_license_org_code" varchar(16) COLLATE "pg_catalog"."default",
                                        "business_license_credit_code" varchar(32) COLLATE "pg_catalog"."default" NOT NULL,
                                        "business_license_legal" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                        "business_license_legal_path" varchar(256) COLLATE "pg_catalog"."default",
                                        "latest_audit_log_id" int8,
                                        "state" int4 NOT NULL,
                                        "auth_type" int4 NOT NULL,
                                        "remark" text COLLATE "pg_catalog"."default",
                                        "updated_at" timestamp(6),
                                        "created_at" timestamp(6),
                                        "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_license" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_license"."idcard_front_path" IS '身份证头像面照片';
COMMENT ON COLUMN "public"."sys_license"."idcard_back_path" IS '身份证国徽面照片';
COMMENT ON COLUMN "public"."sys_license"."idcard_no" IS '身份证号';
COMMENT ON COLUMN "public"."sys_license"."idcard_expired_date" IS '身份证有效期';
COMMENT ON COLUMN "public"."sys_license"."idcard_address" IS '身份证户籍地址';
COMMENT ON COLUMN "public"."sys_license"."person_contact_name" IS '负责人，必须是自然人';
COMMENT ON COLUMN "public"."sys_license"."person_contact_mobile" IS '负责人，联系电话';
COMMENT ON COLUMN "public"."sys_license"."business_license_name" IS '公司全称';
COMMENT ON COLUMN "public"."sys_license"."business_license_address" IS '公司地址';
COMMENT ON COLUMN "public"."sys_license"."business_license_path" IS '营业执照图片地址';
COMMENT ON COLUMN "public"."sys_license"."business_license_scope" IS '经营范围';
COMMENT ON COLUMN "public"."sys_license"."business_license_reg_capital" IS '注册资本';
COMMENT ON COLUMN "public"."sys_license"."business_license_term_time" IS '营业期限';
COMMENT ON COLUMN "public"."sys_license"."business_license_org_code" IS '组织机构代码';
COMMENT ON COLUMN "public"."sys_license"."business_license_credit_code" IS '统一社会信用代码';
COMMENT ON COLUMN "public"."sys_license"."business_license_legal" IS '法人';
COMMENT ON COLUMN "public"."sys_license"."business_license_legal_path" IS '法人证照，如果法人不是自然人，则该项必填';
COMMENT ON COLUMN "public"."sys_license"."latest_audit_log_id" IS '最新的审核记录ID';

-- ----------------------------
-- Records of sys_license
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_logs
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_logs";
CREATE TABLE "public"."sys_logs" (
                                     "id" int8 NOT NULL,
                                     "user_id" int8,
                                     "error" text COLLATE "pg_catalog"."default",
                                     "category" varchar(16) COLLATE "pg_catalog"."default",
                                     "level" int4,
                                     "content" json,
                                     "context" text COLLATE "pg_catalog"."default",
                                     "created_at" timestamp(6),
                                     "updated_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_logs" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_logs"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_logs"."user_id" IS '用户UID';
COMMENT ON COLUMN "public"."sys_logs"."error" IS '错误信息';
COMMENT ON COLUMN "public"."sys_logs"."category" IS '分类';
COMMENT ON COLUMN "public"."sys_logs"."level" IS '等级';
COMMENT ON COLUMN "public"."sys_logs"."content" IS '日志内容';
COMMENT ON COLUMN "public"."sys_logs"."context" IS '上下文数据';

-- ----------------------------
-- Records of sys_logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_menu";
CREATE TABLE "public"."sys_menu" (
                                     "id" int8 NOT NULL,
                                     "path" varchar(128) COLLATE "pg_catalog"."default",
                                     "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                     "redirect" varchar(128) COLLATE "pg_catalog"."default",
                                     "title" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                     "icon" varchar(128) COLLATE "pg_catalog"."default",
                                     "component" varchar(128) COLLATE "pg_catalog"."default",
                                     "parent_id" int8 DEFAULT 0,
                                     "sort" int4 DEFAULT 0,
                                     "state" int2 DEFAULT 1,
                                     "description" varchar(128) COLLATE "pg_catalog"."default",
                                     "created_at" timestamp(6),
                                     "updated_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_menu" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_menu"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_menu"."path" IS '路径';
COMMENT ON COLUMN "public"."sys_menu"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_menu"."redirect" IS '跳转';
COMMENT ON COLUMN "public"."sys_menu"."title" IS '标题';
COMMENT ON COLUMN "public"."sys_menu"."icon" IS '图标';
COMMENT ON COLUMN "public"."sys_menu"."component" IS '组件';
COMMENT ON COLUMN "public"."sys_menu"."parent_id" IS '所属父级';
COMMENT ON COLUMN "public"."sys_menu"."sort" IS '排序';
COMMENT ON COLUMN "public"."sys_menu"."state" IS '状态：0隐藏，1显示';
COMMENT ON COLUMN "public"."sys_menu"."description" IS '描述';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_organization
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_organization";
CREATE TABLE "public"."sys_organization" (
                                             "id" int8 NOT NULL,
                                             "name" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                             "parent_id" int8 DEFAULT 0,
                                             "cascade_deep" int4,
                                             "description" varchar(128) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_organization" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_organization"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_organization"."parent_id" IS '父级ID';
COMMENT ON COLUMN "public"."sys_organization"."cascade_deep" IS '级联深度';
COMMENT ON COLUMN "public"."sys_organization"."description" IS '描述';

-- ----------------------------
-- Records of sys_organization
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_permission
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_permission";
CREATE TABLE "public"."sys_permission" (
                                           "id" int8 NOT NULL,
                                           "parent_id" int8,
                                           "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                           "description" varchar(128) COLLATE "pg_catalog"."default",
                                           "identifier" varchar(128) COLLATE "pg_catalog"."default" NOT NULL,
                                           "type" int4,
                                           "match_mode" int4,
                                           "is_show" int2 DEFAULT 1,
                                           "sort" int4,
                                           "created_at" timestamp(6),
                                           "updated_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_permission" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_permission"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_permission"."parent_id" IS '父级ID';
COMMENT ON COLUMN "public"."sys_permission"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_permission"."description" IS '描述';
COMMENT ON COLUMN "public"."sys_permission"."identifier" IS '标识符';
COMMENT ON COLUMN "public"."sys_permission"."type" IS '类型：1api，2menu';
COMMENT ON COLUMN "public"."sys_permission"."match_mode" IS '匹配模式：ID：0，标识符：1';
COMMENT ON COLUMN "public"."sys_permission"."is_show" IS '是否显示：0不显示 1显示';
COMMENT ON COLUMN "public"."sys_permission"."sort" IS '排序';

-- ----------------------------
-- Records of sys_permission
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947106208184773, 0, '用户管理', '', 'User', 1, 0, 1, 0, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947175853095365, 5947106208184773, '查看用户', '查看某个用户登录账户', 'User::ViewDetail', 1, 0, 1, 0, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947176286288325, 5947106208184773, '用户列表', '查看所有用户', 'User::List', 1, 0, 1, 1, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947177123969477, 5947106208184773, '重置密码', '重置某个用户的登录密码', 'User::ResetPassword', 1, 0, 1, 2, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947176737372613, 5947106208184773, '设置状态', '设置某个用户的状态', 'User::SetState', 1, 0, 1, 3, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947177469213125, 5947106208184773, '修改密码', '修改自己的登录密码', 'User::ChangePassword', 1, 0, 1, 4, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5949854362632262, 5947106208184773, '修改用户名称', '修改用户登录账户名称信息', 'User::Update', 1, 0, 1, 5, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5949854362632264, 5947106208184773, '设置用户角色', '设置某一个用户的角色', 'User::SetUserRole', 1, 0, 1, 6, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5949854362632265, 5947106208184773, '设置用户权限', '设置某一个用户的权限', 'User::SetPermission', 1, 0, 1, 7, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948649344204869, 0, '组织架构', '', 'Organization', 1, 0, 0, 1, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948649434447941, 5948649344204869, '查看组织架构', '查看某个组织架构', 'Organization::ViewDetail', 1, 0, 0, 0, '2023-01-17 21:32:52', '2023-01-17 21:32:52');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948649530392645, 5948649344204869, '组织架构列表', '查看所有组织架构列表', 'Organization::List', 1, 0, 0, 1, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948649642721349, 5948649344204869, '更新组织架构', '更新某个组织架构', 'Organization::Update', 1, 0, 0, 2, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948649739583557, 5948649344204869, '删除组织架构', '删除某个组织架构', 'Organization::Delete', 1, 0, 0, 3, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948649828712517, 5948649344204869, '创建组织架构', '创建组织架构', 'Organization::Create', 1, 0, 0, 4, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948684761759818, 0, '角色管理', '', 'Role', 1, 0, 1, 2, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948684761759813, 5948684761759818, '查看角色', '查看某个角色', 'Role::ViewDetail', 1, 0, 1, 0, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948684761759814, 5948684761759818, '角色列表', '查看所有角色', 'Role::List', 1, 0, 1, 1, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948684761759815, 5948684761759818, '更新角色信息', '更新某个角色信息', 'Role::Update', 1, 0, 1, 2, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948684761759816, 5948684761759818, '删除角色', '删除某个角色', 'Role::Delete', 1, 0, 1, 3, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948684761759817, 5948684761759818, '创建角色', '创建一个新角色', 'Role::Create', 1, 0, 1, 4, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5950451522795973, 5948684761759818, '设置角色成员', '增加或移除角色成员', 'Role::SetMember', 1, 0, 1, 5, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5950452043151813, 5948684761759818, '设置角色权限', '设置某个角色的权限', 'Role::SetPermission', 1, 0, 1, 6, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5950408166668741, 0, '权限管理', '', 'Permission', 1, 0, 1, 3, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948682180886598, 5950408166668741, '查看权限', '查看某个权限', 'Permission::ViewDetail', 1, 0, 1, 0, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948682180886599, 5950408166668741, '权限列表', '查看所有权限', 'Permission::List', 1, 0, 1, 1, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948682180886600, 5950408166668741, '更新权限', '更新某个权限', 'Permission::Update', 1, 0, 1, 2, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948682180886601, 5950408166668741, '删除权限', '删除某个权限', 'Permission::Delete', 1, 0, 1, 3, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948682180886602, 5950408166668741, '创建权限', '创建权限', 'Permission::Create', 1, 0, 1, 4, '2023-01-17 21:32:53', '2023-01-17 21:32:53');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947986066667973, 0, '公司', '', 'company', 1, 0, 1, 0, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988889178, 5947986066667973, '新增', '', 'company::Create', 1, 1, 0, 0, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988889187, 5947986066667973, '查看明细', '', 'company::ViewDetail', 1, 1, 0, 1, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988889200, 5947986066667973, '列表', '', 'company::List', 1, 1, 0, 2, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988889209, 5947986066667973, '更新', '', 'company::Update', 1, 1, 0, 3, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988954698, 5947986066667973, '设置LOGO', '', 'company::SetLogo', 1, 1, 0, 4, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988954709, 5947986066667973, '设置状态', '', 'company::SetState', 1, 1, 0, 5, '2023-01-17 23:38:33', '2023-01-17 23:38:33');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988954720, 5947986066667973, '设置管理员', '', 'company::SetAdminUser', 1, 1, 0, 6, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988954731, 5947986066667973, '查看认证信息', '查看公司认证信息', 'company::ViewLicense', 1, 1, 0, 7, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988954742, 5947986066667973, '审核认证信息', '审核公司认证信息', 'company::AuditLicense', 1, 1, 0, 8, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948221667408325, 0, '公司员工', '', 'employee', 1, 0, 1, 1, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400988954744, 5948221667408325, '详情', '查看员工详情', 'employee::ViewDetail', 1, 1, 0, 0, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989020234, 5948221667408325, '更多详情', '查看员工更多详情含手机号等', 'employee::MoreDetail', 1, 1, 0, 1, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989020247, 5948221667408325, '列表', '查看员工列表', 'employee::List', 1, 1, 0, 2, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989020260, 5948221667408325, '新增', '新增员工信息', 'employee::Create', 1, 1, 0, 3, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989020273, 5948221667408325, '更新', '更新员工信息', 'employee::Update', 1, 1, 0, 4, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989020286, 5948221667408325, '删除', '删除员工信息', 'employee::Delete', 1, 1, 0, 5, '2023-01-17 23:38:34', '2023-01-17 23:38:34');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989085776, 5948221667408325, '设置手机号', '修改员工手机号', 'employee::SetMobile', 1, 1, 0, 6, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989085789, 5948221667408325, '设置头像', '设置员工头像', 'employee::SetAvatar', 1, 1, 0, 7, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989085802, 5948221667408325, '设置状态', '设置员工任职状态', 'employee::SetState', 1, 1, 0, 8, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989085815, 5948221667408325, '查看认证信息', '查看员工认证信息', 'employee::ViewLicense', 1, 1, 0, 9, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151305, 5948221667408325, '审核认证信息', '审核员工认证信息', 'employee::AuditLicense', 1, 1, 0, 10, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151318, 5948221667408325, '更新认证信息', '更新员工认证信息', 'employee::UpdateLicense', 1, 1, 0, 11, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5948221667408326, 0, '公司团队', '', 'team', 1, 0, 1, 2, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151321, 5948221667408326, '新增', '新增团队信息', 'team::Create', 1, 1, 0, 0, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151328, 5948221667408326, '详情', '查看团队详情', 'team::ViewDetail', 1, 1, 0, 1, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151338, 5948221667408326, '列表', '查看团队列表', 'team::List', 1, 1, 0, 2, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151349, 5948221667408326, '更新', '更新团队信息', 'team::Update', 1, 1, 0, 3, '2023-01-17 23:38:35', '2023-01-17 23:38:35');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989151359, 5948221667408326, '删除', '删除团队信息', 'team::Delete', 1, 1, 0, 4, '2023-01-17 23:38:36', '2023-01-17 23:38:36');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989216846, 5948221667408326, '成员详情', '查看团队成员详情', 'team::MemberDetail', 1, 1, 0, 5, '2023-01-17 23:38:36', '2023-01-17 23:38:36');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989216856, 5948221667408326, '设置成员', '设置团队成员', 'team::SetMember', 1, 1, 0, 6, '2023-01-17 23:38:36', '2023-01-17 23:38:36');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989216866, 5948221667408326, '设置管理人', '设置团队管理人，可以不是团队成员', 'team::SetOwner', 1, 1, 0, 7, '2023-01-17 23:38:36', '2023-01-17 23:38:36');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (6018400989216876, 5948221667408326, '设置队长或组长', '设置团队队长或小组组长，必须是团队成员', 'team::SetCaptain', 1, 1, 0, 8, '2023-01-17 23:38:36', '2023-01-17 23:38:36');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5947175853095366, 5947106208184773, '查看更多详情', '含完整手机号', 'User::ViewMoreDetail', 1, 0, 1, 1, '2023-01-19 18:23:46', '2023-01-19 18:23:46');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845333, 0, '资质', '', 'License', 1, 0, 1, 4, '2023-02-06 23:52:11', '2023-02-06 23:52:11');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845328, 5953153121845333, '查看主体信息', '查看某条主体信息', 'License::ViewDetail', 1, 0, 1, 0, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845329, 5953153121845333, '主体列表', '查看所有主体信息', 'License::List', 1, 0, 1, 1, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845330, 5953153121845333, '更新资质审核信息', '更新某条资质审核信息', 'License::Update', 1, 0, 1, 2, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845332, 5953153121845333, '创建主体', '创建主体信息', 'License::Create', 1, 0, 1, 3, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845331, 5953153121845333, '设置主体状态', '设置某主体认证状态', 'License::SetState', 1, 0, 1, 4, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953151699124300, 0, '审核管理', '', 'Audit', 1, 0, 1, 5, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953151699124297, 5953151699124300, '查看资质审核信息', '查看某条资质审核信息', 'Audit::ViewDetail', 1, 0, 1, 0, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953151699124298, 5953151699124300, '资质审核列表', '查看所有资质审核', 'Audit::List', 1, 0, 1, 1, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953151699124299, 5953151699124300, '更新资质审核信息', '更新某条资质审核信息', 'Audit::Update', 1, 0, 1, 2, '2023-02-06 23:52:12', '2023-02-06 23:52:12');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845349, 0, '财务服务', '', 'Financial', 1, 0, 1, 0, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845334, 5953153121845349, '查看发票详情', '查看发票详情信息', 'Financial::ViewDetail', 1, 0, 1, 0, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845335, 5953153121845349, '查看发票抬头信息', '查看发票抬头信息', 'Financial::ViewInvoice', 1, 0, 1, 1, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845336, 5953153121845349, '查看提现账号', '查看提现账号信息', 'Financial::ViewBankCardDetail', 1, 0, 1, 2, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845337, 5953153121845349, '提现账号列表', '查看所有提现账号', 'Financial::BankCardList', 1, 0, 1, 3, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845338, 5953153121845349, '发票抬头列表', '查看所有发票抬头', 'Financial::InvoiceList', 1, 0, 1, 4, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845339, 5953153121845349, '发票详情列表', '查看所有发票详情', 'Financial::InvoiceDetailList', 1, 0, 1, 5, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845340, 5953153121845349, '审核发票', '审核发票申请', 'Financial::AuditInvoiceDetail', 1, 0, 1, 6, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845341, 5953153121845349, '开发票', '添加发票详情记录', 'Financial::MakeInvoiceDetail', 1, 0, 1, 7, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845342, 5953153121845349, '添加发票抬头', '添加发票抬头信息', 'Financial::CreateInvoice', 1, 0, 1, 8, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845343, 5953153121845349, '申请提现账号', '添加提现账号信息', 'Financial::CreateBankCard', 1, 0, 1, 9, '2023-02-17 22:08:37', '2023-02-17 22:08:37');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845344, 5953153121845349, '删除发票抬头', '删除发票抬头信息', 'Financial::DeleteInvoice', 1, 0, 1, 10, '2023-02-17 22:08:38', '2023-02-17 22:08:38');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845345, 5953153121845349, '删除提现账号', '删除提现账号信息', 'Financial::DeleteBankCard', 1, 0, 1, 11, '2023-02-17 22:08:38', '2023-02-17 22:08:38');
INSERT INTO "public"."sys_permission" ("id", "parent_id", "name", "description", "identifier", "type", "match_mode", "is_show", "sort", "created_at", "updated_at") VALUES (5953153121845346, 5953153121845349, '查看余额', '查看账号余额', 'Financial::GetAccountBalance', 1, 0, 1, 12, '2023-02-17 22:08:38', '2023-02-17 22:08:38');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role";
CREATE TABLE "public"."sys_role" (
                                     "id" int8 NOT NULL,
                                     "name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                     "description" varchar(128) COLLATE "pg_catalog"."default",
                                     "is_system" bool,
                                     "union_main_id" int8,
                                     "created_at" timestamp(6),
                                     "updated_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_role" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_role"."name" IS '名称';
COMMENT ON COLUMN "public"."sys_role"."description" IS '描述';
COMMENT ON COLUMN "public"."sys_role"."is_system" IS '是否默认角色，true仅能修改名称，不允许删除和修改';
COMMENT ON COLUMN "public"."sys_role"."union_main_id" IS '主体id';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5492049091493957, '管理员', '', 't', NULL, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5492052201766981, 'BOSS', '', 't', NULL, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5492052621066309, '财务', '', 't', NULL, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5492053050785861, '运营', '', 't', NULL, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5492053367128133, '客服', '', 't', NULL, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5977677758595141, '管理员', '', 'f', 5977677281165381, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5977706705715269, '管理员', '', 'f', 5977706236739653, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5977720232017989, '管理员', '', 'f', 5977719809835077, NULL, NULL);
INSERT INTO "public"."sys_role" ("id", "name", "description", "is_system", "union_main_id", "created_at", "updated_at") VALUES (5987168481509445, '管理员', '', 'f', 5987167197659205, '2023-01-12 11:15:43', '2023-01-12 11:15:43');
COMMIT;

-- ----------------------------
-- Table structure for sys_sms_logs
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_sms_logs";
CREATE TABLE "public"."sys_sms_logs" (
                                         "id" numeric(20,0) NOT NULL,
                                         "type" varchar(32) COLLATE "pg_catalog"."default",
                                         "context" text COLLATE "pg_catalog"."default",
                                         "mobile" varchar(32) COLLATE "pg_catalog"."default",
                                         "state" varchar(32) COLLATE "pg_catalog"."default",
                                         "result" json,
                                         "user_id" int8,
                                         "license_id" int8,
                                         "created_at" timestamp(6),
                                         "updated_at" timestamp(6),
                                         "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_sms_logs" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_sms_logs"."type" IS '短信平台：qyxs：企业信使';
COMMENT ON COLUMN "public"."sys_sms_logs"."context" IS '短信内容';
COMMENT ON COLUMN "public"."sys_sms_logs"."mobile" IS '手机号';
COMMENT ON COLUMN "public"."sys_sms_logs"."state" IS '发送状态';
COMMENT ON COLUMN "public"."sys_sms_logs"."result" IS '短信接口返回内容';
COMMENT ON COLUMN "public"."sys_sms_logs"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_sms_logs"."license_id" IS '主体ID';

-- ----------------------------
-- Records of sys_sms_logs
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_user";
CREATE TABLE "public"."sys_user" (
                                     "id" int8 NOT NULL,
                                     "username" varchar(45) COLLATE "pg_catalog"."default" NOT NULL,
                                     "password" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                     "state" int4,
                                     "type" int4,
                                     "mobile" varchar(16) COLLATE "pg_catalog"."default",
                                     "created_at" timestamp(6),
                                     "updated_at" timestamp(6),
                                     "deleted_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_user" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_user"."username" IS '账号';
COMMENT ON COLUMN "public"."sys_user"."password" IS '密码';
COMMENT ON COLUMN "public"."sys_user"."state" IS '状态：0未激活、1正常、-1封号、-2异常、-3已注销';
COMMENT ON COLUMN "public"."sys_user"."type" IS '用户类型，0匿名，1用户，2微商，4商户、8广告主、16服务商、32运营中心';
COMMENT ON COLUMN "public"."sys_user"."mobile" IS '手机号';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5751143024558149, 'Boss', '9a33fc0f21173b28af9a7a613640a6128eea56a2875150005ec5e93b480dd06f', 1, 32, NULL, NULL, NULL, NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5302581852373061, 'Super.Admin', 'be86579ef22515322f123fd5b62f6735711d337514132363135ff0871acd7773', 1, -1, NULL, NULL, '2023-01-10 15:43:10', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5977706248405061, '1mux4c8hydh', 'ce5c701a7522cae8af50918435bf19d84bdd109518c13533b404336c14b09708', 1, 32, '18661667165', '2023-01-10 19:09:24', '2023-01-10 19:09:24', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5977719818551365, '1muxaknta9x', '6b6542bcde42079fd363b6765d520fa8ff3db86e78f52496108d75fea2a45bfd', 1, 32, '18661667015', '2023-01-10 19:12:51', '2023-01-10 19:12:51', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5977932491456581, '1mv009vrq6d', '77f3834f752900643a4871dbd3044a0f76fe85ffd56f7d1ece9aa5f722a9830a', 1, 32, '18627331021', '2023-01-10 20:06:53', '2023-01-10 20:06:53', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5978295989436485, '1mv4n9h2yit', 'daba11bfd7475b986c3d048e576949bb070d7215b62dad03af4f4405aa7208ab', 1, 32, '19843061824', '2023-01-10 21:39:20', '2023-01-10 21:39:20', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5981433713131589, '1mw8opor8lh', '567530241859ac9079c45b237a53821d1aee87791aaea666ef177addf2fc3347', 1, 32, '18163386366', '2023-01-11 10:57:18', '2023-01-11 10:57:18', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5981495520723013, '1mw9h3vefd1', 'b7a6f1a88b595edfb4269d4bb3987c8f2fcac29a5bc006ee43f4152bbcd762d8', 1, 32, '18163386366', '2023-01-11 11:13:01', '2023-01-11 11:13:01', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5987149735460933, '1my9mm8v96t', '481479c5fb588fb4c1928d4e6eead2b49ac81f6b3ddbd87f1d7d575789ec73fc', 1, 32, '18625837911', '2023-01-12 11:10:59', '2023-01-12 11:10:59', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5987167558369349, '1my9ut061c5', '523646a2c61cb43667e3d9e2e02eadf7b6fd2063708f793b823a59a22ecda65e', 1, 32, '18625837911', '2023-01-12 11:15:42', '2023-01-12 11:15:42', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5987229831856197, '1myanew6zb9', '110cc50879e348570f0625e7da37bfad26b4c388ad3e991156a370b63e962284', 1, 32, '18155476184', '2023-01-12 11:31:35', '2023-01-12 11:31:35', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (5987317698330693, '1mybrs1mgxx', '662b1e37bf8e10cacf9d5ee2a03e49b45d1c15779f527ce399a2a11398bb30d0', 1, 32, '18355476184', '2023-01-12 11:53:40', '2023-01-12 11:53:40', NULL);
INSERT INTO "public"."sys_user" ("id", "username", "password", "state", "type", "mobile", "created_at", "updated_at", "deleted_at") VALUES (6000781163036741, '1n33kt4w6dh', '946849ab55eab8a935a71ba896e78dfbf89b22cf95222fa88392e241a9350fbd', 1, -1, '19890738726', '2023-01-14 20:57:36', '2023-01-14 20:57:36', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_detail
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_user_detail";
CREATE TABLE "public"."sys_user_detail" (
                                            "id" int8 NOT NULL,
                                            "realname" varchar(64) COLLATE "pg_catalog"."default",
                                            "union_main_name" varchar(128) COLLATE "pg_catalog"."default",
                                            "last_login_ip" varchar(64) COLLATE "pg_catalog"."default",
                                            "last_login_area" varchar(128) COLLATE "pg_catalog"."default",
                                            "last_login_at" timestamp(6)
)
;
ALTER TABLE "public"."sys_user_detail" OWNER TO "kysion";
COMMENT ON COLUMN "public"."sys_user_detail"."id" IS 'ID，保持与USERID一致';
COMMENT ON COLUMN "public"."sys_user_detail"."realname" IS '姓名';
COMMENT ON COLUMN "public"."sys_user_detail"."union_main_name" IS '关联主体名称';
COMMENT ON COLUMN "public"."sys_user_detail"."last_login_ip" IS '最后登录IP';
COMMENT ON COLUMN "public"."sys_user_detail"."last_login_area" IS '最后登录地区';
COMMENT ON COLUMN "public"."sys_user_detail"."last_login_at" IS '最后登录时间';

-- ----------------------------
-- Records of sys_user_detail
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_user_detail" ("id", "realname", "union_main_name", "last_login_ip", "last_login_area", "last_login_at") VALUES (5987167558369349, '蓝精灵', '快乐星球', '127.0.0.1', '0|0|0|内网IP|内网IP', '2023-02-13 10:17:33.858676');
INSERT INTO "public"."sys_user_detail" ("id", "realname", "union_main_name", "last_login_ip", "last_login_area", "last_login_at") VALUES (5977706248405061, '小小怪', '快乐星球', '127.0.0.1', '0|0|0|内网IP|内网IP', '2023-02-13 10:22:54.737823');
INSERT INTO "public"."sys_user_detail" ("id", "realname", "union_main_name", "last_login_ip", "last_login_area", "last_login_at") VALUES (5977719818551365, '游乐王子', '魔法城堡', '127.0.0.1', '0|0|0|内网IP|内网IP', '2023-02-13 10:24:43.80244');
COMMIT;

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_area_id_seq"
    OWNED BY "public"."sys_area"."id";
SELECT setval('"public"."sys_area_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_area_id_seq1"
    OWNED BY "public"."sys_area"."id";
SELECT setval('"public"."sys_area_id_seq1"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_area_id_seq2"
    OWNED BY "public"."sys_area"."id";
SELECT setval('"public"."sys_area_id_seq2"', 1, false);

-- ----------------------------
-- Indexes structure for table co_company
-- ----------------------------
CREATE INDEX "pro_facilitator_contact_mobile_idx_copy1" ON "public"."co_company" USING btree (
    "contact_mobile" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE INDEX "pro_facilitator_contact_name_idx_copy1" ON "public"."co_company" USING btree (
    "contact_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "pro_facilitator_id_idx_copy1" ON "public"."co_company" USING btree (
    "id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "pro_facilitator_name_idx_copy1" ON "public"."co_company" USING btree (
    "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table co_company
-- ----------------------------
ALTER TABLE "public"."co_company" ADD CONSTRAINT "pro_facilitator_copy1_name_key" UNIQUE ("name");

-- ----------------------------
-- Primary Key structure for table co_company
-- ----------------------------
ALTER TABLE "public"."co_company" ADD CONSTRAINT "pro_facilitator_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table co_company_employee
-- ----------------------------
CREATE UNIQUE INDEX "pro_facilitator_employee_id_idx_copy2" ON "public"."co_company_employee" USING btree (
    "id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE INDEX "pro_facilitator_employee_mobile_idx_copy2" ON "public"."co_company_employee" USING btree (
    "mobile" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "pro_facilitator_employee_name_facilitator_id_idx_copy2" ON "public"."co_company_employee" USING btree (
    "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "union_main_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table co_company_employee
-- ----------------------------
ALTER TABLE "public"."co_company_employee" ADD CONSTRAINT "pro_facilitator_employee_copy1_name_union_main_id_key" UNIQUE ("name", "union_main_id");

-- ----------------------------
-- Primary Key structure for table co_company_employee
-- ----------------------------
ALTER TABLE "public"."co_company_employee" ADD CONSTRAINT "pro_facilitator_employee_copy1_pkey1" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table co_company_team
-- ----------------------------
CREATE INDEX "pro_facilitator_team_id_idx_copy1" ON "public"."co_company_team" USING btree (
    "id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );
CREATE UNIQUE INDEX "pro_facilitator_team_name_facilitator_id_idx_copy1" ON "public"."co_company_team" USING btree (
    "name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "union_main_id" "pg_catalog"."int8_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table co_company_team
-- ----------------------------
ALTER TABLE "public"."co_company_team" ADD CONSTRAINT "pro_facilitator_team_copy1_name_facilitator_id_key" UNIQUE ("name", "union_main_id");

-- ----------------------------
-- Primary Key structure for table co_company_team
-- ----------------------------
ALTER TABLE "public"."co_company_team" ADD CONSTRAINT "pro_facilitator_team_copy1_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Uniques structure for table co_company_team_member
-- ----------------------------
ALTER TABLE "public"."co_company_team_member" ADD CONSTRAINT "pro_company_team_member_uniques" UNIQUE ("team_id", "employee_id", "union_main_id");

-- ----------------------------
-- Primary Key structure for table co_company_team_member
-- ----------------------------
ALTER TABLE "public"."co_company_team_member" ADD CONSTRAINT "pro_company_team_member_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Auto increment value for sys_area
-- ----------------------------
SELECT setval('"public"."sys_area_id_seq2"', 1, false);

-- ----------------------------
-- Primary Key structure for table sys_file
-- ----------------------------
ALTER TABLE "public"."sys_file" ADD CONSTRAINT "sys_file_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_license
-- ----------------------------
ALTER TABLE "public"."sys_license" ADD CONSTRAINT "pro_partner_license_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_menu
-- ----------------------------
ALTER TABLE "public"."sys_menu" ADD CONSTRAINT "sys_menu_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_organization
-- ----------------------------
ALTER TABLE "public"."sys_organization" ADD CONSTRAINT "sys_organization_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Indexes structure for table sys_permission
-- ----------------------------
CREATE INDEX "sys_permission_identifier_idx" ON "public"."sys_permission" USING btree (
    "identifier" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Uniques structure for table sys_permission
-- ----------------------------
ALTER TABLE "public"."sys_permission" ADD CONSTRAINT "sys_permission_identifier_key" UNIQUE ("identifier");

-- ----------------------------
-- Primary Key structure for table sys_permission
-- ----------------------------
ALTER TABLE "public"."sys_permission" ADD CONSTRAINT "sys_permission_pkey" PRIMARY KEY ("id", "identifier");

-- ----------------------------
-- Primary Key structure for table sys_role
-- ----------------------------
ALTER TABLE "public"."sys_role" ADD CONSTRAINT "sys_role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_sms_logs
-- ----------------------------
ALTER TABLE "public"."sys_sms_logs" ADD CONSTRAINT "sys_sms_logs_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_user
-- ----------------------------
ALTER TABLE "public"."sys_user" ADD CONSTRAINT "sys_user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_user_detail
-- ----------------------------
ALTER TABLE "public"."sys_user_detail" ADD CONSTRAINT "sys_user_detail_pkey" PRIMARY KEY ("id");
