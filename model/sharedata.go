package model

// DataFlow holds the shared configuration data
type DataFlow struct {
	ConfigFile string
	DelayTime  int
	NumThreads int
	OutputDir  string
	URL        string
	Log        bool
	Method     string // Linear, Tree, etc.
}

// SharedData is the global instance of DataFlow
var SharedData DataFlow

type TreeQueue struct {
	Status    bool
	Cmd       string
	LeftChid  *TreeQueue
	RightChid *TreeQueue
}

type LinearQueue struct {
	Status bool
	Cmd    string
	Chid   *LinearQueue
}

type Queue struct {
	LinerProcess []LinearQueue
	TreeProcess  []TreeQueue
}
