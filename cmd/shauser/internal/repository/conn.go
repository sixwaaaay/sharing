/*
 * Copyright (c) 2023-2024 sixwaaaay.
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *     http://www.apache.org/licenses/LICENSE-2.
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package repository

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/sixwaaaay/shauser/internal/config"
)

// NewDB creates a new database connection using the provided configuration.
// It returns a pointer to a gorm.DB object and an error if any occurred during the connection process.
//
// The function first sets up the configuration options for the database connection.
// It then opens a new MySQL database connection using the DSN provided in the configuration.
// If there are any errors during this process, it returns the error.
//
// If the configuration specifies any MySQL replicas, the function sets up these replicas using the replicas function.
// It then uses the dbresolver plugin to register these replicas with the main database connection.
// If there are any errors during this process, it returns the error.
//
// Finally, it returns the database connection and nil for the error.
func NewDB(config *config.Config) (*gorm.DB, error) {
	options := &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
		PrepareStmt: true,
	}
	open := mysql.Open(config.MySQL.DSN)
	db, err := gorm.Open(open, options)
	if err != nil {
		return nil, err
	}
	if len(config.MySQL.Replicas) > 0 {
		c := replicas(config.MySQL.Replicas)
		if err := db.Use(dbresolver.Register(c)); err != nil {
			return nil, err
		}
	}
	return db, nil
}

// replicas is a helper function that sets up the MySQL replicas for the main database connection.
// It takes a slice of strings, where each string is the DSN for a replica.
// It returns a dbresolver.Config object that contains the dialector for each replica.
//
// The function iterates over the provided replicas, opens a new MySQL connection for each one,
// and appends the dialector for the connection to a slice.
// It then returns a dbresolver.Config object that contains the slice of dialectors.
func replicas(replicas []string) dbresolver.Config {
	var dbs []gorm.Dialector
	for _, replica := range replicas {
		dbs = append(dbs, mysql.Open(replica))
	}
	return dbresolver.Config{Replicas: dbs}
}
