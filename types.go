package kuncore

type ConcurrencyCounter interface {
	Increase(key string)
	Decrease(key string)
	Count(key string) int64
}
