FROM python:3.11-alpine
COPY . /src
RUN pip install elasticsearch && \
   pip install mysql-replication && \
    pip install opentelemetry-api && \
    pip install opentelemetry-sdk && \
    pip install opentelemetry-exporter-otlp-proto-http
CMD ["python", "/src/cdc.py"]