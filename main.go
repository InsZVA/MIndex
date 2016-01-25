package main
//为爬取完的MySQL数据库做倒排索引

import (
	"flag"
	"fmt"
	"database/sql"
	"monkey"
	_"github.com/ziutek/mymysql/godrv"
	"github.com/adamzy/sego"
	"strconv"
	"strings"
)

var (
	//MySQLAddr = flag.String("mysql", "127.0.0.1:3306", "MySQL服务器地址及端口号")
	MySQLUser = flag.String("mysqlu", "root", "MySQL用户名")
	MySQLPasswd = flag.String("mysqlp", "root", "MySQL密码")
	MonkeyAddr = flag.String("monkey", "127.0.0.1", "monkeyDB地址")
	MonkeyPort = flag.String("monkeyP", "1517", "monkeyDB端口号")
	MonkeyPasswd = flag.String("monkeyp", "monkey", "monkeyDB登陆口令")
	Database = flag.String("database", "CloudKt", "要倒排的MySQL数据库")
	Table = flag.String("table", "courses", "要倒排的MySQL数据表")
)

var seg sego.Segmenter

func segoInit() {
	seg.LoadDictionary("./dictionary.txt")

	//segments := seg.Segment([]byte(*text))
	//fmt.Println(sego.SegmentsToSlice(segments, true))
}

func main() {
	flag.Parse()
	segoInit()
	db,err := sql.Open("mymysql",fmt.Sprintf("%s/%s/%s",*Database,*MySQLUser,*MySQLPasswd))
	if(err != nil){panic(err)}
	defer db.Close()
	monkeyCli,err := monkey.New(*MonkeyAddr,*MonkeyPort,*MonkeyPasswd)
	if(err != nil){panic(err)}
	monkeyCli.Send([]byte("createdb "+*Database))
	monkeyCli.Send([]byte("switchdb "+*Database))
	result,err := db.Query("select count(*) from "+*Table)
	if err != nil {
		panic(err)
	}
	var count int
	result.Next()
	result.Scan(&count)
	fmt.Println("任务总量：",count)
	for page := 0;page * 100 < count;page++ {
		result,err := db.Query("select id,title from "+*Table+" limit "+strconv.Itoa(page*100)+",100")
		if err != nil {
			panic(err)
		}
		for result.Next() {
			var id int
			var title string
			result.Scan(&id,&title)
			segments := seg.Segment([]byte(title))
			words := sego.SegmentsToSlice(segments,true)
			
			for _,word := range words {
				word = strings.Replace(word, " ", "", -1)
				if word == "" {
					continue
				}
				fmt.Println(word)
				current := monkeyCli.Send([]byte("get "+word))
				buff := []byte{}
				for i := 0;i < len(current);i++{
					if current[i] == 0 {
						break
					}
					buff = append(buff,current[i])
				}
				if len(buff) != 0 {
					buff = append(buff,',')
				}
				cmd := "set "+word+" "+string(buff)+strconv.Itoa(id)
				monkeyCli.Send([]byte(cmd))
				//fmt.Println(string(r))
				//fmt.Println(cmd)
				//fmt.Println("1 - ok")
			}
		}
		fmt.Println(page+1,"00 records finished!")
	}
}