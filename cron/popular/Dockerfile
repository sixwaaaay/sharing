FROM python:3.11-alpine as builder
RUN apk update && apk add --no-cache build-base libpq-dev
WORKDIR /src

# 安装依赖
COPY requirements.txt .
RUN pip install --user --no-warn-script-location -r requirements.txt


FROM python:3.11-alpine
RUN apk update && apk add --no-cache libpq-dev
WORKDIR /src
COPY --from=builder /root/.local /root/.local
ENV PATH=/root/.local/bin:$PATH

COPY compute.py conf.json ./
CMD ["python", "compute.py"]
