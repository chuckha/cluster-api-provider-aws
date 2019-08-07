# Copyright 2019 The Kubernetes Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# Build the manager binary
FROM golang:1.12.7

# default the go proxy
ARG goproxy=https://proxy.golang.org

# run this with docker build --build_arg $(go env GOPROXY) to override the goproxy
ENV GOPROXY=$goproxy

# Copy in the go src
WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY pkg/    pkg/
COPY cmd/    cmd/

COPY third_party/forked/rerun-process-wrapper/start.sh .
COPY third_party/forked/rerun-process-wrapper/restart.sh .

# Build and run
RUN go install -v .
RUN mv /go/bin/manager /manager
ENTRYPOINT ["./start.sh", "/manager"]
