/*
The MIT License (MIT)

Copyright (c) 2017 Chao Ma

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


package main

import "fmt"

var num  = []int{1, 2, 4, 8, 16, 32, 64, 128 , 256 , 512 , 1024 , 2046}//相当于移位
var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
var n int

func main() {
	base()
}

//基本主函数
func base() {
	fmt.Printf("please put in a target number :")
	fmt.Scanf("%d", &n)
	//错误处理
    if n <= 1 || n > 12 {
        fmt.Printf("this is not in the right range !")
    } else {
        //计算全组合总数
        m := counter(n)
        //运行全组合
        combination(m)
        //输出全组合数
        if n > 1 && n <= 12 {
            fmt.Printf("total is : %d", m)
        }
    }
}

//计算全组合数
func counter(n int) (int) {
	m := 1
	for i := 0; i < n; i++ {
		m = m * 2
	}
	m = m - 1
	return m
}

//运用类似移位的算法计算全组合
func combination(m int) {
    for i := 0; i < m + 1; i++ {
        for index , v := range num {
            if index < n {
                temp := i
                if (temp & v) != 0 {
                    fmt.Printf("%d ", list[index])
                }
            }
        }
        println()
    }
}