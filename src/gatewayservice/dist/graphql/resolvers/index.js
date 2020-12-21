"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
const apollo_server_1 = require("apollo-server");
const grpcClient_1 = __importDefault(require("../services/grpcClient"));
const liveMatch_1 = require("./liveMatch/liveMatch");
const prediction_1 = require("./prediction/prediction");
const fixture_1 = require("./fixture/fixture");
const league_1 = require("./league/league");
const home_1 = require("./home/home");
const user_1 = require("./user/user");
const liveMatch = new liveMatch_1.LiveMatch(grpcClient_1.default.liveMatch(), new apollo_server_1.PubSub());
const prediction = new prediction_1.Prediction(grpcClient_1.default.prediction());
const fixture = new fixture_1.Fixture(grpcClient_1.default.fixture());
const league = new league_1.League(grpcClient_1.default.league());
const user = new user_1.User(grpcClient_1.default.user());
const home = new home_1.Home(league, prediction, fixture);
const resolvers = {
    Query: {
        matchSummary: (root, params, context) => {
            return {
                userId: params.request.userId,
                matchId: params.request.matchId,
                id: params.request.matchId,
            };
        },
        predictorData: (root, params, context) => {
            return { userId: params.request.userId };
        },
        homeFeed: (root, params, context) => home.feed(context, params.request),
    },
    Mutation: {
    // addPost: (root: any, params: any) => addPost(root, params)
    },
    Subscription: {
        upcomingMatches: { subscribe: (_obj, args, context) => liveMatch.getUpcomingMatches(_obj, args, context) },
        liveMatchSummary: { subscribe: (_obj, args, context) => liveMatch.getMatchSummary(_obj, args, context) },
        todaysLiveMatches: { subscribe: (_obj, args, context) => liveMatch.getTodaysLiveMatches(_obj, args, context) },
    },
    HomeFeed: {
        score: (parent, args, context) => user.getUserScore(context, parent),
        todaysMatches: (parent, args, context) => liveMatch.listTodaysMatches(context),
        upcomingFixtures: (parent, args, context) => home.upcomingFixturesWithPredictions(context, parent),
    },
    MatchSummary: {
        match: (parent, args, context) => liveMatch.getMatch(context, parent),
        predictionSummary: (parent, args, context) => prediction.getPredictionSummary(context, parent),
        prediction: (parent, args, context) => prediction.getPrediction(context, parent),
    },
    PredictorData: {
        fixtures: (parent, args, context) => fixture.getFixtures(context),
        predictions: (parent, args, context) => prediction.getUserPredictions(context, parent),
        forms: (parent, args, context) => fixture.getTeamForm(context),
    },
};
exports.default = resolvers;
