import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_platform_widgets/flutter_platform_widgets.dart';
import 'package:premier_predictor/constant/home.dart';
import 'package:premier_predictor/ui/education_page.dart';

class EducationList extends StatelessWidget {
  final List<_educationType> educationTypes = [
    _educationType(
      title: "Rules",
      items: rules,
    ),
    _educationType(
      title: "Scoring",
      items: scoring,
    )
  ];

  @override
  Widget build(BuildContext context) {
    return Card(
      elevation: 0.1,
      shape: RoundedRectangleBorder(
        side: BorderSide(color: Colors.grey[300], width: 0.5),
        borderRadius: BorderRadius.circular(5),
      ),
      child: ListView.separated(
        itemCount: educationTypes.length,
        padding: EdgeInsets.zero,
        physics: const NeverScrollableScrollPhysics(),
        shrinkWrap: true,
        itemBuilder: (context, index) {
          return buildListTile(context, educationTypes[index].title,
              educationTypes[index].items);
        },
        separatorBuilder: (context, index) => Divider(
          color: Colors.grey[300],
          height: 0,
        ),
      ),
    );
  }

  ListTile buildListTile(
      BuildContext context, String title, List<EducationItem> items) {
    return ListTile(
      leading: Container(
        padding: EdgeInsets.only(right: 12.0),
        decoration: new BoxDecoration(
          border: new Border(
            right: new BorderSide(width: 1.0, color: Colors.grey[300]),
          ),
        ),
        child: Icon(Icons.sports_soccer, color: Colors.red),
      ),
      title: Text(
        title,
        style: TextStyle(color: Colors.black, fontWeight: FontWeight.bold),
      ),
      trailing:
          Icon(Icons.keyboard_arrow_right, color: Colors.black26, size: 30.0),
      onTap: () {
        Navigator.push(
          context,
          platformPageRoute(
            builder: (context) => EducationPage(
              title: title,
              items: items,
              previousTitle: "Home",
            ),
            context: context,
          ),
        );
      },
    );
  }
}

class _educationType {
  final String title;
  final List<EducationItem> items;

  _educationType({this.title, this.items});
}
