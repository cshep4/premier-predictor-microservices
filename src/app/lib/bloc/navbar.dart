import 'dart:async';

enum NavBarItem { HOME, PREDICTOR, RESULTS, LEAGUE, ACCOUNT }

class BottomNavBarBloc {
  final StreamController<NavBarItem> _navBarController =
      StreamController<NavBarItem>.broadcast();

  NavBarItem defaultItem = NavBarItem.HOME;

  Stream<NavBarItem> get itemStream => _navBarController.stream;

  void pickItem(int i) {
    switch (i) {
      case 0:
        _navBarController.sink.add(NavBarItem.HOME);
        break;
      case 1:
        _navBarController.sink.add(NavBarItem.PREDICTOR);
        break;
      case 2:
        _navBarController.sink.add(NavBarItem.RESULTS);
        break;
      case 3:
        _navBarController.sink.add(NavBarItem.LEAGUE);
        break;
      case 4:
        _navBarController.sink.add(NavBarItem.ACCOUNT);
        break;
    }
  }

  close() {
    _navBarController?.close();
  }
}
