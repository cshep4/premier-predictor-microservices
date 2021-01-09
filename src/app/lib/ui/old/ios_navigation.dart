//import 'package:flutter/cupertino.dart';
//import 'package:flutter/material.dart';
//
//import '../title_page.dart';
//
//class IOSNavigationPage extends StatefulWidget {
//  createState() => _IOSNavigationPageState();
//}
//
//class _IOSNavigationPageState extends State<IOSNavigationPage> {
//  @override
//  void initState() {
//    super.initState();
//  }
//
//  @override
//  void dispose() {
//    super.dispose();
//  }
//
//  @override
//  Widget build(BuildContext context) {
//    return CupertinoTabScaffold(
//      tabBar: CupertinoTabBar(
//        items: [
//          BottomNavigationBarItem(
//              icon: Icon(CupertinoIcons.home), title: Text("Home")),
//          BottomNavigationBarItem(
//              icon: Icon(CupertinoIcons.search), title: Text("Predictor")),
//          BottomNavigationBarItem(
//              icon: Icon(CupertinoIcons.search), title: Text("Results")),
//          BottomNavigationBarItem(
//              icon: Icon(CupertinoIcons.search), title: Text("Leagues")),
//          BottomNavigationBarItem(
//              icon: Icon(CupertinoIcons.person), title: Text("Account")),
//        ],
//      ),
//      tabBuilder: (context, index) {
//        switch (index) {
//          case 0:
//            return _homePage();
//            break;
//          case 1:
//            return _predictorPage();
//            break;
//          case 2:
//            return _resultsPage();
//            break;
//          case 3:
//            return _leaguePage();
//            break;
//          case 4:
//            return _accountArea();
//            break;
//          default:
//            return _accountArea();
//            break;
//        }
//      },
//    );
//  }
//
//  Widget _homePage() {
//    return HomePage(title: 'My Todo List');
//  }
//
//  Widget _predictorPage() {
//    return Center(
//      child: Text(
//        'Predictor Screen',
//        style: TextStyle(
//          fontWeight: FontWeight.w700,
//          color: Colors.red,
//          fontSize: 25.0,
//        ),
//      ),
//    );
//  }
//
//  Widget _resultsPage() {
//    return Center(
//      child: Text(
//        'Results Screen',
//        style: TextStyle(
//          fontWeight: FontWeight.w700,
//          color: Colors.blue,
//          fontSize: 25.0,
//        ),
//      ),
//    );
//  }
//
//  Widget _leaguePage() {
//    return Center(
//      child: Text(
//        'League Screen',
//        style: TextStyle(
//          fontWeight: FontWeight.w700,
//          color: Colors.blue,
//          fontSize: 25.0,
//        ),
//      ),
//    );
//  }
//
//  Widget _accountArea() {
//    return Center(
//      child: Text(
//        'Account Screen',
//        style: TextStyle(
//          fontWeight: FontWeight.w700,
//          color: Colors.blue,
//          fontSize: 25.0,
//        ),
//      ),
//    );
//  }
//}
