from tests.utils import econf_compile
import json

import pytest

def test_udp():
  yaml = """
---
apiVersion: getambassador.io/v3alpha1
kind: Listener
metadata:
  name: emissary-ingress-listener-8080
  namespace: ambassador
spec:
  port: 8080
  protocol: HTTP
  securityModel: XFP
  hostBinding:
    namespace:
      from: ALL
---
apiVersion: getambassador.io/v3alpha1
kind: Listener
metadata:
  name: emissary-ingress-listener-8443
  namespace: ambassador
spec:
  port: 8443
  protocol: HTTPS
  securityModel: XFP
  hostBinding:
    namespace:
      from: ALL
---
apiVersion: getambassador.io/v3alpha1
kind: Listener
metadata:
  name: listener-udp-8443
  namespace: ambassador
spec:
  port: 8443
  protocolStack:
    - TLS
    - HTTP
    - UDP # this triggers a QUIC Socket to be used instead of TCP Socket, which already has TLS built in, no need to have TLS protocol in Stack
  securityModel: XFP
  hostBinding:
    namespace:
      from: ALL
"""
  econf = econf_compile(yaml, envoy_version="V3")
  # write it out
  with open('econf_golden.json', 'w') as outfile:
      json.dump(econf, outfile, indent=2)
  assert 1 == 1
