import 'package:premier_predictor/model/todo.dart';
import 'package:premier_predictor/repository/todo_repository.dart';

import 'dart:async';

class TodoBloc {
  final TodoRepository _todoRepository;

  StreamController<List<Todo>> _todoController;

  get todos => _todoController.stream;

  TodoBloc(this._todoRepository);

  init() {
    if (_todoController == null || _todoController.isClosed) {
      _todoController = StreamController<List<Todo>>.broadcast();
    }
    getTodos();
  }

  getTodos({String query}) async {
    //sink is a way of adding data reactively to the stream
    //by registering a new event
    _todoController.sink.add(await _todoRepository.getAllTodos(query: query));
  }

  addTodo(Todo todo) async {
    await _todoRepository.insertTodo(todo);
    getTodos();
  }

  updateTodo(Todo todo) async {
    await _todoRepository.updateTodo(todo);
    getTodos();
  }

  deleteTodoById(int id) async {
    _todoRepository.deleteTodoById(id);
    getTodos();
  }

  dispose() {
    _todoController?.close();
  }
}
