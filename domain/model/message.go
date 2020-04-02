package model

type Message struct {
	Msg  string `json:"msg"`
	Done bool   `json:"done"`
}

func (messsage *Message) CreateUser(msg string, done bool) (Message, error) {
	return Message{
		Msg:  msg,
		Done: done,
	}, nil
}
