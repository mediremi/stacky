package stacky

type instruction uint8
type instructions []instruction

const (
	instPrint instruction = iota
	instAdd
	instMult
	instLiteral
	instDBGSTK
)
