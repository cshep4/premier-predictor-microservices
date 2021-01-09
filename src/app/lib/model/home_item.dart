import 'package:premier_predictor/model/todo.dart';

abstract class HomeItem {}

class TodoList extends HomeItem {
  List<Todo> todos;

  TodoList({this.todos});
}

class MessageItem extends HomeItem {
  String message;

  MessageItem({this.message});
}

class EducationItem extends HomeItem {}

class TodaysMatchesItem extends HomeItem {}

class UpcomingFixturesItem extends HomeItem {}
