# Copyright 2019 The FoodUnit Authors. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# FoodUnit 3 MariaDB image (Development Version)
# Build command: docker image build -t dbdevimg -f docker/maria/dev.Dockerfile .
# Run command: docker container run --name dbdevctr -v ${pwd}/.mariadb:/var/lib/mysql --rm -p 3307:3306 -e MYSQL_ROOT_PASSWORD=root --network funet dbdevimg

# Currently, the MariaDB image doesn't support mounted host volumes with Docker Desktop.
# As a workaround, this image will base on MySQL.
FROM mysql:8.0

# PACKAGES defines the apt packages to get installed.
ENV PACKAGES openssh-server openssh-client

RUN apt-get update && apt-get upgrade && apt-get install -y ${PACKAGES}
RUN sed -i 's|^#PermitRootLogin.*|PermitRootLogin yes|g' /etc/ssh/sshd_config; \
    echo "root:root" | chpasswd