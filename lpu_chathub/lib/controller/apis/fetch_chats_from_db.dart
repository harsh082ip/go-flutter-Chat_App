import 'dart:convert';
import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:lpu_chathub/const.dart';
import 'package:http/http.dart' as http;
import 'package:lpu_chathub/models/conversation_model.dart';

class FetchChats {
  static Future<ApiResponse> fetchChatsFromDB(String roomID) async {
    String? token = await LocalKeys.getToken();

    String url =
        "${BaseUrl.baseUrl}/user/fetchchatsfromdatabase/$roomID?jwt_key=$token";

    if (token != "") {
      var uri = Uri.parse(url);

      var response = await http.get(uri);

      if (response.statusCode != 200) {
        log(response.body);
        return ApiResponse(messages: []);
      }

      final jsonBody = jsonDecode(response.body);
      ApiResponse apiResponse = ApiResponse.fromJson(jsonBody);
      // log("ApiResponse" + apiResponse.messages.toString());
      return apiResponse;
    } else {
      return ApiResponse(messages: []);
    }
  }
}
