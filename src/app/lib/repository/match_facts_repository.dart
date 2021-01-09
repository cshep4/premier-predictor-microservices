import 'dart:async';

import 'package:premier_predictor/dao/match_facts_dao.dart';
import 'package:premier_predictor/model/match_facts.dart';
import 'package:premier_predictor/service/match_facts_service.dart';

class MatchFactsRepository {
  final MatchFactsDao _matchFactsDao;
  final MatchFactsService _matchFactsService;

  StreamController<MatchFacts> _todayController;
  Stream<MatchFacts> _todaysMatchesStream;

  StreamController<MatchFacts> _matchController;
  Stream<MatchFacts> _matchStream;

  MatchFactsRepository(this._matchFactsDao, this._matchFactsService);

  Stream<MatchFacts> getTodaysMatches({String query}) {
    _todayController = StreamController<MatchFacts>();

    _matchFactsDao.getTodaysMatches().then(
          (matches) => matches.forEach((m) => _todayController.add(m)),
        );

    _todaysMatchesStream = _matchFactsService.getTodaysMatches();

    StreamSubscription<MatchFacts> ss = _todaysMatchesStream.listen(
      (m) => _updateMatches(m, _matchController),
      onDone: () => _todayController.close(),
      onError: (Object error) => _todayController.close(),
      cancelOnError: true,
    );

    _todayController.onCancel = () {
      ss.cancel();
    };

    return _todayController.stream;
  }

  Stream<MatchFacts> get(String id) {
    _matchController = StreamController();
    _matchController.addStream(_matchFactsDao.get(id).asStream());

    _matchStream = _matchFactsService.get(id);
    StreamSubscription<MatchFacts> ss = _matchStream.listen(
      (m) => _updateMatches(m, _matchController),
      onDone: () => _matchController.close(),
      onError: () => _matchController.close(),
      cancelOnError: true,
    );

    _todayController.onCancel = () {
      ss.cancel();
    };

    return _matchController.stream;
  }

  void _updateMatches(MatchFacts m, StreamController<MatchFacts> controller) {
    controller.add(m);
    _matchFactsDao.upsert(m);
  }
}
