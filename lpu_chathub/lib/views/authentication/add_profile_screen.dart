import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'dart:developer';
import 'dart:io';
import 'package:lpu_chathub/controller/profile_pic_controller.dart';
import 'package:lpu_chathub/views/screens/home.dart';

class AddProfileScreen extends StatefulWidget {
  const AddProfileScreen({Key? key});

  @override
  State<AddProfileScreen> createState() => _AddProfileScreenState();
}

class _AddProfileScreenState extends State<AddProfileScreen> {
  File? imgpath;
  bool isImageloading = false;
  final TextEditingController usernameController = TextEditingController();

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromARGB(255, 16, 26, 36),
      appBar: AppBar(
        backgroundColor: const Color.fromARGB(255, 41, 47, 54),
        elevation: 0,
        centerTitle: true,
        title: const Text(
          'Profile Info',
          style: TextStyle(color: Colors.white, fontSize: 25.0, fontWeight: FontWeight.bold),
        ),
      ),
      body: SingleChildScrollView(
        child: Container(
          padding: const EdgeInsets.all(20.0),
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.center,
            children: [
              const SizedBox(height: 20.0),
              const Text(
                'Please provide your name and a profile photo',
                style: TextStyle(color: Colors.white, fontSize: 18.0),
              ),
              const SizedBox(height: 40.0),
              Stack(
                children: [
                  Container(
                    decoration: const BoxDecoration(
                      color: Colors.grey,
                      shape: BoxShape.circle,
                    ),
                    height: 220.0,
                    width: 220.0,
                    child: imgpath == null
                        ? Image.asset('assets/images/user.png')
                        : ClipOval(
                            child: Image.file(
                              imgpath!,
                              fit: BoxFit.cover,
                            ),
                          ),
                  ),
                  Positioned(
                    bottom: -7,
                    right: 12,
                    child: Container(
                      child: IconButton(
                        onPressed: () async {
                          setState(() {
                            isImageloading = true;
                          });
                          final newImgPath = await Profile_Pic.pickFile();
                          if (newImgPath != null) {
                            setState(() {
                              imgpath = File(newImgPath.path);
                              isImageloading = false;
                            });
                          } else {
                            setState(() {
                              isImageloading = false;
                            });
                          }
                        },
                        icon: const Icon(
                          Icons.edit,
                          color: Colors.white,
                          size: 32.0,
                        ),
                      ),
                    ),
                  ),
                ],
              ),
              const SizedBox(height: 20.0),
              Container(
                margin: const EdgeInsets.symmetric(horizontal: 20.0),
                child: TextFormField(
                  controller: usernameController,
                  style: const TextStyle(color: Colors.white),
                  decoration: const InputDecoration(
                    labelText: 'Username',
                    labelStyle:  TextStyle(color: Colors.white, fontSize: 20.0),
                    focusedBorder:  UnderlineInputBorder(
                      borderSide: BorderSide(color: Colors.white),
                    ),
                    enabledBorder:  UnderlineInputBorder(
                      borderSide: BorderSide(color: Colors.white),
                    ),
                  ),
                ),
              ),
              SizedBox(height: MediaQuery.of(context).size.height*0.09,),
              Container(
                padding: const EdgeInsets.all(20.0),
                width: MediaQuery.of(context).size.width,
                child: ElevatedButton(
                  style: ElevatedButton.styleFrom(
                    backgroundColor: Colors.white,
                    shape: const StadiumBorder(),
                  ),
                  onPressed: () {
                    log(imgpath!.path+usernameController.text);
                    Get.off(HomePage());
                  },
                  child: const Text(
                    'Next',
                    style: TextStyle(fontSize: 20.0, color: Color.fromARGB(255, 16, 26, 36),fontWeight: FontWeight.bold ),
                  ),
                ),
              ),
              const SizedBox(height: 20.0),
            ],
          ),
        ),
      ),
    );
  }
}
