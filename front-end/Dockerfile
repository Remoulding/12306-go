FROM node:18 AS builder

# 使用淘宝源（解决 EAI_AGAIN 或 DNS 问题）
ENV YARN_REGISTRY=https://registry.npmmirror.com
RUN yarn config set registry ${YARN_REGISTRY}

WORKDIR /app

COPY package.json yarn.lock ./

# 安装依赖（先安装CA证书）
RUN yarn install --frozen-lockfile  --verbose

COPY . .

EXPOSE 12306

CMD ["yarn", "serve"]
