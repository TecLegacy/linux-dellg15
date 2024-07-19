FROM node:21

RUN mkdir -p /home/app

WORKDIR /home/app

COPY ./package.json ./pnpm-lock.yaml ./

RUN npm install -g pnpm && pnpm install

COPY . .

CMD ["pnpm", "start"]



