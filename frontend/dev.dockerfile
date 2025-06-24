FROM node:24

WORKDIR /cache_modules

COPY package.json package-lock.json ./
RUN npm install

WORKDIR /app
COPY . .

