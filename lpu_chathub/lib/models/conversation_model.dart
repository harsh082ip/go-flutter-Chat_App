class Message {
  String content;
  String username;

  Message({required this.content, required this.username});

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      content: json['content'] ?? "",
      username: json['username'] ?? "",
    );
  }
}

class ApiResponse {
  List<Message> messages;

  ApiResponse({required this.messages});

  factory ApiResponse.fromJson(Map<String, dynamic> json) {
    List<dynamic> messageList = json['messages'] ?? [];
    List<Message> messages =
        messageList.map((message) => Message.fromJson(message)).toList();

    return ApiResponse(messages: messages);
  }
}
