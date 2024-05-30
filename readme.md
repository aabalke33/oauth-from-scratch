# OAuth 2.0 Auth Framework Implementation from Scratch

This demostration provides the Authorization Code implimentation of Oauth 2.0 from scratch in golang.
This includes the client (represented as a print service), and the authorization / resource server (represented as a photo gallery service).

[Video Tutorial](https://www.youtube.com/watch?v=iXDynkSgpZo)

[Specification](https://datatracker.ietf.org/doc/html/rfc6749)

## Abstract

When looking for Oauth implementations, I could only find demonstrations of
the client-side code, not the authorization server. To better understand
the framework, I have implemented a basic version of the Oauth 2.0 Authorization
Code Grant specification as defined in [RFC6749](https://datatracker.ietf.org/doc/html/rfc6749).

This implementation includes the authorization server, client, and resource server.

## Considerations

1. Input validation and errors are not included (see RFC6749 4.1.2.1).
2. Asymmetrical encryption for client id and client secret is not implemented.
3. Code and Access Token expiration is not implemented.
4. Refresh Token is not implemented.
5. Random integers are used for access tokens, state, and codes for simplicity.
