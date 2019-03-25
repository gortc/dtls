
# DTLS

Package dtls implements Datagram Transport Layer Security Version 1.2

Based on TLS implementation from go1.13. 
Work in progress.

## RFCs

The package aims to implement the follwing RFCs. Note that the requirement status is based on the [WebRTC spec](https://tools.ietf.org/html/draft-ietf-rtcweb-overview), focusing on data channels for now.

rfc | status | requirement | description
----|--------|-------------|----
[![RFC6347](https://img.shields.io/badge/RFC-6347-blue.svg)](https://tools.ietf.org/html/rfc6347) | ![status](https://img.shields.io/badge/status-research-orange.svg) | [![status](https://img.shields.io/badge/requirement-MUST-green.svg)](https://tools.ietf.org/html/rfc2119) | Datagram Transport Layer Security Version 1.2
[DTLS-SRTP Mux](https://tools.ietf.org/html/rfc5764#section-5.1.2) | ![status](https://img.shields.io/badge/status-research-orange.svg) | [![status](https://img.shields.io/badge/requirement-MUST-green.svg)](https://tools.ietf.org/html/rfc2119) | DTLS-SRTP Multiplexing
[![RFC7983](https://img.shields.io/badge/RFC-7983-blue.svg)](https://tools.ietf.org/html/rfc7983) | ![status](https://img.shields.io/badge/status-research-orange.svg) | [![status](https://img.shields.io/badge/requirement-MUST-green.svg)](https://tools.ietf.org/html/rfc2119) | Multiplexing Scheme Updates for SRTP Extension for DTLS
