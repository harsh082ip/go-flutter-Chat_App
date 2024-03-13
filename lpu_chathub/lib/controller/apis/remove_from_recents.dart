import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:http/http.dart' as http;

class RemoveFromRecentApis {
  static Future<bool> removeFromRecents(String username) async {
    String uid = await LocalKeys.getUid();
    String token = await LocalKeys.getToken();

    String url =
        "${BaseUrl.baseUrl}/user/removerecentuser/$uid?jwtkey=$token&username=$username";

    if (token != "") {
      var uri = Uri.parse(url);

      var response = await http.get(uri);
      var jsonRes = json.decode(response.body);
      if (response.statusCode != 200) {
        Get.snackbar('Some error Occured :/', jsonRes.toString(),
            snackPosition: SnackPosition.BOTTOM,
            colorText: Colors.white,
            backgroundColor: Colors.red);
        return false;
      }
      Get.snackbar('👍', 'User Removed Successfully',
          snackPosition: SnackPosition.BOTTOM,
          colorText: Colors.white,
          backgroundColor: Colors.blue);
      return true;
    }
    return true;
  }
}
