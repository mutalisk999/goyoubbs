package cronjob

import (
	"github.com/ego008/goutils/splock"
	"github.com/ego008/sdb"
	"goyoubbs/app"
	"time"
)

type BaseHandler struct {
	App *app.Application
}

func (h *BaseHandler) MainCronJob() {

	db := h.App.Db

	spl := splock.SimpleLock{}
	tick1 := time.Tick(3 * time.Minute)   // 清理过期一些 keys
	tick3 := time.Tick(30 * time.Minute)  // 数据库备份
	tick6 := time.Tick(9 * time.Second)   // 从 title 提取 tag
	tick7 := time.Tick(29 * time.Second)  // 设置 topic tag
	tick8 := time.Tick(240 * time.Second) // 向百度、bing、google 提交网址
	tick10 := time.Tick(21 * time.Second) // avatar
	tick11 := time.Tick(39 * time.Second) // send mail

	daySecond := int64(3600 * 24)

	for {
		select {
		case <-tick1:
			lk := spl.Init("tk1")
			if !lk.IsLocked() {
				lk.Lock()
				// 做一些清理工作
				limit := 10
				timeBefore := uint64(time.Now().UTC().Unix() - daySecond)
				scoreStartB := sdb.I2b(timeBefore)
				zbnList := []string{
					"article_detail_token",
					// "user_login_token", // 登录出错限制，未实现
				}
				for _, bn := range zbnList {
					rs := db.Zrscan(bn, nil, scoreStartB, limit)
					if rs.OK() {
						keys := make([][]byte, len(rs.Data)/2)
						j := 0
						for i := 0; i < (len(rs.Data) - 1); i += 2 {
							keys[j] = rs.Data[i]
							j++
						}
						_ = db.Zmdel(bn, keys)
					}
				}
				lk.UnLock()
			}
		case <-tick3:
			if h.App.Cf.Site.AutoDataBackup {
				dataBackup(db, h.App.Cf.Site.DataBackupDir)
			}
		case <-tick6:
			if len(h.App.Cf.Site.GetTagApi) > 0 {
				lk := spl.Init("tk6")
				if !lk.IsLocked() {
					lk.Lock()
					getTagFromTitle(db, h.App.Cf.Site.GetTagApi)
					lk.UnLock()
				}
			}
		case <-tick7:
			lk := spl.Init("tk7")
			if !lk.IsLocked() {
				lk.Lock()
				setArticleTag(h.App.Mc, db)
				lk.UnLock()
			}
		case <-tick8:
			if len(h.App.Cf.Site.BaiduSubUrl) > 0 {
				var lk = spl.Init("tk8")
				if !lk.IsLocked() {
					lk.Lock()
					submitUrl(db, h.App.Cf.Site)
					lk.UnLock()
				}
			}
		case <-tick10:
			lk := spl.Init("tk10")
			if !lk.IsLocked() {
				lk.Lock()
				SetAvatar(db, h.App.Cf.Site.Socks5Proxy)
				lk.UnLock()
			}
		case <-tick11:
			lk := spl.Init("tk11")
			if !lk.IsLocked() {
				lk.Lock()
				sendMail(db, h.App.Cf.Site)
				lk.UnLock()
			}
		default:
			time.Sleep(time.Second)
		}
	}
}
