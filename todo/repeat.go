package todo

const (
	REPEAT_MONTHLY = "monthly"
	REPEAT_WEEKLY  = "weekly"
	REPEAT_DAILY   = "daily"
)

var RepeatStatuses = []string{REPEAT_MONTHLY, REPEAT_WEEKLY, REPEAT_DAILY}

type Repeat struct {
	Types string
	Data  []string
}
