// Package twistededwards leverages https://github.com/dedis/kyber/
// The kyber packages provides extended homogenous coordinate group operations.
// However, due to the nature of their API, it is difficult to use the provided library in a sane way.
// We deviate from the library by exposing previously internal APIs in such a way in order to
// use the twistededwards curve with our group parameters.
// The kyber package does not claim to have constant-time operations, therefore this package is not safe for production.
package twistededwards
