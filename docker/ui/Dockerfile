#
# FoodUnit 3 UI image (Dev version)
#

FROM node:12.2.0-alpine

RUN mkdir -p /app
WORKDIR /app

# PATH overrides the default PATH value.
ENV PATH "/app/node_modules/.bin:$PATH"

# Copy package.json into the app directory.
COPY ui/package.json /app/package.json
# Install all UI dependencies.
RUN npm install
RUN npm install -g react-scripts@3.0.1

# Start the UI.
CMD ["npm", "start"]