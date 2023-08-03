package ThreadPool

type Task struct {
	handler func(v ...interface{})
	param   []interface{}
}
