import 'package:flutter/material.dart';
import 'package:lpu_chathub/const.dart';
import 'package:lpu_chathub/controller/apis/remove_from_recents.dart';
import 'package:lpu_chathub/models/user_model.dart';

class UserListTile extends StatelessWidget {
  final String name;
  final String profilePicUrl;
  final String userEmail;
  final bool displayTrailing;
  final User? user;
  final VoidCallback refreshCallback;

  UserListTile({
    Key? key,
    required this.name,
    required this.userEmail,
    required this.profilePicUrl,
    required this.displayTrailing,
    required this.user,
    required this.refreshCallback,
  }) : super(key: key);

  @override
  Widget build(BuildContext context) {
    return Card(
      color: AppColors.appBarColor,
      elevation: 4,
      margin: EdgeInsets.symmetric(vertical: 8, horizontal: 16),
      shape: RoundedRectangleBorder(
        borderRadius: BorderRadius.circular(12),
      ),
      child: ListTile(
        contentPadding: EdgeInsets.all(16),
        leading: CircleAvatar(
          radius: 24,
          backgroundColor: Colors.teal,
          backgroundImage: NetworkImage(profilePicUrl),
        ),
        title: Text(
          name,
          style: TextStyle(fontWeight: FontWeight.bold, color: Colors.white),
        ),
        subtitle: Text(
          userEmail,
          style: TextStyle(color: Colors.white),
        ),
        trailing: displayTrailing
            ? IconButton(
                onPressed: () async {
                  bool refreshRequired =
                      await RemoveFromRecentApis.removeFromRecents(
                          user!.username.toString());
                  if (refreshRequired) {
                    refreshCallback();
                  }
                },
                icon: Icon(
                  Icons.clear,
                  color: Colors.red,
                ),
              )
            : null,
      ),
    );
  }
}
