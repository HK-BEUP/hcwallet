// Copyright (c) 2013-2015 The btcsuite developers
// Copyright (c) 2016-2017 The Decred developers
// Copyright (c) 2018-2020 The Hc developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package netparams

import "github.com/HcashOrg/hcd/chaincfg"

// Params is used to group parameters for various networks such as the main
// network and test networks.
type Params struct {
	*chaincfg.Params
	JSONRPCClientPort string
	JSONRPCServerPort string
	GRPCServerPort    string
}

// MainNetParams contains parameters specific to running vexonwallet and
// vexond on the main network (wire.MainNet).
var MainNetParams = Params{
	Params:            &chaincfg.MainNetParams,
	JSONRPCClientPort: "18556",
	JSONRPCServerPort: "18557",
	GRPCServerPort:    "18558",
}

// TestNet2Params contains parameters specific to running vexonwallet and
// vexond on the test network (version 2) (wire.TestNet2).
var TestNet2Params = Params{
	Params:            &chaincfg.TestNet2Params,
	JSONRPCClientPort: "28556",
	JSONRPCServerPort: "28557",
	GRPCServerPort:    "28558",
}

// SimNetParams contains parameters specific to the simulation test network
// (wire.SimNet).
var SimNetParams = Params{
	Params:            &chaincfg.SimNetParams,
	JSONRPCClientPort: "13009",
	JSONRPCServerPort: "13010",
	GRPCServerPort:    "13011",
}
