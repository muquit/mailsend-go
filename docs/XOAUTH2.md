# XOAUTH2 support
@XOAUTH2@ support is available in v1.0.11-b1 (Released on Aug-24-2025). 
@MSGO@ itself does not implement full OAuth2 flow because implementing full
OAuth2 would require managing web browser redirects, secure token storage 
across different platforms and maintaing state of token refresh cycles.

Therefore, I've created a companion tool called @OAUTH_HELPER@ for token
management. This separation keeps @MSGO@ simple and secure while
giving you full control over how tokens are obtained and stored.
For automation, just pipe token from @OAUTH_HELPER@ directly into @MSGO@.

Please visit @OAUTH_HELPER@ page for details. It has examples on how to 
integrate with @MSGO@ and @XOAUTH2@.

Please create an @ISSUES@ if you need help or have any questions.
