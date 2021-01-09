class MatchDay {
  DateTime date;
  List<MatchFacts> matches;

  MatchDay({
    this.date,
    this.matches,
  });

  factory MatchDay.fromJson(dynamic data) => MatchDay(
        date: DateTime.parse(data["date"]),
        matches: data["matches"].map((d) => MatchFacts.fromJson(d)).toList(),
      );

  bool isToday() {
    final now = DateTime.now();
    return DateTime(date.year, date.month, date.day) ==
        DateTime(now.year, now.month, now.day);
  }
}

class MatchFacts {
  String id;
  String compId;
  String formattedDate;
  String season;
  String week;
  String venue;
  String venueId;
  String venueCity;
  String venueLatitude;
  String venueLongitude;
  String venueCountry;
  String status;
  String timer;
  String time;
  String localTeamId;
  String localTeamName;
  String localTeamScore;
  String visitorTeamId;
  String visitorTeamName;
  String visitorTeamScore;
  String htScore;
  String ftScore;
  String etScore;
  String penaltyLocal;
  String penaltyVisitor;
  List<Event> events;
  Commentary commentary;
  DateTime matchDate;

  MatchFacts({
    this.id,
    this.compId,
    this.formattedDate,
    this.season,
    this.week,
    this.venue,
    this.venueId,
    this.venueCity,
    this.venueLatitude,
    this.venueLongitude,
    this.venueCountry,
    this.status,
    this.timer,
    this.time,
    this.localTeamId,
    this.localTeamName,
    this.localTeamScore,
    this.visitorTeamId,
    this.visitorTeamName,
    this.visitorTeamScore,
    this.htScore,
    this.ftScore,
    this.etScore,
    this.penaltyLocal,
    this.penaltyVisitor,
    this.events,
    this.commentary,
    this.matchDate,
  });

  factory MatchFacts.fromJson(dynamic data) => MatchFacts(
        id: data["id"],
        compId: data["compId"],
        formattedDate: data["formattedDate"],
        season: data["season"],
        week: data["week"],
        venue: data["venue"],
        venueId: data["venueId"],
        venueCity: data["venueCity"],
        venueLatitude: data["venueLatitude"],
        venueLongitude: data["venueLongitude"],
        venueCountry: data["venueCountry"],
        status: data["status"],
        timer: data["timer"],
        time: data["time"],
        localTeamId: data["localTeamId"],
        localTeamName: data["localTeamName"],
        localTeamScore: data["localTeamScore"],
        visitorTeamId: data["visitorTeamId"],
        visitorTeamName: data["visitorTeamName"],
        visitorTeamScore: data["visitorTeamScore"],
        htScore: data["htScore"],
        ftScore: data["ftScore"],
        etScore: data["etScore"],
        penaltyLocal: data["penaltyLocal"],
        penaltyVisitor: data["penaltyVisitor"],
        events: data["events"].map((d) => Event.fromJson(d)).toList(),
        commentary: data["commentary"],
        matchDate: DateTime.parse(data["matchDate"]),
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "compId": compId,
        "formattedDate": formattedDate,
        "season": season,
        "week": week,
        "venue": venue,
        "venueId": venueId,
        "venueCity": venueCity,
        "venueLatitude": venueLatitude,
        "venueLongitude": venueLongitude,
        "venueCountry": venueCountry,
        "status": status,
        "timer": timer,
        "time": time,
        "localTeamId": localTeamId,
        "localTeamName": localTeamName,
        "localTeamScore": localTeamScore,
        "visitorTeamId": visitorTeamId,
        "visitorTeamName": visitorTeamName,
        "visitorTeamScore": visitorTeamScore,
        "htScore": htScore,
        "ftScore": ftScore,
        "etScore": etScore,
        "penaltyLocal": penaltyLocal,
        "penaltyVisitor": penaltyVisitor,
        "events": events,
        "commentary": commentary,
        "matchDate": matchDate.toIso8601String(),
      };

  String getHomeTeam() {
    if (localTeamName.isEmpty) {
      return "TBC";
    }

    return localTeamName;
  }

  String getAwayTeam() {
    if (visitorTeamName.isEmpty) {
      return "TBC";
    }

    return visitorTeamName;
  }

  String getTimer() {
    if (status == "FT" || status == "HT") {
      return status;
    } else {
      return timer + "'";
    }
  }

  String getHomeScore() {
    if (localTeamScore.isEmpty) {
      return "0";
    } else {
      return localTeamScore;
    }
  }

  String getAwayScore() {
    if (visitorTeamScore.isEmpty) {
      return "0";
    } else {
      return visitorTeamScore;
    }
  }

  bool isPreGame() => status.isEmpty;

  bool isPlaying() => status != "FT" && status != "HT";
}

class Event {
  String id;
  String type;
  String result;
  String minute;
  String extraMin;
  String team;
  String player;
  String playerId;
  String assist;
  String assistId;

  Event({
    this.id,
    this.type,
    this.result,
    this.minute,
    this.extraMin,
    this.team,
    this.player,
    this.playerId,
    this.assist,
    this.assistId,
  });

  factory Event.fromJson(dynamic data) => Event(
        id: data["id"],
        type: data["type"],
        result: data["result"],
        minute: data["minute"],
        extraMin: data["extraMin"],
        team: data["team"],
        player: data["player"],
        playerId: data["playerId"],
        assist: data["assist"],
        assistId: data["assistId"],
      );
}

class Commentary {
  String matchId;
  MatchInfo matchInfo;
  Lineup lineup;
  Lineup subs;
  Substitutions substitutions;
  List<Comment> comments;
  MatchStats matchStats;
  PlayerStats playerStats;

  Commentary({
    this.matchId,
    this.matchInfo,
    this.lineup,
    this.subs,
    this.substitutions,
    this.comments,
    this.matchStats,
    this.playerStats,
  });

  factory Commentary.fromJson(dynamic data) => Commentary(
        matchId: data["matchId"],
        matchInfo: MatchInfo.fromJson(data["matchInfo"]),
        lineup: Lineup.fromJson(data["lineup"]),
        subs: Lineup.fromJson(data["subs"]),
        substitutions: Substitutions.fromJson(data["substitutions"]),
        comments: data["comments"].map((d) => Comment.fromJson(d)).toList(),
        matchStats: MatchStats.fromJson(data["matchStats"]),
        playerStats: PlayerStats.fromJson(data["playerStats"]),
      );
}

class MatchInfo {
  String stadium;
  String attendance;
  String referee;

  MatchInfo({
    this.stadium,
    this.attendance,
    this.referee,
  });

  factory MatchInfo.fromJson(dynamic data) => MatchInfo(
        stadium: data["stadium"],
        attendance: data["attendance"],
        referee: data["referee"],
      );
}

class Lineup {
  List<Position> localTeam;
  List<Position> visitorTeam;

  Lineup({
    this.localTeam,
    this.visitorTeam,
  });

  factory Lineup.fromJson(dynamic data) => Lineup(
        localTeam: data["localTeam"].map((d) => Position.fromJson(d)).toList(),
        visitorTeam:
            data["visitorTeam"].map((d) => Position.fromJson(d)).toList(),
      );
}

class Position {
  String id;
  String number;
  String name;
  String pos;

  Position({
    this.id,
    this.number,
    this.name,
    this.pos,
  });

  factory Position.fromJson(dynamic data) => Position(
        id: data["id"],
        number: data["number"],
        name: data["name"],
        pos: data["pos"],
      );
}

class Substitutions {
  List<Substitution> localTeam;
  List<Substitution> visitorTeam;

  Substitutions({
    this.localTeam,
    this.visitorTeam,
  });

  factory Substitutions.fromJson(dynamic data) => Substitutions(
        localTeam:
            data["localTeam"].map((d) => Substitution.fromJson(d)).toList(),
        visitorTeam:
            data["visitorTeam"].map((d) => Substitution.fromJson(d)).toList(),
      );
}

class Substitution {
  String offName;
  String onName;
  String offId;
  String onId;
  String minute;
  String tableId;

  Substitution({
    this.offName,
    this.onName,
    this.offId,
    this.onId,
    this.minute,
    this.tableId,
  });

  factory Substitution.fromJson(dynamic data) => Substitution(
        offName: data["offName"],
        onName: data["onName"],
        offId: data["offId"],
        onId: data["onId"],
        minute: data["minute"],
        tableId: data["tableId"],
      );
}

class Comment {
  String id;
  String important;
  String goal;
  String minute;
  String comment;

  Comment({
    this.id,
    this.important,
    this.goal,
    this.minute,
    this.comment,
  });

  factory Comment.fromJson(dynamic data) => Comment(
        id: data["id"],
        important: data["important"],
        goal: data["goal"],
        minute: data["minute"],
        comment: data["comment"],
      );
}

class MatchStats {
  List<TeamStats> localTeam;
  List<TeamStats> visitorTeam;

  MatchStats({
    this.localTeam,
    this.visitorTeam,
  });

  factory MatchStats.fromJson(dynamic data) => MatchStats(
        localTeam: data["localTeam"].map((d) => TeamStats.fromJson(d)).toList(),
        visitorTeam:
            data["visitorTeam"].map((d) => TeamStats.fromJson(d)).toList(),
      );
}

class TeamStats {
  String shotsTotal;
  String shotsOnGoal;
  String fouls;
  String corners;
  String offsides;
  String possessionTime;
  String yellowCards;
  String redCards;
  String saves;
  String tableId;

  TeamStats({
    this.shotsTotal,
    this.shotsOnGoal,
    this.fouls,
    this.corners,
    this.offsides,
    this.possessionTime,
    this.yellowCards,
    this.redCards,
    this.saves,
    this.tableId,
  });

  factory TeamStats.fromJson(dynamic data) => TeamStats(
        shotsTotal: data["shotsTotal"],
        shotsOnGoal: data["shotsOnGoal"],
        fouls: data["fouls"],
        corners: data["corners"],
        offsides: data["offsides"],
        possessionTime: data["possessionTime"],
        yellowCards: data["yellowCards"],
        redCards: data["redCards"],
        saves: data["saves"],
        tableId: data["tableId"],
      );
}

class PlayerStats {
  Players localTeam;
  Players visitorTeam;

  PlayerStats({
    this.localTeam,
    this.visitorTeam,
  });

  factory PlayerStats.fromJson(dynamic data) => PlayerStats(
        localTeam: data["localTeam"],
        visitorTeam: data["visitorTeam"],
      );
}

class Players {
  List<Player> player;

  Players({
    this.player,
  });

  factory Players.fromJson(dynamic data) => Players(
        player: data["player"].map((d) => Player.fromJson(d)).toList(),
      );
}

class Player {
  String id;
  String num;
  String name;
  String pos;
  String posX;
  String posY;
  String shotsTotal;
  String shotsOnGoal;
  String goals;
  String assists;
  String offsides;
  String foulsDrawn;
  String foulsCommitted;
  String saves;
  String yellowCards;
  String redCards;
  String penScore;
  String penMiss;

  Player({
    this.id,
    this.num,
    this.name,
    this.pos,
    this.posX,
    this.posY,
    this.shotsTotal,
    this.shotsOnGoal,
    this.goals,
    this.assists,
    this.offsides,
    this.foulsDrawn,
    this.foulsCommitted,
    this.saves,
    this.yellowCards,
    this.redCards,
    this.penScore,
    this.penMiss,
  });

  factory Player.fromJson(dynamic data) => Player(
        id: data["id"],
        num: data["num"],
        name: data["name"],
        pos: data["pos"],
        posX: data["posX"],
        posY: data["posY"],
        shotsTotal: data["shotsTotal"],
        shotsOnGoal: data["shotsOnGoal"],
        goals: data["goals"],
        assists: data["assists"],
        offsides: data["offsides"],
        foulsDrawn: data["foulsDrawn"],
        foulsCommitted: data["foulsCommitted"],
        saves: data["saves"],
        yellowCards: data["yellowCards"],
        redCards: data["redCards"],
        penScore: data["penScore"],
        penMiss: data["penMiss"],
      );
}
