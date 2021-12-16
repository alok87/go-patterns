package behavioural

import "fmt"

/*

Summary:
Strategy pattern is useful when you need to switch between algorithmic logic
in runtime. Clients can switch to different logics to do the same thing at
runtime.

Example:
Loadbalancer can load balance incoming requests based on many strategies
like round robin, weighted round robin, least connection. Stategy pattern
is used to switch the loadbalancing logic to any one of them at runtime.
Interface: Loadbalancing Strategy
Implementations: RoundRobin, Weighted RoundRobin and LeastConn

Benefit:
Enables switching algorithm at runtime.
*/

type Strategy int

const (
	RoundRobin Strategy = iota
	WeightedRoundRobin
	LeastConnection
)

// LoadbalancingStrategy defines what makes a loalbalacning strategy
type LoadbalancingStrategy interface {
	// Next finds the next server to send the request to
	Next()
}

// roundRobin implements LoadbalancingStrategy
// this the default strategy
type roundRobin struct{}

func (r *roundRobin) Next() {
	fmt.Println("roundRobin")
}

// weightedRoundRobin implements LoadbalancingStrategy
type weightedRoundRobin struct{}

func (r *weightedRoundRobin) Next() {
	fmt.Println("weightedRoundRobin")
}

// leastConn implements LoadbalancingStrategy
type leastConn struct{}

func (r *leastConn) Next() {
	fmt.Println("leastConn")
}

// Loadbalancer reverse proxy the incoming request by calling Balance()
// which loadbalancer to the next server by finding the next server
// using the set set strategy. There is
type Loadbalancer struct {
	strategies      map[Strategy]LoadbalancingStrategy
	currentStrategy LoadbalancingStrategy
}

func NewLoadbalancer() *Loadbalancer {
	return &Loadbalancer{
		strategies: map[Strategy]LoadbalancingStrategy{
			RoundRobin: &roundRobin{},
		},
		currentStrategy: &roundRobin{},
	}
}

// SetStrategy gives the client ability to change the logic
// of loadbalancing at runtime, this is the benefit of Strategy pattern
func (l *Loadbalancer) SetStrategy(s Strategy) error {
	v, ok := l.strategies[s]
	if ok {
		l.currentStrategy = v
		return nil
	}

	var newStrategy LoadbalancingStrategy
	switch s {
	case WeightedRoundRobin:
		newStrategy = &weightedRoundRobin{}
		l.strategies[WeightedRoundRobin] = newStrategy
	case LeastConnection:
		newStrategy = &leastConn{}
		l.strategies[LeastConnection] = newStrategy
	default:
		return fmt.Errorf("unsupported strategy")
	}

	l.currentStrategy = newStrategy

	return nil
}

func (l *Loadbalancer) Balance() {
	l.currentStrategy.Next()
	fmt.Println("loadbalanced")
}
