#  Copyright (c) 2024 sixwaaaay.
#  Licensed under the Apache License, Version 2.0 (the "License");
#  you may not use this file except in compliance with the License.
#  You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
#  Unless required by applicable law or agreed to in writing, software
#  distributed under the License is distributed on an "AS IS" BASIS,
#  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
#  See the License for the specific language governing permissions and
#  limitations under the License.

import json
import logging

import mysql.connector
import psycopg2

POPULAR = """
SELECT row_number() OVER () AS order_num,
       id,
       score
FROM (SELECT video_id AS id,
             sum(CASE
                     WHEN event_type = 1 THEN 1
                     WHEN event_type = 2 THEN -1
                     WHEN event_type = 3 THEN 0.01
                 END) AS score
      FROM VIDEO_EVENTS
      WHERE DATE_TRUNC('day', event_time) = DATE_TRUNC('day', CURRENT_DATE)
      GROUP BY video_id
      ORDER BY score DESC) AS score_table
            """

OPERATION = """
INSERT INTO popular_videos (order_num, id, score)
VALUES (%s, %s, %s)
ON DUPLICATE KEY UPDATE score = VALUES(score), id = VALUES(id)
"""


def load_conf():
    with open("conf.json", "r") as f:
        # json.dump(conf, f, indent=4)
        conf = json.load(f)
    return conf


def main():
    # Connect to PostgreSQL database and MySQL database
    conf = load_conf()
    with psycopg2.connect(**conf["postgres"]) as pg_conn, mysql.connector.connect(**conf["mysql"]) as mysql_conn:
        with pg_conn.cursor() as pg_cur, mysql_conn.cursor() as mysql_cur:
            pg_cur.execute(POPULAR)
            while True:
                rows = pg_cur.fetchmany(1000)
                if not rows:
                    break

                mysql_cur.executemany(OPERATION, rows)

                mysql_conn.commit()


if __name__ == '__main__':
    logging.basicConfig(level=logging.INFO)
    try:
        main()
    except KeyboardInterrupt:
        logging.info("KeyboardInterrupt exit")
    except Exception as e:
        logging.exception(e)
