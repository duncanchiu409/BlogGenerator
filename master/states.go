package master

type State string

const (
	unstarted State = "unstarted"
	content   State = "generating-content"
	picture   State = "generating-picture"
	complete  State = "complete"
)
