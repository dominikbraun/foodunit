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

# FoodUnit 3 UI image (Development Version)

# Start build stage.
FROM node:12.2.0-alpine

RUN mkdir -p /app
WORKDIR /app

# PATH overrides the default PATH value.
ENV PATH "/app/node_modules/.bin:$PATH"
# PORT specifies the port the server listens on.
ENV PORT 80

# Copy package.json into the app directory.
COPY ui/package.json /app/package.json
# Install all UI dependencies.
RUN npm install
RUN npm install -g react-scripts@3.0.1

COPY ui /app
RUN npm build

# Start final stage.
FROM nginx:1.16.0-alpine

COPY --from=0 /app/build /usr/share/nginx/html

EXPOSE ${PORT}

# Serve the UI.
CMD ["nginx", "-g", "daemon off;"]