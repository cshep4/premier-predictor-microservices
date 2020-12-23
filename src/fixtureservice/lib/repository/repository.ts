import {FixtureSchema} from "../schema/fixture";
import * as mongoose from 'mongoose';
import {Fixture} from "../model/fixture";

const DATABASE = 'fixture';
const COLLECTION = '2020-2021';

export const ERR_FIXTURES_NOT_FOUND = "fixtures not found";

const FixtureModel = mongoose.model(COLLECTION, FixtureSchema);

export class Repository {
    constructor() {
        this.mongoSetup();
    }

    private mongoSetup(): void {
        let mongoScheme = process.env.MONGO_SCHEME;
        let mongoUser = process.env.MONGO_USERNAME;
        let mongoPassword = process.env.MONGO_PASSWORD;
        let mongoHost = process.env.MONGO_HOST;
        let mongoPort = process.env.MONGO_PORT;

        let mongoUrl: string = mongoScheme + '://';

        if (mongoUser && mongoPassword) {
            mongoUrl += mongoUser + ':' + mongoPassword + '@';
        }

        mongoUrl += mongoHost;

        if (mongoPort) {
            mongoUrl += ':' + mongoPort;
        }

        mongoUrl += '/' + DATABASE + '?retryWrites=true';

        mongoose.Promise = global.Promise;
        mongoose.connect(mongoUrl, {useNewUrlParser: true});
    }

    public getAllFixtures(): Promise<Fixture[]> {
        return new Promise<Fixture[]>((resolve, reject) => {
            FixtureModel.find().then(fixtures => {
                resolve(
                    fixtures.map(f => ({
                        id: f._doc._id,
                        played: f._doc.played,
                        dateTime: f._doc.dateTime,
                        matchday: f._doc.matchday,
                        hTeam: f._doc.hTeam,
                        aTeam: f._doc.aTeam,
                        hGoals: f._doc.hGoals,
                        aGoals: f._doc.aGoals,
                    }))
                );
            }, err => {
                reject(err);
            });
        });
    }

    public getAllPastFixtures(): Promise<Fixture[]> {
        return new Promise<Fixture[]>((resolve, reject) => {
            const filter = {
                'played': 1,
                'hGoals': {
                    '$ne': null
                },
                'aGoals': {
                    '$ne': null
                }
            };
            FixtureModel.find(filter).then(fixtures => {
                resolve(
                    fixtures.map(f => ({
                        id: f._doc._id,
                        played: f._doc.played,
                        dateTime: f._doc.dateTime,
                        matchday: f._doc.matchday,
                        hTeam: f._doc.hTeam,
                        aTeam: f._doc.aTeam,
                        hGoals: f._doc.hGoals,
                        aGoals: f._doc.aGoals,
                    }))
                );
            }, err => {
                reject(err);
            });
        });
    }

    public getAllFutureFixtures(): Promise<Map<string, string>> {
        return new Promise<Map<string, string>>((resolve, reject) => {
            const filter = {
                'dateTime': {
                    '$gte': new Date()
                },
            };
            FixtureModel.find(filter).then(fixtures => {
                const result = fixtures.reduce(function(map, f) {
                    map[f.id] = f.dateTime.toDateString();
                    return map;
                }, {});
                resolve(result);
            }, err => {
                reject(err);
            });
        });
    }

    public getFixtureById(id: string): Promise<Fixture> {
        return new Promise<Fixture>((resolve, reject) => {
            FixtureModel.findById(id).then(f => {
                if (!f) {
                    reject(ERR_FIXTURES_NOT_FOUND);
                    return;
                }

                let fixture = {
                    id: f._doc._id,
                    played: f._doc.played,
                    dateTime: f._doc.dateTime,
                    matchday: f._doc.matchday,
                    hTeam: f._doc.hTeam,
                    aTeam: f._doc.aTeam,
                    hGoals: f._doc.hGoals,
                    aGoals: f._doc.aGoals,
                };

                resolve(fixture);
            }, err => {
                reject(err);
            });
        });
    }
}
