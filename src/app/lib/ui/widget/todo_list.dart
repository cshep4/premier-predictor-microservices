import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:premier_predictor/bloc/home_bloc.dart';
import 'package:premier_predictor/bloc/todo_bloc.dart';
import 'package:premier_predictor/model/home_item.dart';
import 'package:premier_predictor/model/todo.dart';
import 'package:premier_predictor/ui/widget/loading.dart';
import 'package:premier_predictor/ui/widget/todo.dart';

class TodoCardList extends StatelessWidget {
  //We load our Todo BLoC that is used to get
  //the stream of Todo for StreamBuilder
  final TodoBloc todoBloc;

  TodoCardList({Key key, this.todoBloc}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: getTodosWidget(),
    );
  }

  Widget getTodosWidget() {
    /*The StreamBuilder widget,
    basically this widget will take stream of data (todos)
    and construct the UI (with state) based on the stream
    */
    return StreamBuilder(
      stream: todoBloc.todos,
      builder: (BuildContext context, AsyncSnapshot<List<Todo>> snapshot) {
        return getTodoCardWidget(snapshot);
      },
    );
  }

  Widget getTodoCardWidget(AsyncSnapshot<List<Todo>> snapshot) {
    if (!snapshot.hasData) {
      todoBloc.getTodos();
      return Container();
    }

    if (snapshot.data.length == 0) {
      return Container();
    }

    return ListView.builder(
      physics: const NeverScrollableScrollPhysics(),
      shrinkWrap: true,
      itemCount: snapshot.data.length,
      itemBuilder: (context, itemPosition) {
        Todo todo = snapshot.data[itemPosition];
        return new TodoCard(
          todo: todo,
          todoBloc: todoBloc,
        );
      },
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

  dispose() {
    /*close the stream in order
    to avoid memory leaks
    */
    todoBloc.dispose();
  }
}
