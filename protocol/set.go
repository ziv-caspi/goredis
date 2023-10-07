package protocol

type SetParams struct {
	Key        string
	Value      string
	TtlSeconds int
}
type SetReponse struct {
	NewValue string
}
