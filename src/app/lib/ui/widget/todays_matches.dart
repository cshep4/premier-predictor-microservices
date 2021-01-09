import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:premier_predictor/bloc/match_facts_bloc.dart';
import 'package:premier_predictor/model/match_facts.dart';

class TodaysMatches extends StatelessWidget {
  final MatchFactsBloc matchFactsBloc;

  List<MatchFacts> matches;

  TodaysMatches({Key key, this.matchFactsBloc}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    matchFactsBloc.initTodaysMatches();
    matches = List<MatchFacts>();
    matches.add(MatchFacts());
    matches.add(MatchFacts());
    matches.add(MatchFacts());

    matches[0].localTeamId = "6";
    matches[0].localTeamName = "Manchester City";
    matches[0].localTeamScore = "0";
    matches[0].visitorTeamId = "5";
    matches[0].visitorTeamName = "Liverpool";
    matches[0].visitorTeamScore = "3";
    matches[0].status = "90";
    matches[0].timer = "90+5";
    matches[0].time = "15:00";

    matches[1].localTeamId = "7";
    matches[1].localTeamName = "Manchester United";
    matches[1].localTeamScore = "0";
    matches[1].visitorTeamId = "1";
    matches[1].visitorTeamName = "Arsenal";
    matches[1].visitorTeamScore = "1";
    matches[1].status = "FT";

    matches[2].localTeamId = "12";
    matches[2].localTeamName = "West Ham United";
    matches[2].localTeamScore = "";
    matches[2].visitorTeamId = "332";
    matches[2].visitorTeamName = "Leeds United";
    matches[2].visitorTeamScore = "";
    matches[2].status = "";
    matches[2].time = "15:00";

    return Column(
      children: <Widget>[_buildTitle(), _buildTodaysMatchesCard()],
    );
  }

  Widget _buildTitle() {
    return Container(
      padding: EdgeInsets.only(top: 12, bottom: 4.0),
      width: double.infinity,
      child: Padding(
        padding: const EdgeInsets.symmetric(horizontal: 6),
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          mainAxisSize: MainAxisSize.max,
          children: <Widget>[
            RichText(
              textAlign: TextAlign.start,
              text: TextSpan(
                children: [
                  TextSpan(
                    text: 'Todays Matches',
                    style: TextStyle(
                      color: Colors.black,
                      fontFamily: 'RobotoMono',
                      fontWeight: FontWeight.w700,
                      fontSize: 20.0,
                    ),
                  ),
                ],
              ),
            ),
          ],
        ),
      ),
    );
  }

  Widget _buildTodaysMatchesCard() {
    return Card(
      elevation: 0.1,
      shape: RoundedRectangleBorder(
        side: BorderSide(
          color: Colors.grey[300],
          width: 0.5,
        ),
        borderRadius: BorderRadius.circular(5),
      ),
      child: ListView.separated(
        itemCount: matches.length,
        padding: EdgeInsets.zero,
        physics: const NeverScrollableScrollPhysics(),
        shrinkWrap: true,
        itemBuilder: (context, index) {
          return _buildListTile(matches[index]);
        },
        separatorBuilder: (context, index) => Divider(
          color: Colors.grey[300],
          height: 0,
        ),
      ),
    );
  }

  Widget _buildListTile(MatchFacts match) {
    List<Widget> widgets = List<Widget>();

    widgets.add(_buildTeamName(match.getHomeTeam()));
    widgets.add(_buildTeamLogo(match.localTeamId));

    _buildMatchStatus(match).forEach((element) {
      widgets.add(element);
    });

    widgets.add(_buildTeamLogo(match.visitorTeamId));
    widgets.add(_buildTeamName(match.getAwayTeam()));

    return Row(
      children: widgets,
    );
  }

  Expanded _buildTeamLogo(String id) {
    return Expanded(
      flex: 12,
      child: Container(
        height: 50,
        padding: EdgeInsets.only(top: 7.0, bottom: 7.0),
        alignment: Alignment.center, // This is needed
        child: Center(
          child: Image(
            image: AssetImage('assets/emblem/' + id + '.png'),
          ),
        ),
      ),
    );
  }

  Expanded _buildTeamName(String name) {
    return Expanded(
      flex: 30,
      child: _padding(
        Center(
          child: Text(
            name,
            textAlign: TextAlign.center,
            style: TextStyle(
              fontSize: 10.0,
              fontWeight: FontWeight.w400,
              color: Colors.black,
            ),
          ),
        ),
      ),
    );
  }

  List<Widget> _buildMatchStatus(MatchFacts match) {
    if (match.isPreGame()) {
      return [
        Expanded(
          flex: 30,
          child: _padding(
            Center(
              child: Text(
                match.time,
                textAlign: TextAlign.center,
                style: TextStyle(
                  fontSize: 10.0,
                  fontWeight: FontWeight.w400,
                  color: Colors.black,
                ),
              ),
            ),
          ),
        ),
      ];
    }

    return [
      Expanded(
        flex: 10,
        child: _padding(
          Center(
            child: Text(
              match.getHomeScore(),
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: 10.0,
                fontWeight: FontWeight.w400,
                color: getStatusColour(match),
              ),
            ),
          ),
        ),
      ),
      Expanded(
        flex: 10,
        child: Center(
          child: Text(
            match.getTimer(),
            textAlign: TextAlign.center,
            style: TextStyle(
              fontSize: 10.0,
              fontWeight: FontWeight.w400,
              color: getStatusColour(match),
            ),
          ),
        ),
      ),
      Expanded(
        flex: 10,
        child: _padding(
          Center(
            child: Text(
              match.getAwayScore(),
              textAlign: TextAlign.center,
              style: TextStyle(
                fontSize: 10.0,
                fontWeight: FontWeight.w400,
                color: getStatusColour(match),
              ),
            ),
          ),
        ),
      ),
    ];
  }

  Widget _padding(Widget widget) {
    return Container(
      padding: EdgeInsets.all(10.0),
      child: widget,
    );
  }

  Color getStatusColour(MatchFacts match) {
    if (match.isPlaying()) {
      return Colors.red;
    }

    return Colors.black;
  }
}
