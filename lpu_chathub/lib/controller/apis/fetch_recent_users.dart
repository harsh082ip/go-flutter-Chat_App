import 'dart:convert';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:lpu_chathub/shared_res/loggedin_user_singleton.dart';
import 'package:http/http.dart' as http;

class FetchRecents {
  static Future<List<User>> fetchRecentUsers() async {
    // User? user = await LoggedInUserSingleton().getUser();
    // String uid = user!.userid.toString();

    String uid = await LocalKeys.getUid();
    String? token = await LocalKeys.getToken();
    String url = "${BaseUrl.baseUrl}/user/fetchrecentusers/$uid?jwtkey=$token";
    log(url);
    if (token != "") {
      var uri = Uri.parse(url);

      var response = await http.get(uri);
      if (response.statusCode != 200) {
        Get.snackbar('Error Occured :/', response.body,
            backgroundColor: Colors.blue,
            snackPosition: SnackPosition.BOTTOM,
            colorText: Colors.white);
        // return ;
        throw Exception('Failed to load data');
      }

      log(response.body);
      final List<dynamic> data = jsonDecode(response.body);
      return data.map((json) => User.fromJson(json)).toList();
    }

    throw Exception("Token Not Found");
  }
}
