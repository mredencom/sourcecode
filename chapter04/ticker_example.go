package main
// Ticker 和 Timer 类似，区别是：Ticker 中的 runtimeTimer 字段的 period 字段会赋值为 NewTicker(d Duration) 中的 d，表示每间隔 d 纳秒，定时器就会触发一次。
// 除非程序终止前定时器一直需要触发，否则，不需要时应该调用 Ticker.Stop 来释放相关资源。
// 如果程序终止前需要定时器一直触发，可以使用更简单方便的 time.Tick 函数，因为 Ticker 实例隐藏起来了，因此，该函数启动的定时器无法停止。
func main() {

}
