package main

import (
    "fmt"
    "github.com/PuerkitoBio/goquery"
    "github.com/astaxie/beego/orm"
    _ "github.com/go-sql-driver/mysql" // import your used driver
    "github.com/hu17889/go_spider/core/common/page"
    "github.com/hu17889/go_spider/core/pipeline"
    "github.com/hu17889/go_spider/core/spider"
    _ "github.com/lib/pq"
)
type MyPageProcesser struct {
    startNewsId int
}
// Model Struct
type User struct {
    Id   int
    Name string `orm:"size(100)"`
}

func init() {
    //// set default database
    //orm.RegisterDataBase("default", "mysql", "root:root@tcp(192.168.31.75:3306)/aoi?charset=utf8", 30)
    //
    //// register model
    //orm.RegisterModel(new(User))
    //
    //// create table
    //orm.RunSyncdb("default", false, true)

    orm.RegisterDriver("postgres", orm.DRPostgres)

    orm.RegisterDataBase("default", "postgres", "postgres://postgres:baymin1024@rest.yining.site/spider?sslmode=disable")

    orm.RegisterModel(new(User))
    // create table
    orm.RunSyncdb("default", false, true)
}
func NewMyPageProcesser() *MyPageProcesser {
    return &MyPageProcesser{}
}

// Parse html dom here and record the parse result that we want to crawl.
// Package simplejson (https://github.com/bitly/go-simplejson) is used to parse data of json.
func (this *MyPageProcesser) Process(p *page.Page) {
    if !p.IsSucc() {
        println(p.Errormsg())
        return
    }

    query := p.GetHtmlParser()
    query.Find("#sf-item-list-data").Each(func(i int, s *goquery.Selection) {

        href:= s.Text()
        fmt.Printf(href)
    })
    //println(p.GetTargetRequests())

}

func (this *MyPageProcesser) Finish() {
    fmt.Printf("TODO:before end spider \r\n")
}

func main() {

    spider.NewSpider(NewMyPageProcesser(), "fuck").
        AddUrl("https://sf.taobao.com/item_list.htm?city=%C4%FE%B2%A8&category=50025969", "html"). // start url, html is the responce type ("html" or "json" or "jsonp" or "text")
        AddPipeline(pipeline.NewPipelineConsole()).                                                                                   // Print result to std output
        AddPipeline(pipeline.NewPipelineFile("/tmp/sinafile")).                                                                       // Print result in file
        OpenFileLog("/tmp").                                                                                                          // Error info or other useful info in spider will be logged in file of defalt path like "WD/log/log.2014-9-1".
        SetSleepTime("rand", 1000, 3000).                                                                                             // Sleep time between 1s and 3s.
        Run()

    o := orm.NewOrm()

    user := User{Name: "slene"}

    // insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    // update
    user.Name = "astaxie"
    num, err := o.Update(&user)
    fmt.Printf("NUM: %d, ERR: %v\n", num, err)

    // read one
    u := User{Id: user.Id}
    err = o.Read(&u)
    fmt.Printf("ERR: %v\n", err)


    //AddPipeline(pipeline.NewPipelineFile("/tmp/tmpfile")). // print result in file

    // delete
    //num, err = o.Delete(&u)
    //fmt.Printf("NUM: %d, ERR: %v\n", num, err)
}
