import {PubSub} from 'apollo-server';
import grpc from '../services/grpcClient';
import {LiveMatch} from "./liveMatch/liveMatch";
import {Prediction} from "./prediction/prediction";
import {SubscriptionContext, TokenContext} from "../model/context/context";
import {GetUserPredictionRequest, PredictionRequest, PredictionSummaryRequest} from "./prediction/client";
import {GetMatchRequest} from "./liveMatch/client";
import {Fixture} from "./fixture/fixture";
import {League} from "./league/league";
import {Home} from "./home/home";
import {User} from "./user/user";

const liveMatch = new LiveMatch(grpc.liveMatch(), new PubSub());
const prediction = new Prediction(grpc.prediction());
const fixture = new Fixture(grpc.fixture());
const league = new League(grpc.league());
const user = new User(grpc.user());
const home = new Home(league, prediction, fixture);

const resolvers: any = {
    Query: {
        matchSummary: (root: any, params: any, context: TokenContext) => {
            return {
                userId: params.request.userId,
                matchId: params.request.matchId,
                id: params.request.matchId,
            }
        },
        predictorData: (root: any, params: any, context: TokenContext) => {
            return {userId: params.request.userId}
        },
        homeFeed: (root: any, params: any, context: TokenContext) => home.feed(context, params.request),
    },
    Mutation: {
        // addPost: (root: any, params: any) => addPost(root, params)
    },
    Subscription: {
        upcomingMatches: {subscribe: (_obj: any, args: any, context: SubscriptionContext) => liveMatch.getUpcomingMatches(_obj, args, context)},
        liveMatchSummary: {subscribe: (_obj: any, args: any, context: SubscriptionContext) => liveMatch.getLiveMatchSummary(_obj, args, context)},
        todaysLiveMatches: {subscribe: (_obj: any, args: any, context: SubscriptionContext) => liveMatch.getTodaysLiveMatches(_obj, args, context)},
    },
    LiveMatchSummary: {
        predictionSummary: (parent: PredictionSummaryRequest, args: any, context: TokenContext) => prediction.getPredictionSummary(context, parent),
        prediction: (parent: PredictionRequest, args: any, context: TokenContext) => prediction.getPrediction(context, parent),
    },
    HomeFeed: {
        score: (parent: GetMatchRequest, args: any, context: TokenContext) => user.getUserScore(context, parent),
        todaysMatches: (parent: GetMatchRequest, args: any, context: TokenContext) => liveMatch.listTodaysMatches(context),
        upcomingFixtures: (parent: any, args: any, context: TokenContext) => home.upcomingFixturesWithPredictions(context, parent),
    },
    MatchSummary: {
        match: (parent: GetMatchRequest, args: any, context: TokenContext) => liveMatch.getMatch(context, parent),
        predictionSummary: (parent: PredictionSummaryRequest, args: any, context: TokenContext) => prediction.getPredictionSummary(context, parent),
        prediction: (parent: PredictionRequest, args: any, context: TokenContext) => prediction.getPrediction(context, parent),
    },
    PredictorData: {
        fixtures: (parent: any, args: any, context: TokenContext) => fixture.getFixtures(context),
        predictions: (parent: GetUserPredictionRequest, args: any, context: TokenContext) => prediction.getUserPredictions(context, parent),
        forms: (parent: any, args: any, context: TokenContext) => fixture.getTeamForm(context),
    },
};

export default resolvers;
