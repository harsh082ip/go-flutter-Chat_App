import 'dart:async';

import 'package:image_picker/image_picker.dart';

class Profile_Pic {
  static FutureOr<XFile?> pickFile() async {
    final returnedImage =
        await ImagePicker().pickImage(source: ImageSource.gallery);
    if (returnedImage != null) {
      return returnedImage;
    }
    return null;
  }
}