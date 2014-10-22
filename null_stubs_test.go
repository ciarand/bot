package bot

type nullResponder struct {
}

func (n nullResponder) Respond(m Message) {
}

type nullContext struct{}

func (n nullContext) Recv() Message {
	return nullMessage{}
}

func (n nullContext) Send(m Message) {
}

type nullMessage struct{}

func (n nullMessage) String() string {
	return ""
}

func (n nullMessage) Body() string {
	return ""
}

func (n nullMessage) Source() Source {
	return nullSource{}
}

func (n nullMessage) Context() Context {
	return nullContext{}
}

type nullSource struct{}

func (n nullSource) String() string {
	return ""
}

func (n nullSource) ID() string {
	return ""
}
