import 'dart:convert';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:http/http.dart' as http;

class FetchHomeApis {
  static Future<List<User>> fetchHomeData() async {
    String? uid = await LocalKeys.getUid();
    String? token = await LocalKeys.getToken();
    String url = "${BaseUrl.baseUrl}/user/fetchhomedata/$uid?jwtkey=$token";
    print(url);
    if (token != "") {
      var uri = Uri.parse(url);
      var response = await http.get(uri);
      var jsonRes = json.decode(response.body);
      if (response.statusCode != 200) {
        Get.snackbar('Failed to Fetch Home Data', jsonRes.toString(),
            snackPosition: SnackPosition.BOTTOM,
            colorText: Colors.white,
            backgroundColor: Colors.red,
            duration: Duration(seconds: 2));
        throw Exception('Failed to load data');
      }

      log(response.body);
      final List<dynamic> data = jsonDecode(response.body);
      return data.map((json) => User.fromJson(json)).toList();
    }

    throw Exception('Token Not Found');
  }
}
