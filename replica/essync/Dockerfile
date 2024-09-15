FROM python:3.11-alpine
COPY . /src

RUN pip install -r /src/requirements.txt

RUN pip install opentelemetry-api && \
    pip install opentelemetry-sdk && \
    pip install opentelemetry-exporter-otlp-proto-http && \
    pip install opentelemetry-distro

RUN opentelemetry-bootstrap -a install

ENTRYPOINT [ "opentelemetry-instrument", "python", "/src/cdc.py"]