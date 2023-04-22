# nest

Nest.jsでのOpen Telemetryの使用例
参考 https://www.tomray.dev/nestjs-open-telemetry

まずは設定ファイルを作成する`tracing.ts`
spanProcessor → トレースの方法
instrumentations → 自動トレースのライブラリを入れる
[自動で入れることのできる物のリスト](https://opentelemetry.io/ecosystem/registry/?language=js&component=instrumentation)
