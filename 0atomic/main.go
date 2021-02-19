package main

import "log"

/**
原子计数
原子性就是操作是不可中断的,
可以当做为最轻量的锁
尽量不用,善用chan
当多个协程操作同一个变量时 可能会出现问题
比如协程1 和协程2同时对变量i进行+1操作
最后的结果可能就是加1而不是加2
因为他们俩读取变量i的时候都是同一个值,
如果用了原子操作,那么协程1在给变量加值的时候 协程2是不能读取完成的
等到协程1加1操作完成了,协程2才能读取到i的值
*/
func main() {
	i := 200
	log.Println(i % 200)
}
