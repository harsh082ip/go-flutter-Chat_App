import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:lpu_chathub/shared_res/loggedin_user_singleton.dart';
import 'package:http/http.dart' as http;

class AddToRecentApis {
  static Future<void> addToRecents(String username) async {
    Get.snackbar('Here', 'Calling...');
    String? token = await Token.getToken();
    User? currentUser = LoggedInUserSingleton().getUser();
    log(currentUser!.userid.toString());
    String url =
        "${BaseUrl.baseUrl}/user/addtorecentlyviewed/${currentUser.userid.toString()}?jwt_key=$token&username=$username";
    log("Url" + url);

    if (token != "") {
      var uri = Uri.parse(url);

      var response = await http.get(uri);

      if (response.statusCode != 200) {
        Get.snackbar('Error Occured :/', response.body,
            backgroundColor: Colors.blue,
            snackPosition: SnackPosition.BOTTOM,
            colorText: Colors.white);
      } else {
        log('Done');
      }
    }
  }
}
