import 'package:flutter/widgets.dart';
import 'package:premier_predictor/bloc/home_bloc.dart';
import 'package:premier_predictor/bloc/match_facts_bloc.dart';
import 'package:premier_predictor/bloc/todo_bloc.dart';
import 'package:premier_predictor/model/home_item.dart';
import 'package:premier_predictor/ui/widget/education_list.dart';
import 'package:premier_predictor/ui/widget/loading.dart';
import 'package:premier_predictor/ui/widget/message.dart';
import 'package:premier_predictor/ui/widget/refresher.dart';
import 'package:premier_predictor/ui/widget/todays_matches.dart';
import 'package:premier_predictor/ui/widget/todo_list.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';

class HomePage extends StatefulWidget {
  final TodoBloc todoBloc;
  final HomeBloc homeBloc;
  final MatchFactsBloc matchFactsBloc;

  const HomePage({Key key, this.todoBloc, this.homeBloc, this.matchFactsBloc})
      : super(key: key);

  createState() => _HomePageState(todoBloc, homeBloc, matchFactsBloc);
}

class _HomePageState extends State<HomePage>
    with AutomaticKeepAliveClientMixin<HomePage> {
  final TodoBloc todoBloc;
  final HomeBloc homeBloc;
  final MatchFactsBloc matchFactsBloc;

  @override
  bool get wantKeepAlive => true;

  _HomePageState(this.todoBloc, this.homeBloc, this.matchFactsBloc);

  @override
  void initState() {
    todoBloc.init();
    homeBloc.init();
    super.initState();
  }

  @override
  void dispose() {
    print("tool page state being disposed! wantKeepAlive is " +
        wantKeepAlive.toString());
    todoBloc.dispose();
    homeBloc.dispose();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    super.build(context);
    return Container(
      child: getHomeWidget(),
    );
  }

  Widget getHomeWidget() {
    /*The StreamBuilder widget,
    basically this widget will take stream of data (todos)
    and construct the UI (with state) based on the stream
    */
    return StreamBuilder(
      stream: homeBloc.homeItems,
      builder: (BuildContext context, AsyncSnapshot<List<HomeItem>> snapshot) {
        return getHomeItemsWidget(snapshot);
      },
    );
  }

  Widget getHomeItemsWidget(AsyncSnapshot<List<HomeItem>> snapshot) {
    if (!snapshot.hasData) {
      homeBloc.getHomeItems();
      return Center(child: Loading());
    }

    if (snapshot.data.length == 0) {
      return Container(
        child: Center(
          child: noTodoMessageWidget(),
        ),
      );
    }

    return Refresher(
      onRefresh: homeBloc.getHomeItems,
      child: ListView.builder(
        itemCount: snapshot.data.length,
        padding: EdgeInsets.all(5),
        itemBuilder: (context, itemPosition) {
          HomeItem homeItem = snapshot.data[itemPosition];
          if (homeItem is MessageItem) {
            return new Message(message: homeItem.message);
          }
          if (homeItem is EducationItem) {
            return EducationList();
          }
          if (homeItem is TodaysMatchesItem) {
            // return StreamBuilder(
            //   stream: homeBloc.homeItems,
            //   builder: (BuildContext context,
            //       AsyncSnapshot<List<HomeItem>> snapshot) {
            //     return getHomeItemsWidget(snapshot);
            //   },
            // );
            return TodaysMatches(matchFactsBloc: matchFactsBloc);
          }

          return new TodoCardList(todoBloc: todoBloc);
        },
      ),
    );
  }

  Widget noTodoMessageWidget() {
    return Container(
      child: Text(
        "Start adding Todo...",
        style: TextStyle(fontSize: 19, fontWeight: FontWeight.w500),
      ),
    );
  }
}
