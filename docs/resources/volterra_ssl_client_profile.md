---

page_title: "Volterra: ssl_client_profile"

description: "The ssl_client_profile allows CRUD of Ssl Client Profile resource on Volterra SaaS"
-------------------------------------------------------------------------------------------------

Resource volterra_ssl_client_profile
====================================

The Ssl Client Profile allows CRUD of Ssl Client Profile resource on Volterra SaaS

~> **Note:** Please refer to [Ssl Client Profile API docs](https://docs.cloud.f5.com/docs-v2/api/ssl-client-profile) to learn more

Example Usage
-------------

```hcl
resource "volterra_ssl_client_profile" "example" {
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

`allow_dynamic_record_sizing`- (Optional) Enables or disables dynamic application record sizing. Specify enabled when you want to allow dynamic record sizing. The default value is disabled. (`String`).

`allow_expired_crl` - (Optional) Use the specified CRL file even if it has expired. (`String`).

`allow_non_ssl` - (Optional) Enables or disables non-SSL connections. Specify enabled to allow non-SSL connections to pass through the traffic management system as clear text. (`String`).

`app_service` - (Optional) The application service to which the object belongs. (`String`).(Deprecated)

`authenticate` - (Optional) Specifies frequency of authentication. The default value is once. Note that if this is set to always session cache and session ticket will be disabled. (`String`).

`authenticate_depth` - (Optional) Specifies the authenticate depth. This is the client certificate chain maximum traversal depth. (`Int`).

`c3d_drop_unknown_ocsp_status` - (Optional) Specifies the BIG-IP action when the OCSP returns unknown status. The default value is drop, which causes the connection to be dropped. Conversely, you can specify ignore to cause the connection to ignore the unknown status and continue. (`String`).

`cache_size` - (Optional) Specifies the SSL session cache size. For client-side profiles only, you can configure timeout and size values for the SSL session cache. Because each profile maintains a separate SSL session cache, you can configure the values on a per-profile basis. (`Int`).

`cache_timeout` - (Optional) Specifies the SSL session cache timeout value. This specifies the number of usable lifetime seconds of negotiated SSL session IDs. The default timeout value for the SSL session cache is 3600 seconds. Acceptable values are integers greater than or equal to 0 and less than or equal to 86400. (`Int`).

`cert_extension_includes` - (Optional) Specifies the extensions of the web server certificates to be included in the generated certificates using SSL Forward Proxy. The default value is ['basic-constraints', 'subject-alternative-name']. (`List of String`).(Deprecated)

`cert_lookup_by_ipaddr_port` - (Optional) Specifies whether SSL forward proxy lookup certificate by ipaddr/port feature is enabled or not. The default value is disabled. (`String`).

###### One of the arguments from this list "cipher_group, ciphers" can be set

`cipher_group` - (Optional) Specifies an associated cipher group. (`String`).

`ciphers` - (Optional) Specifies a cipher name. (`String`).

`crl` - (Optional) x-displayName: "Crl" (`String`).(Deprecated)

`crl_file` - (Optional) Specifies the certificate revocation list file name. (`String`).(Deprecated)

`data_0rtt` - (Optional) Specifies if TLSv1.3 should accept 0-RTT with early data, with or without anti-replay. To protect against packet replay, F5 recommends that you enable anti-replay. The default value is disabled, which means TLSv1.3 will discard any early data. (`String`).

`forward_proxy_bypass_default_action` - (Optional) Specifies the SSL forward proxy bypass default action. The default value is intercept. (`String`).

`generic_alert` - (Optional) Enables or disables generic-alert which if use generic alert number in Alert message. The default option is enabled. (`String`).

`handshake_timeout` - (Optional) Specifies the handshake timeout in seconds. You can also specify indefinite. (`Int`).

`hello_extension_includes` - (Optional) Specifies the hello extensions received from client to be included in hello extensions sent to the server by SSL Forward Proxy. The default value is [ 'application-layer-protocol-negotiation' ]. (`List of String`).

`inherit_ca_certkeychain` - (Optional) x-displayName: "Inherit CA Certkeychain" (`String`).

`inherit_certkeychain` - (Optional) This is a read only value used internally. (`String`).

`log_ssl_c3d_events` - (Optional) Configures the log level above which system logs ssl c3d events for this profile. (`String`).(Deprecated)

`log_ssl_client_authentication_events` - (Optional) Configures the log level above which system logs ssl client authentication events for this profile. (`String`).(Deprecated)

`log_ssl_forward_proxy_events` - (Optional) Configures the log level above which system logs ssl forward proxy events for this profile. (`String`).(Deprecated)

`log_ssl_handshake_events` - (Optional) Configures the log level above which system logs ssl handshake events for this profile. (`String`).(Deprecated)

`max_active_handshakes` - (Optional) Specifies the maximum allowed active handshakes. The default value is 0. (`Int`).

`max_aggregate_renegotiation_per_minute` - (Optional) Specifies the maximum number of aggregate renegotiation attempts allowed in a minute. The default value is indefinite. (`String`).

`max_renegotiations_per_minute` - (Optional) Specifies the maximum number of renegotiation attempts allowed in a minute. The default value is 5. (`Int`).

`maximum_record_size` - (Optional) Specifies the profile's maximum record size. The range is 128 - 16384. The default value is 16384. (`Int`).

`mod_ssl_methods` - (Optional) Enables or disables ModSSL method emulation. Enable this option when OpenSSL methods are inadequate. For example, you can enable this option when you want to use SSL compression over TLSv1. (`String`).

`mode` - (Optional) Specifies the profile mode, which enables or disables SSL processing. The default value is enabled. (`String`).

`notify_cert_status_to_virtual_server` - (Optional) Specifies whether to propagate the status of the certificates of this clientssl profile to the virtual servers that are using this clientssl profile. (`String`).

`ocsp_stapling` - (Optional) Specifies whether to enable OCSP stapling. (`String`).

`peer_cert_mode` - (Optional) Specifies the peer certificate mode. (`String`).

`peer_no_renegotiate_timeout`- (Optional) Specifies the number of seconds that the system waits for ClientHello before sending Fatal Alert after sending Hello Request. The default is 10 seconds. You can set it to Indefinite which specifies that the system continue to wait for ClientHello for an unlimited time. (`Int`).

`proxy_ca_cert` - (Optional) Specifies the Certification Authority cert for SSL Forward Proxy. (This option is deprecated beginning in v14.0.0. Instead, you can use the cert-key-chain option with the usage argument to add an FWDP CA key/cert.) (`String`).(Deprecated)

`proxy_ca_key` - (Optional) Specifies the Certification Authority key for SSL Forward Proxy. (This option is deprecated beginning in v14.0.0. Instead, you can use the cert-key-chain option with the usage argument to add an FWDP CA key/cert.) (`String`).(Deprecated)

`proxy_ca_passphrase` - (Optional) Specifies the passphrase of the Certification Authority key for SSL Forward Proxy. (This option is deprecated beginning in v14.0.0. Instead, you can use the cert-key-chain option with the usage argument to add an FWDP CA key/cert.) (`String`).(Deprecated)

`proxy_ssl` - (Optional) Enables proxy SSL mode, which requires a corresponding server SSL profile with proxy-ssl enabled to allow for modification of application data within an SSL tunnel. (`String`).

`proxy_ssl_passthrough` - (Optional) Enables proxy SSL passthrough mode, which requires a corresponding server SSL profile with proxy-ssl-passthrough enabled to allow for modification of application data within an SSL tunnel. (`String`).

`renegotiate_max_record_delay` - (Optional) Specifies the maximum number of SSL records that the traffic management system can receive before it renegotiates an SSL session. After the system receives this number of SSL records, it closes the connection. This setting applies to client-side profiles only. The default value is indefinite. (`String`).

`renegotiate_period` - (Optional) Specifies the number of seconds required to renegotiate an SSL session. The default value is indefinite. (`String`).

`renegotiate_size` - (Optional) Specifies the size of the application data, in megabytes, that is transmitted over the secure channel above which the traffic management system must renegotiate the SSL session. The default value is indefinite. (`String`).

`renegotiation` - (Optional) Controls mid-stream renegotiation. The default value is enabled. (`String`).

`retain_certificate` - (Optional) When true, client certificate is retained in SSL session. (`String`).

`secure_renegotiation` - (Optional) Controls secure renegotiation. The default value is require. (`String`).

`server_name` - (Optional) Name matched to TLS/1.1 and above client SSL requests that support the Server Name Indication extension. The default value is empty, which disables support for this extension. (`String`).

`session_mirroring` - (Optional) Enables or disables session mirroring to the high-availability peer. The default option is disabled. (`String`).

`session_ticket` - (Optional) Enables or disables session-ticket. The default option is disabled, see RFC5077. (`String`).

`session_ticket_timeout` - (Optional) Specifies the session ticket timeout. The default value is 0. (`Int`).

`sni_default` - (Optional) When true, this profile is the default SSL profile when a client connection does not specify a known server name, or does not specify any server name at all. The default value is false. (`String`).

`sni_require` - (Optional) When true, SNI support is required for the peer and if a client connection does not specify a known server name, or does not specify any server name at all, the handshake will fail. The default value is false. (`String`).

`ssl_c3d` - (Optional) Enables or disables SSL client certificate constrained delegation. The default option is disabled. Conversely, you can specify B<enabled> to use the SSL client certificate constrained delegation. (`String`).

`ssl_forward_proxy` - (Optional) Specifies whether SSL forward proxy feature is enabled or not. The default value is disabled. (`String`).

`ssl_forward_proxy_bypass` - (Optional) Specifies whether SSL forward proxy bypass feature is enabled or not. The default value is disabled. (`String`).

`ssl_forward_proxy_verified_handshake` - (Optional) x-displayName: "SSL Forward Proxy Verified Handshake" (`String`).

`ssl_sign_hash` - (Optional) SSL sign hash algorithm to sign and verify SSL Server Key Exchange and Certificate Verify messages for the specified SSL profiles. (`String`).

`strict_resume` - (Optional) Enables or disables strict-resume. The default option is disabled, which causes the SSL profile to allow uncleanly shut down SSL sessions to be resumed. Conversely, you can specify enabled to prevent an SSL session from being resumed after an unclean shutdown. (`String`).

`tm_options` - (Optional) Enables options, including some industry-related workarounds. Enter options inside braces, for example, { dont-insert-empty-fragments }. The default value is { dont-insert-empty-fragments no-tlsv1.3 }. (`List of String`).

`unclean_shutdown` - (Optional) By default, the SSL profile performs unclean shutdowns of all SSL connections, which means that underlying TCP connections are closed without exchanging the required SSL shutdown alerts. If you want to force the SSL profile to perform a clean shutdown of all SSL connections, set this option to disabled. (`String`).

`user_spec` - (Optional) User specified properties. (`List of String`).

Attribute Reference
-------------------

*   `id` - This is the id of the configured ssl_client_profile.
