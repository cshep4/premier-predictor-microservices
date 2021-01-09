import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class TitlePage extends StatelessWidget {
  final String title;
  final Widget page;

  TitlePage({Key key, this.title, this.page}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return CupertinoPageScaffold(
      child: NestedScrollView(
        headerSliverBuilder: (BuildContext context, bool innerBoxIsScrolled) {
          return <Widget>[
            CupertinoSliverNavigationBar(
              largeTitle: Text(title),
              backgroundColor: Colors.white,
            )
          ];
        },
        body: page,
      ),
    );
  }
}
