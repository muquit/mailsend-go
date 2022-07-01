# Frequently Asked Questions (FAQ)

**1. How to send mail via smtp.gmail.com?**

From May 30, 2022, Google no longer supports the use of third-party apps to sign in to Google Account using username and password. However, an app-specific password can be set for mailsend-go to send mail via smtp.gmail.com. Here are the steps:

- Login to your gmail account
- Click on the Profile icon at the right side and then click on **Manage your Google Account**
- Click on **Security** link from the list on the left sidebar.
- Now you have to enable 2FA (Two Factor Authentication). Think carefully if you are going to do that for your main account. I used a test account and it does not seem to ask for 2FA code while sending email using smtp.gmail.com. However, it will require to enter 2FA code when you login to gmail.com. By default it sends 2FA code to your phone#, which is not secure. Configure to use Google Authenticator App instead for 2FA.
- After 2FA is enabled, Click on **Security** link again. Generate the app specific password by clicking on the **App Passwords** link. Specify mailsend-go as the app (I don't think it matters).
- Use the username and this app specific password to send mail via smtp.gmail.com. It does not seem to ask for 2FA code.

If there are any gotchas or need more clarification, please send a pull request or update Issue #49 with your experience and I will update the FAQ.

```
-- updated: Jul-01-2022
```
