package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func process(message string,length,tag int) int {
	messByte:=[]byte(message)
	if len(messByte)!=length{  //消息长度不一致
		return -4
	}
	for i := 0; i < len(messByte); i++ {
		if (messByte[i]<='9'&&messByte[i]>='0')||(messByte[i]<='e'&&messByte[i]>='a'){
		}else {
			return -3           //消息类型不正确
		}
	}
	for len(messByte)>0{
		
	}
}
func main() {
	reader := bufio.NewReader(os.Stdin) // 从标准输入生成读对象
	text, _ := reader.ReadString('\n') // 读到换行
    text = strings.TrimSpace(text)
	length,Tag:=0,0
	list:=strings.Split(text," ")
	length,_=strconv.Atoi(list[0])
	Tag,_=strconv.Atoi(list[1])

    text, _ = reader.ReadString('\n') // 读到换行
    text = strings.TrimSpace(text)
	fmt.Println(text)
	// mess:=0
    // fmt.Sscanf(text,"%x",&mess)
	
	
}