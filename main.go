package main

import (
	"github.com/gochargepoint"
)

// NewClient returns a new chargepoint V5 client
funct NewClient(user user, password password) v5.NewChargepointV5Client {
	return v5.NewChargepointV5Client(user, password)
}