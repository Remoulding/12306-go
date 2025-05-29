# front-end/Dockerfile

FROM node:18-alpine AS builder
ARG VUE_APP_API_BASE=http://nginx
ENV VUE_APP_API_BASE=$VUE_APP_API_BASE
WORKDIR /app
COPY package.json yarn.lock ./
RUN yarn install --frozen-lockfile
COPY . .
RUN yarn build