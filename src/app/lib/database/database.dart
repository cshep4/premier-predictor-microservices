import 'dart:async';
import 'dart:io';

import 'package:path/path.dart';
import 'package:path_provider/path_provider.dart';
import 'package:premier_predictor/dao/fixture_dao.dart';
import 'package:premier_predictor/dao/match_facts_dao.dart';
import 'package:sqflite/sqflite.dart';

final matchFactsTABLE = 'MatchFacts';
final fixtureTABLE = 'Fixture';
final todoTABLE = 'Todo';

class DatabaseProvider {
  Database _database;

  Future<Database> get database async {
    if (_database != null) return _database;
    _database = await createDatabase();
    return _database;
  }

  createDatabase() async {
    Directory documentsDirectory = await getApplicationDocumentsDirectory();
    //"PremierPredictor.db is our database instance name
    String path = join(documentsDirectory.path, "PremierPredictor.db");

    var database = await openDatabase(
      path,
      version: 1,
      onCreate: initDB,
      onUpgrade: onUpgrade,
    );
    return database;
  }

  //This is optional, and only used for changing DB schema migrations
  void onUpgrade(Database database, int oldVersion, int newVersion) {
    if (newVersion > oldVersion) {}
  }

  void initDB(Database database, int version) async {
    await database.execute("CREATE TABLE $todoTABLE ("
        "id INTEGER PRIMARY KEY, "
        "description TEXT, "
        /*SQLITE doesn't have boolean type
        so we store isDone as integer where 0 is false
        and 1 is true*/
        "is_done INTEGER "
        ")");

    await database.execute(createMatchFactsTableQuery);

    await database.execute(createFixtureTableQuery);
  }
}
