import 'package:flutter/material.dart';
import 'package:lpu_chathub/views/authentication/signup.dart';

class LoginPage extends StatefulWidget {
  const LoginPage({Key? key}) : super(key: key);

  @override
  State<LoginPage> createState() => _LoginPageState();
}

class _LoginPageState extends State<LoginPage> {
  // Boolean to track if the password is visible or not
  bool isPasswordVisible = false;

  @override
  Widget build(BuildContext context) {
    return Scaffold(
      backgroundColor: const Color.fromARGB(255, 16, 26, 36),
      body: Container(
        height: MediaQuery.of(context).size.height,
        width: MediaQuery.of(context).size.width,
        child: SingleChildScrollView(
          child: Column(
            crossAxisAlignment: CrossAxisAlignment.start,
            children: [
              // Header section
              SizedBox(
                height: MediaQuery.of(context).size.height * 0.25,
                width: MediaQuery.of(context).size.width,
                child: const Row(
                  mainAxisAlignment: MainAxisAlignment.center,
                  children: [
                    Column(
                      mainAxisAlignment: MainAxisAlignment.center,
                      children: [
                        Text(
                          'CYBER',
                          style: TextStyle(
                              fontSize: 30.0,
                              fontWeight: FontWeight.bold,
                              color: Colors.white),
                        ),
                        Text('SEC',
                            style: TextStyle(
                                fontSize: 30.0,
                                fontWeight: FontWeight.bold,
                                color: Colors.white)),
                      ],
                    ),
                    SizedBox(
                      width: 15.0,
                    ),
                    Text('SYMPOSIUM',
                        style: TextStyle(
                            fontSize: 38.0,
                            fontWeight: FontWeight.bold,
                            color: Colors.white))
                  ],
                ),
              ),
              // Login form section
              Container(
                padding: const EdgeInsets.all(20.0),
                width: MediaQuery.of(context).size.width,
                child: Column(
                  crossAxisAlignment: CrossAxisAlignment.start,
                  children: [
                    const Text(
                      'Login',
                      style: TextStyle(
                          fontSize: 34.0,
                          color: Colors.white,
                          fontWeight: FontWeight.bold),
                    ),
                    SizedBox(
                      height: MediaQuery.of(context).size.height * 0.05,
                    ),
                    // Username TextFormField
                    TextFormField(
                      style: const TextStyle(color: Colors.white),
                      decoration: const InputDecoration(
                        focusedBorder: UnderlineInputBorder(
                            borderSide: BorderSide(color: Colors.white)),
                        enabledBorder: UnderlineInputBorder(
                            borderSide: BorderSide(color: Colors.white)),
                        enabled: true,
                        labelText: 'Username',
                        labelStyle:
                            TextStyle(color: Colors.white, fontSize: 20.0),
                        prefixIcon: Icon(
                          color: Colors.white,
                          Icons.person,
                          size: 30.0,
                        ),
                        hoverColor: Colors.white,
                      ),
                    ),
                    SizedBox(
                      height: MediaQuery.of(context).size.height * 0.05,
                    ),
                    // Password TextFormField
                    TextFormField(
                      style: const TextStyle(color: Colors.white),
                      obscureText: !isPasswordVisible,
                      decoration: InputDecoration(
                        hoverColor: Colors.white,
                        focusedBorder: const UnderlineInputBorder(borderSide: BorderSide(color: Colors.white)),
                        enabledBorder: const UnderlineInputBorder(
                            borderSide: BorderSide(color: Colors.white)),
                        enabled: true,
                        labelText: 'Password',
                        labelStyle: const TextStyle(
                            color: Colors.white, fontSize: 20.0),
                        prefixIcon: IconButton(
                            color: Colors.white,
                            onPressed: () {
                              setState(() {
                                isPasswordVisible = !isPasswordVisible;
                              });
                            },
                            icon: isPasswordVisible
                                ? const Icon(Icons.visibility)
                                : const Icon(Icons.visibility_off)),
                      ),
                    ),
                    SizedBox(
                      height: MediaQuery.of(context).size.height * 0.03,
                    ),
                    // Forgot Password section
                     Row(
                      mainAxisAlignment: MainAxisAlignment.end,
                      children: [
                        TextButton(
                          onPressed:(){},
                          child:const Text('Forgot Password ?',
                          style: TextStyle(
                              fontSize: 16.0,
                              fontWeight: FontWeight.bold,
                              color: Colors.white),)
                        ),
                      ],
                    ),
                  ],
                ),
              ),
              // Login Button and Sign Up section
              SizedBox(
                height: MediaQuery.of(context).size.height * 0.31,
                child: Column(children: [
                  Container(
                    margin: const EdgeInsets.all(10.0),
                    width: MediaQuery.of(context).size.width,
                    // Login Button
                    child: ElevatedButton(
                      onPressed: () {
                       
                      },
                      style: ElevatedButton.styleFrom(
                          shape: const StadiumBorder(),
                          backgroundColor: Colors.white),
                      child: const Text(
                        'LOGIN',
                        style: TextStyle(
                            fontSize: 24.0,
                            fontWeight: FontWeight.bold,
                            color: Color.fromARGB(255, 16, 26, 36)),
                      ),
                    ),
                  ),
                  SizedBox(
                    height: MediaQuery.of(context).size.height * 0.09,
                  ),
                  // Sign Up link section
                  Row(
                    mainAxisAlignment: MainAxisAlignment.center,
                    children: [
                      const Text(
                        'Don\'t have an account ?',
                        style: TextStyle(color: Colors.white, fontSize: 18.0),
                      ),
                      TextButton(
                          onPressed: () {
                             Navigator.pushReplacementNamed(context, '/signup');
                          },
                          child: const Text(
                            'Sign Up',
                            style: TextStyle(
                                fontSize: 22.0, fontWeight: FontWeight.bold),
                          ))
                    ],
                  )
                ]),
              ),
            ],
          ),
        ),
      ),
    );
  }
}
