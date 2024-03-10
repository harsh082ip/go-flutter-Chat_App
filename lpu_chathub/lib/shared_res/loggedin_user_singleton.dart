import 'package:lpu_chathub/models/user_model.dart';

class LoggedInUserSingleton {
  static LoggedInUserSingleton? _instance;
  User? _logged_user;

  LoggedInUserSingleton._internal();
  
  factory LoggedInUserSingleton() {
    if (_instance == null) {
      _instance = LoggedInUserSingleton._internal();
    }
    return _instance!;
  }

  void setUser(User logged_user) {
    _logged_user = logged_user;
  }

  User? getUser() {
    return _logged_user;
  }
}
