package executors

type Executor interface {
	RecoverDump() map[string][][]byte
}
