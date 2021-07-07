package kit

type PreLockEvent struct {
	Receipt       []byte
	Proof         []byte
	AppchainIndex uint64
	BlockNumber   uint64
}

type PreLock struct {
	preLockEvent *PreLockEvent
	priority     int
}

func PreLockConstructor(preLockEvent *PreLockEvent, priority int) PreLock {
	return PreLock{preLockEvent, priority}
}

func (p *PreLock) GetPreLockEvent() *PreLockEvent {
	return p.preLockEvent
}

func (p *PreLock) GetPriority() int {
	return p.priority
}

type PreLockQueue []PreLock

func (t *PreLockQueue) Len() int {
	return len(*t) //
}

func (t *PreLockQueue) Less(i, j int) bool {
	return (*t)[i].priority < (*t)[j].priority
}

func (t *PreLockQueue) Swap(i, j int) {
	(*t)[i], (*t)[j] = (*t)[j], (*t)[i]
}

func (t *PreLockQueue) Push(x interface{}) {
	*t = append(*t, x.(PreLock))
}

func (t *PreLockQueue) Pop() interface{} {
	n := len(*t)
	x := (*t)[n-1]
	*t = (*t)[:n-1]
	return x
}
