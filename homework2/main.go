package main

import (
	"flag"
	"fmt"
)
func main() {
	var num int
	flag.IntVar(&num, "type", -1,  "本程序接受h,type参数。输入h时候,命令行显示使用帮助; type值为1时输出一句唐诗,为2时输出一句宋词")
	
	flag.Parse() 
	
	Test(num)
}

//实现函数
func Test(res int)  {

    if  res==1 {
        fmt.Println("相思似海深，旧事如天远")
	}else if res==2{
        fmt.Println("汗马嘶风，边鸿翻月，垅上铁衣寒早。")
	}
}
