package boot

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_consts"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/kysion/base-library/utility/downloader"
	"github.com/lionsoul2014/ip2region/binding/golang/xdb"
	"log"
	"os"
)

func InitIp2region() {
	workFolder, _ := os.Getwd()
	folderPath := workFolder + "/resource/assets/static"

	if !gfile.Exists(folderPath) {
		_ = gfile.Mkdir(folderPath)
	}

	ip2regionPath := g.Cfg().MustGet(
		context.Background(),
		"service.ip2region.xdbPath",
		folderPath+"/ip2region.xdb",
	).String()

	if ip2regionPath == "" || gfile.Size(ip2regionPath) <= 0 {
		log.Println("开始下载IP信息库资源")
		d := downloader.NewDownloader(
			"https://mirror.ghproxy.com/https://github.com/lionsoul2014/ip2region/raw/master/data/ip2region.xdb",
			gfile.Basename(ip2regionPath),
			gfile.Abs(gfile.Dir(ip2regionPath)),
			10,
		)
		if err := d.Download(); err != nil {
			panic("ip2region 获取失败")
		}
	}
	if gfile.Size(ip2regionPath) <= 0 {
		panic("ip2region 校验失败")
	}

	cBuff, err := xdb.LoadContentFromFile(ip2regionPath)
	if err != nil {
		panic("ip2region 初始化失败")
	}
	sys_consts.Global.Searcher, _ = xdb.NewWithBuffer(cBuff)
}
