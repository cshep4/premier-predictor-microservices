import {schema as schemaPublic} from './graphql';
import express from 'express';
import {ApolloServer} from 'apollo-server-express'
import * as http from "http";

const server = new ApolloServer({
    schema: schemaPublic,
    subscriptions: {
        onConnect: (connectionParams: any, webSocket) => {
            return {token: connectionParams.token}
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
});

const app = express();
server.applyMiddleware({app});

const httpServer = http.createServer(app);
server.installSubscriptionHandlers(httpServer);

httpServer.listen({port: 4000}, () => {
    console.log(`🚀 Server ready at http://localhost:4000${server.graphqlPath}`);
    console.log(`🚀 Subscriptions ready at ws://localhost:4000${server.subscriptionsPath}`);
});
