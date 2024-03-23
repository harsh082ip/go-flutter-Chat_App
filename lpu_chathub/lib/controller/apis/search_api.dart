import 'dart:convert';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:lpu_chathub/shared_res/loggedin_user_singleton.dart';
import 'package:shared_preferences/shared_preferences.dart';

class SearchApis {
  static String baseUrl = "http://192.168.135.132:8006";

  static Future<User?> searchUser(String username) async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    String token = prefs.getString('token') ?? "";

    if (token != "") {
      Map<String, dynamic> decodeToken = JwtDecoder.decode(token);
      String? logged_in_username = decodeToken['username'];
      print(logged_in_username);
      log(decodeToken.toString());
      if (logged_in_username != username) {
        var uri = Uri.parse(
            "${BaseUrl.baseUrl}/user/getuserbyusername/$username?jwtkey=$token");
        print(uri.toString());

        var response = await http.get(uri);
        if (response.statusCode == 200) {
          Get.back();

          final jsonUser = jsonDecode(response.body);

          User user = User.fromJson(jsonUser);
          log(jsonUser.toString());

          return user;
        } else {
          print(response.body);
          print(response.body);
          Get.back();

          final resDecode = jsonDecode(response.body);

          Get.snackbar('Error whie Searching :/', resDecode['error'],
              backgroundColor: Colors.blue,
              snackPosition: SnackPosition.BOTTOM,
              colorText: Colors.white);
          return null;
        }
      } else {
        Get.back();
        print('cannot search yourself');
        Get.snackbar('Are you Serious?🙄', 'You cannot search yourself',
            backgroundColor: Colors.blue,
            snackPosition: SnackPosition.BOTTOM,
            colorText: Colors.white);
      }
    } else {
      print("Unauthorized: Token not found");

      print('hey');
      return null;
    }
  }
}
