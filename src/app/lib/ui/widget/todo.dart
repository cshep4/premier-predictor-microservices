import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:premier_predictor/bloc/todo_bloc.dart';
import 'package:premier_predictor/model/todo.dart';

class TodoCard extends StatelessWidget {
  final DismissDirectionCallback onDismissed;
  final GestureTapCallback onTap;
  final Todo todo;
  final TodoBloc todoBloc;

  final DismissDirection _dismissDirection = DismissDirection.horizontal;

  const TodoCard(
      {Key key, this.onDismissed, this.todo, this.onTap, this.todoBloc})
      : super(key: key);

  @override
  Widget build(BuildContext context) {
    return new Dismissible(
      background: Container(
        child: Padding(
          padding: EdgeInsets.only(left: 10),
          child: Align(
            alignment: Alignment.centerLeft,
            child: Text(
              "Deleting",
              style: TextStyle(color: Colors.white),
            ),
          ),
        ),
        color: Colors.redAccent,
      ),
      onDismissed: (direction) {
        // delete Todo item by ID whenever the card is dismissed
        todoBloc.deleteTodoById(todo.id);
      },
      direction: _dismissDirection,
      key: new ObjectKey(todo),
      child: Card(
          shape: RoundedRectangleBorder(
            side: BorderSide(color: Colors.grey[200], width: 0.5),
            borderRadius: BorderRadius.circular(5),
          ),
          color: Colors.white,
          child: ListTile(
            leading: InkWell(
              onTap: () {
                //Reverse the value
                todo.isDone = !todo.isDone;
                // This will update Todo isDone with either completed or nots
                todoBloc.updateTodo(todo);
              },
              child: Container(
                //decoration: BoxDecoration(),
                child: Padding(
                  padding: const EdgeInsets.all(15.0),
                  child: todo.isDone
                      ? Icon(
                          Icons.done,
                          size: 26.0,
                          color: Colors.indigoAccent,
                        )
                      : Icon(
                          Icons.check_box_outline_blank,
                          size: 26.0,
                          color: Colors.tealAccent,
                        ),
                ),
              ),
            ),
            title: Text(
              todo.description,
              style: TextStyle(
                  fontSize: 16.5,
                  fontFamily: 'RobotoMono',
                  fontWeight: FontWeight.w500,
                  decoration: todo.isDone
                      ? TextDecoration.lineThrough
                      : TextDecoration.none),
            ),
          )),
    );
  }
}
