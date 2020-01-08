package rx

//Subject 可以手动发送数据
func Subject(input <-chan interface{}) Observable {
	return FromChan(input).Share()
}

//FromSlice 把Slice转成Observable
func FromSlice(slice []interface{}) Observable {
	return func(sink *Observer) error {
		for _, data := range slice {
			sink.Next(data)
			if sink.IsDisposed() {
				return nil
			}
		}
		return nil
	}
}

//Of 发送一系列值
func Of(array ...interface{}) Observable {
	return FromSlice(array)
}

//FromChan 把一个chan转换成事件流
func FromChan(source <-chan interface{}) Observable {
	return func(sink *Observer) error {
		for {
			select {
			case <-sink.Done():
				return nil
			case data, ok := <-source:
				if ok {
					sink.Next(data)
				} else {
					return nil
				}
			}
		}
	}
}

//Range 产生一段范围内的整数序列
func Range(start int, count uint) Observable {
	end := start + int(count)
	return func(sink *Observer) error {
		for i := start; i < end && !sink.IsDisposed(); i++ {
			sink.Next(i)
		}
		return nil
	}
}

//Never 永不回答
func Never() Observable {
	return func(sink *Observer) error {
		<-sink.Done()
		return nil
	}
}

//Empty 不会发送任何数据，直接完成
func Empty() Observable {
	return func(sink *Observer) error {
		return nil
	}
}

//Throw 直接抛出一个错误
func Throw(err error) Observable {
	return func(sink *Observer) error {
		return err
	}
}
