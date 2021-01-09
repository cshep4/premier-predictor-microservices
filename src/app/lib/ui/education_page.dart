import 'package:flutter/material.dart';
import 'package:flutter/widgets.dart';
import 'package:flutter_platform_widgets/flutter_platform_widgets.dart';
import 'package:premier_predictor/constant/home.dart';

class EducationPage extends StatelessWidget {
  final String previousTitle;
  final String title;
  final List<EducationItem> items;

  const EducationPage({Key key, this.items, this.title, this.previousTitle})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return PlatformScaffold(
      appBar: PlatformAppBar(
        title: Text(title),
        cupertino: (context, platform) => CupertinoNavigationBarData(
          previousPageTitle: previousTitle,
          actionsForegroundColor: Colors.red,
        ),
      ),
      body: _buildList(context),
      iosContentPadding: true,
    );
  }

  Widget _buildList(BuildContext context) {
    return ListView.builder(
      physics: const RangeMaintainingScrollPhysics(),
      itemCount: items.length,
      padding: EdgeInsets.only(top: 5, left: 5, right: 5),
      itemBuilder: (context, index) => Card(
        color: Colors.grey[50],
        elevation: 0.1,
        shape: RoundedRectangleBorder(
          side: BorderSide(color: Colors.grey[300], width: 0.5),
          borderRadius: BorderRadius.circular(5),
        ),
        child: ListTile(
          contentPadding: EdgeInsets.all(15.0),
          leading: Container(
            child: items[index].icon,
          ),
          title:
              Text(items[index].text, style: TextStyle(color: Colors.black54)),
        ),
      ),
    );
  }
}
