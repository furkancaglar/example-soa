package user

type Dependencies struct {
	// postgresql
	// bankid
}

type service struct {
	// http transport layer
	// gRPC transport layer
	// cli transport layer
}

// NewUserService This is where all service level gluing, dependency injection, happens.
func NewUserService(cnf Config) service {
	// Here we initialize and inject everything.
	// Note all the dependencies of the service which are required by
	// individual components of the service, such as injecting DB into repository ect.
	// We only want to export the API of the service and nothing else.
	return service{}
}
