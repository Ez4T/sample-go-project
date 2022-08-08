FROM node:16

WORKDIR /app

COPY . .

RUN yarn install

RUN yarn add express

RUN yarn build

EXPOSE 3080
CMD ["node", "server.js"]
