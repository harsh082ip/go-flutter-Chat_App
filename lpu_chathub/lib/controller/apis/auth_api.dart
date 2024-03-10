import 'dart:convert';
import 'dart:developer';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:lpu_chathub/models/user_model.dart';
import 'package:lpu_chathub/shared_res/loggedin_user_singleton.dart';
import 'package:lpu_chathub/views/authentication/login.dart';
import 'package:lpu_chathub/views/screens/home.dart';
import 'package:shared_preferences/shared_preferences.dart';

/// Class containing API methods related to authentication.
class AuthApis {


  static String baseUrl = "http://192.168.135.132:8006";

  /// Method for signing up a user.
  ///
  /// Parameters:
  /// - [imgPath]: Path to the user's profile picture.
  /// - [name]: User's name.
  /// - [email]: User's email address.
  /// - [password]: User's password.
  /// - [username]: User's chosen username.
  static Future<void> signUp(String imgPath, String name, String email, String password, String username) async{
    
    // Change the baseUrl Accordingly
    
    var url = Uri.parse("$baseUrl/users/signup");
    print(url.toString());

    var request = http.MultipartRequest("POST", url);
    request.fields['data'] = json.encode({
      "name": name,
      "email": email,
      "password": password,
      "username": username,
    });

       // Check if the image path is not empty
    if (imgPath != "") {
      // Create a File object from the image path
      File imageFile = File(imgPath);
      // Open a byte stream to read the image file
      final stream = http.ByteStream(imageFile.openRead());
      // Get the length of the image file
      var length = await imageFile.length();
      // Create a MultipartFile object with the image file stream
      var multipartFile = http.MultipartFile(
        'profile_picture', // Field name for the file
        stream, // Byte stream of the file
        length, // Length of the file
        filename: imageFile.path.split('/').last // Filename extracted from the file path
      );

      // Add the multipart file to the request
      request.files.add(multipartFile);
    }

    try {
      var streamedResponse = await request.send();
      var response = await http.Response.fromStream(streamedResponse);
      if (response.statusCode == 200) {
        // Successful sign-up
        print('Sign-up successful');
        Get.defaultDialog(
          title: 'Sign Up Successful',
          content: Text('Please login now'),
          confirm: ElevatedButton(onPressed: () {}, child: Text('OK'))
        );
        Get.offAll(LoginPage());
        Get.snackbar('Sign Up Successful', 'Please login now', backgroundColor: Colors.blue, snackPosition: SnackPosition.BOTTOM, colorText: Colors.white, );
      } else {
        // Failed sign-up
        print('Sign-up failed with status code: ${response.statusCode}');
        print('Response body: ${response.body}');
        Get.back();
        Get.defaultDialog(
          title: 'Error While SignUp',
          content: Center(
            child: Text(response.body),
          )
        );
      }
    } catch (error) {
      print('Error occurred during sign-up: $error');
      Get.back();
      Get.defaultDialog(
          title: 'Error While SignUp',
          content: Center(
            child: Text(error.toString()),
          )
        );
    }
  }


  static Future<void> login(String username, String password) async{

    var url = Uri.parse("$baseUrl/users/login");
    print(url.toString());
   
    var response = await http.post(url,
      headers: <String, String> {
        'Content-Type': 'application/json; charset=UTF-8',
      },
      body: json.encode({
        "username": username,
        "password": password,
      }));

      if (response.statusCode == 200) {
        print(response.body);
        SharedPreferences prefs = await SharedPreferences.getInstance();
        var res = jsonDecode(response.body);
        var token = res['Jwt_Token'];
        prefs.setString('token', token);


    //     Map<String, dynamic> jsonResponse = jsonDecode(response.body);
    // User user = User.fromJson(jsonResponse);
    // LoggedInUserSingleton().setUser(user);
    
    // User? current = LoggedInUserSingleton().getUser();
    // if (current != null) {
    //   log(current.email);
    // }
    
        Get.offAll(HomePage());
      } else {

        Get.back();
        Get.defaultDialog(
          title: 'Error While Login',
          content: Center(
            child: Text('Error: ${response.body}'),
          ),
        );
        print('Failed to login: ${response.body}');
        print(response.statusCode);
      }
  }
}
