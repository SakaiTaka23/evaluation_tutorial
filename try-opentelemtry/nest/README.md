# nest

Nest.jsでのOpen Telemetryの使用例
参考 https://www.tomray.dev/nestjs-open-telemetry

まずは設定ファイルを作成する`tracing.ts`
spanProcessor → トレースの方法
instrumentations → 自動トレースのライブラリを入れる
[自動で入れることのできる物のリスト](https://opentelemetry.io/ecosystem/registry/?language=js&component=instrumentation)
設定ファイルができたら起動時の**一番最初**に関数を配置する

なぜか`@opentelemetry/api`がインストールされないので手動で追加

