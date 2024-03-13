import 'dart:developer';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/controller/apis/add_to_recents.dart';
import 'package:lpu_chathub/controller/apis/add_user_to_chats.dart';
import 'package:lpu_chathub/models/user_model.dart';

class DisplayProfile extends StatefulWidget {
  final User? user;
  DisplayProfile({super.key, required this.user});

  @override
  State<DisplayProfile> createState() => _DisplayProfileState();
}

class _DisplayProfileState extends State<DisplayProfile> {
  @override
  void initState() {
    // TODO: implement initState
    addToRecentlyViewed();

    super.initState();
  }

  addToRecentlyViewed() {
    WidgetsBinding.instance.addPostFrameCallback((_) {
      AddToRecentApis.addToRecents(widget.user!.username.toString());
    });
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        title: Text(
          'Profile',
          style: TextStyle(color: Colors.white),
        ),
        toolbarHeight: 50,
        centerTitle: true,
        backgroundColor: AppColors.primaryColor,
        iconTheme: IconThemeData(color: Colors.white),
        // elevation: 8,
      ),
      body: Container(
        height: MediaQuery.of(context).size.height,
        width: MediaQuery.of(context).size.width,
        color: AppColors.primaryColor,
        child: SingleChildScrollView(
          child: Column(
            children: [
              Container(
                margin: EdgeInsets.only(top: 20.0),
                decoration: BoxDecoration(
                    borderRadius: BorderRadius.all(Radius.circular(50.0))),
                height: 200.0,
                width: 200.0,
                child: ClipOval(
                    child: Image.network(
                  widget.user!.profilePicUrl.toString(),
                  fit: BoxFit.cover,
                )),
              ),

              // User Details
              Container(
                // color: Colors.red,
                // width: MediaQuery.of(context).size.width,
                margin: EdgeInsets.only(left: 20.0),
                alignment: Alignment.topLeft,
                child: Column(
                  mainAxisAlignment: MainAxisAlignment.start,
                  children: [
                    Align(
                      alignment: Alignment.topLeft,
                      child: Text(
                        'Name: ' + widget.user!.name.toString(),
                        style: TextStyle(color: Colors.white, fontSize: 22.0),
                      ),
                    ),
                    Align(
                      alignment: Alignment.topLeft,
                      child: Text(
                        'E-mail: ' + widget.user!.email.toString(),
                        style: TextStyle(color: Colors.white, fontSize: 22.0),
                      ),
                    ),
                    Align(
                      alignment: Alignment.topLeft,
                      child: Text(
                        'Username: ' + widget.user!.username.toString(),
                        style: TextStyle(color: Colors.white, fontSize: 22.0),
                      ),
                    ),
                    Align(
                      alignment: Alignment.topLeft,
                      child: Text(
                        'UserId: ' + widget.user!.userid.toString(),
                        style: TextStyle(color: Colors.white, fontSize: 22.0),
                      ),
                    ),
                    SizedBox(
                      height: 30.0,
                    ),
                    ElevatedButton(
                        onPressed: () {
                          log('Button Pressed');
                          Get.snackbar('Adding', '',
                              snackPosition: SnackPosition.BOTTOM,
                              colorText: Colors.white,
                              backgroundColor: Colors.blue);
                          AddToHome.addUserToChats(
                              widget.user!.username.toString());
                        },
                        child: Text('Add to Chats'))
                  ],
                ),
              )
            ],
          ),
        ),
      ),
    );
  }
}
