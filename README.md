# User Microservice
this is a user microservice made for a new app im building, Giftal.

i'm still working on it. The goal is to finish the following checklist
- [x] implement TOTP
- [x] implement resend time limits
- [x] implement expiry
- [x] implement rate limit by ip address
- [x] implement WhatsApp send
- [ ] REMOVE THE DEBUG KEY OVERRIDE IN THE app_limitter middleware
- [ ] VALIDATE PHONE NUMBER ACCRINDG TO E.164 BEFORE REQUEST (HIGH PRIORITY)
> feel free to re-use this repo
 
note: to re-use it you need to create 2 files in the root directory
- .firebase.json
- .config

where .firebase.json contains the service account json file (obtainable in the firebase console)
and .config contains config env variables. 
