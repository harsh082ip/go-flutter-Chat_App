import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:http/http.dart' as http;

class RoomID {
  static Future<String> getRoomIdByUsernames(String username2) async {
    String? username1 = await LocalKeys.getUsername();
    String? token = await LocalKeys.getToken();
    String url =
        "${BaseUrl.baseUrl}/user/getroomidbyusernames?username1=$username1&username2=$username2&jwtkey=$token";

    if (token != "") {
      var uri = Uri.parse(url);
      var response = await http.get(uri);
      var jsonRes = jsonDecode(response.body);
      if (response.statusCode != 200) {
        Get.snackbar('Error Occured :/', 'Cannot Fetch Room Id',
            snackPosition: SnackPosition.BOTTOM, colorText: Colors.white);
        return "";
      }
      return jsonRes['roomID'].toString();
    }
    throw Exception('Auth Token Not Found! :/');
  }
}
