FROM sandrokeil/typescript:latest as base

FROM base as builder

# Some packages (e.g. @google-cloud/profiler) require additional
# deps for post-install scripts
RUN apk add --update --no-cache \
    python \
    make \
    g++

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install --only=production

FROM base

WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/node_modules ./node_modules

COPY . .

RUN ls

RUN npm run build

ENV PROTO_PATH=/../proto/

EXPOSE 3006 8086

ENTRYPOINT [ "node", "dist/server.js" ]
#ENTRYPOINT [ "/usr/local/bin/npm", "run", "start" ]