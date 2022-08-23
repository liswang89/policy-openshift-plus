#!/bin/bash
# Copyright Contributors to the Open Cluster Management project

set -e

echo "Running test..."
CGO_ENABLED=0 ./bin/ginkgo -v --fail-fast -focus "policy" --junit-report=integration.xml --output-dir=test-output tests/integration
