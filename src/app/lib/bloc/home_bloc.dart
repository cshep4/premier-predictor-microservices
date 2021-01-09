import 'package:premier_predictor/model/home_item.dart';

import 'dart:async';

import 'package:premier_predictor/repository/todo_repository.dart';

class HomeBloc {
  //Get instance of the Repository
  final TodoRepository _todoRepository;

  //Stream controller is the 'Admin' that manages
  //the state of our stream of data like adding
  //new data, change the state of the stream
  //and broadcast it to observers/subscribers
  StreamController<List<HomeItem>> _homeController;

  get homeItems => _homeController.stream;

  HomeBloc(this._todoRepository);

  init() {
    if (_homeController == null || _homeController.isClosed) {
      _homeController = StreamController<List<HomeItem>>.broadcast();
    }
    getHomeItems();
  }

  getHomeItems({String query}) async {
    var homeItems = List<HomeItem>();
    var todos = await _todoRepository.getAllTodos(query: query);

//    homeItems.add(MessageItem(message: 'We are currently experiencing issues with storing predictions, please try again later.'));
    homeItems.add(EducationItem());
    homeItems.add(TodaysMatchesItem());

    if (todos.isNotEmpty) {
      homeItems.add(TodoList(todos: todos));
    }

    //sink is a way of adding data reactively to the stream by registering a new event
    _homeController.sink.add(homeItems);
  }

  dispose() {
    _homeController.close();
  }
}
