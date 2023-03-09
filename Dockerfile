# Copyright 2023 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

FROM golang:1.19 as builder

ADD ./apiclient /go/src/integrationcli/apiclient
ADD ./client /go/src/integrationcli/client
ADD ./cmd /go/src/integrationcli/cmd
ADD ./cloudkms /go/src/integrationcli/cloudkms
ADD ./secmgr /go/src/integrationcli/secmgr

COPY main.go /go/src/integrationcli/main.go
COPY go.mod go.sum /go/src/integrationcli/
WORKDIR /go/src/integrationcli

ENV GO111MODULE=on
RUN go mod tidy
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -trimpath -a -ldflags='-s -w -extldflags "-static"' -o /go/bin/integrationcli /go/src/integrationcli/main.go

FROM gcr.io/distroless/static-debian11
COPY --from=builder /go/bin/integrationcli /
CMD ["/integrationcli"]
