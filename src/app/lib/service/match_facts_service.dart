import 'package:graphql/client.dart';
import "package:graphql/internal.dart";
import 'dart:async';

import 'package:premier_predictor/model/match_facts.dart';

class MatchFactsService {
  static final WebSocketLink _webSocketLink = WebSocketLink(
    url: 'wss://premierpredictor.uk/gateway/graphql',
    config: SocketClientConfig(
      autoReconnect: true,
      delayBetweenReconnectionAttempts: Duration(seconds: 1),
      initPayload: () => {
        "Authorization":
            "",
        "params": {"tenant": "Test"}
      },
    ),
  );

  final GraphQLClient _client = GraphQLClient(
    link: _webSocketLink,
    cache: InMemoryCache(),
  );

  Stream<MatchFacts> getTodaysMatches() {
    Operation operation = Operation(
      operationName: 'todaysLiveMatches',
      documentNode: gql(_todaysMatchesQuery),
    );
    operation.setContext(<String, Map<String, String>>{
      'headers': {"Authorization": "1234"}
    });
    return _client.subscribe(operation).map((event) =>
        MatchFacts.fromJson(event.data["todaysLiveMatches"]["match"]));
  }

  Stream<List<MatchDay>> getUpcoming() {
    Operation operation = Operation(
      operationName: 'upcomingMatches',
      documentNode: gql(_upcomingMatchesQuery),
    );
    return _client
        .subscribe(operation)
        .map((event) => event.data.map((d) => MatchDay.fromJson(d)));
  }

  Stream<MatchFacts> get(String id) {
    Operation operation = Operation(
      operationName: 'matchSummary',
      documentNode: gql(_matchSummaryQuery),
      variables: {
        'request': id,
      },
    );
    return _client
        .subscribe(operation)
        .map((event) => MatchFacts.fromJson(event.data));
  }
}

final String _todaysMatchesQuery = '''
subscription todaysLiveMatches {
    todaysLiveMatches {
        match {
            id
            compId
            formattedDate
            season
            week
            venue
            venueId
            venueCity
            venueLatitude
            venueLongitude
            venueCountry
            status
            timer
            time
            localTeamId
            localTeamName
            localTeamScore
            visitorTeamId
            visitorTeamName
            visitorTeamScore
            htScore
            ftScore
            etScore
            penaltyLocal
            penaltyVisitor
            events {
                id
                type
                result
                minute
                extraMin
                team
                player
                playerId
                assist
                assistId
            }
            commentary {
                matchId
                matchInfo {
                    stadium
                    attendance
                    referee
                }
                lineup {
                    localTeam {
                        id
                        number
                        name
                        pos
                    }
                    visitorTeam {
                        id
                        number
                        name
                        pos
                    }
                }
                subs {
                    localTeam {
                        id
                        number
                        name
                        pos
                    }
                    visitorTeam {
                        id
                        number
                        name
                        pos
                    }
                }
                substitutions {
                    localTeam {
                        offName
                        onName
                        offId
                        onId
                        minute
                        tableId
                    }
                    visitorTeam {
                        offName
                        onName
                        offId
                        onId
                        minute
                        tableId
                    }
                }
                comments {
                    id
                    important
                    goal
                    minute
                    comment
                }
                matchStats {
                    localTeam {
                        shotsTotal
                        shotsOnGoal
                        fouls
                        corners
                        offsides
                        possessionTime
                        yellowCards
                        redCards
                        saves
                        tableId
                    }
                    visitorTeam {
                        shotsTotal
                        shotsOnGoal
                        fouls
                        corners
                        offsides
                        possessionTime
                        yellowCards
                        redCards
                        saves
                        tableId
                    }
                }
                playerStats {
                    localTeam {
                        player {
                            id
                            num
                            name
                            pos
                            posX
                            posY
                            shotsTotal
                            shotsOnGoal
                            goals
                            assists
                            offsides
                            foulsDrawn
                            foulsCommitted
                            saves
                            yellowCards
                            redCards
                            penScore
                            penMiss
                        }
                    }
                    visitorTeam {
                        player {
                            id
                            num
                            name
                            pos
                            posX
                            posY
                            shotsTotal
                            shotsOnGoal
                            goals
                            assists
                            offsides
                            foulsDrawn
                            foulsCommitted
                            saves
                            yellowCards
                            redCards
                            penScore
                            penMiss
                        }
                    }
                }
            }
            matchDate
        }
      
    }
}
''';

final String _upcomingMatchesQuery = '''
    subscription upcomingMatches {
      upcomingMatches {
        matches {
          date
          matches {
            id
            compId
            formattedDate
            season
            week
            venue
            venueId
            venueCity
            status
            timer
            time
            localTeamId
            localTeamName
            localTeamScore
            visitorTeamId
            visitorTeamName
            visitorTeamScore
            htScore
            ftScore
            etScore
            penaltyLocal
            penaltyVisitor
            events {
              id
              type
              result
              minute
              extraMin
              team
              player
              playerId
              assist
              assistId
            }
            commentary {
              matchId
              matchInfo {
                stadium
                attendance
                referee
              }
              lineup {
                localTeam {
                  id
                  number
                  name
                  pos
                }
                visitorTeam {
                  id
                  number
                  name
                  pos
                }
              }
              subs {
                localTeam {
                  id
                  number
                  name
                  pos
                }
                visitorTeam {
                  id
                  number
                  name
                  pos
                }
              }
              substitutions {
                localTeam {
                  offName
                  onName
                  offId
                  onId
                  minute
                  tableId
                }
                visitorTeam {
                  offName
                  onName
                  offId
                  onId
                  minute
                  tableId
                }
              }
              comments {
                id
                important
                goal
                minute
                comment
              }
              matchStats {
                localTeam {
                  shotsTotal
                  shotsOnGoal
                  fouls
                  corners
                  offsides
                  possessionTime
                  yellowCards
                  redCards
                  saves
                  tableId
                }
                visitorTeam {
                  shotsTotal
                  shotsOnGoal
                  fouls
                  corners
                  offsides
                  possessionTime
                  yellowCards
                  redCards
                  saves
                  tableId
                }
              }
              playerStats {
                localTeam {
                  player {
                    id
                    num
                    name
                    pos
                    posX
                    posY
                    shotsTotal
                    shotsOnGoal
                    goals
                    assists
                    offsides
                    foulsDrawn
                    foulsCommitted
                    saves
                    yellowCards
                    redCards
                    penScore
                    penMiss
                  }
                }
                visitorTeam {
                  player {
                    id
                    num
                    name
                    pos
                    posX
                    posY
                    shotsTotal
                    shotsOnGoal
                    goals
                    assists
                    offsides
                    foulsDrawn
                    foulsCommitted
                    saves
                    yellowCards
                    redCards
                    penScore
                    penMiss
                  }
                }
              }
            }
            matchDate
          }
        }
      }
    }
''';

final String _matchSummaryQuery = '''
  subscription matchSummary(\$request: MatchSummaryInput!) {
    matchSummary(request: \$request) {
      match {
        id
        compId
        formattedDate
        season
        week
        venue
        venueId
        venueCity
        status
        timer
        time
        localTeamId
        localTeamName
        localTeamScore
        visitorTeamId
        visitorTeamName
        visitorTeamScore
        htScore
        ftScore
        etScore
        penaltyLocal
        penaltyVisitor
        events {
          id
          type
          result
          minute
          extraMin
          team
          player
          playerId
          assist
          assistId
        }
        commentary {
          matchId
          matchInfo {
            stadium
            attendance
            referee
          }
          lineup {
            localTeam {
              id
              number
              name
              pos
            }
            visitorTeam {
              id
              number
              name
              pos
            }
          }
          subs {
            localTeam {
              id
              number
              name
              pos
            }
            visitorTeam {
              id
              number
              name
              pos
            }
          }
          substitutions {
            localTeam {
              offName
              onName
              offId
              onId
              minute
              tableId
            }
            visitorTeam {
              offName
              onName
              offId
              onId
              minute
              tableId
            }
          }
          comments {
            id
            important
            goal
            minute
            comment
          }
          matchStats {
            localTeam {
              shotsTotal
              shotsOnGoal
              fouls
              corners
              offsides
              possessionTime
              yellowCards
              redCards
              saves
              tableId
            }
            visitorTeam {
              shotsTotal
              shotsOnGoal
              fouls
              corners
              offsides
              possessionTime
              yellowCards
              redCards
              saves
              tableId
            }
          }
          playerStats {
            localTeam {
              player {
              id
              num
              name
              pos
              posX
              posY
              shotsTotal
              shotsOnGoal
              goals
              assists
              offsides
              foulsDrawn
              foulsCommitted
              saves
              yellowCards
              redCards
              penScore
              penMiss
              }
            }
            visitorTeam {
              player {
                id
                num
                name
                pos
                posX
                posY
                shotsTotal
                shotsOnGoal
                goals
                assists
                offsides
                foulsDrawn
                foulsCommitted
                saves
                yellowCards
                redCards
                penScore
                penMiss
              }
            }
          }
        }
        matchDate
      }
      matchPredictionSummary {
        homeWin
        draw
        awayWin
      }
      prediction {
        userId
        matchId
        hGoals
        aGoals
      }
    }
  }
''';
