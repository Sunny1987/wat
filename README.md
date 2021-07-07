# WAT (Web Accessibility Testing) v1.0.0
An open source project on accessibility testing

Purpose: Current implementation of world-wide usage guidelines WCAG is under process for a high quality free web pages testing.
The repo is deployed in heroku platform. The services can be currently availed for WCAG 1.1.1 techniques.

*****************************************
**To login using guest credentials:**

POST https://watsab.herokuapp.com/login

body:
{
  "username":"guest",
  "password":"guest"
}
*****************************************
**To run accessibility scan:**

POST https://watsab.herokuapp.com/scan

body:
{
  "url" : "https://www.google.com",
  "depth" : 1
}
*****************************************
**To Check Last 10 scans** 

GET https://watsab.herokuapp.com/getscans

*****************************************
**To upload a html file for scan**

POST https://watsab.herokuapp.com/uploadhtml
*****************************************

Improvements:
1. Include depth assessment for websites to extract links

Last Major Enhancement:
1. DataBase support to log scans and view upto last 5 scans

Future Enhancements:
1. Add scope for WCAG 2.1 A conformance level guidelines.

