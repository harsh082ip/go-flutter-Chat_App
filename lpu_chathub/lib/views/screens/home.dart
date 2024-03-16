import 'dart:developer';

import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:liquid_pull_to_refresh/liquid_pull_to_refresh.dart';
import 'package:lpu_chathub/controller/apis/fetch_home.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:lpu_chathub/views/screens/chats/chat_screen.dart';
import 'package:lpu_chathub/views/screens/search_screen.dart';
import 'package:lpu_chathub/widgets/user_list_tile.dart';

class HomePage extends StatefulWidget {
  const HomePage({super.key});

  @override
  State<HomePage> createState() => _HomePageState();
}

class _HomePageState extends State<HomePage> {
  late Future<List<User>> _futureUsers = Future.value([]);
  late ScrollController _scrollController;
  @override
  void initState() {
    // TODO: implement initState
    fetchHomeDataFromApi();
    // _scrollController = ScrollController();
    // _scrollController.addListener(_scrollListener);
    super.initState();
  }

  fetchHomeDataFromApi() {
    setState(() {
      _futureUsers = FetchHomeApis.fetchHomeData();
    });
  }

  void _scrollListener() {
    if (_scrollController.offset <=
            _scrollController.position.minScrollExtent &&
        !_scrollController.position.outOfRange) {
      // Scrolled to the top
      _refreshScreen();
    }
  }

  Future<void> _refreshScreen() async {
    log('Refresh Required');
    setState(() {
      _futureUsers = FetchHomeApis.fetchHomeData();
    });
    // return await Future.delayed(Duration(seconds: 2));
  }

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromARGB(255, 16, 26, 36),
      appBar: AppBar(
        title: Text(
          'CyberSec',
          style: TextStyle(color: Colors.white),
        ),
        backgroundColor: const Color.fromARGB(255, 32, 41, 58),
        actions: [
          IconButton(
            onPressed: () {
              Get.to(SearchScreen());
            },
            icon: const Icon(Icons.search),
            iconSize: 30,
            color: Colors.white,
          ),
          IconButton(
            onPressed: () {},
            icon: const Icon(Icons.notification_important),
            iconSize: 30,
            color: Colors.white,
          ),
          IconButton(
            onPressed: () {},
            icon: const Icon(Icons.more_vert),
            iconSize: 30,
            color: Colors.white,
          ),
        ],
      ),
      body: Container(
        child: Column(
          children: [
            Container(
              alignment: Alignment.topLeft,
              margin: EdgeInsets.only(top: 10.0, left: 15.0, bottom: 15.0),
              child: Text(
                'Your Chats',
                style: TextStyle(
                  color: Colors.white,
                  fontSize: 18.0,
                  fontWeight: FontWeight.bold,
                ),
              ),
            ),
            Expanded(
              child: FutureBuilder<List<User>>(
                future: _futureUsers,
                builder: (context, snapshot) {
                  if (snapshot.connectionState == ConnectionState.waiting) {
                    return Center(
                      child: CircularProgressIndicator(color: Colors.blue),
                    );
                  }

                  if (snapshot.data == null || snapshot.data!.isEmpty) {
                    return Center(
                      child: Text(
                        "No User Found",
                        style: TextStyle(color: Colors.white),
                      ),
                    );
                  } else if (snapshot.hasError) {
                    return Center(
                      child: Text(
                        'Error: ${snapshot.error}',
                        style: TextStyle(color: Colors.white),
                      ),
                    );
                  }

                  final users = snapshot.data!;
                  return LiquidPullToRefresh(
                    animSpeedFactor: 6.0,
                    onRefresh: _refreshScreen,
                    color: Color.fromARGB(255, 44, 50, 56),
                    height: 200,
                    showChildOpacityTransition: true,
                    // animSpeedFactor: 10,
                    child: ListView.builder(
                      itemCount: users.length,
                      itemBuilder: (context, index) {
                        final fetchuser = users[index];
                        return InkWell(
                          onTap: () => Get.to(ChatScreen(user: fetchuser)),
                          child: UserListTile(
                            user: fetchuser,
                            name: fetchuser.name.toString(),
                            userEmail: fetchuser.email.toString(),
                            profilePicUrl: fetchuser.profilePicUrl,
                            displayTrailing: false,
                            refreshCallback: () => _refreshScreen(),
                          ),
                        );
                      },
                    ),
                  );
                },
              ),
            ),
            // Text(
            //   'data',
            //   style: TextStyle(color: Colors.white),
            // )
          ],
        ),
      ),
    );
  }
}
