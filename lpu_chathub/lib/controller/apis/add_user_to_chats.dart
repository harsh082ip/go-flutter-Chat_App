import 'dart:convert';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:http/http.dart' as http;

class AddToHome {
  static Future<void> addUserToChats(String clientToAdd) async {
    String? token = await LocalKeys.getToken();
    String? userId = await LocalKeys.getUid();
    String? presentClient = await LocalKeys.getUsername();
    String url =
        "${BaseUrl.baseUrl}/user/addtorecentchats/$userId?jwtkey=$token&client1username=$presentClient&client2username=$clientToAdd";
    print(url);
    if (token != "") {
      var uri = Uri.parse(url);

      var response = await http.get(uri);
      var jsonRes = json.decode(response.body);
      if (response.statusCode != 200) {
        Get.snackbar('Cannot Add User to chats :/', jsonRes.toString(),
            snackPosition: SnackPosition.BOTTOM,
            colorText: Colors.white,
            backgroundColor: Colors.red);
        return;
      }

      Get.snackbar(
          'User added to Your Chats 😁', 'Please Navigate to Home Screen Now',
          snackPosition: SnackPosition.BOTTOM,
          colorText: Colors.white,
          backgroundColor: Colors.blue);
      return;
    }
    log('token empty');
    return;
  }
}
