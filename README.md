# gcpmetadata
GCPのmetadata周りを扱うutility

GCP上では [Metadata Server](https://cloud.google.com/compute/docs/storing-retrieving-metadata) を参照し、Localでは環境変数を利用し、設定値を取得する。
GCPとLocalで同じロジックを使いつつ、設定値を持ち回るために作成した

## 動作確認した環境

* Google App Engine Standard for Go 1.11
* Google Compute Engine
* Google Kubernetes Engine