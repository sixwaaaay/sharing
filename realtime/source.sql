/*
 * Copyright (c) 2024 sixwaaaay.
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

/* content database */
CREATE
SOURCE mysql_content WITH (
  connector = 'mysql-cdc',
  hostname = '127.0.0.1', /* MySQL server host */
  port = '3306',
  username = 'user', /* user name with replication privileges */
  password = 'password',
  database.name = 'content', /* database name for change data capture */
  server.id = 5888
);

/* users database */
CREATE
SOURCE mysql_users WITH (
  connector = 'mysql-cdc',
  hostname = '127.0.0.1', /* MySQL server host */
  port = '3306',
  username = 'user', /* user name with replication privileges */
  password = 'password',
  database.name = 'users',
  server.id = 5889
);

/* comments database */
CREATE
SOURCE mysql_comments WITH (
  connector = 'mysql-cdc',
  hostname = '127.0.0.1', /* MySQL server host */
  port = '3306',
  username = 'user', /* user name with replication privileges */
  password = 'password',
  database.name = 'comments',
  server.id = 5890
);

/* vote database */
CREATE
SOURCE mysql_vote WITH (
  connector = 'mysql-cdc',
  hostname = '127.0.0.1', /* MySQL server host */
  port = '3306',
  username = 'user', /* user name with replication privileges */
  password = 'password',
  database.name = 'vote',
  server.id = 5891
);