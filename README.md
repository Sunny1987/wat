An open source project on accessibility testing

Purpose: Current implementation of world-wide usage guidelines WCAG is under process for a high quality free web pages testing. The repo is deployed in heroku platform. The services can be currently availed for WCAG 1.1.1 techniques.

GET https://watsab.herokuapp.com/scan

body: { "url" : "https://www.google.com", "depth" : 1 }

Future Enhancements:

1. Include depth assessment for websites to extract links
2. Add scope for WCAG 2.1 A conformance level guidelines.