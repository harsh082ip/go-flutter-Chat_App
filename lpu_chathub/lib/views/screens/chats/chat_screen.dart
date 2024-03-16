import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/controller/apis/get_room_id.dart';
import 'package:lpu_chathub/models/user_model.dart';
import 'package:web_socket_channel/web_socket_channel.dart';
import 'package:web_socket_channel/status.dart' as status;

class ChatScreen extends StatefulWidget {
  final User user;
  const ChatScreen({super.key, required this.user});

  @override
  State<ChatScreen> createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  String roomID = "ROOMID";
  late WebSocketChannel channel;
  List<String> messages = [];
  String username = "username";
  bool isCopyTriggered = false;

  @override
  void initState() {
    // TODO: implement initState
    getRoomIDandEnableWebSocketConn();
    super.initState();
  }

  getRoomIDandEnableWebSocketConn() async {
    roomID = await RoomID.getRoomIdByUsernames(widget.user.username.toString());
    runWebSocket(roomID);
  }

  runWebSocket(String roomId) async {
    String? uid = await LocalKeys.getUid();
    username = await LocalKeys.getUsername();
    String url =
        "ws://192.168.135.132:8006/ws/joinroom/$roomID?uid=$uid&username=$username";
    print(url);
    final wsUrl = Uri.parse(url);
    channel = WebSocketChannel.connect(wsUrl);
    roomID = "RoomID";
    channel.stream.listen((event) {
      setState(() {
        messages.insert(0, event);
      });
    });
  }

  @override
  void dispose() {
    channel.sink.close();
    super.dispose();
  }

  // Future<String> getUsername() async {
  //   String? presntUsername = await LocalKeys.getUsername();
  //   return presntUsername;
  // }

  final TextEditingController messageController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
        appBar: AppBar(
          iconTheme: IconThemeData(color: Colors.white),
          backgroundColor: AppColors.appBarColor,
          actions: [
            isCopyTriggered
                ? IconButton(
                    onPressed: () {},
                    icon: Icon(
                      Icons.copy,
                      color: Colors.white,
                    ))
                : Container()
          ],
          title: Row(
            children: [
              CircleAvatar(
                backgroundColor: AppColors.primaryColor,
                backgroundImage: NetworkImage(widget.user.profilePicUrl),
              ),
              SizedBox(
                width: 12.0,
              ),
              Text(
                widget.user.name.toString(),
                style: TextStyle(color: Colors.white),
              )
            ],
          ),
        ),
        body: roomID == "RoomID"
            ? Container(
                color: AppColors.primaryColor,
                height: MediaQuery.of(context).size.height,
                width: MediaQuery.of(context).size.width,
                child: Column(
                  children: [
                    Expanded(
                      child: Container(
                        color: AppColors.primaryColor,
                        child: ListView.builder(
                          reverse: true,
                          itemCount: messages.length,
                          itemBuilder: (context, index) {
                            final jsonMsg = jsonDecode(messages[index]);
                            final msg = jsonMsg['content'];
                            final msgUsername = jsonMsg['username'];

                            if (msg != "A new user has joined the room") {
                              // Check for the specific message
                              return username == msgUsername
                                  ? InkWell(
                                      onLongPress: () {
                                        setState(() {
                                          isCopyTriggered = true;
                                        });
                                      },
                                      child: Container(
                                        alignment: Alignment.topRight,
                                        margin: const EdgeInsets.only(
                                          top: 8.0,
                                          right: 8.0,
                                        ), // Add margin for spacing
                                        child: Container(
                                          padding: const EdgeInsets.symmetric(
                                            vertical: 8.0,
                                            horizontal: 12.0,
                                          ), // Add padding for spacing
                                          decoration: const BoxDecoration(
                                            color: Colors.red,
                                            borderRadius: BorderRadius.only(
                                              topLeft: Radius.circular(16.0),
                                              bottomLeft: Radius.circular(16.0),
                                              bottomRight:
                                                  Radius.circular(16.0),
                                            ),
                                          ),
                                          child: Text(
                                            msg,
                                            style: const TextStyle(
                                              color: Colors.white,
                                              fontWeight: FontWeight.bold,
                                              fontSize: 16.0,
                                            ),
                                          ),
                                        ),
                                      ),
                                    )
                                  : Container(
                                      alignment: Alignment.topLeft,
                                      margin: const EdgeInsets.only(
                                        top: 8.0,
                                        left: 8.0,
                                      ), // Add margin for spacing
                                      child: Container(
                                        padding: EdgeInsets.symmetric(
                                          vertical: 8.0,
                                          horizontal: 12.0,
                                        ), // Add padding for spacing
                                        decoration: BoxDecoration(
                                          color: Colors.blue,
                                          borderRadius: BorderRadius.only(
                                            topRight: Radius.circular(16.0),
                                            bottomLeft: Radius.circular(16.0),
                                            bottomRight: Radius.circular(16.0),
                                          ),
                                        ),
                                        child: Text(
                                          msg,
                                          style: TextStyle(
                                            color: Colors.white,
                                            fontWeight: FontWeight.bold,
                                            fontSize: 16.0,
                                          ),
                                        ),
                                      ),
                                    );
                            } else {
                              return Container(); // Return an empty container for the excluded message
                            }
                          },
                        ),
                      ),
                    ),
                    Container(
                      // color: Colors.teal,
                      margin: EdgeInsets.all(12.0),
                      child: TextFormField(
                        controller: messageController,
                        style: TextStyle(color: Colors.white),
                        decoration: InputDecoration(
                          focusColor: Colors.white,
                          fillColor: Colors.red,

                          hintText: 'Message',
                          hintStyle: TextStyle(
                            color: Colors.white,
                          ),
                          enabledBorder: OutlineInputBorder(
                            borderSide: const BorderSide(
                              color: Color.fromARGB(255, 122, 129, 148),
                            ),
                            borderRadius: BorderRadius.circular(23.0),
                          ),
                          focusedBorder: OutlineInputBorder(
                            borderSide: const BorderSide(
                              color: Color.fromARGB(255, 122, 129, 148),
                            ),
                            borderRadius: BorderRadius.circular(23.0),
                          ),
                          border: OutlineInputBorder(
                            borderRadius: BorderRadius.circular(23.0),
                          ),
                          suffixIcon: InkWell(
                            onTap: () {
                              channel.sink.add(messageController.text);
                              messageController.clear();
                            },
                            child: Icon(
                              Icons.send,
                              color: Colors.white,
                            ),
                          ),
                          // prefixIcon: Icon(
                          //   Icons.emoji_emotions,
                          //   color: Colors.white,
                          // ),
                        ),
                      ),
                    )
                  ],
                ),
              )
            : Center(
                child: CircularProgressIndicator(color: Colors.blue),
              ));
  }
}
