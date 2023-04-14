#!/bin/bash

export OTEL_TRACES_EXPORTER=otlp
export OTEL_METRICS_EXPORTER=otlp
export OTEL_EXPORTER_OTLP_ENDPOINT=localhost:5555
export OTEL_RESOURCE_ATTRIBUTES=service.name=todo,service.version=1.0
export SERVICE_NAME=todo

#docker run -v $(pwd)/configs/collector-config-test.yaml:/etc/otelcol/config.yaml -p 5555:5555 otel/opentelemetry-collector:latest

go run main.go