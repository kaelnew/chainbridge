// Copyright 2020 Stafi Protocol
// SPDX-License-Identifier: LGPL-3.0-only

package ethtest

import (
	"testing"

	utils "github.com/kaelnew/chainbridge/shared/ethereum"
	"github.com/kaelnew/chainbridge-utils/msg"
	"github.com/ethereum/go-ethereum/common"
)

func RegisterResource(t *testing.T, client *utils.Client, bridge, handler common.Address, rId msg.ResourceId, addr common.Address) {
	err := utils.RegisterResource(client, bridge, handler, rId, addr)
	if err != nil {
		t.Fatal(err)
	}
}
