import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';

class EducationItem {
  final Widget icon;
  final String text;

  EducationItem({this.icon, this.text});
}

List<EducationItem> get rules => [
      EducationItem(
        icon: Icon(Icons.sports_soccer, color: Colors.red),
        text:
            "Predictions for each match can be added/updated any time before the match kicks off.",
      ),
      EducationItem(
        icon: Icon(Icons.sports_soccer, color: Colors.red),
        text: "Predictions can be changed as often as you wish.",
      ),
      EducationItem(
        icon: Icon(Icons.sports_soccer, color: Colors.red),
        text:
            "Once a match kicks off, the prediction for that game will be locked in and cannot be changed.",
      ),
      EducationItem(
          icon: Icon(Icons.sports_soccer, color: Colors.red),
          text:
              "User scores will be updated within 24 hours of the full time whistle."),
    ];

List<EducationItem> get scoring => [
      EducationItem(
        icon: Text("1 Pt.", style: TextStyle(color: Colors.red)),
        text:
            "Correct amount of goals for a team. (e.g. prediction of 2-1, final score is 2-3, 1 point will be awarded. Prediction of 3-1, final score is 2-2, 0 points will be awarded).",
      ),
      EducationItem(
        icon: Text("3 Pt.", style: TextStyle(color: Colors.red)),
        text:
            "Correct result, plus points for correct amount of goals (if any). (e.g. prediction of 2-1, final score is 3-2, 3 points will be awarded for predicting the correct team to win).",
      ),
      EducationItem(
        icon: Text("3 Pt.", style: TextStyle(color: Colors.red)),
        text:
            "Correct scoreline, plus points for a correct result and goals. (e.g. prediction of 2-1, final score is 2-1, 8 points will be awarded - 3 for correct scoreline, 3 for correct result, 1 for each team's correct amount of goals).",
      ),
      EducationItem(
          icon: Text("5 Pt.", style: TextStyle(color: Colors.red)),
          text:
              "Each correct team finishing position in the league (based on your predictions)."),
      EducationItem(
          icon: Text("20 Pt.", style: TextStyle(color: Colors.red)),
          text: "The predicted winner chosen at sign-up finish as champions."),
    ];
