package actiontype

type Action int

const QUERY Action = iota

type Type int

const City Type = iota

type EventInfo struct {
	Action Action
	Type   Type
	Body   interface{}
}
