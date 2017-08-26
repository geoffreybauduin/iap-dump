# iap-dump
[![Build Status](https://travis-ci.org/geoffreybauduin/iap-dump.svg?branch=master)](https://travis-ci.org/geoffreybauduin/iap-dump)
[![GitHub release](https://img.shields.io/github/release/geoffreybauduin/iap-dump.svg)](https://github.com/geoffreybauduin/iap-dump/releases)

A command-line application dumping data from an InAppPurchase receipt

# Installation

`go install github.com/geoffreybauduin/iap-dump`

# Usage

## Dump an InAppPurchase from the Appstore

```
$ iap-dump appstore < iap.txt
{"status":21002,"environment":"","receipt":{"receipt_type":"","adam_id":0,"app_item_id":0,"bundle_id":"","application_version":"","download_id":0,"version_external_identifier":0,"original_application_version":"","in_app":null,"receipt_creation_date":"","receipt_creation_date_ms":"","receipt_creation_date_pst":"","request_date":"","request_date_ms":"","request_date_pst":"","original_purchase_date":"","original_purchase_date_ms":"","original_purchase_date_pst":""},"latest_receipt_info":null,"latest_receipt":"","pending_renewal_info":null,"is-retryable":false}
```

# Contributing

Contributions are welcome,

Missing features:

- Handle PlayStore
- Handle Amazon store

# License

MIT License

Copyright (c) 2017 Geoffrey Bauduin

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
