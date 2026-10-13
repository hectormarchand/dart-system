FROM node:24-alpine AS build

WORKDIR /build

COPY dart-frontend/package*.json .

RUN npm ci --prefer-offline --only=production --silent

COPY dart-frontend/ ./

RUN npm run build

FROM nginx:1.29.0-alpine AS app

COPY --from=build /build/dist /usr/share/nginx/html

COPY docker/rasp/frontend/nginx/nginx.conf /etc/nginx/nginx.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
