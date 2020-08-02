import addPost from './addPost';
import listPosts, {Params} from './listPosts';
import {PubSub} from 'apollo-server';
import grpc from '../services/grpcClient';
import {LiveMatch} from "./liveMatch/liveMatch";

const liveMatch = new LiveMatch(grpc.liveMatch(), new PubSub());

const resolvers: any = {
    Mutation: {
        addPost: (root: any, params: any) => addPost(root, params)
    },
    Query: {
        listPosts: (root: any, params: Params) => listPosts(root, params)
    },
    Subscription: {
        upcomingMatches: {subscribe: (_obj: any, args: any, context: any) => liveMatch.getUpcomingMatches(_obj, args, context)},
    },
};

export default resolvers;
