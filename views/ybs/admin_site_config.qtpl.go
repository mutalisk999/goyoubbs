// Code generated by qtc from "admin_site_config.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/ybs/admin_site_config.qtpl:1
package ybs

//line views/ybs/admin_site_config.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/ybs/admin_site_config.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/ybs/admin_site_config.qtpl:1
func (p *AdminSiteConfig) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/ybs/admin_site_config.qtpl:1
	qw422016.N().S(`
<div class="index">
    <div class="markdown-body entry-content">
        <h1>`)
//line views/ybs/admin_site_config.qtpl:4
	qw422016.E().S(p.Title)
//line views/ybs/admin_site_config.qtpl:4
	qw422016.N().S(`</h1>

        <div class="pure-button-group" role="group">
            <a href="/admin/site/router" class="pure-button">自定义路由</a>
            <a href="/admin/site/download/cur/db" class="pure-button">打包下载当前数据</a>
            <a href="/admin/site/download/cur/img" class="pure-button">打包下载图片</a>
        </div>

        <form action="" method="post" class="pure-form pure-form-aligned">
            <fieldset>
                <legend>网站设置</legend>
                <div class="pure-control-group">
                    <label for="Name">名称</label>
                    <input type="text" id="Name" name="Name" value="`)
//line views/ybs/admin_site_config.qtpl:17
	qw422016.E().S(p.SiteConf.Name)
//line views/ybs/admin_site_config.qtpl:17
	qw422016.N().S(`" placeholder="名称" />
                </div>
                <div class="pure-control-group">
                    <label for="Desc">描述</label>
                    <textarea id="Desc" name="Desc" class="pure-input-2-3" placeholder="描述">`)
//line views/ybs/admin_site_config.qtpl:21
	qw422016.E().S(p.SiteConf.Desc)
//line views/ybs/admin_site_config.qtpl:21
	qw422016.N().S(`</textarea>
                </div>
                <div class="pure-control-group">
                    <label for="MainDomain">主域名</label>
                    <input type="text" id="MainDomain" name="MainDomain" value="`)
//line views/ybs/admin_site_config.qtpl:25
	qw422016.E().S(p.SiteConf.MainDomain)
//line views/ybs/admin_site_config.qtpl:25
	qw422016.N().S(`" placeholder="主域名" />
                </div>
                <div class="pure-control-group">
                    <label for="HeaderPartCon">Header内容</label>
                    <textarea id="HeaderPartCon" name="HeaderPartCon" class="pure-input-2-3" placeholder="直接显示在页面 Header 里的内容">`)
//line views/ybs/admin_site_config.qtpl:29
	qw422016.E().S(p.SiteConf.HeaderPartCon)
//line views/ybs/admin_site_config.qtpl:29
	qw422016.N().S(`</textarea>
                </div>
                <div class="pure-control-group">
                    <label for="GoogleAutoAdJs">Google Adsense js</label>
                    <textarea id="GoogleAutoAdJs" name="GoogleAutoAdJs" class="pure-input-2-3" placeholder="放在Header 里的google adsense js 代码">`)
//line views/ybs/admin_site_config.qtpl:33
	qw422016.E().S(p.SiteConf.GoogleAutoAdJs)
//line views/ybs/admin_site_config.qtpl:33
	qw422016.N().S(`</textarea>
                </div>
                <div class="pure-control-group">
                    <label for="FooterPartHtml">页脚 html 内容</label>
                    <textarea id="FooterPartHtml" name="FooterPartHtml" class="pure-input-2-3" placeholder="直接显示在页脚的 html 内容，如备案信息及统计js代码">`)
//line views/ybs/admin_site_config.qtpl:37
	qw422016.E().S(p.SiteConf.FooterPartHtml)
//line views/ybs/admin_site_config.qtpl:37
	qw422016.N().S(`</textarea>
                </div>
                <div class="pure-control-group">
                    <label for="TimeZone">时区</label>
                    <input type="number" id="TimeZone" name="TimeZone" value="`)
//line views/ybs/admin_site_config.qtpl:41
	qw422016.N().D(p.SiteConf.TimeZone)
//line views/ybs/admin_site_config.qtpl:41
	qw422016.N().S(`" placeholder="时区" />
                    <span class="pure-form-message-inline">-12 ～ 12</span>
                </div>
                <div class="pure-control-group">
                    <label for="PageShowNum">每页显示条数</label>
                    <input type="number" id="PageShowNum" name="PageShowNum" value="`)
//line views/ybs/admin_site_config.qtpl:46
	qw422016.N().D(p.SiteConf.PageShowNum)
//line views/ybs/admin_site_config.qtpl:46
	qw422016.N().S(`" placeholder="每页显示文章数" />
                </div>
                <div class="pure-control-group">
                    <label for="TopRateNum">最近浏览条数</label>
                    <input type="number" id="TopRateNum" name="TopRateNum" value="`)
//line views/ybs/admin_site_config.qtpl:50
	qw422016.N().D(p.SiteConf.TopRateNum)
//line views/ybs/admin_site_config.qtpl:50
	qw422016.N().S(`" placeholder="侧栏最近浏览显示文章数" />
                </div>
                <div class="pure-control-group">
                    <label for="RecentCommentNum">最近评论条数</label>
                    <input type="number" id="RecentCommentNum" name="RecentCommentNum" value="`)
//line views/ybs/admin_site_config.qtpl:54
	qw422016.N().D(p.SiteConf.RecentCommentNum)
//line views/ybs/admin_site_config.qtpl:54
	qw422016.N().S(`" placeholder="侧栏显示最近评论条数" />
                </div>
                <div class="pure-control-group">
                    <label for="TitleMaxLen">标题最多字数</label>
                    <input type="number" id="TitleMaxLen" name="TitleMaxLen" value="`)
//line views/ybs/admin_site_config.qtpl:58
	qw422016.N().D(p.SiteConf.TitleMaxLen)
//line views/ybs/admin_site_config.qtpl:58
	qw422016.N().S(`" placeholder="标题最多字数" />
                </div>
                <div class="pure-control-group">
                    <label for="TopicConMaxLen">主贴内容最大字数</label>
                    <input type="number" id="TopicConMaxLen" name="TopicConMaxLen" value="`)
//line views/ybs/admin_site_config.qtpl:62
	qw422016.N().D(p.SiteConf.TopicConMaxLen)
//line views/ybs/admin_site_config.qtpl:62
	qw422016.N().S(`" placeholder="主贴内容最大字数" />
                </div>
                <div class="pure-control-group">
                    <label for="CommentConMaxLen">评论内容最大字数</label>
                    <input type="number" id="CommentConMaxLen" name="CommentConMaxLen" value="`)
//line views/ybs/admin_site_config.qtpl:66
	qw422016.N().D(p.SiteConf.CommentConMaxLen)
//line views/ybs/admin_site_config.qtpl:66
	qw422016.N().S(`" placeholder="评论内容最大字数" />
                </div>
                <div class="pure-control-group">
                    <label for="AutoDataBackup">自动备份数据库</label>
                    <input type="checkbox" id="AutoDataBackup" name="AutoDataBackup" value="true" `)
//line views/ybs/admin_site_config.qtpl:70
	if p.SiteConf.AutoDataBackup {
//line views/ybs/admin_site_config.qtpl:70
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:70
	}
//line views/ybs/admin_site_config.qtpl:70
	qw422016.N().S(` />
                </div>
                <div class="pure-control-group">
                    <label for="DataBackupDir">存放备份数据库目录</label>
                    <input type="text" class="pure-input-1-3" id="DataBackupDir" name="DataBackupDir" value="`)
//line views/ybs/admin_site_config.qtpl:74
	qw422016.E().S(p.SiteConf.DataBackupDir)
//line views/ybs/admin_site_config.qtpl:74
	qw422016.N().S(`" placeholder="存放备份数据库目录" />
                </div>
                <div class="pure-control-group">
                    <label for="Authorized">浏览网站需要登录</label>
                    <input type="checkbox" id="Authorized" name="Authorized" value="true" `)
//line views/ybs/admin_site_config.qtpl:78
	if p.SiteConf.Authorized {
//line views/ybs/admin_site_config.qtpl:78
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:78
	}
//line views/ybs/admin_site_config.qtpl:78
	qw422016.N().S(` /> 需要登录才能浏览页面
                </div>
                <div class="pure-control-group">
                    <label for="AllowNameReg">允许用户名注册</label>
                    <input type="checkbox" id="AllowNameReg" name="AllowNameReg" value="true" `)
//line views/ybs/admin_site_config.qtpl:82
	if p.SiteConf.AllowNameReg {
//line views/ybs/admin_site_config.qtpl:82
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:82
	}
//line views/ybs/admin_site_config.qtpl:82
	qw422016.N().S(` /> 若不勾选则只允许第三方登录
                </div>
                <div class="pure-control-group">
                    <label for="RegReview">用户注册审核</label>
                    <input type="checkbox" id="RegReview" name="RegReview" value="true" `)
//line views/ybs/admin_site_config.qtpl:86
	if p.SiteConf.RegReview {
//line views/ybs/admin_site_config.qtpl:86
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:86
	}
//line views/ybs/admin_site_config.qtpl:86
	qw422016.N().S(` />
                </div>
                <div class="pure-control-group">
                    <label for="CloseReg">关闭新用户注册</label>
                    <input type="checkbox" id="CloseReg" name="CloseReg" value="true" `)
//line views/ybs/admin_site_config.qtpl:90
	if p.SiteConf.CloseReg {
//line views/ybs/admin_site_config.qtpl:90
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:90
	}
//line views/ybs/admin_site_config.qtpl:90
	qw422016.N().S(` />
                </div>
                <div class="pure-control-group">
                    <label for="CloseReply">关闭评论</label>
                    <input type="checkbox" id="CloseReply" name="CloseReply" value="true" `)
//line views/ybs/admin_site_config.qtpl:94
	if p.SiteConf.CloseReply {
//line views/ybs/admin_site_config.qtpl:94
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:94
	}
//line views/ybs/admin_site_config.qtpl:94
	qw422016.N().S(` />
                </div>
                <div class="pure-control-group">
                    <label for="PostReview">发布审核</label>
                    <input type="checkbox" id="PostReview" name="PostReview" value="true" `)
//line views/ybs/admin_site_config.qtpl:98
	if p.SiteConf.PostReview {
//line views/ybs/admin_site_config.qtpl:98
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:98
	}
//line views/ybs/admin_site_config.qtpl:98
	qw422016.N().S(` /> 发帖、回复审核
                </div>
                <div class="pure-control-group">
                    <label for="ResetCookieKey">重设cookie key</label>
                    <input type="checkbox" id="ResetCookieKey" name="ResetCookieKey" value="true" `)
//line views/ybs/admin_site_config.qtpl:102
	if p.SiteConf.ResetCookieKey {
//line views/ybs/admin_site_config.qtpl:102
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:102
	}
//line views/ybs/admin_site_config.qtpl:102
	qw422016.N().S(` /> 立刻/每次重启 让用户重新登录
                </div>
                <div class="pure-control-group">
                    <label for="GetTagApi">分词URL</label>
                    <input type="text" class="pure-input-2-3" id="GetTagApi" name="GetTagApi" value="`)
//line views/ybs/admin_site_config.qtpl:106
	qw422016.E().S(p.SiteConf.GetTagApi)
//line views/ybs/admin_site_config.qtpl:106
	qw422016.N().S(`" placeholder="分词URL 对帖子标题提取标签的接口URL" />
                </div>
                <div class="pure-control-group">
                    <label for="UploadLimit">只允许管理员上传图片</label>
                    <input type="checkbox" id="UploadLimit" name="UploadLimit" value="true" `)
//line views/ybs/admin_site_config.qtpl:110
	if p.SiteConf.UploadLimit {
//line views/ybs/admin_site_config.qtpl:110
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:110
	}
//line views/ybs/admin_site_config.qtpl:110
	qw422016.N().S(` /> 若勾选则只允许管理员上传
                </div>
                <div class="pure-control-group">
                    <label for="UploadDir">存放用户上传图片目录</label>
                    <input type="text" class="pure-input-1-3" id="UploadDir" name="UploadDir" value="`)
//line views/ybs/admin_site_config.qtpl:114
	qw422016.E().S(p.SiteConf.UploadDir)
//line views/ybs/admin_site_config.qtpl:114
	qw422016.N().S(`" placeholder="存放用户上传图片目录" /> 一般填写后不需修改，否则以前上传的文件不可访问
                </div>
                <div class="pure-control-group">
                    <label for="UploadMaxSize">上传图片大小限制</label>
                    <input type="number" id="UploadMaxSize" name="UploadMaxSize" value="`)
//line views/ybs/admin_site_config.qtpl:118
	qw422016.N().D(p.SiteConf.UploadMaxSize)
//line views/ybs/admin_site_config.qtpl:118
	qw422016.N().S(`" placeholder="上传图片大小限制" />
                    <span class="pure-form-message-inline">M</span>
                </div>
                <div class="pure-control-group">
                    <label for="CachedSize">缓存大小</label>
                    <input type="number" id="CachedSize" name="CachedSize" value="`)
//line views/ybs/admin_site_config.qtpl:123
	qw422016.N().D(p.SiteConf.CachedSize)
//line views/ybs/admin_site_config.qtpl:123
	qw422016.N().S(`" placeholder="缓存大小" />
                    <span class="pure-form-message-inline">M</span>
                </div>
                <div class="pure-control-group">
                    <label for="SaveTopicIcon">保存九宫格图片</label>
                    <input type="checkbox" id="SaveTopicIcon" name="SaveTopicIcon" value="true" `)
//line views/ybs/admin_site_config.qtpl:128
	if p.SiteConf.SaveTopicIcon {
//line views/ybs/admin_site_config.qtpl:128
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:128
	}
//line views/ybs/admin_site_config.qtpl:128
	qw422016.N().S(` /> 帖子九宫格图片保存到数据库（以空间换CPU）
                </div>
                <div class="pure-control-group">
                    <label for="RemotePostPw">管理员远程发布密码</label>
                    <input type="text" class="pure-input-1-3" id="RemotePostPw" name="RemotePostPw" value="`)
//line views/ybs/admin_site_config.qtpl:132
	qw422016.E().S(p.SiteConf.RemotePostPw)
//line views/ybs/admin_site_config.qtpl:132
	qw422016.N().S(`" placeholder="请填写强类型密码" /> 管理员远程发布帖子、评论密码
                </div>
                <div class="pure-control-group">
                    <label for="BaiduSubUrl">百度提交网址</label>
                    <input type="text" class="pure-input-2-3" id="BaiduSubUrl" name="BaiduSubUrl" value="`)
//line views/ybs/admin_site_config.qtpl:136
	qw422016.E().S(p.SiteConf.BaiduSubUrl)
//line views/ybs/admin_site_config.qtpl:136
	qw422016.N().S(`" placeholder="百度提交网址, 发新帖自动提交到 百度" />
                </div>
                <div class="pure-control-group">
                    <label for="BingSubUrl">Bing 提交网址</label>
                    <input type="text" class="pure-input-2-3" id="BingSubUrl" name="BingSubUrl" value="`)
//line views/ybs/admin_site_config.qtpl:140
	qw422016.E().S(p.SiteConf.BingSubUrl)
//line views/ybs/admin_site_config.qtpl:140
	qw422016.N().S(`" placeholder="Bing 提交网址, 发新帖自动提交到 Bing" />
                </div>
                <div class="pure-control-group">
                    <label for="GoogleJWTConf">GoogleApi 配置</label>
                    <textarea id="GoogleJWTConf" name="GoogleJWTConf" class="pure-input-2-3" placeholder="GoogleApi JWT config json，发新帖自动提交到 Google">`)
//line views/ybs/admin_site_config.qtpl:144
	qw422016.E().S(p.SiteConf.GoogleJWTConf)
//line views/ybs/admin_site_config.qtpl:144
	qw422016.N().S(`</textarea>
                </div>
                <div class="pure-control-group">
                    <label for="QQClientID">QQClientID</label>
                    <input type="text" id="QQClientID" name="QQClientID" value="`)
//line views/ybs/admin_site_config.qtpl:148
	qw422016.E().S(p.SiteConf.QQClientID)
//line views/ybs/admin_site_config.qtpl:148
	qw422016.N().S(`" placeholder="QQClientID" />
                </div>
                <div class="pure-control-group">
                    <label for="QQClientSecret">QQClientSecret</label>
                    <input type="text" id="QQClientSecret" name="QQClientSecret" value="`)
//line views/ybs/admin_site_config.qtpl:152
	qw422016.E().S(p.SiteConf.QQClientSecret)
//line views/ybs/admin_site_config.qtpl:152
	qw422016.N().S(`" placeholder="QQClientSecret" />
                </div>
                <div class="pure-control-group">
                    <label for="WeiboClientID">WeiboClientID</label>
                    <input type="text" id="WeiboClientID" name="WeiboClientID" value="`)
//line views/ybs/admin_site_config.qtpl:156
	qw422016.E().S(p.SiteConf.WeiboClientID)
//line views/ybs/admin_site_config.qtpl:156
	qw422016.N().S(`" placeholder="WeiboClientID" />
                </div>
                <div class="pure-control-group">
                    <label for="WeiboClientSecret">WeiboClientSecret</label>
                    <input type="text" id="WeiboClientSecret" name="WeiboClientSecret" value="`)
//line views/ybs/admin_site_config.qtpl:160
	qw422016.E().S(p.SiteConf.WeiboClientSecret)
//line views/ybs/admin_site_config.qtpl:160
	qw422016.N().S(`" placeholder="WeiboClientSecret" />
                </div>
                <div class="pure-control-group">
                    <label for="GithubClientID">GithubClientID</label>
                    <input type="text" id="GithubClientID" name="GithubClientID" value="`)
//line views/ybs/admin_site_config.qtpl:164
	qw422016.E().S(p.SiteConf.GithubClientID)
//line views/ybs/admin_site_config.qtpl:164
	qw422016.N().S(`" placeholder="GithubClientID" />
                </div>
                <div class="pure-control-group">
                    <label for="GithubClientSecret">GithubClientSecret</label>
                    <input type="text" id="GithubClientSecret" name="GithubClientSecret" value="`)
//line views/ybs/admin_site_config.qtpl:168
	qw422016.E().S(p.SiteConf.GithubClientSecret)
//line views/ybs/admin_site_config.qtpl:168
	qw422016.N().S(`" placeholder="GithubClientSecret" />
                </div>
                <div class="pure-control-group">
                    <label for="SendEmail">发送Email通知</label>
                    <input type="checkbox" id="SendEmail" name="SendEmail" value="true" `)
//line views/ybs/admin_site_config.qtpl:172
	if p.SiteConf.SendEmail {
//line views/ybs/admin_site_config.qtpl:172
		qw422016.N().S(`checked`)
//line views/ybs/admin_site_config.qtpl:172
	}
//line views/ybs/admin_site_config.qtpl:172
	qw422016.N().S(` /> 有待验证帖子、回复是否发邮件，若需要则需正确填写下面 5 个信息 ↓
                </div>
                <div class="pure-control-group">
                    <label for="SmtpHost">SmtpHost</label>
                    <input type="text" id="SmtpHost" name="SmtpHost" value="`)
//line views/ybs/admin_site_config.qtpl:176
	qw422016.E().S(p.SiteConf.SmtpHost)
//line views/ybs/admin_site_config.qtpl:176
	qw422016.N().S(`" placeholder="smtp.126.com" />
                </div>
                <div class="pure-control-group">
                    <label for="SmtpPort">SmtpPort</label>
                    <input type="text" id="SmtpPort" name="SmtpPort" value="`)
//line views/ybs/admin_site_config.qtpl:180
	qw422016.N().D(p.SiteConf.SmtpPort)
//line views/ybs/admin_site_config.qtpl:180
	qw422016.N().S(`" placeholder="465" />
                </div>
                <div class="pure-control-group">
                    <label for="SmtpEmail">SmtpEmail</label>
                    <input type="text" id="SmtpEmail" name="SmtpEmail" value="`)
//line views/ybs/admin_site_config.qtpl:184
	qw422016.E().S(p.SiteConf.SmtpEmail)
//line views/ybs/admin_site_config.qtpl:184
	qw422016.N().S(`" placeholder="发件人邮箱 abc@126.com" />
                </div>
                <div class="pure-control-group">
                    <label for="SmtpPassword">SmtpPassword</label>
                    <input type="text" id="SmtpPassword" name="SmtpPassword" value="`)
//line views/ybs/admin_site_config.qtpl:188
	qw422016.E().S(p.SiteConf.SmtpPassword)
//line views/ybs/admin_site_config.qtpl:188
	qw422016.N().S(`" placeholder="发件人邮箱密码 xxxx" />
                </div>
                <div class="pure-control-group">
                    <label for="SendToEmail">SendToEmail</label>
                    <input type="text" id="SendToEmail" name="SendToEmail" value="`)
//line views/ybs/admin_site_config.qtpl:192
	qw422016.E().S(p.SiteConf.SendToEmail)
//line views/ybs/admin_site_config.qtpl:192
	qw422016.N().S(`" placeholder="收件人邮箱 123@qq.com" />
                </div>
                <div class="pure-controls">
                    <button type="submit" class="pure-button pure-button-primary">提交</button>
                </div>
            </fieldset>
        </form>

    </div>
</div>

`)
//line views/ybs/admin_site_config.qtpl:203
}

//line views/ybs/admin_site_config.qtpl:203
func (p *AdminSiteConfig) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/ybs/admin_site_config.qtpl:203
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/ybs/admin_site_config.qtpl:203
	p.StreamMainBody(qw422016)
//line views/ybs/admin_site_config.qtpl:203
	qt422016.ReleaseWriter(qw422016)
//line views/ybs/admin_site_config.qtpl:203
}

//line views/ybs/admin_site_config.qtpl:203
func (p *AdminSiteConfig) MainBody() string {
//line views/ybs/admin_site_config.qtpl:203
	qb422016 := qt422016.AcquireByteBuffer()
//line views/ybs/admin_site_config.qtpl:203
	p.WriteMainBody(qb422016)
//line views/ybs/admin_site_config.qtpl:203
	qs422016 := string(qb422016.B)
//line views/ybs/admin_site_config.qtpl:203
	qt422016.ReleaseByteBuffer(qb422016)
//line views/ybs/admin_site_config.qtpl:203
	return qs422016
//line views/ybs/admin_site_config.qtpl:203
}
