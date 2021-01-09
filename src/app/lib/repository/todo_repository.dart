import 'dart:async';

import 'package:graphql/client.dart';
import 'package:premier_predictor/dao/todo_dao.dart';
import 'package:premier_predictor/model/todo.dart';

class TodoRepository {
  final TodoDao _todoDao;

  TodoRepository(this._todoDao);

  Stream<FetchResult> _logStream;

  Future getAllTodos({String query}) => _todoDao.getTodos(query: query);

  Stream<List<Todo>> get({String query}) {
    StreamController<List<Todo>> controller = StreamController<List<Todo>>();

    controller.addStream(_todoDao.getTodos(query: query).asStream());

    Stream<List<Todo>> ss = _logStream.map((event) => null);

    ss.listen((event) {}, onDone: () {
      controller.close();
    }, onError: () {
      controller.close();
    }, cancelOnError: true);

    controller.addStream(ss);

    return controller.stream;
  }

  Future insertTodo(Todo todo) => _todoDao.createTodo(todo);

  Future updateTodo(Todo todo) => _todoDao.updateTodo(todo);

  Future deleteTodoById(int id) => _todoDao.deleteTodo(id);

  Future deleteAllTodos() => _todoDao.deleteAllTodos();
}
