import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:flutter_platform_widgets/flutter_platform_widgets.dart';
import 'package:premier_predictor/bloc/navbar.dart';

class NavigationPage extends StatefulWidget {
  final Widget homePage;

  const NavigationPage({Key key, this.homePage}) : super(key: key);

  createState() => _NavigationPageState(homePage);
}

class _NavigationPageState extends State<NavigationPage> {
  final Widget homePage;

  BottomNavBarBloc _bottomNavBarBloc;

  int _selectedTabIndex = 0;

  _NavigationPageState(this.homePage);

  @override
  void initState() {
    super.initState();
    _bottomNavBarBloc = BottomNavBarBloc();
  }

  @override
  void dispose() {
    _bottomNavBarBloc.close();
    super.dispose();
  }

  @override
  Widget build(BuildContext context) {
    return PlatformScaffold(
//      appBar: AppBar(
//        title: Text('Bottom NavBar Navigation'),
//      ),
      body: Builder(
        builder: (context) {
          switch (_selectedTabIndex) {
            case 0:
              return homePage;
            case 1:
              return _predictorPage();
            case 2:
              return _resultsPage();
            case 3:
              return _leaguePage();
            case 4:
              return _accountArea();
            default:
              return _accountArea();
          }
        },
      ),
      bottomNavBar: PlatformNavBar(
        backgroundColor: Colors.white,
        currentIndex: _selectedTabIndex,
        itemChanged: (index) => setState(
          () {
            _selectedTabIndex = index;
          },
        ),
        items: [
          BottomNavigationBarItem(
              icon: Icon(CupertinoIcons.home), label: "Home"),
          BottomNavigationBarItem(
              icon: Icon(CupertinoIcons.search), label: "Predictor"),
          BottomNavigationBarItem(
              icon: Icon(CupertinoIcons.search), label: "Results"),
          BottomNavigationBarItem(
              icon: Icon(CupertinoIcons.search), label: "Leagues"),
          BottomNavigationBarItem(
              icon: Icon(CupertinoIcons.person), label: "Account"),
        ],
      ),
    );
  }

  Widget _predictorPage() {
    return Center(
      child: Text(
        'Predictor Screen',
        style: TextStyle(
          fontWeight: FontWeight.w700,
          color: Colors.red,
          fontSize: 25.0,
        ),
      ),
    );
  }

  Widget _resultsPage() {
    return Center(
      child: Text(
        'Results Screen',
        style: TextStyle(
          fontWeight: FontWeight.w700,
          color: Colors.blue,
          fontSize: 25.0,
        ),
      ),
    );
  }

  Widget _leaguePage() {
    return Center(
      child: Text(
        'League Screen',
        style: TextStyle(
          fontWeight: FontWeight.w700,
          color: Colors.blue,
          fontSize: 25.0,
        ),
      ),
    );
  }

  Widget _accountArea() {
    return Center(
      child: Text(
        'Account Screen',
        style: TextStyle(
          fontWeight: FontWeight.w700,
          color: Colors.blue,
          fontSize: 25.0,
        ),
      ),
    );
  }
}
