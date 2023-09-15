package PromethoniXTrie

type ActionType int

const (
	Insert ActionType = iota
	Update
	Delete
)

// ActionLogEntry in log, [key, old value, new value]
type ActionLogEntry struct {
	Action  ActionType
	Key     Hash
	OldData Data
	NewData Data
}

type ActionLog struct {
	ActionLogEntries   []*ActionLogEntry
	IsActionLogEnabled bool
}
