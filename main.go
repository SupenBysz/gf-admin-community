package component

import (
	"github.com/SupenBysz/gf-admin-community/internal/cmd"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/joho/godotenv"
	"os"
	"strings"

	_ "github.com/SupenBysz/gf-admin-community/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"
)

func init() {
	godotenv.Load()

	dbconfig := gdb.ConfigNode{
		Host:                 os.Getenv("DB_HOST"),
		Port:                 os.Getenv("DB_PORT"),
		User:                 os.Getenv("DB_USER"),
		Pass:                 os.Getenv("DB_PASS"),
		Name:                 os.Getenv("DB_NAME"),
		Type:                 os.Getenv("DB_TYPE"),
		Link:                 os.Getenv("DB_LINK"),
		Extra:                os.Getenv("DB_EXTRA"),
		Role:                 os.Getenv("DB_ROLE"),
		Debug:                gconv.Bool(os.Getenv("DB_DEBUG")),
		Prefix:               os.Getenv("DB_PREFIX"),
		DryRun:               gconv.Bool(os.Getenv("DB_DRYRUN")),
		Weight:               gconv.Int(os.Getenv("DB_WEIGHT")),
		Charset:              os.Getenv("DB_CHARSET"),
		Protocol:             os.Getenv("DB_PROTOCOL"),
		Timezone:             os.Getenv("DB_TIMEZONE"),
		MaxIdleConnCount:     gconv.Int(os.Getenv("DB_MAX_IDLE_CONN_COUNT")),
		MaxOpenConnCount:     gconv.Int(os.Getenv("DB_MAX_OPEN_CONN_COUNT")),
		MaxConnLifeTime:      gconv.Duration(os.Getenv("DB_MAX_CONN_LIFE_TIME")),
		QueryTimeout:         gconv.Duration(os.Getenv("DB_QUERY_TIMEOUT")),
		ExecTimeout:          gconv.Duration(os.Getenv("DB_EXEC_TIMEOUT")),
		TranTimeout:          gconv.Duration(os.Getenv("DB_TRAN_TIMEOUT")),
		PrepareTimeout:       gconv.Duration(os.Getenv("DB_PREPARE_TIMEOUT")),
		CreatedAt:            "",
		UpdatedAt:            "",
		DeletedAt:            "",
		TimeMaintainDisabled: gconv.Bool(os.Getenv("DB_TIME_MAINTAIN_DISABLED")),
	}

	if strings.Contains(dbconfig.Link, ":") || (dbconfig.Name != "" && dbconfig.User != "") {
		if strings.HasPrefix(dbconfig.Link, "postgres") {
			dbconfig.Type = "pgsql"
		} else {
			dbconfig.Type = strings.Split(dbconfig.Link, ":")[0]
		}

		gdb.SetConfig(gdb.Config{
			gdb.DefaultGroupName: gdb.ConfigGroup{
				dbconfig,
			},
		})
	}
}

func main() {
	cmd.Main.Run(gctx.New())
}
