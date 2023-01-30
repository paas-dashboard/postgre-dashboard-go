#
# Licensed to the Apache Software Foundation (ASF) under one
# or more contributor license agreements.  See the NOTICE file
# distributed with this work for additional information
# regarding copyright ownership.  The ASF licenses this file
# to you under the Apache License, Version 2.0 (the
# "License"); you may not use this file except in compliance
# with the License.  You may obtain a copy of the License at
#
#   http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing,
# software distributed under the License is distributed on an
# "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
# KIND, either express or implied.  See the License for the
# specific language governing permissions and limitations
# under the License.
#

FROM shoothzj/compile:go AS build
COPY . /opt/compile
WORKDIR /opt/compile
RUN go build -o postgre-dashboard .

FROM shoothzj/base

WORKDIR /opt/postgre-dashboard

COPY --from=build /opt/compile/postgre-dashboard /opt/postgre-dashboard/postgre-dashboard

RUN wget -q https://github.com/paas-dashboard/postgre-dashboard-portal/releases/download/latest/postgre-dashboard-portal.tar.gz && \
    tar -xzf postgre-dashboard-portal.tar.gz && \
    rm -rf postgre-dashboard-portal.tar.gz

EXPOSE 10014

CMD ["/usr/bin/dumb-init", "/opt/postgre-dashboard/postgre-dashboard"]
