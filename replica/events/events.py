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
import os
from contextlib import closing
from datetime import datetime

import psycopg2
from pymysqlreplication import BinLogStreamReader
from pymysqlreplication.event import GtidEvent
from pymysqlreplication.row_event import DeleteRowsEvent, UpdateRowsEvent, WriteRowsEvent


def save_position(gtid: str):
    uid, x = gtid.split(":")
    x = int(x)
    try:
        with open('gtid.json', 'r') as f:
            data = json.load(f)
            a, b = map(int, data['gtid'].split(":")[1].split("-"))
            if x < a:
                a = x
            if x > b:
                b = x
            data['gtid'] = f"{uid}:{a}-{b}"
    except FileNotFoundError:
        data = {
            'gtid': f"{uid}:{1}-{x}"
        }
    with open('gtid.json', 'w') as f:
        json.dump(data, f)
    logging.info(f'Saved gtid: {data}')


def load_position():
    try:
        with open('gtid.json', 'r') as f:
            data = json.load(f)
            return data['gtid']
    except FileNotFoundError:
        return None


def load_conf():
    conf_path = os.environ.get("CONF_PATH", "config.json")
    with open(conf_path) as f:
        conf = json.load(f)
    return conf


def events(stream: BinLogStreamReader):
    last_gtid: str | None = None
    for binlog_event in stream:
        if isinstance(binlog_event, GtidEvent):
            if last_gtid is not None:
                save_position(last_gtid)
            last_gtid = binlog_event.gtid
            continue
        for row in binlog_event.rows:
            if isinstance(binlog_event, DeleteRowsEvent):
                vals = row["values"]
                yield vals["id"], 1, vals.get("created_at", datetime.now())
            elif isinstance(binlog_event, UpdateRowsEvent):
                bf_view_count = row["before_values"].get("view_count", None)
                if bf_view_count is None:
                    continue
                af_view_count = row["after_values"].get("view_count", None)
                if bf_view_count != af_view_count:
                    yield row["after_values"]["id"], 3, row["after_values"].get("created_at", datetime.now())
            elif isinstance(binlog_event, WriteRowsEvent):
                vals = row["values"]
                yield vals["id"], 2, vals.get("created_at", datetime.now())


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    try:
        conf = load_conf()
        stream_reader = BinLogStreamReader(
            only_events=[GtidEvent, DeleteRowsEvent, WriteRowsEvent, UpdateRowsEvent],
            auto_position=load_position(),
            **conf["replica"],
        )
        with closing(psycopg2.connect(**conf["postgres"])) as conn, conn.cursor() as cur, closing(
                stream_reader) as streamer:
            for val in events(streamer):
                # Insert into the video_events table
                cur.execute("INSERT INTO video_events (video_id, event_type, event_time) VALUES (%s, %s, %s)", val)
                # Commit the transaction
                conn.commit()
    except KeyboardInterrupt:
        logging.info('KeyboardInterrupt exit')
