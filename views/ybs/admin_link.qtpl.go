// Code generated by qtc from "admin_link.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

//line views/ybs/admin_link.qtpl:1
package ybs

//line views/ybs/admin_link.qtpl:1
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line views/ybs/admin_link.qtpl:1
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line views/ybs/admin_link.qtpl:1
func (p *AdminLink) StreamMainBody(qw422016 *qt422016.Writer) {
//line views/ybs/admin_link.qtpl:1
	qw422016.N().S(`
<div class="index">
    <div class="markdown-body entry-content">
    <h1>`)
//line views/ybs/admin_link.qtpl:4
	qw422016.E().S(p.Title)
//line views/ybs/admin_link.qtpl:4
	qw422016.N().S(`</h1>

    <form action="" method="post" class="pure-form pure-form-stacked">
        <fieldset>
            <legend>`)
//line views/ybs/admin_link.qtpl:8
	qw422016.E().S(p.Act)
//line views/ybs/admin_link.qtpl:8
	qw422016.N().S(` 链接</legend>

            <p>Score: 显示排序，大排前</p>

            <div class="pure-g">
                <div class="pure-u-1 pure-u-sm-1-5">
                    <label for="Name">Name</label>
                    <input id="Name" name="Name" class="pure-u-23-24" type="text" value="`)
//line views/ybs/admin_link.qtpl:15
	qw422016.E().S(p.Link.Name)
//line views/ybs/admin_link.qtpl:15
	qw422016.N().S(`" required>
                </div>

                <div class="pure-u-1 pure-u-sm-1-5">
                    <label for="Score">Score</label>
                    <input id="Score" name="Score" class="pure-u-23-24" type="number" value="`)
//line views/ybs/admin_link.qtpl:20
	qw422016.N().D(p.Link.Score)
//line views/ybs/admin_link.qtpl:20
	qw422016.N().S(`" required>
                </div>

                <div class="pure-u-1 pure-u-sm-2-5">
                    <label for="Url">Url</label>
                    <input id="Url" name="Url" class="pure-u-23-24" type="text" value="`)
//line views/ybs/admin_link.qtpl:25
	qw422016.E().S(p.Link.Url)
//line views/ybs/admin_link.qtpl:25
	qw422016.N().S(`" required>
                </div>

            </div>

            <button type="submit" class="pure-button pure-button-primary">提交</button>
        </fieldset>
    </form>

    <h2>链接列表</h2>

    <ul>
        <li class="bot-line">
            ID - Name - Score - Url
        </li>
        `)
//line views/ybs/admin_link.qtpl:40
	for _, v := range p.LinkLst {
//line views/ybs/admin_link.qtpl:40
		qw422016.N().S(`
        <li class="bot-line">
            `)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().DUL(v.ID)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().S(` - <a href="/admin/link?id=`)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().DUL(v.ID)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().S(`">`)
//line views/ybs/admin_link.qtpl:42
		qw422016.E().S(v.Name)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().S(`</a> - `)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().D(v.Score)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().S(` - `)
//line views/ybs/admin_link.qtpl:42
		qw422016.E().S(v.Url)
//line views/ybs/admin_link.qtpl:42
		qw422016.N().S(`
        </li>
        `)
//line views/ybs/admin_link.qtpl:44
	}
//line views/ybs/admin_link.qtpl:44
	qw422016.N().S(`
    </ul>

    </div>
</div>

`)
//line views/ybs/admin_link.qtpl:50
}

//line views/ybs/admin_link.qtpl:50
func (p *AdminLink) WriteMainBody(qq422016 qtio422016.Writer) {
//line views/ybs/admin_link.qtpl:50
	qw422016 := qt422016.AcquireWriter(qq422016)
//line views/ybs/admin_link.qtpl:50
	p.StreamMainBody(qw422016)
//line views/ybs/admin_link.qtpl:50
	qt422016.ReleaseWriter(qw422016)
//line views/ybs/admin_link.qtpl:50
}

//line views/ybs/admin_link.qtpl:50
func (p *AdminLink) MainBody() string {
//line views/ybs/admin_link.qtpl:50
	qb422016 := qt422016.AcquireByteBuffer()
//line views/ybs/admin_link.qtpl:50
	p.WriteMainBody(qb422016)
//line views/ybs/admin_link.qtpl:50
	qs422016 := string(qb422016.B)
//line views/ybs/admin_link.qtpl:50
	qt422016.ReleaseByteBuffer(qb422016)
//line views/ybs/admin_link.qtpl:50
	return qs422016
//line views/ybs/admin_link.qtpl:50
}
