FROM python:3.11-alpine
COPY . /src
RUN pip install elasticsearch
RUN pip install mysql-replication
CMD ["python", "/src/cdc.py"]