FROM sandrokeil/typescript:latest as base

FROM base as builder

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

FROM base

COPY --from=builder /usr/src/app/node_modules ./node_modules

COPY . .

ENV PROTO_PATH=../../../../proto/

EXPOSE 4000

ENTRYPOINT [ "/usr/local/bin/npm", "run", "start" ]