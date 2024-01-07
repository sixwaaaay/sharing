#
# Copyright (c) 2024 sixwaaaay.
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#     http://www.apache.org/licenses/LICENSE-2.
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

import json
import logging
import os
from contextlib import closing

from elasticsearch import Elasticsearch
from elasticsearch.helpers import bulk
from pymysqlreplication import BinLogStreamReader
from pymysqlreplication.event import RotateEvent
from pymysqlreplication.row_event import DeleteRowsEvent, UpdateRowsEvent, WriteRowsEvent


def load_conf():
    conf_path = os.environ.get("CONF_PATH", "config.json")
    with open(conf_path) as f:
        conf = json.load(f)
    return conf


def save_position(next_binlog, position):
    data = {
        'next_binlog': next_binlog,
        'position': position
    }
    with open('position.json', 'w') as f:
        json.dump(data, f)
    logging.info(f'Saved position: {data}')


def load_position():
    try:
        with open('position.json', 'r') as f:
            data = json.load(f)
            return data['next_binlog'], data['position']
    except FileNotFoundError:
        return None, None


def sync(stream: BinLogStreamReader, client: Elasticsearch):
    for binlog_event in stream:
        if isinstance(binlog_event, RotateEvent):
            save_position(binlog_event.next_binlog, binlog_event.position)
            continue
        table = binlog_event.table
        bulk_data = []
        for row in binlog_event.rows:
            if isinstance(binlog_event, DeleteRowsEvent):
                vals = row["values"]
                action = {
                    "_op_type": 'delete',
                    "_index": table,
                    "_id": identity(vals)
                }
            elif isinstance(binlog_event, UpdateRowsEvent):
                vals = row["after_values"]
                action = {
                    "_op_type": 'index',
                    "_index": table,
                    "_id": identity(vals),
                    "_source": vals
                }
            elif isinstance(binlog_event, WriteRowsEvent):
                vals = row["values"]
                action = {
                    "_op_type": 'index',
                    "_index": table,
                    "_id": identity(vals),
                    "_source": vals
                }
            else:
                action = None
            if action:
                bulk_data.append(action)
        if bulk_data:
            bulk(client, bulk_data)


def identity(vals):
    return str(vals["id"])


def main():
    conf = load_conf()
    log_file, log_pos = load_position()
    stream_reader = BinLogStreamReader(
        connection_settings=conf["mysql"],
        server_id=conf["server_id"],
        only_events=[DeleteRowsEvent, WriteRowsEvent, UpdateRowsEvent, RotateEvent],
        only_schemas=conf["databases"] or None,  # only capture events for this table
        only_tables=conf["tables"] or None,  # only capture events for this schema
        log_pos=log_pos or None,
        log_file=log_file or None,
        blocking=True,
    )
    client = Elasticsearch(**conf["elasticsearch"])
    try:
        with closing(stream_reader) as stream_reader, closing(client) as client:
            sync(stream_reader, client)
    except KeyboardInterrupt:
        stream_reader.close()
        client.close()
        logging.info("KeyboardInterrupt exit")


if __name__ == "__main__":
    logging.basicConfig(level=logging.INFO)
    main()
