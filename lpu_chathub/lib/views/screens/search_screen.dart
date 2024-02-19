import 'package:flutter/material.dart';
import 'package:get/get.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({Key? key}) : super(key: key);

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromARGB(255, 16, 26, 36),
      body: SafeArea(
        child: Column(
          crossAxisAlignment: CrossAxisAlignment.start,
          children: [
            Row(
              children: [
                IconButton(
                  onPressed: () {
                    Get.back();
                  },
                  icon: const Icon(Icons.arrow_back),
                  iconSize: 27.0,
                  color: Colors.grey,
                ),
                Container(
                  height: MediaQuery.of(context).size.height * 0.06,
                  width: MediaQuery.of(context).size.width * 0.82,
                  padding: const EdgeInsets.symmetric(horizontal: 18.0),
                  decoration: BoxDecoration(
                      borderRadius: BorderRadius.circular(20.0),
                      color: const Color.fromARGB(255, 32, 41, 58)),
                  margin: const EdgeInsets.symmetric(
                      horizontal: 10.0, vertical: 18.0),
                  child: SingleChildScrollView( // Added SingleChildScrollView
                    child: TextFormField(
                      style: const TextStyle(color: Colors.white),
                      decoration:  const InputDecoration(
                        hintText: 'Search',
                        hintStyle:  TextStyle(color: Colors.white, fontSize: 18.0),
                        prefixIcon:  Icon(
                          Icons.search,
                          color: Colors.white,
                        ),
                        border: InputBorder.none, // Removed OutlineInputBorder
                      ),
                    ),
                  ),
                ),
              ],
            ),
            Expanded(
              child: Container(
                margin: const EdgeInsets.all(15.0),
                width: MediaQuery.of(context).size.width,
                child: const Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    Text('Recents ', style: TextStyle(color: Colors.white, fontSize: 22.0))
                  ],
                ),
              ),
            )
          ],
        ),
      ),
    );
  }
}
