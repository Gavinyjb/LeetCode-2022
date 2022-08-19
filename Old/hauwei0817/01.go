package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func process(weizhi [][]string,ganrans []string,m,n int)(first,second []string)  {
	if len(ganrans)==0||ganrans==nil{
		return nil,nil
	}
	yangMap:=make(map[string]bool)  //储存感染者列表,便于查询
	for _, v := range ganrans {
		yangMap[v]=true
	}
	//生成[][]byte 图 '1'为阳性 
	helps:=make([][]byte,m)
	for i := 0; i < m; i++ {
		list:=make([]byte,n)
		for j := 0; j <n ; j++ {
			if yangMap[weizhi[i][j]]==true{
				list[j]='1'
			}else {
				list[j]='0'
			}
		}
		helps = append(helps, list)
	} 
	
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helps[i][j]=='1'{  //i,j为感染者
				if i-1>=0&&helps[i-1][j]=='0'{  //上边
					helps[i-1][j]='2'
				}
				if i+1<m&&helps[i+1][j]=='1'{  //下边
					helps[i+1][j]='2'
				}
				if j-1>=0&&helps[i][j-1]=='0'{  //左边
					helps[i][j-1]='2'
				}
				if j+1<n&&helps[i][j+1]=='0'{    //右边
					helps[i][j+1]='2'
				}

				if i-1>=0&&j-1>=0&&helps[i-1][j-1]=='0'{  //左上边
					helps[i-1][j-1]='2'
				}
				if i+1<m&&j+1<n&&helps[i+1][j+1]=='1'{  //右下边
					helps[i+1][j+1]='2'
				}
				if i+1<m&&j-1>=0&&helps[i+1][j-1]=='0'{  //左下边
					helps[i+1][j-1]='2'
				}
				if i-1>=0&&j+1<n&&helps[i-1][j+1]=='0'{    //右上边
					helps[i-1][j+1]='2'
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helps[i][j]=='2'{  //i,j为密接
				if i-1>=0&&helps[i-1][j]=='0'{  //上边
					helps[i-1][j]='3'
				}
				if i+1<m&&helps[i+1][j]=='0'{  //下边
					helps[i+1][j]='3'
				}
				if j-1>=0&&helps[i][j-1]=='0'{  //左边
					helps[i][j-1]='3'
				}
				if j+1<n&&helps[i][j+1]=='0'{    //右边
					helps[i][j+1]='3'
				}
				if i-1>=0&&j-1>=0&&helps[i-1][j-1]=='0'{  //左上边
					helps[i-1][j-1]='3'
				}
				if i+1<m&&j+1<n&&helps[i+1][j+1]=='1'{  //右下边
					helps[i+1][j+1]='3'
				}
				if i+1<m&&j-1>=0&&helps[i+1][j-1]=='0'{  //左下边
					helps[i+1][j-1]='3'
				}
				if i-1>=0&&j+1<n&&helps[i-1][j+1]=='0'{    //右上边
					helps[i-1][j+1]='3'
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if helps[i][j]=='2'{
				first = append(first, weizhi[i][j])
			}
			if helps[i][j]=='3'{
				second = append(second, weizhi[i][j])
			}
		}
	}
	fmt.Println(first)
	fmt.Println(second)
	return first,second
}

func main() {

// L
// 6 5
// A B C D E
// F G H I J
// K L M N O
// P Q R S T
// U V W X Y
// Z 0 1 2 3

	inputreader := bufio.NewReader(os.Stdin)
	input,_:=inputreader.ReadString('\n')
	input=strings.Trim(input,"\n")

	ganrans:=strings.Split(input," ")   //感染者列表
	
	m, n := 0, 0 //m*n矩阵
	// fmt.Scanf("%d %d\n",&m,&n)
	inputMN,_:=inputreader.ReadString('\n')
	mn:=strings.Split(inputMN," ")
	m,_=strconv.Atoi(mn[0])
	n,_=strconv.Atoi(mn[1])

	weizhi := make([][]string, m)   //所有人座位
	for i := 0; i < m; i++ {
		list:=make([]string,n) //每一行
		listStr,_:=inputreader.ReadString('\n')
		nums:=strings.Split(listStr," ")
		for j := 0; j < n; j++ {
			list[i]=nums[i]
		}
		weizhi = append(weizhi, list)
	}
	first,second:=process(weizhi,ganrans,m,n)
	sort.Strings(first)
	sort.Strings(second)
	res:=append(first,second...)
	for i := 0; i < len(res); i++ {
		fmt.Print(res[i])
		fmt.Print(" ")
	}
}