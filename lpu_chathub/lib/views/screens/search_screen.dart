import 'dart:developer';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:get/get.dart';
import 'package:lpu_chathub/controller/apis/fetch_recent_users.dart';
import 'package:lpu_chathub/controller/apis/search_api.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:lpu_chathub/views/screens/display_profile.dart';
import 'package:lpu_chathub/widgets/user_list_tile.dart';

class SearchScreen extends StatefulWidget {
  const SearchScreen({Key? key}) : super(key: key);

  @override
  State<SearchScreen> createState() => _SearchScreenState();
}

class _SearchScreenState extends State<SearchScreen> {
  final TextEditingController searchedText = TextEditingController();
  bool isSearchTriggered = false;
  User? user;
  late Future<List<User>> _futureUsers = Future.value([]);

  void refreshScreen() {
    log('Refreshing...');
    setState(() {
      _futureUsers = FetchRecents.fetchRecentUsers();
    }); // Trigger state change to refresh the screen
  }

  @override
  void initState() {
    // TODO: implement initState
    fetchRecentsUsersCall();
    super.initState();
  }

  fetchRecentsUsersCall() {
    // WidgetsBinding.instance.addPostFrameCallback((_) {
    _futureUsers = FetchRecents.fetchRecentUsers();
    // });
  }

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
                  child: SingleChildScrollView(
                    // Added SingleChildScrollView
                    child: TextFormField(
                      controller: searchedText,
                      style: const TextStyle(color: Colors.white),
                      decoration: InputDecoration(
                        hintText: 'Search',
                        hintStyle:
                            TextStyle(color: Colors.white, fontSize: 18.0),
                        prefixIcon: Visibility(
                          visible: isSearchTriggered,
                          child: InkWell(
                            onTap: () {
                              setState(() {
                                isSearchTriggered = false;
                              });
                              searchedText.clear();
                            },
                            child: Icon(
                              CupertinoIcons.multiply,
                              color: Colors.white,
                            ),
                          ),
                        ),
                        suffixIcon: InkWell(
                          onTap: () async {
                            if (searchedText.text != "") {
                              setState(() {
                                isSearchTriggered = true;
                              });
                              // Show the default dialog at the bottom
                              Get.defaultDialog(
                                title: 'Please Wait...',
                                titleStyle: TextStyle(color: Colors.white),
                                backgroundColor: Colors.blue,
                                content: Container(
                                  alignment: Alignment
                                      .bottomCenter, // Align content to the bottom
                                  child: CircularProgressIndicator(),
                                ),
                              );

                              user = await SearchApis.searchUser(
                                  searchedText.text);
                              if (user != null) {
                                log(user!.email.toString());
                                log(user!.name.toString());
                                log(user!.profilePicUrl.toString());
                              }
                            } else {
                              print("Empty");
                            }
                          },
                          child: Icon(
                            Icons.search,
                            color: Colors.white,
                          ),
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
                child: isSearchTriggered != true
                    ? Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          const Text('Recents ',
                              style: TextStyle(
                                  color: Colors.white, fontSize: 22.0)),

                          // Display Recents chat Tiles
                          // _futureUsers != Future.value([]) ?
                          FutureBuilder<List<User>>(
                              future: _futureUsers,
                              builder: ((context, snapshot) {
                                if (snapshot.connectionState ==
                                    ConnectionState.waiting) {
                                  return Center(
                                    child: CircularProgressIndicator(
                                        color: Colors.blue),
                                  );
                                }
                                // Check if snapshot.data is null
                                if (snapshot.data == null ||
                                    snapshot.data!.isEmpty) {
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
                                  ));
                                }

                                final users = snapshot.data!;

                                return Expanded(
                                  child: ListView.builder(
                                      // physics: BouncingScrollPhysics(),
                                      itemCount: users.length,
                                      itemBuilder: (context, index) {
                                        final fetchuser = users[index];
                                        return UserListTile(
                                          user: fetchuser,
                                          name: fetchuser.name.toString(),
                                          userEmail: fetchuser.email.toString(),
                                          profilePicUrl:
                                              fetchuser.profilePicUrl,
                                          displayTrailing: true,
                                          refreshCallback: () =>
                                              refreshScreen(),
                                        );
                                      }),
                                );
                              }))
                          // : Container()
                        ],
                      )
                    : Column(
                        crossAxisAlignment: CrossAxisAlignment.start,
                        children: [
                          if (user !=
                              null) // Add a null check before accessing user properties
                            InkWell(
                              onTap: () {
                                print('Pressed...');
                                Get.to(() => DisplayProfile(
                                      user: user,
                                    ));
                              },
                              child: UserListTile(
                                user: user,
                                displayTrailing: false,
                                name: user!.name.toString(),
                                userEmail: user!.email.toString(),
                                profilePicUrl: user!.profilePicUrl.toString(),
                                refreshCallback: () => refreshScreen(),
                              ),
                            ),
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
