import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_platform_widgets/flutter_platform_widgets.dart';
import 'package:premier_predictor/bloc/navbar.dart';
import 'package:premier_predictor/bloc/todo_bloc.dart';
import 'package:premier_predictor/model/todo.dart';

class TitlePage extends StatelessWidget {
  final StatelessWidget page;
  final String title;

  TitlePage({Key key, this.title, this.page}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return PlatformScaffold(
      appBar: PlatformAppBar(
        material: (_, __) => MaterialAppBarData(
          leading: Icon(PlatformIcons(context).book),
//          trailing: Icon(PlatformIcons(context).book),
          title: new Text(title),
        ),
      ),
      body: SafeArea(
        child: Container(
          color: Colors.white,
          padding: const EdgeInsets.only(left: 2.0, right: 2.0, bottom: 2.0),
          child: page,
        ),
      ),
    );
  }
}
