import {schema as schemaPublic} from './graphql';
import express from 'express';
import {ApolloServer} from 'apollo-server-express'
import * as http from "http";
import WebSocket from "ws";

const server = new ApolloServer({
    schema: schemaPublic,
    subscriptions: {
        onConnect: (connectionParams: any, webSocket) => {
            if (connectionParams.authorization) {
                return {
                    connectionParams: connectionParams,
                    webSocket: webSocket,
                    token: connectionParams.authorization,
                }
            }
            return {
                connectionParams: connectionParams,
                webSocket: webSocket,
                token: connectionParams.token,
            }
        },
        onDisconnect: (webSocket, context) => {
            console.log("onDisconnect");
        },
    },
    context: async ({req, connection}) => {
        if (connection) {
            return connection.context;
        } else {
            const token = req.headers.authorization || "";
            return {token};
        }
    },
    playground: {
        subscriptionEndpoint: 'wss://35.246.124.255/gateway/graphql',
    },
});

const app = express();
server.applyMiddleware({app});

const httpServer = http.createServer(app);
server.installSubscriptionHandlers(httpServer);

httpServer.listen({port: 4000}, () => {
    console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`);
    console.log(`🚀 Subscriptions ready at ws://localhost:4000${server.subscriptionsPath}`);
});
