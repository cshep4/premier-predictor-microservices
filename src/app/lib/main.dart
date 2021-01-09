import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter/services.dart';
import 'package:flutter_platform_widgets/flutter_platform_widgets.dart';
import 'package:premier_predictor/bloc/home_bloc.dart';
import 'package:premier_predictor/bloc/match_facts_bloc.dart';
import 'package:premier_predictor/dao/match_facts_dao.dart';
import 'package:premier_predictor/repository/match_facts_repository.dart';
import 'package:premier_predictor/repository/todo_repository.dart';
import 'package:premier_predictor/service/match_facts_service.dart';
import 'package:premier_predictor/ui/home_page.dart';
import 'package:flutter/foundation.dart' as foundation;
import 'package:premier_predictor/ui/navigation.dart';
import 'package:premier_predictor/ui/widget/title_page.dart';

import 'bloc/todo_bloc.dart';
import 'dao/todo_dao.dart';
import 'database/database.dart';

bool get isIos =>
    foundation.defaultTargetPlatform == foundation.TargetPlatform.iOS;

void main() => runApp(MyApp());

class MyApp extends StatelessWidget {
  @override
  Widget build(BuildContext context) {
    SystemChrome.setPreferredOrientations([
      DeviceOrientation.portraitUp,
      DeviceOrientation.portraitDown,
    ]);
    return PlatformApp(
      title: 'Premier Predictor',
      home: _buildApp(),
    );
  }

  Widget _buildApp() {
    var db = DatabaseProvider();
    var todoDao = TodoDao(db);
    var matchFactsDao = MatchFactsDao(db);
    var matchFactsService = MatchFactsService();
    var todoRespository = TodoRepository(todoDao);
    var matchFactsRespository =
        MatchFactsRepository(matchFactsDao, matchFactsService);
    var todoBloc = TodoBloc(todoRespository);
    var homeBloc = HomeBloc(todoRespository);
    var matchFactsBloc = MatchFactsBloc(matchFactsRespository);
    var homePage = HomePage(
      todoBloc: todoBloc,
      homeBloc: homeBloc,
      matchFactsBloc: matchFactsBloc,
    );

    return NavigationPage(
      homePage: TitlePage(title: "Premier Predictor", page: homePage),
    );
  }
}
