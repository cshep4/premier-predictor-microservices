import 'dart:async';
import 'package:premier_predictor/database/database.dart';
import 'package:premier_predictor/model/todo.dart';

class TodoDao {
  final DatabaseProvider _dbProvider;

  TodoDao(this._dbProvider);

  Future<int> createTodo(Todo todo) async {
    final db = await _dbProvider.database;
    var result = db.insert(todoTABLE, todo.toDatabaseJson());
    return result;
  }

  //Searches if query string was passed
  Future<List<Todo>> getTodos({List<String> columns, String query}) async {
    final db = await _dbProvider.database;

    List<Map<String, dynamic>> result;
    if (query != null && query.isNotEmpty) {
      result = await db.query(
        todoTABLE,
        columns: columns,
        where: 'description LIKE ?',
        whereArgs: ["%$query%"],
      );
    } else {
      result = await db.query(
        todoTABLE,
        columns: columns,
      );
    }

    List<Todo> todos = result.isNotEmpty
        ? result.map((item) => Todo.fromDatabaseJson(item)).toList()
        : [];
    return todos;
  }

  Future<int> updateTodo(Todo todo) async {
    final db = await _dbProvider.database;

    return await db.update(
      todoTABLE,
      todo.toDatabaseJson(),
      where: "id = ?",
      whereArgs: [todo.id],
    );
  }

  Future<int> deleteTodo(int id) async {
    final db = await _dbProvider.database;

    return await db.delete(
      todoTABLE,
      where: 'id = ?',
      whereArgs: [id],
    );
  }

  Future deleteAllTodos() async {
    final db = await _dbProvider.database;
    return await db.delete(todoTABLE);
  }
}
