class User {
  String id;
  String name;
  String email;
  String password;
  String username;
  String userid;
  String profilePicUrl;

  User({
    required this.id,
    required this.name,
    required this.email,
    required this.password,
    required this.username,
    required this.userid,
    required this.profilePicUrl,
  });

  // Factory constructor to create a User object from JSON
  factory User.fromJson(Map<String, dynamic> json) {
    String id = json['_id'] != null ? json['_id']['\$oid'] : "";
    return User(
      id: id,
      name: json['name'] ?? "",
      email: json['email'] ?? "",
      password: json['password'] ?? "",
      username: json['username'] ?? "",
      userid: json['userid'] ?? "",
      profilePicUrl: json['profile_pic_url'] ?? "",
    );
  }

  // Static method to convert User object to JSON format
  static Map<String, dynamic> toJson(User user) {
    return {
      'id': user.id,
      'name': user.name,
      'email': user.email,
      'password': user.password,
      'username': user.username,
      'userid': user.userid,
      'profile_pic_url': user.profilePicUrl,
    };
  }
}
