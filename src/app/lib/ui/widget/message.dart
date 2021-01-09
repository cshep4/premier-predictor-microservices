import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';

class Message extends StatelessWidget {
  final String message;

  const Message({Key key, this.message}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Container(
      child: _buildMessageCard(),
    );
  }

  Card _buildMessageCard() {
    return Card(
      elevation: 0.1,
      shape: RoundedRectangleBorder(
        side: BorderSide(color: Colors.grey[300], width: 0.5),
        borderRadius: BorderRadius.circular(5),
      ),
      color: Colors.grey[50],
      child: ListTile(
        contentPadding: const EdgeInsets.all(20.0),
        leading: Container(
          padding: EdgeInsets.only(right: 12.0),
          child: Icon(
            Icons.info_outline,
            color: Colors.red,
//              color: Colors.blueAccent[100],
            size: 30.0,
          ),
        ),
        title: Text(
          message,
          style: TextStyle(
              fontSize: 14.5,
              color: Colors.black54,
              fontFamily: 'RobotoMono',
              fontWeight: FontWeight.w300),
        ),
      ),
    );
  }
}
