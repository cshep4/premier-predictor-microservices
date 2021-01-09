//import 'package:flutter/cupertino.dart';
//import 'package:flutter/material.dart';
//import 'package:premier_predictor/bloc/navbar.dart';
//
//import '../title_page.dart';
//
//class AndroidNavigationPage extends StatefulWidget {
//  createState() => _AndroidNavigationPageState();
//}
//
//class _AndroidNavigationPageState extends State<AndroidNavigationPage> {
//  BottomNavBarBloc _bottomNavBarBloc;
//
//  @override
//  void initState() {
//    super.initState();
//    _bottomNavBarBloc = BottomNavBarBloc();
//  }
//
//  @override
//  void dispose() {
//    _bottomNavBarBloc.close();
//    super.dispose();
//  }
//
//  @override
//  Widget build(BuildContext context) {
//    return Scaffold(
////      appBar: AppBar(
////        title: Text('Bottom NavBar Navigation'),
////      ),
//      body: StreamBuilder<NavBarItem>(
//        stream: _bottomNavBarBloc.itemStream,
//        initialData: _bottomNavBarBloc.defaultItem,
//        builder: (BuildContext context, AsyncSnapshot<NavBarItem> snapshot) {
//          switch (snapshot.data) {
//            case NavBarItem.HOME:
//              return _homePage();
//            case NavBarItem.PREDICTOR:
//              return _predictorPage();
//            case NavBarItem.RESULTS:
//              return _resultsPage();
//            case NavBarItem.LEAGUE:
//              return _leaguePage();
//            case NavBarItem.ACCOUNT:
//              return _accountArea();
//          }
//
//          return _accountArea();
//        },
//      ),
//      bottomNavigationBar: StreamBuilder(
//        stream: _bottomNavBarBloc.itemStream,
//        initialData: _bottomNavBarBloc.defaultItem,
//        builder: (BuildContext context, AsyncSnapshot<NavBarItem> snapshot) {
//          return BottomNavigationBar(
//            fixedColor: Colors.blueAccent,
//            currentIndex: snapshot.data.index,
//            onTap: _bottomNavBarBloc.pickItem,
//            items: [
//              BottomNavigationBarItem(
//                title: Text('Home'),
//                icon: Icon(Icons.home),
//              ),
//              BottomNavigationBarItem(
//                title: Text('Predictor'),
//                icon: Icon(Icons.notifications),
//              ),
//              BottomNavigationBarItem(
//                title: Text('Results'),
//                icon: Icon(Icons.notifications),
//              ),
//              BottomNavigationBarItem(
//                title: Text('Leagues'),
//                icon: Icon(Icons.notifications),
//              ),
//              BottomNavigationBarItem(
//                title: Text('Settings'),
//                icon: Icon(Icons.settings),
//              ),
//            ],
//          );
//        },
//      ),
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
