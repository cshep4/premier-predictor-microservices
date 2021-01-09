import 'dart:async';

import 'package:premier_predictor/database/database.dart';
import 'package:premier_predictor/model/match_facts.dart';
import 'package:premier_predictor/model/todo.dart';
import 'package:sqflite/sqflite.dart';

class MatchFactsDao {
  final DatabaseProvider _dbProvider;

  MatchFactsDao(this._dbProvider);

  Future<int> create(Todo todo) async {
    final db = await _dbProvider.database;
    var result = db.insert(matchFactsTABLE, todo.toDatabaseJson());
    return result;
  }

  Future<List<MatchFacts>> getTodaysMatches(
      {List<String> columns, int limit}) async {
    final db = await _dbProvider.database;

    List<Map<String, dynamic>> result = await db.query(
      matchFactsTABLE,
      columns: columns,
      where: '''
            matchDate >= date('now') AND 
            matchDate < date('now', '+1 days')
            ''',
      limit: limit,
    );

    return result.isNotEmpty
        ? result.map((item) => MatchFacts.fromJson(item)).toList()
        : [];
  }

  Future<MatchFacts> get(String id, {List<String> columns}) async {
    final db = await _dbProvider.database;

    List<Map<String, dynamic>> result = await db.query(
      matchFactsTABLE,
      columns: columns,
      where: "id = ?",
      whereArgs: [id],
      limit: 1,
    );

    if (result.isEmpty) {
      return null;
    }

    return result.map((item) => MatchFacts.fromJson(item)).toList()[0];
  }

  Future<int> upsert(MatchFacts matchFacts) async {
    final db = await _dbProvider.database;

    return await db.insert(
      matchFactsTABLE,
      matchFacts.toJson(),
      conflictAlgorithm: ConflictAlgorithm.replace,
    );
  }

  Future<int> delete(int id) async {
    final db = await _dbProvider.database;

    return await db.delete(
      matchFactsTABLE,
      where: 'id = ?',
      whereArgs: [id],
    );
  }

  Future deleteAllTodos() async {
    final db = await _dbProvider.database;
    return await db.delete(matchFactsTABLE);
  }
}

String createMatchFactsTableQuery = '''
  CREATE TABLE $matchFactsTABLE (
      id TEXT PRIMARY KEY,
      compId TEXT,
      formattedDate TEXT,
      season TEXT,
      week TEXT,
      venue TEXT,
      venueId TEXT,
      venueCity TEXT,
      venueLatitude TEXT,
      venueLongitude TEXT,
      venueCountry TEXT,
      status TEXT,
      timer TEXT,
      time TEXT,
      localTeamId TEXT,
      localTeamName TEXT,
      localTeamScore TEXT,
      visitorTeamId TEXT,
      visitorTeamName TEXT,
      visitorTeamScore TEXT,
      htScore TEXT,
      ftScore TEXT,
      etScore TEXT,
      penaltyLocal TEXT,
      penaltyVisitor TEXT,
      events BLOB,
      commentary BLOB,
      matchDate TEXT
  )
''';
