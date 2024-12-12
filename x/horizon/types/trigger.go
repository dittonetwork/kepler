package types

type TriggerType string

const (
	EveryBlock TriggerType = "EveryBlock"
	TimeBased  TriggerType = "TimeBased"
)

type Trigger struct {
	Type       TriggerType
	EveryBlock struct{}
	TimeBased  [194]bool
}
