import 'package:flutter/material.dart';
import 'package:shared_preferences/shared_preferences.dart';

class AppColors {
  static const Color primaryColor = Color.fromARGB(255, 16, 26, 36);
  static const Color appBarColor = Color.fromARGB(255, 32, 41, 58);
}

class BaseUrl {
  static const String baseUrl = "http://192.168.135.132:8006";
}

class LocalKeys {
  static Future<String> getToken() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String token = prefs.getString('token') ?? "";

    return token;
  }

  static Future<String> getUid() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String uid = prefs.getString('uid') ?? "";

    return uid;
  }

  static Future<String> getUsername() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String username = prefs.getString('username') ?? "";

    return username;
  }

  static Future<String> getProfilePicUrl() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String profile_pic_url = prefs.getString('profile_pic_url') ?? "";

    return profile_pic_url;
  }
}
