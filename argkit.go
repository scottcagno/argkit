/*
 *	Argument Parser (os.Args wrapper)
 *	---------------------------------
 * 	Copyright (c) 2013, Scott Cagno, All rights reserved. 
 *	Use of this source code is governed by a BSD-style
 *	license that can be found in the LICENSE file.
 */

package argkit

import (
	"os"
	"strings"
)

// argument parser
type ArgStore struct {
	Prog string
	Argv []string
	Argc int
}

// return new argument parser instance
func NewArgStore() *ArgStore {
	return &ArgStore{os.Args[0], os.Args[1:], len(os.Args)}
}

// return arg at index provided
func (ap *ArgStore) ArgAtIndex(idx int) string {
	return ap.Argv[idx-1]
}

// return index of arg provided
func (ap *ArgStore) IndexOfArg(arg string) int {
	for k, v := range ap.Argv {
		if v == arg {
			return k + 1
		}
	}
	return 0
}

// return slice of args in the index range provided
func (ap *ArgStore) ArgsInRange(start, end int) []string {
	if end < ap.Argc {
		return ap.Argv[start-1 : end]
	}
	return nil
}

// return slice of args 
// slice is determined starting with the arg name provided
// and ending with the number of arguments to trail after the arg name
func (ap *ArgStore) ParseN(arg string, count int) (args []string) {
	idx := ap.IndexOfArg(arg)
	if idx+count < ap.Argc {
		for i := idx - 1; i < idx+count; i++ {
			args = append(args, ap.Argv[i])
		}
		return args
	}
	return nil
}

// return slice of argument values
// slice is first split by '=', it is chopped
// down further if it contains any more csv values
func (ap *ArgStore) ParseArg(arg string) []string {
	nargs := strings.Split(arg, "=")
	for k, v := range nargs {
		if strings.Contains(v, ",") {
			csv := strings.Split(v, ",")
			nargs = append(nargs[:k], csv...)
		}
	}
	return nargs
}
