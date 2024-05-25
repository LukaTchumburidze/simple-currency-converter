package repository

var Aggregator AggregatorImpl

type AggregatorImpl struct {
	RequestRepository
	ResponseRepository
}
