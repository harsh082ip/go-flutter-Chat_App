import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AppColors {
  static const Color primaryColor = Color.fromARGB(255, 16, 26, 36);
}

class BaseUrl {
  static const String baseUrl = "http://192.168.135.132:8006";
}

class Token {
  static Future<String> getToken() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String token = prefs.getString('token') ?? "";

    return token;
  }
}
