package managers

type IPubCallback interface {
	Process()
	OnTimeout()
}
