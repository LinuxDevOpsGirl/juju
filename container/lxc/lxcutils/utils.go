// Copyright 2013 Canonical Ltd.
// Licensed under the AGPLv3, see LICENCE file for details.

package lxcutils

// RunningInsideLXC reports whether or not we are running inside an
// LXC container.
func RunningInsideLXC() (bool, error) {
	return runningInsideLXC()
}
