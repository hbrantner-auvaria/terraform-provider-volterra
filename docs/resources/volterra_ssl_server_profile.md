---

page_title: "Volterra: ssl_server_profile"

description: "The ssl_server_profile allows CRUD of Ssl Server Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_ssl_server_profile
====================================

The Ssl Server Profile allows CRUD of Ssl Server Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ssl Server Profile API docs](https://docs.cloud.f5.com/docs-v2/api/ssl-server-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ssl_server_profile" "example" {
  name      = "acmecorp-web"
  namespace = "staging"
}
```

Argument Reference
------------------

### Metadata Argument Reference

`annotations` - (Optional) queryable and should be preserved when modifying objects. (`String`).

`description` - (Optional) Human readable description for the object (`String`).

`disable` - (Optional) A value of true will administratively disable the object (`Bool`).

`labels` - (Optional) by selector expression (`String`).

`name` - (Required) The value of name has to follow DNS-1035 format. (`String`).

`namespace` - (Optional) Must be a DNS_LABEL format. For a namespace object itself, namespace value will be "" (`String`).

### Spec Argument Reference

`alert_timeout` - (Optional) Specifies the alert timeout in seconds. You can also specify indefinite or immediate. (`String`).

`allow_expired_crl` - (Optional) Use the specified CRL file even if it has expired. (`String`).

`authenticate` - (Optional) Specifies frequency of authentication. The default value is once. Note that if this is set to always session cache and session ticket will be disabled. (`String`).

`authenticate_depth` - (Optional) Specifies the client certificate chain maximum traversal depth. (`Int`).

`authenticate_name` - (Optional) Specifies a Common Name (CN) that is embedded in a server certificate. The system authenticates a server based on the specified CN. (`String`).

`bypass_on_client_cert_fail` - (Optional) Specifies whether bypass SSL forward proxy traffic when fails to retrieve the Client Certificate from the BIGIP that Server ask for. The default value is disabled. (`String`).

`bypass_on_handshake_alert` - (Optional) Specifies whether bypass SSL forward proxy traffic when receiving handshake_failure(40) Alert. The default value is disabled. (`String`).

`c3d_cert_lifespan` - (Optional) Specifies the lifespan of the certificate generated using SSL client certificate constrained delegation. The default value is 24 hours. (`Int`).

`cache_size` - (Optional) Specifies the SSL session cache size. For client-side profiles only, you can configure timeout and size values for the SSL session cache. Because each profile maintains a separate SSL session cache, you can configure the values on a per-profile basis. (`Int`).

`cache_timeout` - (Optional) Specifies the SSL session cache timeout value, which is the usable lifetime seconds of negotiated SSL session IDs. The default value is 3600 seconds. Acceptable values are integers greater than or equal to 0 and less than or equal to 86400. (`Int`).

###### One of the arguments from this list "cipher_group, ciphers" can be set

`cipher_group` - (Optional) Specifies an associated cipher group. (`String`).

`ciphers` - (Optional) Specifies a cipher name. (`String`).

`data_0rtt` - (Optional) Specifies if TLSv1.3 should send 0-RTT early data when available. The default value is disabled. (`String`).

`expire_cert_response_control` - (Optional) Specifies the action for the BIG-IP system to take when the server certificate has expired. The default value is drop, which causes the connection to be dropped. Conversely, you can specify ignore to cause the connection to ignore the error and continue or you can specify mask in case of SSL forward proxy to mask server certificate errors and continue with handshake and forge a good certificate on client-side. Note that drop works only if the certificate is trusted. (`String`).

`generic_alert` - (Optional) Enables or disables generic-alert which if use generic alert number in Alert message. The default option is enabled. (`String`).

`handshake_timeout` - (Optional) Specifies the handshake timeout in seconds. You can also specify indefinite. (`Int`).

`max_active_handshakes` - (Optional) Specifies the maximum allowed active handshakes. The default value is 0. (`Int`).

`mod_ssl_methods` - (Optional) Enables or disables ModSSL method emulation. Use enabled when OpenSSL methods are inadequate. For example, you can enable ModSSL method emulation when you want to use SSL compression over TLSv1. (`String`).

`mode` - (Optional) Enables or disables SSL processing. The default value is enabled. (`String`).

`ocsp` - (Optional) Specifies the name of ocsp profile for purpose of validating status of server certificate. Specifying none disables ocsp validation of server certificate. The default value is none.. See [ref](#ref) below for details.

`peer_cert_mode` - (Optional) Specifies the peer certificate mode. (`String`).

`proxy_ssl` - (Optional) Enables proxy SSL mode, which requires a corresponding client SSL profile with proxy-ssl enabled to allow for modification of application data within an SSL tunnel. (`String`).

`proxy_ssl_passthrough` - (Optional) Enables proxy SSL passthrough mode, which requires a corresponding client SSL profile with proxy-ssl-passthrough enabled to allow for modification of application data within an SSL tunnel. (`String`).

`renegotiate_period` - (Optional) Specifies the number of seconds from the initial connect time after which the system renegotiates an SSL session. The default value is indefinite meaning that you do not want the system to renegotiate SSL sessions. Each time the session renegotiation is successful, a new connection is started. Therefore, the system attempts to renegotiate the session again, in the specified amount of time following the successful session renegotiation. For example, setting the Renegotiate Period to 3600 seconds triggers session renegotiation at least once an hour. (`String`).

`renegotiate_size` - (Optional) Specifies a throughput size, in bytes, of SSL renegotiation. This setting forces the traffic management system to renegotiate an SSL session based on the size, in megabytes, of application data that is transmitted over the secure channel. The default value is indefinite specifying that you do not want a throughput size. (`String`).

`renegotiation` - (Optional) Controls mid-stream renegotiation. The default value is enabled. (`String`).

`retain_certificate` - (Optional) When true, server certificate is retained in SSL session. (`String`).

`revoked_cert_status_response_control` - (Optional) Specifies the system action when the server certificate status is revoked. The default value is drop, which causes the connection to be dropped. You can specify ignore to cause the connection to ignore the error and continue handshake. You can specify mask in case of SSL forward proxy to mask server certificate status error and continue handshake. (`String`).

`secure_renegotiation` - (Optional) Controls secure renegotiation. The default value is require-strict. (`String`).

`server_name` - (Optional) Name matched to TLS/1.1 and above client SSL requests that support the Server Name Indication extension. The default value is empty, which disables support for this extension. (`String`).

`session_mirroring` - (Optional) Enables or disables session mirroring to the high-availability peer. The default option is disabled. (`String`).

`session_ticket` - (Optional) Enables or disables session-ticket. The default option is disabled, see RFC5077. (`String`).

`sni_default` - (Optional) When true, this profile is the default SSL profile when a client connection does not specify a known server name, or does not specify any server name at all. The default value is false. (`String`).

`sni_require` - (Optional) When true, SNI support is required for the peer. If a client connection does not specify a known server name, or does not specify any server name at all, the connection will be rejected. The default value is false. (`String`).

`ssl_c3d` - (Optional) Enables or disables SSL client certificate constrained delegation. The default option is disabled. Conversely, you can specify enabled to use the SSL client certificate constrained delegation. (`String`).

`ssl_forward_proxy` - (Optional) Specifies if the SSL Forward Proxy feature is enabled. The default value is disabled. (`String`).

`ssl_forward_proxy_bypass` - (Optional) Specifies if the SSL Forward Proxy Bypass feature is enabled. The default value is disabled. (`String`).

`ssl_forward_proxy_verified_handshake` - (Optional) x-displayName: "SSL Forward Proxy Verified Handshake" (`String`).

`ssl_sign_hash` - (Optional) SSL sign hash algorithm to sign and verify SSL Server Key Exchange and Certificate Verify messages for the specified SSL profiles. (`String`).

`strict_resume` - (Optional) You can enable or disable the resumption of SSL sessions after an unclean shutdown. The default value is disabled, which indicates that the SSL profile refuses to resume SSL sessions after an unclean shutdown. (`String`).

`unclean_shutdown` - (Optional) By default, the SSL profile performs unclean shutdowns of all SSL connections, which means that underlying TCP connections are closed without exchanging the required SSL shutdown alerts. If you want to force the SSL profile to perform a clean shutdown of all SSL connections, you can disable this. (`String`).

`unknown_cert_status_response_control` - (Optional) Specifies the system action when the server certificate status is unknown. The default value is ignore, which causes the connection to ignore the error and continue handshake. You can specify drop which causes the connection to be dropped. You can specify mask in case of SSL forward proxy to mask server certificate status error and continue handshake. (`String`).

`untrusted_cert_response_control` - (Optional) Specifies the system action when the server certificate has untrusted CA. The default value is drop, which causes the connection to be dropped. Conversely, you can specify ignore to cause the connection to ignore the error and continue or you can specify mask in case of SSL forward proxy to mask server certificate errors and continue with handshake and forge a good certificate on client-side. (`String`).

`user_spec` - (Optional) User specified properties. (`List of String`).

### Ref

Reference to another volterra object is shown like below

name - (Required) then name will hold the referred object's(e.g. route's) name. (String).

namespace - (Optional) then namespace will hold the referred object's(e.g. route's) namespace. (String).

tenant - (Optional) then tenant will hold the referred object's(e.g. route's) tenant. (String).

Attribute Reference
-------------------

*   `id` - This is the id of the configured ssl_server_profile.
