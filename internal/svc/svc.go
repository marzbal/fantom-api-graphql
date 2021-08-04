// Package svc implements blockchain data processing services.
package svc

// Svc defines the interface required for a service
// to be manageable by the orchestrator.
type Svc interface {
	// name returns the name of the service
	name() string

	// init initializes the service
	init()

	// run executes the service
	run()

	// close signals the service to terminate
	close()
}

// service implements general base for services implementing svc interface.
type service struct {
	mgr     *ServiceManager
	sigStop chan bool
}

// init prepares the account dispatcher to perform its function.
func (s *service) init() {
	s.sigStop = make(chan bool, 1)
}

// close terminates the block dispatcher.
func (s *service) close() {
	if s.sigStop != nil {
		s.sigStop <- true
	}
}
