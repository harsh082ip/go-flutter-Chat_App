import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:jwt_decoder/jwt_decoder.dart';
import 'package:lpu_chathub/views/authentication/login.dart';
import 'package:lpu_chathub/views/screens/home.dart';
import 'package:shared_preferences/shared_preferences.dart';

void main() {
  runApp(const MyApp());
}

class MyApp extends StatefulWidget {
  const MyApp({Key? key}) : super(key: key);

  @override
  State<MyApp> createState() => _MyAppState();
}

class _MyAppState extends State<MyApp> {
  late String token = "";
  late bool isExpired = false;

  @override
  void initState() {
    super.initState();
    isLoggedIn();
  }

  isLoggedIn() async {
    SharedPreferences prefs = await SharedPreferences.getInstance();
    print("This is what i have ${prefs.getString('token')}");
    token = prefs.getString('token') ?? ""; // Use null-aware operator
    if(token != "") {
      if (!JwtDecoder.isExpired(token)) {
      setState(() {
        isExpired = false;
      });
    } else {
      setState(() {
        isExpired = true;
      });
    }
    }
    else {
      setState(() {
        isExpired = true;
      });
    }

  }

  @override
  Widget build(BuildContext context) {
    Widget home;
    if (isExpired) {
      home = LoginPage();
    } else {
      home = HomePage();
    }
    return GetMaterialApp(
      debugShowCheckedModeBanner: false,
      home: home,
    );
  }
}
