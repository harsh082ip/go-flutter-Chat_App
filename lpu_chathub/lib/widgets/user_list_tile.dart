import 'dart:developer';

import 'package:flutter/cupertino.dart';
import 'package:flutter/material.dart';
import 'package:lpu_chathub/controller/apis/remove_from_recents.dart';
import 'package:lpu_chathub/models/user_model.dart';

class UserListTile extends StatelessWidget {
  final String name;
  final String profile_pic_url;
  final String user_email;
  final bool diplayTrailing;
  final User? user;
  final VoidCallback refreshCallBack;
  UserListTile(
      {super.key,
      required this.name,
      required this.user_email,
      required this.profile_pic_url,
      required this.diplayTrailing,
      required this.user,
      required this.refreshCallBack});

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
        trailing: diplayTrailing
            ? IconButton(
                onPressed: () async {
                  bool refreshRequired =
                      await RemoveFromRecentApis.removeFromRecents(
                          user!.username.toString());
                  if (refreshRequired) {
                    // setState(() {
                    //   // refreshCallback();
                    // });
                    log('called');
                    refreshCallBack();
                  }
                },
                icon: Icon(CupertinoIcons.multiply))
            : null,
      ),
    );
  }
}
