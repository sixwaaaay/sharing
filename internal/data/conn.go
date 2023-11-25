/*
 * Copyright (c) 2023 sixwaaaay.
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

package data

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"

	"github.com/sixwaaaay/shauser/internal/config"
)

func NewData(config *config.Config) (*gorm.DB, error) {
	options := &gorm.Config{
		SkipDefaultTransaction: true,
		QueryFields:            true,
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

func replicas(replicas []string) dbresolver.Config {
	var dbs []gorm.Dialector
	for _, replica := range replicas {
		dbs = append(dbs, mysql.Open(replica))
	}
	return dbresolver.Config{Replicas: dbs}
}
