import 'package:flutter/widgets.dart';
import 'package:premier_predictor/ui/android/title_page.dart' as android;
import 'package:premier_predictor/ui/ios/title_page.dart' as ios;
import 'package:flutter/foundation.dart' as foundation;

class TitlePage extends StatelessWidget {
  final platform = foundation.defaultTargetPlatform;

  final String title;
  final Widget page;

  TitlePage({Key key, this.title, this.page}) : super(key: key);

  @override
  Widget build(BuildContext context) {
    if (platform == foundation.TargetPlatform.iOS) {
      return ios.TitlePage(title: title, page: page);
    }

    return android.TitlePage(title: title, page: page);
  }
}
