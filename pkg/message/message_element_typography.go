package message

import "fmt"

type MessageElmentBr struct {
	*noAliasMessageElement
}

func (e *MessageElmentBr) Tag() string {
	return "br"
}

func (e *MessageElmentBr) Stringify() string {
	return fmt.Sprintln()
}

type MessageElmentP struct {
	*noAliasMessageElement
	*childrenMessageElement
}

func (e *MessageElmentP) Tag() string {
	return "p"
}

func (e *MessageElmentP) Stringify() string {
	return e.stringifyByTag(e.Tag())
}

type MessageElementMessage struct {
	*noAliasMessageElement
	*childrenMessageElement
}

func (e *MessageElementMessage) Tag() string {
	return "message"
}

func (e *MessageElementMessage) Stringify() string {
	return e.stringifyByTag(e.Tag())
}
