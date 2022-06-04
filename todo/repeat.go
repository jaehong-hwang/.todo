package todo

const (
	REPEAT_MONTHLY = "monthly"
	REPEAT_WEEKLY  = "weekly"
	REPEAT_DAILY   = "daily"
)

type Repeat struct {
	Types string
	Data  []string
}
