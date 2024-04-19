/*
 Navicat Premium Data Transfer

 Source Server         : 筷满客-Dev
 Source Server Type    : PostgreSQL
 Source Server Version : 140008 (140008)
 Source Host           : 10.68.74.250:5432
 Source Catalog        : kuaimkdb
 Source Schema         : public

 Target Server Type    : PostgreSQL
 Target Server Version : 140008 (140008)
 File Encoding         : 65001

 Date: 19/04/2024 16:59:11
*/


-- ----------------------------
-- Table structure for sys_industry
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_industry";
CREATE TABLE "public"."sys_industry" (
                                         "id" int8 NOT NULL,
                                         "category_id" int8,
                                         "category_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                         "category_desc" varchar(255) COLLATE "pg_catalog"."default",
                                         "rate" int4,
                                         "parent_id" int8,
                                         "sort" int2,
                                         "state" int2,
                                         "created_at" timestamptz(6),
                                         "updated_at" timestamptz(6),
                                         "deleted_at" timestamptz(6)
)
;
ALTER TABLE "public"."sys_industry" OWNER TO "kuaimk";
COMMENT ON COLUMN "public"."sys_industry"."id" IS 'ID';
COMMENT ON COLUMN "public"."sys_industry"."category_id" IS '行业ID';
COMMENT ON COLUMN "public"."sys_industry"."category_name" IS '行业名称';
COMMENT ON COLUMN "public"."sys_industry"."category_desc" IS '行业描述';
COMMENT ON COLUMN "public"."sys_industry"."rate" IS '费率';
COMMENT ON COLUMN "public"."sys_industry"."parent_id" IS '父级ID';
COMMENT ON COLUMN "public"."sys_industry"."sort" IS '排序';
COMMENT ON COLUMN "public"."sys_industry"."state" IS '状态：0隐藏，1显示';

-- ----------------------------
-- Records of sys_industry
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149861818437, 0, '餐饮/零售', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:22.538835+08', '2024-04-19 16:54:22.538835+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149864702021, 1, '餐饮', '', NULL, 8610149861818437, 0, 1, '2024-04-19 16:54:22.582857+08', '2024-04-19 16:54:22.582857+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149866537029, 3, '零售', '', NULL, 8610149861818437, 1, 1, '2024-04-19 16:54:22.610379+08', '2024-04-19 16:54:22.610379+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149868240965, 4, '食品生鲜', '', NULL, 8610149861818437, 2, 1, '2024-04-19 16:54:22.63688+08', '2024-04-19 16:54:22.63688+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149870075973, 47, '文物经营/文物复制品销售', '', NULL, 8610149861818437, 3, 1, '2024-04-19 16:54:22.664285+08', '2024-04-19 16:54:22.664285+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149871845445, 63, '服饰鞋包', '', NULL, 8610149861818437, 4, 1, '2024-04-19 16:54:22.691251+08', '2024-04-19 16:54:22.691251+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149873614917, 64, '美妆日化', '', NULL, 8610149861818437, 5, 1, '2024-04-19 16:54:22.718185+08', '2024-04-19 16:54:22.718185+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149875318853, 65, '健身运动服务', '', NULL, 8610149861818437, 6, 1, '2024-04-19 16:54:22.744848+08', '2024-04-19 16:54:22.744848+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149876629573, 0, '交通/加油', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:22.764418+08', '2024-04-19 16:54:22.764418+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149878333509, 5, '快递', '', NULL, 8610149876629573, 0, 1, '2024-04-19 16:54:22.790986+08', '2024-04-19 16:54:22.790986+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149880168517, 6, '物流', '', NULL, 8610149876629573, 1, 1, '2024-04-19 16:54:22.818296+08', '2024-04-19 16:54:22.818296+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149881872453, 20, '加油/加气', '', NULL, 8610149876629573, 2, 1, '2024-04-19 16:54:22.844836+08', '2024-04-19 16:54:22.844836+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149883707461, 21, '铁路客运', '', NULL, 8610149876629573, 3, 1, '2024-04-19 16:54:22.872413+08', '2024-04-19 16:54:22.872413+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149885476933, 22, '高速公路收费', '', NULL, 8610149876629573, 4, 1, '2024-04-19 16:54:22.899428+08', '2024-04-19 16:54:22.899428+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149887246405, 23, '城市公共交通', '', NULL, 8610149876629573, 5, 1, '2024-04-19 16:54:22.926919+08', '2024-04-19 16:54:22.926919+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149889146949, 24, '船舶/海运服务', '', NULL, 8610149876629573, 6, 1, '2024-04-19 16:54:22.955951+08', '2024-04-19 16:54:22.955951+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149891113029, 26, '机票/票务代理', '', NULL, 8610149876629573, 7, 1, '2024-04-19 16:54:22.985416+08', '2024-04-19 16:54:22.985416+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149892882501, 56, '停车缴费', '', NULL, 8610149876629573, 8, 1, '2024-04-19 16:54:23.012998+08', '2024-04-19 16:54:23.012998+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149894258757, 0, '生活/娱乐', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:23.03341+08', '2024-04-19 16:54:23.03341+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149896093765, 7, '咨询/娱乐票务', '', NULL, 8610149894258757, 0, 1, '2024-04-19 16:54:23.061376+08', '2024-04-19 16:54:23.061376+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149897863237, 12, '共享服务', '', NULL, 8610149894258757, 1, 1, '2024-04-19 16:54:23.08829+08', '2024-04-19 16:54:23.08829+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149899698245, 13, '休闲娱乐/旅游服务', '', NULL, 8610149894258757, 2, 1, '2024-04-19 16:54:23.116424+08', '2024-04-19 16:54:23.116424+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149901467717, 14, '游艺厅/KTV', '', NULL, 8610149894258757, 3, 1, '2024-04-19 16:54:23.143967+08', '2024-04-19 16:54:23.143967+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149903302725, 15, '网吧', '', NULL, 8610149894258757, 4, 1, '2024-04-19 16:54:23.171843+08', '2024-04-19 16:54:23.171843+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149905203269, 16, '院线影城', '', NULL, 8610149894258757, 5, 1, '2024-04-19 16:54:23.200328+08', '2024-04-19 16:54:23.200328+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149906972741, 17, '演出赛事', '', NULL, 8610149894258757, 6, 1, '2024-04-19 16:54:23.227896+08', '2024-04-19 16:54:23.227896+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149908742213, 18, '居民生活服务', '', NULL, 8610149894258757, 7, 1, '2024-04-19 16:54:23.254923+08', '2024-04-19 16:54:23.254923+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149910577221, 19, '景区/酒店', '', NULL, 8610149894258757, 8, 1, '2024-04-19 16:54:23.282862+08', '2024-04-19 16:54:23.282862+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149912412229, 25, '旅行社', '', NULL, 8610149894258757, 9, 1, '2024-04-19 16:54:23.310243+08', '2024-04-19 16:54:23.310243+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149913722949, 0, '互联网服务', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:23.330385+08', '2024-04-19 16:54:23.330385+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149915492421, 2, '电商平台', '', NULL, 8610149913722949, 0, 1, '2024-04-19 16:54:23.357858+08', '2024-04-19 16:54:23.357858+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149917261893, 8, '婚介平台/就业信息平台/其他信息服务平台', '', NULL, 8610149913722949, 1, 1, '2024-04-19 16:54:23.384464+08', '2024-04-19 16:54:23.384464+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149919096901, 27, '在线图书/视频/音乐', '', NULL, 8610149913722949, 2, 1, '2024-04-19 16:54:23.412909+08', '2024-04-19 16:54:23.412909+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149920931909, 28, '门户论坛/网络广告及推广/软件开发/其他互联网服务', '', NULL, 8610149913722949, 3, 1, '2024-04-19 16:54:23.44019+08', '2024-04-19 16:54:23.44019+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149922635845, 29, '游戏', '', NULL, 8610149913722949, 4, 1, '2024-04-19 16:54:23.466574+08', '2024-04-19 16:54:23.466574+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149924339781, 30, '网络直播/直播平台', '', NULL, 8610149913722949, 5, 1, '2024-04-19 16:54:23.493027+08', '2024-04-19 16:54:23.493027+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149926109253, 45, '互联网募捐信息平台', '', NULL, 8610149913722949, 6, 1, '2024-04-19 16:54:23.519726+08', '2024-04-19 16:54:23.519726+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149927419973, 0, '房产/金融', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:23.539183+08', '2024-04-19 16:54:23.539183+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149929123909, 9, '房地产', '', NULL, 8610149927419973, 0, 1, '2024-04-19 16:54:23.56572+08', '2024-04-19 16:54:23.56572+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149930893381, 10, '房产中介', '', NULL, 8610149927419973, 1, 1, '2024-04-19 16:54:23.592213+08', '2024-04-19 16:54:23.592213+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149932662853, 43, '财经/股票类资讯', '', NULL, 8610149927419973, 2, 1, '2024-04-19 16:54:23.619408+08', '2024-04-19 16:54:23.619408+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149934497861, 44, '保险业务', '', NULL, 8610149927419973, 3, 1, '2024-04-19 16:54:23.647363+08', '2024-04-19 16:54:23.647363+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149936332869, 46, '信用还款', '', NULL, 8610149927419973, 4, 1, '2024-04-19 16:54:23.67546+08', '2024-04-19 16:54:23.67546+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149938167877, 48, '典当', '', NULL, 8610149927419973, 5, 1, '2024-04-19 16:54:23.703266+08', '2024-04-19 16:54:23.703266+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149939937349, 60, '银行还款', '银行信贷还款', NULL, 8610149927419973, 6, 1, '2024-04-19 16:54:23.730356+08', '2024-04-19 16:54:23.730356+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149941248069, 0, '民生缴费', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:23.750452+08', '2024-04-19 16:54:23.750452+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149942952005, 38, '虚拟充值', '', NULL, 8610149941248069, 0, 1, '2024-04-19 16:54:23.776979+08', '2024-04-19 16:54:23.776979+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149944721477, 39, '有线电视缴费', '', NULL, 8610149941248069, 1, 1, '2024-04-19 16:54:23.80397+08', '2024-04-19 16:54:23.80397+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149946490949, 40, '其他缴费', '', NULL, 8610149941248069, 2, 1, '2024-04-19 16:54:23.830757+08', '2024-04-19 16:54:23.830757+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149948260421, 41, '水电煤气缴费', '', NULL, 8610149941248069, 3, 1, '2024-04-19 16:54:23.857846+08', '2024-04-19 16:54:23.857846+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149950029893, 54, '党费', '', NULL, 8610149941248069, 4, 1, '2024-04-19 16:54:23.884255+08', '2024-04-19 16:54:23.884255+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149951275077, 0, '医疗', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:23.903976+08', '2024-04-19 16:54:23.903976+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149953044549, 11, '宠物医院', '', NULL, 8610149951275077, 0, 1, '2024-04-19 16:54:23.930799+08', '2024-04-19 16:54:23.930799+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149954748485, 34, '保健器械/医疗器械/非处方药品', '', NULL, 8610149951275077, 1, 1, '2024-04-19 16:54:23.956753+08', '2024-04-19 16:54:23.956753+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149956452421, 35, '其他私立/民营医院/诊所', '', NULL, 8610149951275077, 2, 1, '2024-04-19 16:54:23.982692+08', '2024-04-19 16:54:23.982692+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149958287429, 62, '私立/民营口腔及眼科、医疗美容医院/诊所', '', NULL, 8610149951275077, 3, 1, '2024-04-19 16:54:24.010691+08', '2024-04-19 16:54:24.010691+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149959598149, 0, '教育', NULL, NULL, 0, NULL, 1, '2024-04-19 16:54:24.030659+08', '2024-04-19 16:54:24.030659+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149961302085, 31, '培训机构', '', NULL, 8610149959598149, 0, 1, '2024-04-19 16:54:24.056756+08', '2024-04-19 16:54:24.056756+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149963006021, 32, '民办大学及院校', '', NULL, 8610149959598149, 1, 1, '2024-04-19 16:54:24.082793+08', '2024-04-19 16:54:24.082793+08', NULL);
INSERT INTO "public"."sys_industry" ("id", "category_id", "category_name", "category_desc", "rate", "parent_id", "sort", "state", "created_at", "updated_at", "deleted_at") VALUES (8610149964841029, 33, '民办学校（非全国高等学校）', '', NULL, 8610149959598149, 2, 1, '2024-04-19 16:54:24.110343+08', '2024-04-19 16:54:24.110343+08', NULL);
COMMIT;

-- ----------------------------
-- Primary Key structure for table sys_industry
-- ----------------------------
ALTER TABLE "public"."sys_industry" ADD CONSTRAINT "sys_industry_pkey" PRIMARY KEY ("id");
