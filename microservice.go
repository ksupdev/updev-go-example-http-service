package main

type IMicroservice interface {
	Start() error
	Cleanup() error

	//HTTP Services
	GET(path string, h ServiceHandleFunc)
	POST(path string, h ServiceHandleFunc)
	PUT(path string, h ServiceHandleFunc)
	PATCH(path string, h ServiceHandleFunc)
	DELETE(path string, h ServiceHandleFunc)
}

type ServiceHandleFunc func(ctx IContext) error
