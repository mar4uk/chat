FROM node:12 as build

WORKDIR /usr/src/app

COPY package.json ./
COPY package-lock.json ./

RUN npm install
COPY . ./
RUN npm run build

FROM nginx:1.19.1-alpine
COPY --from=build /usr/src/app/build /usr/share/nginx/html
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
