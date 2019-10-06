package rx

//Count 计算元素总数然后返回
func (ob Observable) Count() Observable {
	return func(sink *Observer) error {
		count := 0
		defer func() {
			sink.Next(count)
			sink.Stop()
		}()
		return ob.subscribe(NextFunc(func(event *Event) {
			count++
		}), sink.stop)
	}
}

//Max 完成时返回最大值
func (ob Observable) Max() Observable {
	return func(sink *Observer) error {
		var max int
		defer func() {
			sink.Next(max)
			sink.Stop()
		}()
		return ob.subscribe(NextFunc(func(event *Event) {
			if data := event.Data.(int); data > max {
				max = data
			}
		}), sink.stop)
	}
}

//Min 完成时返回最小值
func (ob Observable) Min() Observable {
	return func(sink *Observer) error {
		var min int
		defer func() {
			sink.Next(min)
			sink.Stop()
		}()
		return ob.subscribe(NextFunc(func(event *Event) {
			if data := event.Data.(int); data < min {
				min = data
			}
		}), sink.stop)
	}
}

//Reduce 完成时返回累加值
func (ob Observable) Reduce(f func(interface{}, interface{}) interface{}) Observable {
	return func(sink *Observer) error {
		var aac interface{}
		defer func() {
			sink.Next(aac)
			sink.Stop()
		}()
		aacNext := func(event *Event) {
			aac = f(aac, event.Data)
		}
		return ob.subscribe(NextFunc(func(event *Event) {
			aac = event.Data
			event.Target.next = NextFunc(aacNext)
		}), sink.stop)
	}
}
