package executors

type Command interface {
	execute() error
}
