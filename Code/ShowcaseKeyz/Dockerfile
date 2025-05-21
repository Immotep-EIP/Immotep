FROM node:22-alpine AS builder

WORKDIR /app

COPY package*.json .

RUN npm install

ARG VITE_API_URL

ENV VITE_API_URL=$VITE_API_URL

COPY . .

RUN npm run build

FROM node:22-alpine AS runner

WORKDIR /app

COPY --from=builder /app/dist .

ENV NODE_ENV=production

EXPOSE 3003

RUN npm install -g serve

CMD ["serve", "-s", ".", "-l", "3003"]
