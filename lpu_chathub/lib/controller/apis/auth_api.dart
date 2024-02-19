import 'dart:convert';
import 'dart:io';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:http/http.dart' as http;
import 'package:lpu_chathub/views/screens/home.dart';

/// Class containing API methods related to authentication.
class AuthApis {

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
    var baseUrl = "http://192.168.117.132:8006/";
    var url = Uri.parse("$baseUrl/signup");

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
        Get.offAll(HomePage());
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
}
