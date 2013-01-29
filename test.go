// Copyright (c) 2012-2013 Jason McVetta.  This is Free Software, released under
// the terms of the GPL v3.  See http://www.gnu.org/copyleft/gpl.html for details.

/*******************************************************************************

The Neo4j Manual section numbers quoted in the test suite refer to the manual
for milestone release 1.8.  http://docs.neo4j.org/chunked/1.8/

*******************************************************************************/

/*

To run these tests, a Neo4j 1.8 instance must be running on localhost on the
default port.  

If the test suite complete successfully, all new nodes & relationships created
for testing will have been deleted in cleanup.  However if the test suite fails
to complete due to a panic, the cleanup code may not get called.  Therefore it
is not recommended to run this test suite on a db containing valuable data - run
it on a throwaway testing db instead! It's possible we could reduce this risk by
using defer() for cleanup.

*/

package neo4j

import (
	"github.com/jmcvetta/randutil"
	"log"
	"testing"
)

// Database connection used by all tests
var db *Database

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
	//
	var err error
	db, err = Connect("http://localhost:7474/db/data")
	if err != nil {
		log.Println(err)
	}
}

func rndStr(t *testing.T) string {
	name, err := randutil.AlphaString(12)
	if err != nil {
		t.Fatal(err)
	}
	return name
}