/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
package mysql

import (
	"log"
	"time"

	"github.com/weaviate/weaviate/connectors"
)

type Mysql struct{}

type Object struct {
	Uuid         string // uuid, also used in Object's id
	Owner        string // uuid of the owner
	RefType      string // type, as defined
	CreateTimeMs int64  // creation time in ms
	Object       string // the JSON object, id will be collected from current uuid
	Deleted      bool   // if true, it does not exsist anymore
}

func (f *Mysql) Connect() error {
	return nil
}

func (f *Mysql) Add(owner string, refType string, object string) (string, error) {
	log.Fatalf("Connecting to Mysql DB - NOTE ONLY FOR DEMO PURPOSE")

	return "IM NOT USED", nil
}

func (f *Mysql) Get(Uuid string) (dbinit.Object, error) {

	task := dbinit.Object{
		Uuid:         "temp",
		Owner:        "temp",
		RefType:      "temp",
		CreateTimeMs: time.Now().UnixNano() / int64(time.Millisecond),
		Object:       "temp",
		Deleted:      false,
	}

	return task, nil
}
