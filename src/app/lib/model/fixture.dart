class Fixture {
  String id;
  String hTeam;
  String aTeam;
  int hGoals;
  int aGoals;
  int played;
  DateTime dateTime;
  int matchday;

  Fixture({
    this.id,
    this.hTeam,
    this.aTeam,
    this.hGoals,
    this.aGoals,
    this.played,
    this.dateTime,
    this.matchday,
  });

  factory Fixture.fromJson(dynamic data) => Fixture(
        id: data["id"],
        hTeam: data["hTeam"],
        aTeam: data["aTeam"],
        hGoals: data["hGoals"],
        aGoals: data["aGoals"],
        played: data["played"],
        dateTime: DateTime.parse(data["dateTime"]),
        matchday: data["matchday"],
      );

  Map<String, dynamic> toJson() => {
        "id": id,
        "hTeam": hTeam,
        "aTeam": aTeam,
        "hGoals": hGoals,
        "aGoals": aGoals,
        "played": played,
        "dateTime": dateTime.toIso8601String(),
        "matchday": matchday,
      };
}
