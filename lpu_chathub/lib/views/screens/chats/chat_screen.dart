import 'package:flutter/material.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/models/user_model.dart';

class ChatScreen extends StatefulWidget {
  final User user;
  const ChatScreen({super.key, required this.user});

  @override
  State<ChatScreen> createState() => _ChatScreenState();
}

class _ChatScreenState extends State<ChatScreen> {
  final TextEditingController messageController = TextEditingController();
  @override
  Widget build(BuildContext context) {
    return Scaffold(
      appBar: AppBar(
        iconTheme: IconThemeData(color: Colors.white),
        backgroundColor: AppColors.appBarColor,
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
      body: Container(
        color: AppColors.primaryColor,
        height: MediaQuery.of(context).size.height,
        width: MediaQuery.of(context).size.width,
        child: Column(
          children: [
            Expanded(
              child: Container(
                color: AppColors.primaryColor,
                child: ListView.builder(
                  itemCount: 50,
                  itemBuilder: (context, index) {
                    // return Text(
                    //   'Hello' + index.toString(),
                    //   style: TextStyle(color: Colors.white),
                    // );
                  },
                ),
              ),
            ),
            Container(
              // color: Colors.teal,
              margin: EdgeInsets.all(12.0),
              child: TextFormField(
                controller: messageController,
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
                      print('Hello...');
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
      ),
    );
  }
}
