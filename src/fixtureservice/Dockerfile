## Docker Parent Image with Node and Typescript
#FROM sandrokeil/typescript:latest
#
## Create Directory for the Container
#WORKDIR /app
#
## Copy the files we need to our new Directory
#ADD . /app
#
## Expose the port outside of the container
#EXPOSE 3006 8086
#
##RUN rm -fr node_modules/
#
## Grab dependencies and transpile src directory to dist
#RUN npm install && tsc
#
## Start the server
#ENTRYPOINT ["node", "dist/server.js"]
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

RUN tsc

FROM base

#RUN GRPC_HEALTH_PROBE_VERSION=v0.2.0 && \
#    wget -qO/bin/grpc_health_probe https://github.com/grpc-ecosystem/grpc-health-probe/releases/download/${GRPC_HEALTH_PROBE_VERSION}/grpc_health_probe-linux-amd64 && \
#    chmod +x /bin/grpc_health_probe
#
#WORKDIR /usr/src/app

COPY --from=builder /usr/src/app/node_modules ./node_modules

COPY . .

EXPOSE 3006 8086

ENTRYPOINT [ "node", "dist/server.js" ]