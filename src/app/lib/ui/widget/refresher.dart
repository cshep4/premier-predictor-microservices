import 'package:flutter/widgets.dart';
import 'package:pull_to_refresh/pull_to_refresh.dart';

class Refresher extends StatelessWidget {
  final RefreshController _refreshController =
      RefreshController(initialRefresh: false);

  final Widget child;
  final VoidCallback onRefresh;

  Refresher({Key key, this.onRefresh, this.child}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return SmartRefresher(
      enablePullDown: true,
      header: ClassicHeader(
        refreshStyle: RefreshStyle.Follow,
      ),
      controller: _refreshController,
      onRefresh: () {
        onRefresh();
        _refreshController.refreshCompleted();
      },
      child: child,
    );
  }
}
