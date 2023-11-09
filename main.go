package main

import (
	_ "embed"
	"fmt"
	"github.com/go-ego/gse"
	"time"
)

//go:embed zh/s_1.txt
var zhs string

//go:embed zh/t_1.txt
var zht string

func main() {
	var sg gse.Segmenter
	// 嵌入字典
	err := sg.LoadDictEmbed(zhs + "\n" + zht)
	if err != nil {
		panic(err)
	}
	// 从文件读字典
	//err := sg.LoadDict("zh/s_1.txt,zh/t_1.txt")
	//if err != nil {
	//	panic(err)
	//}
	str := "复仇者联盟3：无限战争》是全片使用IMAX摄影机拍摄制作的的科幻片."
	cut := sg.Cut(str, true)
	fmt.Println(cut, len(cut))
	time.Sleep(time.Minute)
}
