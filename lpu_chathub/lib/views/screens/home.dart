import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/views/screens/search_screen.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromARGB(255, 16, 26, 36),
      appBar: AppBar(
        backgroundColor: const Color.fromARGB(255, 32, 41, 58),
        // title: const Text('Chat Screen'),
        actions: [
          IconButton(onPressed: (){
            Get.to(SearchScreen());
          }, icon: const Icon(Icons.search), iconSize: 30,color: Colors.white,),
          IconButton(onPressed: (){}, icon: const Icon(Icons.notification_important),iconSize: 30,color: Colors.white),
          IconButton(onPressed: (){}, icon: const Icon(Icons.more_vert), iconSize: 30,color: Colors.white,),
        ],
      ),
    );
  }
}