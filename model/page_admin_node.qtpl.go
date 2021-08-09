// Code generated by qtc from "page_admin_node.qtpl". DO NOT EDIT.
// See https://github.com/valyala/quicktemplate for details.

// admin_login
//

//line model/page_admin_node.qtpl:3
package model

//line model/page_admin_node.qtpl:3
import (
	qtio422016 "io"

	qt422016 "github.com/valyala/quicktemplate"
)

//line model/page_admin_node.qtpl:3
var (
	_ = qtio422016.Copy
	_ = qt422016.AcquireByteBuffer
)

//line model/page_admin_node.qtpl:4
type AdminNode struct {
	BasePage
	Act  string // 行为名称，添加/编辑
	Node Node   // 分区
}

//line model/page_admin_node.qtpl:11
func (p *AdminNode) StreamMainBody(qw422016 *qt422016.Writer) {
//line model/page_admin_node.qtpl:11
	qw422016.N().S(`
<div class="index">
    <div class="markdown-body entry-content">
    <h1>`)
//line model/page_admin_node.qtpl:14
	qw422016.E().S(p.Title)
//line model/page_admin_node.qtpl:14
	qw422016.N().S(`</h1>

    <form action="" method="post" class="pure-form pure-form-stacked">
        <fieldset>
            <legend>`)
//line model/page_admin_node.qtpl:18
	qw422016.E().S(p.Act)
//line model/page_admin_node.qtpl:18
	qw422016.N().S(` 节点</legend>

            <p>Score: 显示排序，大排前</p>

            <div class="pure-g">
                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="Name">Name</label>
                    <input id="Name" name="Name" class="pure-u-23-24" type="text" value="`)
//line model/page_admin_node.qtpl:25
	qw422016.E().S(p.Node.Name)
//line model/page_admin_node.qtpl:25
	qw422016.N().S(`" required>
                </div>

                <div class="pure-u-1 pure-u-sm-1-6">
                    <label for="Score">Score</label>
                    <input id="Score" name="Score" class="pure-u-23-24" type="number" value="`)
//line model/page_admin_node.qtpl:30
	qw422016.N().D(p.Node.Score)
//line model/page_admin_node.qtpl:30
	qw422016.N().S(`" required>
                </div>

            </div>

            <div class="pure-g">
                <div class="pure-u-1 pure-u-sm-1-1">
                    <label for="About">About</label>
                    <textarea id="About" name="About" class="pure-input-1">`)
//line model/page_admin_node.qtpl:38
	qw422016.E().S(p.Node.About)
//line model/page_admin_node.qtpl:38
	qw422016.N().S(`</textarea>
                </div>
            </div>


            <button type="submit" class="pure-button pure-button-primary">提交</button>
        </fieldset>
    </form>

    <h2>节点列表</h2>
    <ul>
        <li class="bot-line">
            ID - Name - Score - About
        </li>
        `)
//line model/page_admin_node.qtpl:52
	for _, v := range p.NodeLst {
//line model/page_admin_node.qtpl:52
		qw422016.N().S(`
        <li class="bot-line">
            `)
//line model/page_admin_node.qtpl:54
		qw422016.N().DUL(v.ID)
//line model/page_admin_node.qtpl:54
		qw422016.N().S(` - <a href="/admin/node?id=`)
//line model/page_admin_node.qtpl:54
		qw422016.N().DUL(v.ID)
//line model/page_admin_node.qtpl:54
		qw422016.N().S(`">`)
//line model/page_admin_node.qtpl:54
		qw422016.E().S(v.Name)
//line model/page_admin_node.qtpl:54
		qw422016.N().S(`</a> - `)
//line model/page_admin_node.qtpl:54
		qw422016.N().D(v.Score)
//line model/page_admin_node.qtpl:54
		qw422016.N().S(` - `)
//line model/page_admin_node.qtpl:54
		qw422016.E().S(v.About)
//line model/page_admin_node.qtpl:54
		qw422016.N().S(`
        </li>
        `)
//line model/page_admin_node.qtpl:56
	}
//line model/page_admin_node.qtpl:56
	qw422016.N().S(`
    </ul>


    </div>
</div>

`)
//line model/page_admin_node.qtpl:63
}

//line model/page_admin_node.qtpl:63
func (p *AdminNode) WriteMainBody(qq422016 qtio422016.Writer) {
//line model/page_admin_node.qtpl:63
	qw422016 := qt422016.AcquireWriter(qq422016)
//line model/page_admin_node.qtpl:63
	p.StreamMainBody(qw422016)
//line model/page_admin_node.qtpl:63
	qt422016.ReleaseWriter(qw422016)
//line model/page_admin_node.qtpl:63
}

//line model/page_admin_node.qtpl:63
func (p *AdminNode) MainBody() string {
//line model/page_admin_node.qtpl:63
	qb422016 := qt422016.AcquireByteBuffer()
//line model/page_admin_node.qtpl:63
	p.WriteMainBody(qb422016)
//line model/page_admin_node.qtpl:63
	qs422016 := string(qb422016.B)
//line model/page_admin_node.qtpl:63
	qt422016.ReleaseByteBuffer(qb422016)
//line model/page_admin_node.qtpl:63
	return qs422016
//line model/page_admin_node.qtpl:63
}