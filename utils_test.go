/*
 * Copyright (c) 2023 Zander Schwid & Co. LLC.
 * SPDX-License-Identifier: BUSL-1.1
 */

package raftmod_test

import (
	"fmt"
	"github.com/sprintframework/raftmod"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLocalIP(t *testing.T) {

	ip, err := raftmod.PrivateIP()
	require.NoError(t, err)

	fmt.Printf("ip=%v\n", ip)


}
