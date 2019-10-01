package rx

type (
	Event struct {
		Data    interface{}
		Control *Control
	}
	Observer interface {
		OnNext(*Event)
	}
	Observable   func(*Control) error
	Operator     func(Observable) Observable
	ControlSet   map[*Control]interface{}
	ObserverFunc func(*Event)
	ObserverChan chan *Event
)

func (observer ObserverFunc) OnNext(event *Event) {
	observer(event)
}
func (observer ObserverChan) OnNext(event *Event) {
	observer <- event
}

var (
	EmptyObserver = ObserverFunc(func(event *Event) {})
)

func (set ControlSet) add(ctrl *Control) {
	set[ctrl] = nil
}
func (set ControlSet) remove(ctrl *Control) {
	delete(set, ctrl)
}
func (set ControlSet) isEmpty() bool {
	return len(set) == 0
}
func (ob Observable) Pipe(cbs ...Operator) Observable {
	for _, cb := range cbs {
		ob = cb(ob)
	}
	return ob
}
