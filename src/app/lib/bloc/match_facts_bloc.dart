import 'dart:async';

import 'package:premier_predictor/model/match_facts.dart';
import 'package:premier_predictor/repository/match_facts_repository.dart';

class MatchFactsBloc {
  final MatchFactsRepository _matchFactsRepository;

  StreamController<MatchFacts> _todayController =
      StreamController<MatchFacts>.broadcast();
  StreamController<MatchFacts> _matchController =
      StreamController<MatchFacts>.broadcast();

  get today => _todayController.stream;

  get match => _matchController.stream;

  MatchFactsBloc(this._matchFactsRepository);

  initTodaysMatches() {
    if (_todayController == null || _todayController.isClosed) {
      _todayController = StreamController<MatchFacts>.broadcast();
    }
    getTodaysMatches();
  }

  initMatch(String id) {
    if (_matchController == null || _matchController.isClosed) {
      _matchController = StreamController<MatchFacts>.broadcast();
    }
    getMatch(id);
  }

  getTodaysMatches() async {
    _todayController.sink.addStream(_matchFactsRepository.getTodaysMatches());
  }

  getMatch(String id) async {
    _matchController.sink.addStream(_matchFactsRepository.get(id));
  }

  disposeTodaysMatches() {
    _todayController.close();
  }

  disposeMatch() {
    _matchController.close();
  }
}
