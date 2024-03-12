import 'package:flutter/material.dart';

class UserListTile extends StatelessWidget {
  String name;
  String profile_pic_url;
  String user_email;
  UserListTile(
      {super.key,
      required this.name,
      required this.user_email,
      required this.profile_pic_url});

  @override
  Widget build(BuildContext context) {
    return Container(
      color: Colors.white,
      child: ListTile(
        title: Text(name), // Use user!.name with null check operator
        leading: CircleAvatar(
          backgroundColor: Colors.teal,
          backgroundImage: NetworkImage(profile_pic_url),
        ),
        subtitle: Text(user_email),
      ),
    );
  }
}
